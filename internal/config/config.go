package config

import (
	"time"

	"weavelab.xyz/monorail/shared/wlib/config"
	"weavelab.xyz/monorail/shared/wlib/werror"
)

const (
	primaryConnString = "pg-primary-connect-string"
	replicaConnString = "pg-replica-connect-string"

	dbPrimaryAddr = "db-primary-addr"
	dbReplicaAddr = "db-replica-addr"
	dbName        = "db-name"

	maxIdleConnections    = "max-idle-connections"
	maxOpenConnections    = "max-open-connections"
	maxConnectionLifetime = "max-connection-lifetime"
	logQueries            = "log-queries"
)

var (
	PrimaryConnString string
	ReplicaConnString string

	DBPrimaryAddr string
	DBReplicaAddr string
	DBName        string

	MaxIdleConnections    int
	MaxOpenConnections    int
	MaxConnectionLifetime time.Duration
	LogQueries            bool
)

func init() {
	config.Add(primaryConnString, "", "connection string to the primary db")
	config.Add(replicaConnString, "", "connection string to the replica db")

	config.Add(dbPrimaryAddr, "", "primary database server host:port")
	config.Add(dbReplicaAddr, "", "replica database server host:port")
	config.Add(dbName, "", "database name")

	config.Add(maxIdleConnections, "0", "maximum number of connections in the idle connection pool")
	config.Add(maxOpenConnections, "10", "maximum number of open connections to the database")
	config.Add(maxConnectionLifetime, "15m", "maximum amount of time a connection may be reused")
	config.Add(logQueries, "false", "include query logging")

}

func Init() error {
	config.Parse()

	var err error

	PrimaryConnString = config.Get(primaryConnString)
	ReplicaConnString = config.Get(replicaConnString)

	primaryAddr := config.Get(dbPrimaryAddr)
	if primaryAddr != "" {
		DBPrimaryAddr, err = config.GetAddress(dbPrimaryAddr, false)
		if err != nil {
			return werror.Wrap(err, "unable to get primary database address")
		}
	}

	replicaAddr := config.Get(dbReplicaAddr)
	if replicaAddr != "" {
		DBReplicaAddr, err = config.GetAddress(dbReplicaAddr, false)
		if err != nil {
			return werror.Wrap(err, "unable to get replica database address")
		}
	}

	DBName = config.Get(dbName)

	MaxIdleConnections, err = config.GetInt(maxIdleConnections, false)
	if err != nil {
		return werror.Wrap(err, "error getting maxIdleConnections")
	}

	MaxOpenConnections, err = config.GetInt(maxOpenConnections, false)
	if err != nil {
		return werror.Wrap(err, "error getting maxOpenConnections")
	}

	MaxConnectionLifetime, err = config.GetDuration(maxConnectionLifetime, false)
	if err != nil {
		return werror.Wrap(err, "error getting maxConnectionLifetime")
	}

	LogQueries, err = config.GetBool(logQueries, false)
	if err != nil {
		return werror.Wrap(err, "error getting logQueries")
	}

	return nil
}
