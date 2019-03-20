package wtracer

import (
	"context"
	"fmt"
	"net/http"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"weavelab.xyz/monorail/shared/wlib/version"
	"weavelab.xyz/monorail/shared/wlib/werror"
	"weavelab.xyz/monorail/shared/wlib/wlog"
)

const (
	tracerBodyBaggage = "include-body"
	TraceBodyHeader   = jaeger.TraceBaggageHeaderPrefix + tracerBodyBaggage

	RequestIDTag      = "RequestID"
	HTTPURLPatternTag = "Pattern"

	UberTraceIDHeader = "Uber-Trace-Id"

	MaxLogFieldSize = 64000

	lowerBoundSamplingRate = 10.0 / 60 // at least once every a minute
	samplingRate           = 0.001     // random sampling rate

)

type (
	Stopper func(context.Context) error
)

var tracingConfigured = false

func DefaultTracer() (opentracing.Tracer, error) {

	if tracingConfigured == false {

		t, _, err := New(DefaultSampler)
		if err != nil {
			return nil, werror.Wrap(err)
		}

		opentracing.SetGlobalTracer(t)
		tracingConfigured = true
	}

	t := opentracing.GlobalTracer()

	return t, nil
}

type SamplerType int

const (
	DefaultSampler = SamplerType(iota)
	RemoteControlledSampler
	AlwaysSampler
	NeverSampler
)

func New(samplerType SamplerType) (opentracing.Tracer, Stopper, error) {
	/*
		zipkinPropagator := zipkin.NewZipkinB3HTTPHeaderPropagator()

		zipkinInjector := jaeger.TracerOptions.Injector(opentracing.HTTPHeaders, zipkinPropagator)
		zipkinExtractor := jaeger.TracerOptions.Extractor(opentracing.HTTPHeaders, zipkinPropagator)

		// Zipkin shares span ID between client and server spans; it must be enabled via the following option.
		zipkinSharedRPCSpan := jaeger.TracerOptions.ZipkinSharedRPCSpan(true)
	*/

	metrics := newMetrics()
	logger := newLogger()

	appName := version.Info().Name

	var sampler jaeger.Sampler
	switch samplerType {
	case NeverSampler:
		sampler = jaeger.NewConstSampler(false)

	case AlwaysSampler:
		sampler = jaeger.NewConstSampler(true)

	case RemoteControlledSampler:
		samplerMetrics := jaeger.SamplerOptions.Metrics(metrics)
		samplerLogger := jaeger.SamplerOptions.Logger(logger)
		sampler = jaeger.NewRemotelyControlledSampler(appName, samplerMetrics, samplerLogger)

	default:
		// TODO: make this configurable
		lowerBound := lowerBoundSamplingRate
		samplingRate := samplingRate // sample rate
		var err error
		sampler, err = jaeger.NewGuaranteedThroughputProbabilisticSampler(lowerBound, samplingRate)
		if err != nil {
			return nil, nil, werror.Wrap(err, "unable to configure sampler")
		}
	}

	// use default port (6831) and default max packet length (65000)
	sender, err := jaeger.NewUDPTransport("", 0)
	if err != nil {
		return nil, nil, werror.Wrap(err, "unable to create UDP sender for reporter")
	}

	reportMetrics := jaeger.ReporterOptions.Metrics(metrics)
	reportLogger := jaeger.ReporterOptions.Logger(logger)
	reporter := jaeger.NewRemoteReporter(sender, reportMetrics, reportLogger)

	// create Jaeger tracer
	tracer, closer := jaeger.NewTracer(
		appName,
		sampler,
		reporter,
		//		zipkinInjector,
		//		zipkinExtractor,
		//		zipkinSharedRPCSpan,
	)

	stopper := func(_ context.Context) error {
		err := closer.Close()
		if err != nil {
			return werror.Wrap(err, "unable to close Jaeger tracer")
		}

		return nil
	}

	return tracer, stopper, nil

}

type logger struct{}

func newLogger() *logger {
	return &logger{}
}

func (l *logger) Error(msg string) {
	wlog.WError(werror.New(msg))
}

// Infof is a formatted print, it could be improved to handle the structured logging better
func (l *logger) Infof(msg string, args ...interface{}) {
	wlog.Info(fmt.Sprintf(msg, args...))
}

func newMetrics() *jaeger.Metrics {
	return jaeger.NewNullMetrics()
}

func ShouldLogBodies(ctx opentracing.SpanContext) bool {

	l := false

	ctx.ForeachBaggageItem(func(k, v string) bool {
		// returns whether or not we not to look at more items

		if k == tracerBodyBaggage {
			l = true
			return false // we found what we were looking for, loop no more
		}

		return true
	})

	return l
}

// SetOutgoingTraceID -- grabs the traceID from the context and adds it to the outgoing context
// Particularly useful for reverse proxies
func SetOutgoingTraceID(ctx context.Context, r *http.Request) {

	span := opentracing.SpanFromContext(ctx)
	if span == nil {
		return
	}

	spanCtx := span.Context()

	sc, ok := spanCtx.(jaeger.SpanContext)
	if !ok {
		return
	}

	r.Header.Add(UberTraceIDHeader, sc.String())

}
