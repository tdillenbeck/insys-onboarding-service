package wgrpcclient

import (
	"context"
	"fmt"
	"net"
	"strings"
	"time"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"weavelab.xyz/monorail/shared/go-utilities/null"
	"weavelab.xyz/monorail/shared/wlib/version"
	"weavelab.xyz/monorail/shared/wlib/werror"
	"weavelab.xyz/monorail/shared/wlib/wmetrics"
)

const (
	grpcStatsConnPrefix = "grpc_client_conn"
)

func init() {
	wmetrics.SetLabels(grpcStatsConnPrefix, "target")
}

func defaultUnaryMiddleware() ([]grpc.UnaryClientInterceptor, error) {

	tracingMiddleware, err := NewUnaryTracerInterceptor(nil)
	if err != nil {
		return nil, werror.Wrap(err, "unable to load default unary middleware").SetCode(werror.CodeInternal)
	}
	m := []grpc.UnaryClientInterceptor{UnaryRequestID, tracingMiddleware, UnaryToken}

	return m, nil
}

func defaultStreamingMiddleware() ([]grpc.StreamClientInterceptor, error) {
	m := []grpc.StreamClientInterceptor{}
	return m, nil
}

// New creates a new gRPC client with the default middleware and any other middleware passed in. The defaults are added after the custom ones passed in.
func New(ctx context.Context, target string, unaryMiddleware []grpc.UnaryClientInterceptor, streamMiddleware []grpc.StreamClientInterceptor, opt ...grpc.DialOption) (*grpc.ClientConn, error) {

	target, err := setScheme(target)
	if err != nil {
		return nil, werror.Wrap(err, "unable to set target scheme")
	}

	unaryM, err := defaultUnaryMiddleware()
	if err != nil {
		return nil, werror.Wrap(err)
	}

	streamingM, err := defaultStreamingMiddleware()
	if err != nil {
		return nil, werror.Wrap(err)
	}

	//Add list of passed in middlewares to defaults
	unaryMiddleware = append(unaryMiddleware, unaryM...)
	streamMiddleware = append(streamMiddleware, streamingM...)

	opt = append(opt, grpc.WithUnaryInterceptor(grpc_middleware.ChainUnaryClient(unaryMiddleware...)), grpc.WithStreamInterceptor(grpc_middleware.ChainStreamClient(streamMiddleware...)))

	info := version.Info()
	ua := fmt.Sprintf("%s/%s-%s", info.Name, info.Version, null.Truncate(info.GitHash, 8))
	defaultOpt := []grpc.DialOption{
		grpc.WithBackoffMaxDelay(time.Second * 2),
		grpc.WithUserAgent(ua),
		grpc.WithInsecure(),
		grpc.WithBalancerName(roundrobin.Name),
	}

	opt = append(defaultOpt, opt...)

	wmetrics.Incr(1, grpcStatsConnPrefix, strings.Replace(target, ".", "_", -1))

	return grpc.DialContext(ctx, target, opt...)
}

// NewVanilla creates a gRPC client without the default middleware.
func NewVanilla(ctx context.Context, target string, unaryMiddleWare []grpc.UnaryClientInterceptor, streamMiddleware []grpc.StreamClientInterceptor, opt ...grpc.DialOption) (*grpc.ClientConn, error) {

	opt = append(opt, grpc.WithUnaryInterceptor(grpc_middleware.ChainUnaryClient(unaryMiddleWare...)), grpc.WithStreamInterceptor(grpc_middleware.ChainStreamClient(streamMiddleware...)))

	wmetrics.Incr(1, grpcStatsConnPrefix, strings.Replace(target, ".", "_", -1))

	return grpc.DialContext(ctx, target, opt...)
}

// NewDefault creates a gRPC client with defaults
func NewDefault(ctx context.Context, target string, opt ...grpc.DialOption) (*grpc.ClientConn, error) {
	return New(ctx, target, nil, nil, opt...)
}

// NewRetry creates a gRPC client using defaults + grpc_retry middleware.
// Sample:
// retryOptions := []grpc_retry.CallOption {
//     grpc_retry.WithMax(3),
//     grpc_retry.WithBackoff(grpc_retry.BackoffLinear(100 * time.Millisecond),
//     grpc_retry.WithCodes(codes.Unavailable),
// }
// client, err := wgrpcclient.NewRetry(ctx, target, retryOptions...)
// See https://github.com/grpc-ecosystem/go-grpc-middleware/tree/master/retry for options
func NewRetry(ctx context.Context, target string, opt ...grpc_retry.CallOption) (*grpc.ClientConn, error) {
	return New(ctx, target, []grpc.UnaryClientInterceptor{grpc_retry.UnaryClientInterceptor(opt...)}, []grpc.StreamClientInterceptor{grpc_retry.StreamClientInterceptor(opt...)})
}

func setScheme(target string) (string, error) {

	if hasScheme(target) {
		return target, nil
	}

	// assume that we have only host[:port] as the target
	headless, err := isHeadless(target)
	if err != nil {
		return "", werror.Wrap(err)
	}
	// default resolver is the passthrough resolver, adding
	// the dns:/// prefix will switch it to the DNS resolver
	// for now, only use it with headless services
	// in the future, it may be okay to make it the default schema
	if headless || true {
		// https://godoc.org/google.golang.org/grpc#DialContext
		target = "dns:///" + target
	}

	return target, nil

}

func hasScheme(target string) bool {
	// assume the presence of // indicates a url scheme is present
	if strings.Contains(target, "//") {
		return true
	}
	return false
}

// returns whether or not a list of IPs are available for a given target
func isHeadless(target string) (bool, error) {

	host, _, err := net.SplitHostPort(target)
	if err != nil {
		if strings.Contains(err.Error(), "missing port in address") {
			host = target
		} else {
			return false, werror.Wrap(err)
		}
	}

	// lookup how many records are returned
	ips, err := net.LookupIP(host)
	if err != nil {
		return false, werror.Wrap(err)
	}

	if len(ips) > 1 {
		return true, nil
	}

	return false, nil
}
