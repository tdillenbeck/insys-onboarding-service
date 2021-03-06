package wgrpcserver

import (
	"context"

	"google.golang.org/grpc"
	"weavelab.xyz/monorail/shared/wiggum"
	"weavelab.xyz/monorail/shared/wlib/wgrpc"
	"weavelab.xyz/monorail/shared/wlib/wgrpc/wgrpcserver/wrapstream"
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

func UnaryAudienceInterceptor(audience string) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		token, found := wiggum.ContextToken(ctx)
		if !found {
			return nil, wiggum.NotAuthorizedError.Here("no token found for audience check")
		}
		if token.ACLType() == wiggum.WeaveACLType {
			if !token.HasAudience(audience) {
				return nil, wiggum.NotAuthorizedError.Here("mismatched audience")
			}
		}
		return handler(ctx, req)
	}
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
