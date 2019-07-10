package wgrpcserver

import (
	"context"
	"time"

	"google.golang.org/grpc"
)

//UnaryTimeout will set the context's timeout for every incoming request
func UnaryTimeout(timeout time.Duration) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		ctx, done := context.WithTimeout(ctx, timeout)
		defer done()

		return handler(ctx, req)
	}
}
