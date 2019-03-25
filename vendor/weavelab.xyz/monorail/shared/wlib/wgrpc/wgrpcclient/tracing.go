package wgrpcclient

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
	"weavelab.xyz/monorail/shared/wlib/wgrpc"
	"weavelab.xyz/monorail/shared/wlib/wtracer"
)

func NewUnaryTracerInterceptor(tracer opentracing.Tracer) (grpc.UnaryClientInterceptor, error) {

	if tracer == nil {
		var err error
		tracer, err = wtracer.DefaultTracer()
		if err != nil {
			return nil, err
		}
	}

	i := openTracingClientInterceptor(tracer)

	return i, nil
}

func openTracingClientInterceptor(tracer opentracing.Tracer) grpc.UnaryClientInterceptor {

	return func(
		ctx context.Context,
		method string,
		req, resp interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {

		var err error
		var parentCtx opentracing.SpanContext
		if parent := opentracing.SpanFromContext(ctx); parent != nil {
			parentCtx = parent.Context()
		}

		name := "gRPC " + method

		clientSpan := tracer.StartSpan(name, opentracing.ChildOf(parentCtx), ext.SpanKindRPCClient, wgrpc.GRPCComponentTag)
		defer clientSpan.Finish()

		requestLogged := false
		logPayloads := wtracer.ShouldLogBodies(clientSpan.Context())
		// check to see if the request should be logged or not
		if logPayloads {
			clientSpan.LogFields(log.String("gRPC.request", marshalJSON(req)))
			requestLogged = true
		}

		md, ok := metadata.FromOutgoingContext(ctx)
		if !ok {
			md = metadata.New(nil)
		} else {
			md = md.Copy()
		}

		mdWriter := wgrpc.MetadataReaderWriter{MD: md}
		err = tracer.Inject(clientSpan.Context(), opentracing.HTTPHeaders, mdWriter)
		// We have no better place to record an error than the Span itself :-/
		if err != nil {
			clientSpan.LogFields(log.String("event", "Tracer.Inject() failed"), log.Error(err))
		}

		ctx = metadata.NewOutgoingContext(ctx, md)

		err = invoker(ctx, method, req, resp, cc, opts...)
		if err != nil {
			if requestLogged == false {
				clientSpan.LogFields(log.String("gRPC.request", marshalJSON(req)))
			}
			clientSpan.LogFields(
				log.Error(err),
			)
		}

		if logPayloads {
			out := marshalJSON(resp)
			if len(out) > wtracer.MaxLogFieldSize {
				out = "body too large"
			}

			clientSpan.LogFields(log.String("gRPC.response", out))
		}

		wgrpc.SetSpanTags(clientSpan, err, true)

		clientSpan.SetTag(wtracer.RequestIDTag, wcontext.RequestID(ctx))

		return err
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

	// TODO: need to scrub the marshalled JSON for secrets

	return string(s)
}
