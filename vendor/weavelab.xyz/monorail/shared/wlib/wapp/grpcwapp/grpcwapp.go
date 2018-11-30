//grpcwapp is used to start a single grpc service
package grpcwapp

import (
	"net"

	"google.golang.org/grpc"

	"weavelab.xyz/monorail/shared/wlib/config"
	"weavelab.xyz/monorail/shared/wlib/wapp"
	"weavelab.xyz/monorail/shared/wlib/werror"
	"weavelab.xyz/monorail/shared/wlib/wlog"
	"weavelab.xyz/monorail/shared/wlib/wlog/tag"
	"weavelab.xyz/monorail/shared/wlib/wmetrics"
)

const (
	grpcAddrConfig = "grpc-addr"
)

func init() {
	config.Add(grpcAddrConfig, ":9000", "The address on which to serve the gRPC service")
}

// BootstrapFunc must return a gRPC server, and an error
type BootstrapFunc func() (srv *grpc.Server, err error)

// Bootstrap calls the function to get a grpc.Server and then serves the gRPC server on the configured address
func Bootstrap(b BootstrapFunc) wapp.StartFunc {

	return func() (wapp.StopFunc, error) {
		wmetrics.Incr(1, "wapp", "grpcwapp")

		addr := config.Get(grpcAddrConfig)
		if addr == "" {
			return nil, werror.New("value required for config:" + grpcAddrConfig)
		}

		srv, err := b()
		if err != nil {
			return nil, werror.Wrap(err, "error starting grpcwapp")
		}

		lis, err := net.Listen("tcp", addr)
		if err != nil {
			return nil, err
		}

		go func() {
			wlog.Info("Serving gRPC service", tag.String("addr", addr))
			err = srv.Serve(lis)
			if err != nil {
				wapp.Exit(werror.Wrap(err, "error from grpc Serve"))
			}
		}()

		return wapp.WrapSimpleStopFunc(srv.GracefulStop), nil
	}
}
