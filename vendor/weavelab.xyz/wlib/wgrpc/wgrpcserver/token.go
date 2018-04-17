package wgrpcserver

import (
	"context"

	"weavelab.xyz/wiggum"
	"weavelab.xyz/wlib/wgrpc"
	"weavelab.xyz/wlib/wgrpc/wgrpcserver/wrapstream"
	"google.golang.org/grpc"
)

/*
	UnaryToken gets a token from the metadata in the context and adds it as a value to the context.
	If no token is given, it is up to the handler to decide whether or not to continue.
	If an invalid token string is given, an error will be returned.
*/
func UnaryToken(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {

	ctx, err = addTokenToContext(ctx)
	if err != nil {
		return nil, err
	}

	return handler(ctx, req)

}

/*
	StreamToken gets a token from the metadata in the context and adds it as a value to the context.
	If no token is given, it is up to the handler to decide whether or not to continue.
	If an invalid token string is given, an error will be returned.
*/
func StreamToken(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	wrappedStream := wrapstream.WrapServerStream(ss)

	ctx, err := addTokenToContext(wrappedStream.Context())
	if err != nil {
		return err
	}

	wrappedStream.WrappedContext = ctx

	return handler(srv, wrappedStream)
}

func addTokenToContext(ctx context.Context) (context.Context, error) {
	t, ok := wgrpc.IncomingMetadata(ctx, wgrpc.TokenMetadataKey)
	if !ok {
		return ctx, nil
	}

	ctx, _, err := wiggum.ContextAddTokenString(ctx, t)
	if err != nil {
		return nil, err
	}

	return ctx, nil
}
