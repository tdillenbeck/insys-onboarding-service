package wgrpcserver

import (
	"context"
	"encoding/json"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/opentracing/opentracing-go/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"weavelab.xyz/monorail/shared/wlib/wcontext"
	"weavelab.xyz/monorail/shared/wlib/werror"
	"weavelab.xyz/monorail/shared/wlib/wgrpc"
	"weavelab.xyz/monorail/shared/wlib/wlog"
	"weavelab.xyz/monorail/shared/wlib/wtracer"
)

func NewUnaryTracerInterceptor(tracer opentracing.Tracer) (grpc.UnaryServerInterceptor, error) {

	if tracer == nil {
		var err error
		tracer, err = wtracer.DefaultTracer()
		if err != nil {
			return nil, err
		}
	}

	i := openTracingServerInterceptor(tracer)

	return i, nil
}

func openTracingServerInterceptor(tracer opentracing.Tracer) grpc.UnaryServerInterceptor {

	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {

		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			md = metadata.New(nil)
		}

		spanContext, err := tracer.Extract(opentracing.HTTPHeaders, wgrpc.MetadataReaderWriter{MD: md})
		if err != nil && err != opentracing.ErrSpanContextNotFound {
			wlog.WError(werror.Wrap(err, "span context not found"))
		}

		name := "gRPC " + info.FullMethod
		serverSpan := tracer.StartSpan(name, ext.RPCServerOption(spanContext), wgrpc.GRPCComponentTag)
		defer serverSpan.Finish()

		// check to see if the request should be logged or not
		requestLogged := false
		logPayloads := wtracer.ShouldLogBodies(serverSpan.Context())
		if logPayloads {
			serverSpan.LogFields(log.String("gRPC.request", marshalJSON(req)))
			requestLogged = true
		}
		if logPayloads {
			ctx = wlog.WithDebugPrint(ctx, true)
		}

		ctx = opentracing.ContextWithSpan(ctx, serverSpan)

		resp, err = handler(ctx, req)
		if err != nil {
			if requestLogged == false {
				serverSpan.LogFields(log.String("gRPC.request", marshalJSON(req)))
			}
			serverSpan.LogFields(
				log.Error(err),
			)
		}

		if logPayloads {
			out := marshalJSON(resp)
			if len(out) > wtracer.MaxLogFieldSize {
				out = "body too large"
			}
			serverSpan.LogFields(log.String("gRPC.response", out))
		}

		wgrpc.SetSpanTags(serverSpan, err, false)

		serverSpan.SetTag(wtracer.RequestIDTag, wcontext.RequestID(ctx))

		return resp, err
	}
}

var marshaler = jsonpb.Marshaler{
	EmitDefaults: true,
	OrigName:     true,
}

func marshalJSON(o interface{}) string {
	var s string

	m, ok := o.(proto.Message)
	if ok {
		s, _ = marshaler.MarshalToString(m)
	} else {
		b, _ := json.Marshal(o)
		s = string(b)
	}

	return string(s)
}
