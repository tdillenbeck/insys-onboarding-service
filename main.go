package main

import (
	"context"

	"google.golang.org/grpc"
	"weavelab.xyz/insys-onboarding/config"
	"weavelab.xyz/insys-onboarding/server"
	"weavelab.xyz/protorepo/dist/go/services/insys/onboarding"
	"weavelab.xyz/wlib/wapp"
	"weavelab.xyz/wlib/wapp/grpcwapp"
	"weavelab.xyz/wlib/werror"
	"weavelab.xyz/wlib/wgrpc/wgrpcserver"
	"weavelab.xyz/wlib/wlog"
)

func main() {
	err := config.Init()
	if err != nil {
		wapp.Exit(werror.Wrap(err, "error initializing config values"))
	}

	grpcStarter := grpcwapp.Bootstrap(grpcBootstrap())

	wapp.ProbesAddr = ":4444"
	wapp.Up(
		context.Background(),
		grpcStarter,
	)

	// whenever wapp gets the signal to shutdown it will stop all of your "starters" in reverse order and then return
	wlog.Info("done")
}

func grpcBootstrap() grpcwapp.BootstrapFunc {
	return func() (*grpc.Server, error) {
		s := wgrpcserver.NewDefault()
		srv := &server.OnboardingService{}
		onboarding.RegisterOnboardingServer(s, srv)

		return s, nil
	}
}
