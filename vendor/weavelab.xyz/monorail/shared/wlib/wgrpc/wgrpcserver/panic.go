package wgrpcserver

import (
	"context"
	"fmt"
	"os"

	"weavelab.xyz/monorail/shared/wlib/werror"
	"weavelab.xyz/monorail/shared/wlib/wgrpc/wgrpcserver/wrapstream"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

var HandleError = func(ctx context.Context, werr *werror.Error) {
	Logger.WErrorC(ctx, werr)
	fmt.Fprintln(os.Stderr, werr.PrintStack())
}

//UnaryPanicRecover captures a panic
func UnaryPanicRecover(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	defer func() {
		if r := recover(); r != nil {
			werr := werror.New(fmt.Sprint("panic occurred while processing gRPC request: ", r))
			HandleError(ctx, werr)

			err = grpc.Errorf(codes.Internal, werr.Error())
		}
	}()

	return handler(ctx, req)

}

//StreamPanicRecover captures panics
func StreamPanicRecover(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) (err error) {
	defer func() {
		if r := recover(); r != nil {
			werr := werror.New(fmt.Sprint("panic occurred while processing gRPC stream: ", r))
			ctx := ss.Context()
			HandleError(ctx, werr)

			err = grpc.Errorf(codes.Internal, werr.Error())
		}
	}()

	wrappedStream := wrapstream.WrapServerStream(ss)

	err = handler(srv, wrappedStream)

	return err
}
