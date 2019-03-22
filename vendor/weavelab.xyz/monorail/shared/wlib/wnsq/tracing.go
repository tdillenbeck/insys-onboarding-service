package wnsq

import (
	"strings"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"weavelab.xyz/monorail/shared/wlib/wtracer"
)

var (
	nsqComponentTag = opentracing.Tag{Key: string(ext.Component), Value: "NSQ"}

	tracer, _ = wtracer.DefaultTracer()
)

// metadataReaderWriter satisfies both the opentracing.TextMapReader and
// opentracing.TextMapWriter interfaces.
type MetadataReaderWriter map[string]string // should this be a map[string][]string???, the proto spec currently doesn't allow it

func (w MetadataReaderWriter) Set(key, val string) {
	if w == nil {
		w = make(map[string]string)
	}

	key = strings.ToLower(key)
	w[key] = val
}

func (w MetadataReaderWriter) ForeachKey(handler func(key, val string) error) error {
	if w == nil {
		w = make(map[string]string)
	}

	for k, v := range w {
		if err := handler(k, v); err != nil {
			return err
		}
	}

	return nil
}
