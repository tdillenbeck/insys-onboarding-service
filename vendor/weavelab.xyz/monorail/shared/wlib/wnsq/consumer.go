package wnsq

import (
	"context"
	"strings"

	"github.com/nsqio/go-nsq"
	"weavelab.xyz/monorail/shared/wlib/werror"
	"weavelab.xyz/monorail/shared/wlib/wlog"
	"weavelab.xyz/monorail/shared/wlib/wmetrics"
)

const (
	metricWNSQ = "wnsq"
)

func init() {
	wmetrics.SetLabels(metricWNSQ, "topic", "channel", "status")
}

//Consumer wraps an NSQ Consumer with Weave feautures such as DDebugging
type Consumer struct {
	*nsq.Consumer
	handler Handler
	topic   string
	channel string

	metricsTopic   string // cleaned versions of topic
	metricsChannel string // and channel for use with metrics
}

//Handler allows passing a ctx with the nsq message
type Handler interface {
	HandleMessage(ctx context.Context, msg *nsq.Message) error
}

// HandlerFunc is a convenience type to avoid having to declare a struct
// to implement the Handler interface, it can be used like this:
//
// 	consumer.AddHandler(nsq.HandlerFunc(func(m *Message) error {
// 		// handle the message
// 	}))
type HandlerFunc func(ctx context.Context, message *nsq.Message) error

// HandleMessage implements the Handler interface
func (h HandlerFunc) HandleMessage(ctx context.Context, m *nsq.Message) error {
	return h(ctx, m)
}

//NewConsumer creates a new NSQ Consumer that wraps the handler with a DDebug handler. Each message is deserialized into an NSQDebugStruct, if a value exists for DebugHeader then the message will be logged
func NewConsumer(topic string, channel string, config *nsq.Config) (*Consumer, error) {

	consumer, err := nsq.NewConsumer(topic, channel, config)
	if err != nil {
		return nil, werror.Wrap(err, "error creating consumer").Add("topic", topic).Add("channel", channel)
	}

	consumer.SetLogger(infoLogger, nsq.LogLevelInfo)
	consumer.SetLogger(warningLogger, nsq.LogLevelWarning)
	consumer.SetLogger(errorLogger, nsq.LogLevelError)

	// dot's get separated into separate tags, remove them
	mTopic := strings.Replace(topic, ".", "_", -1)
	mChannel := strings.Replace(channel, ".", "_", -1)

	wConsumer := Consumer{
		Consumer: consumer,
		topic:    topic,
		channel:  channel,

		metricsTopic:   mTopic,
		metricsChannel: mChannel,
	}

	return &wConsumer, nil
}

//AddHandler adds a handler to the underlying NSQ consumer. It is wrapped by a handler to take care of any DDebug messages
func (w *Consumer) AddHandler(h Handler) {
	w.handler = h
	w.Consumer.AddHandler(w)
}

//AddConcurrentHandlers adds a handler to the underlying NSQ consumer. It is wrapped by a handler to take care of any DDebug messages
func (w *Consumer) AddConcurrentHandlers(h Handler, concurrency int) {
	w.handler = h
	w.Consumer.AddConcurrentHandlers(w, concurrency)
}

//HandleMessage handles any necessary Open Tracing debugging and then passes the message to the underlying handler.
func (w *Consumer) HandleMessage(message *nsq.Message) error {

	if w.handler == nil {
		err := werror.New("no handler has been added to the NSQ consumer")
		wlog.WError(err)
		return err
	}

	stop := wmetrics.StartTimer(metricWNSQ, w.metricsTopic, w.metricsChannel)

	// call the tracing middleware
	err := w.tracingConsumerMiddleware(message)
	if err != nil {
		stop("error")
		return werror.Wrap(err)
	}

	stop("success")
	return nil
}
