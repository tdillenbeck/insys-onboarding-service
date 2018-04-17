package wgrpcserver

import (
	"context"

	"weavelab.xyz/wlib/wcontext"
	"weavelab.xyz/wlib/wgrpc"
	"weavelab.xyz/wlib/wgrpc/wgrpcserver/wrapstream"
	"google.golang.org/grpc"
)

type keys int

//UnaryRequestID gets a request ID from the metadata in the context and adds it as a value to the context, or creates one if it doesn't exist
func UnaryRequestID(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	ctx = addRequestIDToContext(ctx)

	return handler(ctx, req)
}

//StreamRequestID gets a request ID from the metadata in the context and adds it as a value to the context, or creates one if it doesn't exist
func StreamRequestID(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	wrappedStream := wrapstream.WrapServerStream(ss)

	ctx := addRequestIDToContext(wrappedStream.Context())
	wrappedStream.WrappedContext = ctx

	return handler(srv, wrappedStream)
}

func addRequestIDToContext(ctx context.Context) context.Context {
	id, _ := wgrpc.IncomingMetadata(ctx, wgrpc.RequestIDMetadataKey)

	// if id is empty string, it will generate a new id for us
	return wcontext.NewWithRequestID(ctx, id)
}
