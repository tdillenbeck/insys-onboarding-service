package wgrpcserver

import (
	"time"

	"github.com/mwitkow/go-grpc-middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"
	"weavelab.xyz/wlib/werror"
)

var (
	// Load balancing docs
	// https://github.com/grpc/grpc/blob/master/doc/load-balancing.md
	// https://blog.bugsnag.com/envoy/

	// TODO: make these configurable is a better way than as global variables
	// Server side connection management
	// https://github.com/grpc/proposal/blob/master/A9-server-side-conn-mgt.md

	MaxConnectionAge      = time.Minute * 15
	MaxConnectionAgeGrace = time.Minute * 5
	MaxConnectionIdle     = time.Minute * 10
)

func defaultUnaryMiddleware() ([]grpc.UnaryServerInterceptor, error) {

	tracingMiddleware, err := NewUnaryTracerInterceptor(nil)
	if err != nil {
		return nil, werror.Wrap(err, "unable to load default unary middleware").SetCode(werror.CodeInternal)
	}

	m := []grpc.UnaryServerInterceptor{
		UnaryPanicRecover,
		UnaryRequestID,
		UnaryLogging,
		UnaryStats,
		tracingMiddleware,
	}

	return m, nil
}

func defaultStreamingMiddleware() ([]grpc.StreamServerInterceptor, error) {
	m := []grpc.StreamServerInterceptor{
		StreamPanicRecover,
		StreamRequestID,
		StreamLogging,
		StreamStats,
	}

	return m, nil
}

// NewDefault just gives you a new server with the default middleware
func NewDefault(opt ...grpc.ServerOption) (*grpc.Server, error) {
	return New(nil, nil, opt...)
}

// New creates a new gRPC server with the default middleware and any other middleware passed in. The defaults are added after the custom ones passed in.
func New(unaryMiddleware []grpc.UnaryServerInterceptor, streamMiddleware []grpc.StreamServerInterceptor, opt ...grpc.ServerOption) (*grpc.Server, error) {

	//Add list of passed in middlewares to defaults

	defaultUnaryM, err := defaultUnaryMiddleware()
	if err != nil {
		return nil, werror.Wrap(err)
	}

	defaultStreamingM, err := defaultStreamingMiddleware()
	if err != nil {
		return nil, werror.Wrap(err)
	}

	unaryMiddleware = append(defaultUnaryM, unaryMiddleware...)
	streamMiddleware = append(defaultStreamingM, streamMiddleware...)

	opt = append(opt, grpc_middleware.WithUnaryServerChain(unaryMiddleware...), grpc_middleware.WithStreamServerChain(streamMiddleware...))

	// https://godoc.org/google.golang.org/grpc/keepalive#ServerParameters
	keepaliveOpt := grpc.KeepaliveParams(keepalive.ServerParameters{

		MaxConnectionIdle: MaxConnectionIdle,
		// useful to force connection rebalancing
		// so one server doesn't end up with all of the connections
		MaxConnectionAge:      MaxConnectionAge,
		MaxConnectionAgeGrace: MaxConnectionAgeGrace,

		Time: time.Second * 30,
		// Timeout: // if going through HA Proxy, there is a 50 second timeout on both RX/TX
	})

	opt = append(opt, keepaliveOpt)

	//grpc_middleware has to be used because grpc.Server actually only allows one interceptor
	s := grpc.NewServer(opt...)

	reflection.Register(s)

	return s, nil
}

// NewVanilla creates a gRPC server without the default middleware.
func NewVanilla(unaryMiddleWare []grpc.UnaryServerInterceptor, streamMiddleware []grpc.StreamServerInterceptor, opt ...grpc.ServerOption) *grpc.Server {

	//grpc_middleware has to be used because grpc.Server actually only allows one interceptor
	opt = append(opt, grpc_middleware.WithUnaryServerChain(unaryMiddleWare...), grpc_middleware.WithStreamServerChain(streamMiddleware...))

	s := grpc.NewServer(opt...)
	return s
}
