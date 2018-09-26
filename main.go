package main

import (
	"context"
	"time"

	cgrpc "google.golang.org/grpc"

	"weavelab.xyz/insys-onboarding/internal/config"
	"weavelab.xyz/insys-onboarding/internal/grpc"
	"weavelab.xyz/insys-onboarding/internal/psql"
	"weavelab.xyz/protorepo/dist/go/services/insys"
	"weavelab.xyz/wlib/wapp"
	"weavelab.xyz/wlib/wapp/grpcwapp"
	"weavelab.xyz/wlib/werror"
	"weavelab.xyz/wlib/wgrpc/wgrpcserver"
	"weavelab.xyz/wlib/wlog"
	"weavelab.xyz/wlib/wsql"
	"weavelab.xyz/wlib/wvault/wvaultdb"
)

type databaseConnectionOptions struct {
	MaxOpenConnections    int
	MaxIdleConnections    int
	MaxConnectionLifetime time.Duration
	LogQueries            bool
}

func main() {
	err := config.Init()
	if err != nil {
		wapp.Exit(werror.Wrap(err, "error initializing config values"))
	}

	ctx := context.Background()

	dbOptions := &databaseConnectionOptions{
		MaxOpenConnections:    config.MaxOpenConnections,
		MaxIdleConnections:    config.MaxIdleConnections,
		MaxConnectionLifetime: config.MaxConnectionLifetime,
		LogQueries:            config.LogQueries,
	}

	var db *wsql.PG

	if config.PrimaryConnString != "" && config.ReplicaConnString != "" {
		db, err = initDBConnectionFromConnString(ctx, config.PrimaryConnString, config.ReplicaConnString, dbOptions)
	} else {
		db, err = initDBConnectionFromVault(ctx, config.DBPrimaryAddr, config.DBReplicaAddr, config.DBName, dbOptions)
	}
	if err != nil {
		wapp.Exit(werror.Wrap(err, "error establishing database connection"))
	}

	categoryService := &psql.CategoryService{DB: db}
	taskInstanceService := &psql.TaskInstanceService{DB: db}
	onboarderService := &psql.OnboarderService{DB: db}
	onboardersLocationService := &psql.OnboardersLocationService{DB: db}

	onboardingServer := grpc.NewOnboardingServer(categoryService, taskInstanceService)
	onboarderServer := grpc.NewOnboarderServer(onboarderService)
	onboardersLocationServer := grpc.NewOnboardersLocationServer(onboardersLocationService)

	grpcStarter := grpcwapp.Bootstrap(grpcBootstrap(onboardingServer, onboarderServer, onboardersLocationServer))

	wapp.ProbesAddr = ":4444"
	wapp.Up(
		ctx,
		grpcStarter,
	)

	// whenever wapp gets the signal to shutdown it will stop all of your "starters" in reverse order and then return
	wlog.InfoC(ctx, "done")
}

func grpcBootstrap(onboardingServer *grpc.OnboardingServer, onboarderServer *grpc.OnboarderServer, onboardersLocationServer *grpc.OnboardersLocationServer) grpcwapp.BootstrapFunc {
	return func() (*cgrpc.Server, error) {
		gs, err := wgrpcserver.NewDefault()
		if err != nil {
			wapp.Exit(werror.Wrap(err, "error getting a new default wgrpc server"))
		}

		insys.RegisterOnboardingServer(gs, onboardingServer)
		insys.RegisterOnboarderServer(gs, onboarderServer)
		insys.RegisterOnboardersLocationServer(gs, onboardersLocationServer)

		return gs, nil
	}
}

func initDBConnectionFromConnString(ctx context.Context, primaryConnString, replicaConnString string, options *databaseConnectionOptions) (*wsql.PG, error) {
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

func initDBConnectionFromVault(ctx context.Context, primaryHost, replicaHost, dbName string, options *databaseConnectionOptions) (*wsql.PG, error) {
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
