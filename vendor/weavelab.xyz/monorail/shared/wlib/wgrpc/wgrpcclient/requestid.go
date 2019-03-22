package wgrpcclient

import (
	"context"
	"weavelab.xyz/monorail/shared/wlib/wcontext"
	"weavelab.xyz/monorail/shared/wlib/wgrpc"
	"google.golang.org/grpc"
)

//UnaryRequestID gets a request ID from the metadata in the context and adds it as a value to the context, or creates one if it doesn't exist
func UnaryRequestID(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	ctx = addRequestIDToMetadata(ctx)

	return invoker(ctx, method, req, reply, cc, opts...)
}

func addRequestIDToMetadata(ctx context.Context) context.Context {

	id := wcontext.RequestID(ctx)
	if id == "" {
		return ctx
	}

	ctx = wgrpc.SetOutgoingContext(ctx, wgrpc.RequestIDMetadataKey, id)

	return ctx
}
