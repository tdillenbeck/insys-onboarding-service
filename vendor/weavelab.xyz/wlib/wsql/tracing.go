package wsql

import (
	"context"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/opentracing/opentracing-go/log"
	"weavelab.xyz/wlib/wcontext"
	"weavelab.xyz/wlib/wtracer"
)

var (
	// Morally a const:
	sqlComponentTag = opentracing.Tag{Key: string(ext.Component), Value: "SQL"}
)

func (p *PG) openTracingInterceptor(ctx context.Context, method string, query string) func() {

	var parentCtx opentracing.SpanContext
	if parent := opentracing.SpanFromContext(ctx); parent != nil {
		parentCtx = parent.Context()
	}

	name := "SQL " + method

	clientSpan := p.tracer.StartSpan(name, opentracing.ChildOf(parentCtx), ext.SpanKindConsumer, sqlComponentTag)

	done := func() {
		clientSpan.LogFields(log.String("query", query))
		clientSpan.SetTag(wtracer.RequestIDTag, wcontext.RequestID(ctx))
		clientSpan.Finish()
	}

	return done

}
