/*

Package wsql implements a postgres interface with logging, metrics, and tracing middleware.

*/
package wsql

import (
	"database/sql"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/opentracing/opentracing-go"
	"weavelab.xyz/wlib/werror"
	"weavelab.xyz/wlib/wlog"
	"weavelab.xyz/wlib/wlog/tag"
	"weavelab.xyz/wlib/wmetrics"
	"weavelab.xyz/wlib/wtracer"
)

const (
	pgPrimaryName = "primary"
	pgReplicaName = "replica"
)

type PG struct {
	db        *DB
	dbReplica *DB

	tracer opentracing.Tracer

	LogQueries bool
	setupLock  sync.Mutex

	settings *Settings
}

type DB struct {
	stopMetrics chan struct{}

	Name string
	xdb  *sqlx.DB
}

type Settings struct {
	PrimaryConnectString ConnectString
	ReplicaConnectString ConnectString

	MaxIdleConnections    int
	MaxOpenConnections    int
	MaxConnectionLifetime time.Duration

	LogQueries bool
}

func New(s *Settings) (*PG, error) {

	t, err := wtracer.DefaultTracer()
	if err != nil {
		return nil, werror.Wrap(err)
	}

	var p = PG{
		LogQueries: s.LogQueries,
		settings:   s,
		tracer:     t,
	}

	err = p.setupDatabase(s)
	if err != nil {
		return nil, err
	}

	return &p, nil
}

func (p *PG) setupDatabase(s *Settings) error {
	p.setupLock.Lock()
	defer p.setupLock.Unlock()

	// don't log database credentials
	wlog.Info("Connecting to primary database")

	//---------------------------------------------
	// setup a read-write connection to the primary
	//---------------------------------------------
	dbPrimary, err := p.connect(s, pgPrimaryName, s.PrimaryConnectString)
	if err != nil {
		return werror.Wrap(err, "unable to connect to primary database server")
	}

	// initially set replica connection to primary connection
	dbReplica := dbPrimary

	//--------------------------------------------
	// setup a read-only connection to the replica
	//--------------------------------------------
	rcs := s.ReplicaConnectString.String()
	if rcs != "" {
		wlog.Info("Connecting to replica database")

		dbReplica, err = p.connect(s, pgReplicaName, s.ReplicaConnectString)
		if err != nil {
			return werror.Wrap(err, "unable to connect to replica database server")
		}
	}

	// swap in the new database connections
	// shut down metric reporting on the old connections
	same := false
	if p.db == p.dbReplica {
		same = true
	}

	if p.db != nil {
		p.db.SetMaxIdleConns(0)
		close(p.db.stopMetrics)

		atomic.StorePointer((*unsafe.Pointer)(unsafe.Pointer(&p.db)), unsafe.Pointer(dbPrimary))
	} else {
		p.db = dbPrimary
	}

	if p.dbReplica != nil {
		p.dbReplica.SetMaxIdleConns(0)
		if same == false {
			close(p.dbReplica.stopMetrics)
		}
		atomic.StorePointer((*unsafe.Pointer)(unsafe.Pointer(&p.dbReplica)), unsafe.Pointer(dbReplica))
	} else {
		p.dbReplica = dbReplica
	}

	return nil

}

func (p *PG) connect(s *Settings, name string, cs ConnectString) (*DB, error) {

	css := cs.String()

	if cs.Driver == "" {
		cs.Driver = "postgres"
	}

	conn, err := sqlx.Connect(cs.Driver, css)
	if err != nil {
		return nil, werror.Wrap(err).Add("csHost", cs.Host).Add("csSet", cs.connectString != "")
	}

	db := &DB{
		xdb:         conn,
		Name:        name,
		stopMetrics: make(chan struct{}),
	}

	if s.MaxConnectionLifetime == 0 {
		s.MaxConnectionLifetime = time.Minute * 5
	}

	primaryLifetime := randomLifetime(s.MaxConnectionLifetime)

	wlog.Info("Setting db connection settings",
		tag.Int("max_idle", s.MaxIdleConnections),
		tag.Int("max_open", s.MaxOpenConnections),
		tag.Duration("max_connection_lifetime", primaryLifetime))

	db.SetMaxOpenConns(s.MaxOpenConnections)
	db.SetMaxIdleConns(s.MaxIdleConnections)
	db.SetConnectionMaxLifetime(primaryLifetime)

	go db.SendConnectionStatistics()

	return db, nil
}

func randomLifetime(desired time.Duration) time.Duration {

	// take a random number from [0 and 7.5% of the desired lifetime]
	// and subtract it from the lifetime so that all of the connection
	// instances across pods aren't synchronized
	r := rand.Intn(int(float64(desired) * 0.075))

	n := time.Duration(int(desired) - r)

	return n
}

func (p *DB) SendConnectionStatistics() {

	ticker := time.NewTicker(time.Second * 10)
	defer ticker.Stop()

	description := "connection_pool." + p.Name + ".open"

loop:
	for {
		select {
		case <-ticker.C:
			s := p.Stats()

			wmetrics.Gauge(s.OpenConnections, description)
		case <-p.stopMetrics:
			break loop
		}
	}

}

func (p *DB) SetMaxIdleConns(i int) {
	p.xdb.SetMaxIdleConns(i)
}

func (p *DB) SetMaxOpenConns(i int) {
	p.xdb.SetMaxOpenConns(i)
}

func (p *DB) SetConnectionMaxLifetime(maxLifetime time.Duration) {
	p.xdb.SetConnMaxLifetime(maxLifetime)
}

func (p *DB) Stats() sql.DBStats {
	return p.xdb.Stats()
}

func (p *PG) Ping() error {
	return p.db.xdb.Ping()
}

func (p *PG) Close() error {
	p.db.xdb.Close()
	p.dbReplica.xdb.Close()

	return nil
}

func (p *PG) UpdateCredentials(username string, password string) error {

	if p.settings.PrimaryConnectString.connectString != "" {
		return werror.New("primary connection string can not be changed")
	}

	if p.settings.ReplicaConnectString.connectString != "" {
		return werror.New("replica connection string can not be changed")
	}

	p.settings.PrimaryConnectString.Username = username
	p.settings.PrimaryConnectString.Password = password

	p.settings.ReplicaConnectString.Username = username
	p.settings.ReplicaConnectString.Password = password

	err := p.setupDatabase(p.settings)
	if err != nil {
		return werror.Wrap(err, "unable to update connection strings")
	}

	return nil
}

func (p *PG) rw() *sqlx.DB {
	db := atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&p.db)))

	return (*DB)(db).xdb
}

func (p *PG) r() *sqlx.DB {
	db := atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&p.dbReplica)))

	return (*DB)(db).xdb
}
