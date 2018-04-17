package wnsq

import (
	"context"
	"github.com/nsqio/go-nsq"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/opentracing/opentracing-go/log"
	"weavelab.xyz/wlib/werror"
	"weavelab.xyz/wlib/wlog"
	"weavelab.xyz/wlib/wtracer"
)

func (w *Consumer) tracingConsumerMiddleware(msg *nsq.Message) error {

	ctx := context.Background()

	if tracer == nil {
		return w.handler.HandleMessage(ctx, msg)
	}

	md, err := extractMetaData(msg.Body)
	if err != nil {
		// if we're unable to extract meta data, don't fail
		// TODO: figure out something better to do than ignore the error
	}

	spanContext, err := tracer.Extract(opentracing.HTTPHeaders, MetadataReaderWriter(md))
	if err != nil && err != opentracing.ErrSpanContextNotFound {
		wlog.WError(werror.Wrap(err, "span context not found"))
	}

	name := "NSQ Consumer " + w.topic + "/" + w.channel

	span := tracer.StartSpan(name, ext.RPCServerOption(spanContext), nsqComponentTag)
	defer span.Finish()

	ctx = opentracing.ContextWithSpan(ctx, span)

	logPayloads := wtracer.ShouldLogBodies(span.Context())
	if logPayloads {
		span.LogFields(log.Object("nsq.body", msg.Body))
	}

	err = w.handler.HandleMessage(ctx, msg)
	if err != nil {
		span.LogFields(log.String("event", "error"), log.String("message", err.Error()))
		return err
	}

	span.SetTag("nsq.topic", w.topic)
	span.SetTag("nsq.channel", w.channel)

	return nil
}

//unwrap checks if the message needs to be unwrapped to a wnsqproto.NSQMessage instance and sets the correct data on the message body if it does.
func extractMetaData(b []byte) (map[string]string, error) {

	// peek and see if the message looks like JSON
	var err error
	var md map[string]string
	if isJSON(b) {
		md, err = extractJSON(b)
	} else {
		md, err = extractProto(b)
	}
	if err != nil {
		// what to do with error?
		return nil, nil
	}

	return md, nil
}
