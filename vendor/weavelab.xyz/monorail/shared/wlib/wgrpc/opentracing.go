// orignal source from github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc

package wgrpc

import (
	"strings"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

var (
	// Morally a const:
	GRPCComponentTag = opentracing.Tag{Key: string(ext.Component), Value: "gRPC"}
)

// metadataReaderWriter satisfies both the opentracing.TextMapReader and
// opentracing.TextMapWriter interfaces.
type MetadataReaderWriter struct {
	metadata.MD
}

func (w MetadataReaderWriter) Set(key, val string) {
	// The GRPC HPACK implementation rejects any uppercase keys here.
	//
	// As such, since the HTTP_HEADERS format is case-insensitive anyway, we
	// blindly lowercase the key (which is guaranteed to work in the
	// Inject/Extract sense per the OpenTracing spec).
	key = strings.ToLower(key)
	w.MD[key] = append(w.MD[key], val)
}

func (w MetadataReaderWriter) ForeachKey(handler func(key, val string) error) error {
	for k, vals := range w.MD {
		for _, v := range vals {
			if err := handler(k, v); err != nil {
				return err
			}
		}
	}

	return nil
}

// A Class is a set of types of outcomes (including errors) that will often
// be handled in the same way.
type Class string

const (
	Unknown Class = "0xx"
	// Success represents outcomes that achieved the desired results.
	Success Class = "2xx"
	// ClientError represents errors that were the client's fault.
	ClientError Class = "4xx"
	// ServerError represents errors that were the server's fault.
	ServerError Class = "5xx"
)

// ErrorClass returns the class of the given error
func ErrorClass(code codes.Code) Class {

	switch code {
	// Success or "success"
	case codes.OK, codes.Canceled:
		return Success

	// Client errors
	case codes.InvalidArgument, codes.NotFound, codes.AlreadyExists,
		codes.PermissionDenied, codes.Unauthenticated, codes.FailedPrecondition,
		codes.OutOfRange:
		return ClientError

	// Server errors
	case codes.DeadlineExceeded, codes.ResourceExhausted, codes.Aborted,
		codes.Unimplemented, codes.Internal, codes.Unavailable, codes.DataLoss:
		return ServerError

	// Not sure
	case codes.Unknown:
		return Unknown
	default:
		return Unknown
	}

}

// SetSpanTags sets one or more tags on the given span according to the
// error.
func SetSpanTags(span opentracing.Span, err error, client bool) {

	code := codes.Unknown
	s, ok := status.FromError(err)
	if ok {
		code = s.Code()
	}

	class := ErrorClass(code)

	span.SetTag("response_code", code)
	span.SetTag("response_class", class)

	if err == nil {
		return
	}

	if client || class == ServerError {
		ext.Error.Set(span, true)
	}
}
