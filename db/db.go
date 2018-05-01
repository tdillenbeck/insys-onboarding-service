package db

import (
	"context"
	"time"

	"weavelab.xyz/wlib/wsql"
	"weavelab.xyz/wlib/wvault/wvaultdb"
)

var Conn *wsql.PG

type DatabaseConnectionOptions struct {
	MaxOpenConnections    int
	MaxIdleConnections    int
	MaxConnectionLifetime time.Duration
	LogQueries            bool
}

func InitConnectionFromConnString(ctx context.Context, primaryConnString, replicaConnString string, options *DatabaseConnectionOptions) error {
	var err error

	settings := wsql.Settings{
		MaxOpenConnections:    options.MaxOpenConnections,
		MaxIdleConnections:    options.MaxIdleConnections,
		MaxConnectionLifetime: options.MaxConnectionLifetime,
		LogQueries:            options.LogQueries,
	}

	settings.PrimaryConnectString.SetConnectString(primaryConnString)
	settings.ReplicaConnectString.SetConnectString(replicaConnString)

	Conn, err = wsql.New(&settings)
	if err != nil {
		return err
	}

	return nil
}

func InitConnectionFromVault(ctx context.Context, primaryHost, replicaHost, dbName string, options *DatabaseConnectionOptions) error {
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

	Conn, err = wvaultdb.New(ctx, role, target, &settings)
	if err != nil {
		return err
	}

	return nil
}
