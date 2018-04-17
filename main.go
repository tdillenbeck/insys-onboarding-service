package main

import (
	"context"
	"net/http"
	"time"

	"google.golang.org/grpc"
	"weavelab.xyz/go-utilities/responsewrapper"
	"weavelab.xyz/insys-onboarding/config"
	"weavelab.xyz/insys-onboarding/exampleproto"
	"weavelab.xyz/insys-onboarding/server"
	"weavelab.xyz/wlib/wapp"
	"weavelab.xyz/wlib/wapp/grpcwapp"
	"weavelab.xyz/wlib/wapp/httpwapp"
	"weavelab.xyz/wlib/werror"
	"weavelab.xyz/wlib/wgoji"
	"weavelab.xyz/wlib/wgoji/whttptracing"
	"weavelab.xyz/wlib/wgrpc/wgrpcserver"
	"weavelab.xyz/wlib/wlog"
	"weavelab.xyz/wlib/wlog/tag"
	"weavelab.xyz/wlib/wmetrics"
)

var (
	// config variables
	delayConfig = "delay-time"

	// metric names
	metricSimpleHandler = "simple-handler"
)

func init() {
	wmetrics.SetLabels(metricSimpleHandler, "delay-time")
}

func main() {

	returnsErr()

	err := config.Init()
	if err != nil {
		wapp.Exit(werror.Wrap(err, "error initializing config values"))
	}

	myHttpHandler := simpleHandler(config.DelayConfig)

	// add appropriate middleware from the wlib/wgoji package
	statsWrappedHandler := wgoji.StatsMiddleware(myHttpHandler)
	tracingMiddleware, err := whttptracing.New(nil)
	if err != nil {
		wapp.Exit(werror.Wrap(err, "error creating tracing middleware"))
	}
	tracingWrappedHandler := tracingMiddleware(statsWrappedHandler)

	// create a WApp HTTP Starter using the httpwapp package
	httpStarter := httpwapp.Starter(":8000", tracingWrappedHandler)

	// create a grpc starter using the grpcwapp package
	grpcStarter := grpcwapp.Bootstrap(grpcBootstrap(config.DelayConfig))

	// pass the starters to wapp.Up to start your application
	wapp.Up(
		context.Background(),
		httpStarter,
		grpcStarter,
	)

	// whenever wapp gets the signal to shutdown it will stop all of your "starters" in reverse order and then return
	wlog.Info("done")
}

func returnsErr() error {
	return nil
}

//////////// HTTP

func simpleHandler(delay time.Duration) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		rw := responsewrapper.New(w, r)

		// here we actually send metrics based on what we set up in init--the number of parameters passed in here must match the number passed to wmetrics.SetLabels in init
		stop := wmetrics.StartTimer(metricSimpleHandler, delay.String())
		defer stop()

		// structured logging with wlog.Info
		wlog.Info("got request, waiting for delay before responding", tag.String("delay", delay.String()))

		time.Sleep(delay)

		rw.Status(http.StatusOK)
		_, _ = w.Write([]byte("hi"))
	}
}

//////////// GRPC

func grpcBootstrap(delay time.Duration) grpcwapp.BootstrapFunc {
	return func() (*grpc.Server, error) {

		s := wgrpcserver.NewDefault()

		exampleImplementation := &server.ServerImpl{
			Delay: delay,
		}

		exampleproto.RegisterExampleAPIServer(s, exampleImplementation)

		return s, nil
	}
}
