package wvault

import (
	"context"

	"github.com/opentracing/opentracing-go"
	"weavelab.xyz/monorail/shared/wlib/wtracer"
)

var (
	AlwaysTraceRenewRequests = true
	alwaysTracer             opentracing.Tracer
)

func WithNewTracingParent(ctx context.Context) (context.Context, func()) {
	if AlwaysTraceRenewRequests == false {
		return ctx, func() {}
	}

	// create a new parent span and force tracing

	// make sure tracer is configured
	if alwaysTracer == nil {
		var err error
		alwaysTracer, _, err = wtracer.New(wtracer.AlwaysSampler)
		if err != nil {
			return ctx, func() {}
		}
	}

	clientSpan := alwaysTracer.StartSpan("Vault")

	ctx = opentracing.ContextWithSpan(ctx, clientSpan)

	return ctx, clientSpan.Finish
}
