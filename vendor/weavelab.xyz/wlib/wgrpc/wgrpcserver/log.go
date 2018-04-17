package wgrpcserver

import (
	"context"
	"time"

	"weavelab.xyz/wlib/wcontext"
	"weavelab.xyz/wlib/werror"
	"weavelab.xyz/wlib/wgrpc/wgrpcserver/wrapstream"
	"weavelab.xyz/wlib/wlog"
	"weavelab.xyz/wlib/wlog/tag"
	"google.golang.org/grpc"
)

const (
	inbound  = "inbound"
	outbound = "outbound"
)

type logger interface {
	InfoC(c context.Context, msg string, tags ...tag.Tag)
	WErrorC(context.Context, *werror.Error)
}

//Logger can be replaced to customize the log middleware's behavor
var Logger logger = wlog.NewWLogger(wlog.WlogdLogger)

//UnaryLogging for handling logging for unary gRPC endpoints
func UnaryLogging(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	start := time.Now()
	//TODO get request id

	// What info can we and should we log here
	Logger.InfoC(
		ctx,
		"making gRPC request",
		tag.String("FullMethod", info.FullMethod),
		tag.String("t", time.Now().String()))

	resp, err = handler(ctx, req)

	//Log after handler is finished
	Logger.InfoC(
		ctx,
		"finished gRPC request",
		tag.String("FullMethod", info.FullMethod),
		tag.String("t", time.Now().String()),
		tag.String("duration", time.Since(start).String()),
		tag.String("grpcStatus", grpc.Code(err).String()))

	return resp, err
}

//StreamLogging for handling logging for unary gRPC endpoints
func StreamLogging(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	start := time.Now()
	ctx := ss.Context()

	//Do logging before streaming starts
	preLogMsg(ctx, info, "Initiating gRPC stream")

	inboundCounter := 0
	outboundCounter := 0

	wrappedStream := wrapstream.WrapServerStream(ss)
	wrappedStream.RegisterRecvMiddleware(logStreamCounterMiddleware(ctx, &inboundCounter))
	wrappedStream.RegisterSendMiddleware(logStreamCounterMiddleware(ctx, &outboundCounter))

	err := handler(srv, wrappedStream)
	if err != nil {
		return err
	}

	postLogMsg(ctx, info, err, start, "Closing gRPC stream",
		tag.Int("inboundMessages", inboundCounter),
		tag.Int("outboundMessages", outboundCounter),
	)

	return nil
}

func logStreamCounterMiddleware(ctx context.Context, counter *int) func(inner wrapstream.StreamHandler) wrapstream.StreamHandler {

	return func(inner wrapstream.StreamHandler) wrapstream.StreamHandler {
		mw := func(m interface{}) error {

			err := inner.Stream(m)
			if err != nil {
				return err
			}

			*counter = *counter + 1

			return nil
		}

		return wrapstream.StreamFunc(mw)
	}
}

func preLogMsg(ctx context.Context, info *grpc.StreamServerInfo, msg string, tags ...tag.Tag) {
	tags = append(tags, tag.String("FullMethod", info.FullMethod),
		tag.String("t", time.Now().String()),
		tag.Bool("isClientStream", info.IsClientStream),
		tag.Bool("isServerStream", info.IsServerStream),
		tag.String("requestID", wcontext.RequestID(ctx)))

	Logger.InfoC(
		ctx,
		msg,
		tags...,
	)
}

func postLogMsg(ctx context.Context, info *grpc.StreamServerInfo, err error, startTime time.Time, msg string, tags ...tag.Tag) {
	tags = append(tags, tag.String("FullMethod", info.FullMethod),
		tag.String("t", time.Now().String()),
		tag.String("duration", time.Since(startTime).String()),
		tag.String("grpcStatus", grpc.Code(err).String()),
		tag.String("requestID", wcontext.RequestID(ctx)))

	Logger.InfoC(
		ctx,
		msg,
		tags...,
	)
}
