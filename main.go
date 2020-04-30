package main

import (
	"context"

	cgrpc "google.golang.org/grpc"

	"weavelab.xyz/insys-onboarding-service/internal/config"
	"weavelab.xyz/insys-onboarding-service/internal/grpc"
	"weavelab.xyz/insys-onboarding-service/internal/nsq/consumers"
	"weavelab.xyz/insys-onboarding-service/internal/nsq/producers"
	"weavelab.xyz/insys-onboarding-service/internal/psql"
	"weavelab.xyz/insys-onboarding-service/internal/zapier"

	"weavelab.xyz/monorail/shared/grpc-clients/client-grpc-clients/authclient"
	"weavelab.xyz/monorail/shared/grpc-clients/client-grpc-clients/featureflagsclient"
	"weavelab.xyz/monorail/shared/protorepo/dist/go/services/insys"
	"weavelab.xyz/monorail/shared/wlib/wapp"
	"weavelab.xyz/monorail/shared/wlib/wapp/grpcwapp"
	"weavelab.xyz/monorail/shared/wlib/wapp/nsqwapp"
	"weavelab.xyz/monorail/shared/wlib/wdns"
	"weavelab.xyz/monorail/shared/wlib/werror"
	"weavelab.xyz/monorail/shared/wlib/wgrpc/wgrpcclient"
	"weavelab.xyz/monorail/shared/wlib/wgrpc/wgrpcserver"
	"weavelab.xyz/monorail/shared/wlib/wlog"
	"weavelab.xyz/monorail/shared/wlib/wsql"
)

func main() {
	err := config.Init()
	if err != nil {
		wapp.Exit(werror.Wrap(err, "error initializing config values"))
	}

	ctx := context.Background()

	// setup database connection
	dbOptions := &psql.ConnectionOptions{
		MaxOpenConnections:    config.MaxOpenConnections,
		MaxIdleConnections:    config.MaxIdleConnections,
		MaxConnectionLifetime: config.MaxConnectionLifetime,
		LogQueries:            config.LogQueries,
	}

	var db *wsql.PG
	if config.DBDriver == wsql.CloudSQLDriver {
		db, err = psql.ConnectionFromCloudDriver(ctx, config.DBPrimaryAddr, config.DBReplicaAddr, config.DBName, config.DBUsername, config.DBPassword, dbOptions)
	} else if config.PrimaryConnString != "" && config.ReplicaConnString != "" {
		db, err = psql.ConnectionFromConnString(ctx, config.PrimaryConnString, config.ReplicaConnString, dbOptions)
	} else {
		db, err = psql.ConnectionFromVault(ctx, config.DBPrimaryAddr, config.DBReplicaAddr, config.DBName, dbOptions)
	}
	if err != nil {
		wapp.Exit(werror.Wrap(err, "error establishing database connection"))
	}

	// setup grpc clients
	authClient, err := authclient.New(ctx, config.AuthServiceAddr)
	if err != nil {
		wapp.Exit(werror.Wrap(err, "error setting up auth client"))
	}

	featureFlagsClient, err := featureflagsclient.New(ctx, config.FeatureFlagsAddr)
	if err != nil {
		wapp.Exit(werror.Wrap(err, "error setting up feature flags client"))
	}

	portingDataClient, err := initPortingDataClient(ctx, config.PortingDataGRPCAddr)
	if err != nil {
		wapp.Exit(werror.Wrap(err, "error setting up porting data client"))
	}

	provisioningClient, err := initProvisioningClient(ctx, config.ProvisioningGRPCAddress)
	if err != nil {
		wapp.Exit(werror.Wrap(err, "error setting up provisioning client"))
	}

	// setup nsq publishers
	producers.Init(config.NSQDAddr)
	chiliPiperScheduleEventCreatedPublisher := producers.NewChiliPiperScheduleEventCreatedPublisher(config.NSQChiliPiperScheduleEventCreatedTopic)

	// setup zapier client
	zapierClient := zapier.New(config.ZapierURL)

	// setup grpc
	categoryService := &psql.CategoryService{DB: db}
	chiliPiperScheduleEventsService := &psql.ChiliPiperScheduleEventService{DB: db}
	rescheduleTrackingService := &psql.RescheduleTrackingEventService{DB: db}
	taskInstanceService := &psql.TaskInstanceService{DB: db}
	onboarderService := &psql.OnboarderService{DB: db}
	onboardersLocationService := &psql.OnboardersLocationService{DB: db}
	handoffSnapshotService := &psql.HandoffSnapshotService{DB: db}

	chiliPiperScheduleEventServer := grpc.NewChiliPiperScheduleEventServer(chiliPiperScheduleEventCreatedPublisher, chiliPiperScheduleEventsService)
	rescheduleTrackingEventServer := grpc.NewRescheduleEventServer(rescheduleTrackingService)
	onboardingServer := grpc.NewOnboardingServer(categoryService, taskInstanceService, portingDataClient)
	onboarderServer := grpc.NewOnboarderServer(onboarderService)
	onboardersLocationServer := grpc.NewOnboardersLocationServer(onboardersLocationService, taskInstanceService)
	handoffSnapshotServer := grpc.NewHandoffSnapshotServer(handoffSnapshotService)

	// setup nsq consumers
	nsqConfig := nsqwapp.NewConfig()
	nsqConfig.ConcurrentHandlers = config.NSQConcurrentHandlers
	nsqConfig.NSQConfig.MaxInFlight = config.NSQMaxInFlight

	chiliPiperScheduleEventCreatedSubscriber := consumers.NewChiliPiperScheduleEventCreatedSubscriber(chiliPiperScheduleEventsService, onboarderService, rescheduleTrackingService, onboardersLocationServer, onboardingServer, featureFlagsClient)
	portingDataRecordCreatedSubscriber := consumers.NewPortingDataRecordCreatedSubscriber(ctx, taskInstanceService)
	loginEventCreatedSubscriber := consumers.NewLogInEventCreatedSubscriber(ctx, authClient, featureFlagsClient, provisioningClient, zapierClient)

	grpcStarter := grpcwapp.Bootstrap(grpcBootstrap(chiliPiperScheduleEventServer, onboardingServer, onboarderServer, onboardersLocationServer, handoffSnapshotServer, rescheduleTrackingEventServer))

	wapp.ProbesAddr = ":4444"
	wapp.Up(
		ctx,
		grpcStarter,
		nsqwapp.Bootstrap(config.NSQChiliPiperScheduleEventCreatedTopic, config.NSQChannel, config.NSQLookupAddrs, nsqConfig, chiliPiperScheduleEventCreatedSubscriber),
		nsqwapp.Bootstrap(config.NSQPortingDataRecordCreatedTopic, config.NSQChannel, config.NSQLookupAddrs, nsqConfig, portingDataRecordCreatedSubscriber),
		nsqwapp.Bootstrap(config.NSQLoginEventCreatedTopic, config.NSQChannel, config.NSQLookupAddrs, nsqConfig, loginEventCreatedSubscriber),
	)

	// whenever wapp gets the signal to shutdown it will stop all of your "starters" in reverse order and then return
	wlog.InfoC(ctx, "done")
}

func grpcBootstrap(
	chiliPiperScheduleEventServer *grpc.ChiliPiperScheduleEventServer,
	onboardingServer *grpc.OnboardingServer, onboarderServer *grpc.OnboarderServer,
	onboardersLocationServer *grpc.OnboardersLocationServer,
	handoffSnapshotServer *grpc.HandoffSnapshotServer,
	rescheduleTrackingEventServer *grpc.RescheduleTrackingEventServer,
) grpcwapp.BootstrapFunc {
	return func() (*cgrpc.Server, error) {
		gs, err := wgrpcserver.NewDefault()
		if err != nil {
			wapp.Exit(werror.Wrap(err, "error getting a new default wgrpc server"))
		}

		insys.RegisterChiliPiperScheduleEventServer(gs, chiliPiperScheduleEventServer)
		insys.RegisterOnboardingServer(gs, onboardingServer)
		insys.RegisterOnboarderServer(gs, onboarderServer)
		insys.RegisterOnboardersLocationServer(gs, onboardersLocationServer)
		insys.RegisterHandoffSnapshotServer(gs, handoffSnapshotServer)
		insys.RegisterRescheduleTrackingEventServer(gs, rescheduleTrackingEventServer)

		return gs, nil
	}
}

func initPortingDataClient(ctx context.Context, grpcAddr string) (insys.PortingDataServiceClient, error) {
	defaultPortingDataGrpcAddress := "insys-porting-data.insys.svc.cluster.local.:grpc"

	if grpcAddr == "" {
		var err error

		grpcAddr, err = wdns.ResolveAddress(defaultPortingDataGrpcAddress)
		if err != nil {
			return nil, werror.Wrap(err, "unable to use default settings address for porting data client")
		}
	}

	g, err := wgrpcclient.NewDefault(ctx, grpcAddr)
	if err != nil {
		return nil, werror.Wrap(err, "unable to setup PortingData grpc client")
	}

	return insys.NewPortingDataServiceClient(g), nil
}

func initProvisioningClient(ctx context.Context, grpcAddr string) (insys.ProvisioningClient, error) {
	g, err := wgrpcclient.NewDefault(ctx, grpcAddr)
	if err != nil {
		return nil, werror.Wrap(err, "unable to setup Provisioning grpc client")
	}

	return insys.NewProvisioningClient(g), nil
}
