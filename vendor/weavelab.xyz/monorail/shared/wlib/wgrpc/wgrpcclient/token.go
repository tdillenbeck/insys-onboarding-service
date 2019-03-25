package wgrpcclient

import (
	"weavelab.xyz/monorail/shared/wiggum"
	"weavelab.xyz/monorail/shared/wlib/wgrpc"

	"context"
	"google.golang.org/grpc"
)

/*
	UnaryToken adds the token in the context to the request metadata, if a token is present.
*/
func UnaryToken(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {

	ctx = addTokenToMetadata(ctx)

	return invoker(ctx, method, req, reply, cc, opts...)
}

/*
	StreamToken adds the token from the context to the request, if a token is present.
*/
func StreamToken(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {

	ctx = addTokenToMetadata(ctx)

	return streamer(ctx, desc, cc, method, opts...)
}

func addTokenToMetadata(ctx context.Context) context.Context {

	t, ok := wiggum.ContextToken(ctx)
	if !ok {
		return ctx
	}

	s := t.String()

	ctx = wgrpc.SetOutgoingContext(ctx, wgrpc.TokenMetadataKey, s)

	return ctx
}
