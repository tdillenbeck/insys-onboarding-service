package wgrpc

import (
	"context"

	"weavelab.xyz/wlib/werror"
	"weavelab.xyz/wlib/wlog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type direction int

const (
	incoming = direction(1)
	outgoing = direction(2)
)

/*
	SetMetadata adds a key and value to the request metadata
*/
func SetIncomingContext(ctx context.Context, key string, value string) context.Context {
	return setContext(ctx, key, value, incoming)
}
func SetOutgoingContext(ctx context.Context, key string, value string) context.Context {
	return setContext(ctx, key, value, outgoing)
}

func setContext(ctx context.Context, key string, value string, d direction) context.Context {

	if ctx == nil {
		ctx = context.Background()
	}

	var md metadata.MD
	var ok bool
	if d == incoming {
		md, ok = metadata.FromIncomingContext(ctx)
	} else {
		md, ok = metadata.FromOutgoingContext(ctx)
	}
	if !ok {
		// if no metadata exists in context then we'll add it with the debugID
		md = metadata.Pairs(key, value)
		if d == incoming {
			return metadata.NewIncomingContext(ctx, md)
		} else {
			return metadata.NewOutgoingContext(ctx, md)
		}
	}

	md[key] = []string{value}

	if d == incoming {
		ctx = metadata.NewIncomingContext(ctx, md)
	} else {
		ctx = metadata.NewOutgoingContext(ctx, md)
	}

	return ctx
}

/*
	GetMetadata returns a value from the metadata for the given key
*/
func IncomingMetadata(ctx context.Context, key string) (string, bool) {
	return getMetadata(ctx, key, incoming)
}

func OutgoingMetadata(ctx context.Context, key string) (string, bool) {
	return getMetadata(ctx, key, outgoing)
}

func getMetadata(ctx context.Context, key string, d direction) (string, bool) {

	if ctx == nil {
		return "", false
	}

	var md metadata.MD
	var ok bool

	if d == incoming {
		md, ok = metadata.FromIncomingContext(ctx)
	} else {
		md, ok = metadata.FromOutgoingContext(ctx)
	}
	if !ok {
		return "", false
	}

	v := md[key]
	if len(v) != 1 {
		return "", false
	}

	return v[0], true

}

type Code = werror.Code

const (
	CodeInvalidArgument  = Code(codes.InvalidArgument)
	CodePermissionDenied = Code(codes.PermissionDenied)
	CodeUnauthenticated  = Code(codes.Unauthenticated)

	CodeInternal      = Code(codes.Internal)
	CodeNotFound      = Code(codes.NotFound)
	CodeUnimplemented = Code(codes.Unimplemented)
)

// Error attaches a code to the werror and logs it if it is not InvalidArgument, PermissionDenied, or Unauthenticated
func Error(code Code, werr *werror.Error) error {

	werrorCode := werr.Code()
	if werrorCode != 0 {
		if code != 0 {
			// we prioritize the code in the werror since it is more specific, but we'll still add the grpc code that would have been returned to the error
			werr.Add("grpcCode", code)
		}

		code = werrorCode
	}

	if code == 0 {
		code = Code(codes.Unknown)
	}

	if werr == nil {
		werr = werror.New("err was nil")
	}

	class := ErrorClass(codes.Code(code))
	switch class {
	case Success, ClientError:
		return status.Error(codes.Code(code), werr.Error())

	case ServerError:
	default:
	}

	// log the error and return
	wlog.WError(werr.Add("code", code))

	return status.Error(codes.Code(code), werr.Error())
}

func IsCode(code Code, err error) bool {
	status, ok := status.FromError(err)
	if !ok {
		return false
	}

	return codes.Code(code) == status.Code()
}

func Wrap(err error, message string) *werror.Error {

	code := codes.Unknown
	status, ok := status.FromError(err)
	if ok {
		code = status.Code()
	}

	w := werror.Wrap(err, message).SetCode(werror.Code(code))

	return w
}
