package psql

import (
	"context"
	"time"

	"weavelab.xyz/monorail/shared/wlib/wsql"
	"weavelab.xyz/monorail/shared/wlib/wvault/wvaultdb"
)

type ConnectionOptions struct {
	MaxOpenConnections    int
	MaxIdleConnections    int
	MaxConnectionLifetime time.Duration
	LogQueries            bool
}

// ConnectionFromConnString sets up the primary and replica wsql connection using the provided connection strings.
func ConnectionFromConnString(ctx context.Context, primaryConnString, replicaConnString string, options *ConnectionOptions) (*wsql.PG, error) {
	var err error

	settings := wsql.Settings{
		MaxOpenConnections:    options.MaxOpenConnections,
		MaxIdleConnections:    options.MaxIdleConnections,
		MaxConnectionLifetime: options.MaxConnectionLifetime,
		LogQueries:            options.LogQueries,
	}

	settings.PrimaryConnectString.SetConnectString(primaryConnString)
	settings.ReplicaConnectString.SetConnectString(replicaConnString)

	conn, err := wsql.New(&settings)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

// ConnectionFromVault will use our instance of Vault by Hashicorp to create a database connection.
func ConnectionFromVault(ctx context.Context, primaryHost, replicaHost, dbName string, options *ConnectionOptions) (*wsql.PG, error) {
	var err error

	settings := wsql.Settings{
		PrimaryConnectString: wsql.ConnectString{
			Host:     primaryHost,
			Database: dbName,
		},
		ReplicaConnectString: wsql.ConnectString{
			Host:     replicaHost,
			Database: dbName,
		},
		MaxOpenConnections:    options.MaxOpenConnections,
		MaxIdleConnections:    options.MaxIdleConnections,
		MaxConnectionLifetime: options.MaxConnectionLifetime,
		LogQueries:            options.LogQueries,
	}

	role := "db_insys_onboarding"
	target := "db_insys_onboarding"

	conn, err := wvaultdb.New(ctx, role, target, &settings)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

// ConnectionfromCloudDriiver will initialize the primary and replica database connections using the connection settings passed in and use the cloud dql driver.
func ConnectionFromCloudDriver(ctx context.Context, primaryHost, replicaHost, dbName, username, password string, options *ConnectionOptions) (*wsql.PG, error) {
	settings := wsql.Settings{
		PrimaryConnectString: wsql.ConnectString{
			Driver: wsql.CloudSQLDriver,

			Database: dbName,
			Host:     primaryHost,
			Password: password,
			Username: username,
		},
		ReplicaConnectString: wsql.ConnectString{
			Driver: wsql.CloudSQLDriver,

			Database: dbName,
			Host:     replicaHost,
			Password: password,
			Username: username,
		},
		MaxOpenConnections:    options.MaxOpenConnections,
		MaxIdleConnections:    options.MaxIdleConnections,
		MaxConnectionLifetime: options.MaxConnectionLifetime,
		LogQueries:            options.LogQueries,
	}

	conn, err := wsql.New(&settings)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
