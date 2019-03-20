package wnsq

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/opentracing/opentracing-go/log"
	"weavelab.xyz/monorail/shared/wlib/wtracer"
)

func tracingProducerMiddleware(ctx context.Context, topic string, data []byte, p func(topic string, data []byte) error) error {

	var parentCtx opentracing.SpanContext
	if parent := opentracing.SpanFromContext(ctx); parent != nil {
		parentCtx = parent.Context()
	}

	name := "NSQ Producer " + topic

	clientSpan := tracer.StartSpan(name, opentracing.ChildOf(parentCtx), ext.SpanKindRPCClient, nsqComponentTag)
	defer clientSpan.Finish()

	logs := make([]log.Field, 0, 3)

	mdWriter := MetadataReaderWriter(make(map[string]string))

	err := tracer.Inject(clientSpan.Context(), opentracing.HTTPHeaders, mdWriter)
	if err != nil {
		// We have no better place to record an error than the Span itself :-/
		logs = append(logs, log.String("event", "Tracer.Inject() failed"), log.Error(err))
	}

	logPayloads := wtracer.ShouldLogBodies(clientSpan.Context())
	if logPayloads {
		logs = append(logs, log.Object("nsq.body", data))
	}

	// add the metadata to the outgoing request
	data, err = inject(mdWriter, data)
	if err != nil {
		logs = append(logs, log.String("error.inject", err.Error()))
	}

	if len(logs) > 0 {
		clientSpan.LogFields(logs...)
	}

	// publish the message
	err = p(topic, data)
	if err != nil {
		clientSpan.LogFields(log.String("event", "error"), log.String("message", err.Error()))
	}

	clientSpan.SetTag("nsq.topic", topic)

	return err
}
