package config

import (
	"time"

	"weavelab.xyz/monorail/shared/wlib/config"
	"weavelab.xyz/monorail/shared/wlib/werror"
)

const (
	// Database Settings
	primaryConnString = "pg-primary-connect-string"
	replicaConnString = "pg-replica-connect-string"

	dbPrimaryAddr = "db-primary-addr"
	dbReplicaAddr = "db-replica-addr"
	dbName        = "db-name"

	dbDriver   = "db-driver"
	dbPassword = "db-password"
	dbUsername = "db-username"

	maxIdleConnections    = "max-idle-connections"
	maxOpenConnections    = "max-open-connections"
	maxConnectionLifetime = "max-connection-lifetime"
	logQueries            = "log-queries"

	// NSQ Settings
	nsqConcurrentHandlersConfig = "nsq-concurrent-handlers"
	nsqMaxInFlightConfig        = "nsq-max-in-flight"
	nsqdAddrConfig              = "nsqd-addr"
	nsqLookupdAddrsConfig       = "nsq-lookupd-addrs"
	nsqChannelConfig            = "nsq-listen-channel"

	nsqChiliPiperScheduleEventCreatedTopic = "nsq-chili-piper-schedule-event-created-topic"
	nsqPortingDataRecordCreatedTopic       = "nsq-porting-data-record-created-topic"
	nsqLoginEventCreatedTopic              = "nsq-login-event-created-topic"

	// grpc client settings
	authServiceAddress      = "auth-service-address"
	featureFlagsAddress     = "feature-flags-address"
	portingDataGRPCAddress  = "porting-data-grpc-address"
	provisioningGRPCAddress = "provisioning-grpc-address"

	// zapier
	zapierURL = "zapier-url"
)

var (
	// Database Settings
	PrimaryConnString string
	ReplicaConnString string

	DBPrimaryAddr string
	DBReplicaAddr string
	DBName        string

	DBDriver   string
	DBPassword string
	DBUsername string

	MaxIdleConnections    int
	MaxOpenConnections    int
	MaxConnectionLifetime time.Duration
	LogQueries            bool

	// NSQ Settings
	NSQDAddr       string
	NSQLookupAddrs []string
	NSQChannel     string

	NSQChiliPiperScheduleEventCreatedTopic string
	NSQPortingDataRecordCreatedTopic       string
	NSQLoginEventCreatedTopic              string

	NSQMaxInFlight        int
	NSQConcurrentHandlers int

	AuthServiceAddr         string
	FeatureFlagsAddr        string
	PortingDataGRPCAddr     string
	ProvisioningGRPCAddress string

	ZapierURL string
)

func init() {
	// Database Settings
	config.Add(primaryConnString, "", "connection string to the primary db")
	config.Add(replicaConnString, "", "connection string to the replica db")

	config.Add(dbPrimaryAddr, "", "primary database server host:port")
	config.Add(dbReplicaAddr, "", "replica database server host:port")
	config.Add(dbName, "", "database name")

	config.Add(dbDriver, "", "database driver")
	config.Add(dbPassword, "", "database user password")
	config.Add(dbUsername, "", "database username")

	config.Add(maxIdleConnections, "0", "maximum number of connections in the idle connection pool")
	config.Add(maxOpenConnections, "10", "maximum number of open connections to the database")
	config.Add(maxConnectionLifetime, "15m", "maximum amount of time a connection may be reused")
	config.Add(logQueries, "false", "include query logging")

	// NSQ Settings
	config.Add(nsqChannelConfig, "Onboarding", "The channel on which to consume")
	config.Add(nsqdAddrConfig, "nsqd.nsq.svc.cluster.local.:4150", "nsqd addresses")
	config.Add(nsqLookupdAddrsConfig, "lookupd-0.lookupd.nsq.svc.cluster.local.:4161;lookupd-1.lookupd.nsq.svc.cluster.local.:4161;lookupd-2.lookupd.nsq.svc.cluster.local.:4161", "NSQ lookupd addresses")

	config.Add(nsqChiliPiperScheduleEventCreatedTopic, "ChiliPiperScheduleEventCreated", "The topic NSQ to consume for chili piper created events")
	config.Add(nsqPortingDataRecordCreatedTopic, "PortingDataCreated", "The topic NSQ to consume for porting data record created events")
	config.Add(nsqLoginEventCreatedTopic, "LoginEvent", "Platform Auth's Login Event is published whenever a client logs in")

	config.Add(nsqConcurrentHandlersConfig, "100", "Number of concurrent handlers")
	config.Add(nsqMaxInFlightConfig, "1000", "NSQ config number of times to attempt a message")

	config.Add(authServiceAddress, "auth-api.auth.svc.cluster.local.:grpc", "The grpc address of the auth service")
	config.Add(featureFlagsAddress, "client-feature-flags.client.svc.cluster.local.:grpc", "The grpc address of the feature flags service")
	config.Add(portingDataGRPCAddress, "insys-porting-data.insys.svc.cluster.local.:grpc", "The grpc address of the Porting Data service")
	config.Add(provisioningGRPCAddress, "insys-provisioning.insys.svc.cluster.local.:grpc", "The grpc address of the Provisioning Service")

	config.Add(zapierURL, "https://hooks.zapier.com/hooks/catch/883949/o246fjf", "The address zapier webhook")
}

func Init() error {
	config.Parse()

	var err error

	PrimaryConnString = config.Get(primaryConnString)
	ReplicaConnString = config.Get(replicaConnString)

	DBPrimaryAddr = config.Get(dbPrimaryAddr)
	DBReplicaAddr = config.Get(dbReplicaAddr)

	DBName = config.Get(dbName)
	DBPassword = config.Get(dbPassword)
	DBUsername = config.Get(dbUsername)
	DBDriver = config.Get(dbDriver)

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

	// NSQ Settings
	NSQLookupAddrs, err = config.GetAddressArray(nsqLookupdAddrsConfig, false)
	if err != nil {
		return werror.Wrap(err, "error getting lookupd-addrs")
	}

	NSQDAddr, err = config.GetAddress(nsqdAddrConfig, false)
	if err != nil {
		return werror.Wrap(err, "error getting nsqd-addr")
	}

	NSQChiliPiperScheduleEventCreatedTopic = config.Get(nsqChiliPiperScheduleEventCreatedTopic)
	NSQPortingDataRecordCreatedTopic = config.Get(nsqPortingDataRecordCreatedTopic)
	NSQLoginEventCreatedTopic = config.Get(nsqLoginEventCreatedTopic)

	NSQChannel = config.Get(nsqChannelConfig)

	NSQMaxInFlight, err = config.GetInt(nsqMaxInFlightConfig, false)
	if err != nil {
		return werror.Wrap(err, "error getting max attempts")
	}

	NSQConcurrentHandlers, err = config.GetInt(nsqConcurrentHandlersConfig, false)
	if err != nil {
		return werror.Wrap(err, "error getting concurrent handlers")
	}

	AuthServiceAddr, err = config.GetAddress(authServiceAddress, false)
	if err != nil {
		return werror.Wrap(err, "error getting auth service grpc address")
	}

	FeatureFlagsAddr, err = config.GetAddress(featureFlagsAddress, false)
	if err != nil {
		return werror.Wrap(err, "error getting feature flags grpc address")
	}

	PortingDataGRPCAddr, err = config.GetAddress(portingDataGRPCAddress, false)
	if err != nil {
		return werror.Wrap(err, "error getting porting data grpc address")
	}

	ProvisioningGRPCAddress, err = config.GetAddress(provisioningGRPCAddress, false)
	if err != nil {
		return werror.Wrap(err, "error getting provisioning grpc address")
	}

	ZapierURL = config.Get(zapierURL)
	if ZapierURL == "" {
		return werror.Wrap(err, "error getting zapier URL")
	}

	return nil
}
