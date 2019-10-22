package wlog

import (
	"context"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"

	"weavelab.xyz/monorail/shared/wlib/wlog/tag"
)

var _ LogMiddlewareFunc = TracingMiddleware

func TracingMiddleware(ctx *context.Context, mType *LogMsgType, msg *string, tags *[]tag.Tag) {

	// don't try if there is no context given
	if ctx == nil || *ctx == nil {
		return
	}

	span := opentracing.SpanFromContext(*ctx)
	if span == nil {
		// if there's no span, don't start one
		return
	}

	if mType == nil {
		t := INFO
		mType = &t
	}

	// decide if we should add the log message to the span
	include := shouldLog(span, *mType)
	if include == false {
		return
	}

	m := ""
	if msg != nil {
		m = *msg
	}

	tagString := ""
	if tags != nil {
		tagString = tagsString(*tags)
	}

	mTypeString := mType.String()

	span.LogFields(
		log.String(mTypeString, m+" "+tagString),
	)

}

func tagsString(tags []tag.Tag) string {
	var tagsString string
	s := ""
	for i, t := range tags {
		if i == 1 {
			s = ", "
		}
		tagsString += s + t.String()
	}

	return tagsString
}

const includeBodyBaggageKey = "include-body"

func shouldLog(span opentracing.Span, logLevel LogMsgType) bool {

	// Error and Info should always be attached to spans
	// Debug should only be logged iff include-body is set
	switch logLevel {
	case ERROR:
		return true
	case INFO:
		return true
	}

	include := span.BaggageItem(includeBodyBaggageKey)
	if include != "" {
		return true
	}

	return false
}
