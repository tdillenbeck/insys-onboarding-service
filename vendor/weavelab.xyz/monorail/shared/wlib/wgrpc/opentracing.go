// orignal source from github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc

package wgrpc

import (
	"strings"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"weavelab.xyz/monorail/shared/wlib/werror"
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

// SetSpanTags sets one or more tags on the given span according to the
// error.
func SetSpanTags(span opentracing.Span, err error, client bool) {

	code := codes.Unknown
	s, ok := status.FromError(err)
	if ok {
		code = s.Code()
	}

	class := werror.ErrorClass(werror.Code(code))

	span.SetTag("response_code", code)
	span.SetTag("response_class", class)

	if err == nil {
		return
	}

	if client || class == werror.ServerError {
		ext.Error.Set(span, true)
	}
}
