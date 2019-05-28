package wnsq

import (
	"time"

	"context"

	"github.com/nsqio/go-nsq"
	"weavelab.xyz/monorail/shared/wlib/werror"
)

//Producer is an interface used for sending messages to a given queue and topic
type Producer struct {
	*nsq.Producer
}

//NewProducer wraps an NSQ Producer with a Weave implementation that allows publishing a message with a context.
func NewProducer(addr string, config *nsq.Config) (*Producer, error) {

	producer, err := nsq.NewProducer(addr, config)
	if err != nil {
		return nil, werror.Wrap(err, "error creating NSQ producer").Add("addr", addr)
	}

	producer.SetLogger(infoLogger, nsq.LogLevelInfo)
	producer.SetLogger(warningLogger, nsq.LogLevelWarning)
	producer.SetLogger(errorLogger, nsq.LogLevelError)

	return &Producer{Producer: producer}, nil
}

//DeferredPublishWithContext wraps the data with ddebug info if necessary and called DeferredPublish.
func (mq *Producer) DeferredPublish(ctx context.Context, topic string, delay time.Duration, data []byte) error {

	p := func(topic string, data []byte) error {
		err := mq.Producer.DeferredPublish(topic, delay, data)
		if err != nil {
			return werror.Wrap(err, "could not publish message").Add("topic", topic)
		}
		return nil
	}

	err := tracingProducerMiddleware(ctx, topic, data, p)
	if err != nil {
		return err
	}

	return nil
}

//PublishWithContext is used to publish to NSQ. The context will be used to extract values such as the DebugID for distributed debugging.
func (mq *Producer) Publish(ctx context.Context, topic string, data []byte) error {

	p := func(topic string, data []byte) error {

		err := pub(mq.Producer, topic, data)
		if err != nil {
			return werror.Wrap(err, "could not publish message").Add("topic", topic)
		}

		return nil
	}

	err := tracingProducerMiddleware(ctx, topic, data, p)
	if err != nil {
		return err
	}

	return nil
}

var pub = func(producer *nsq.Producer, topic string, msg []byte) error {
	return producer.Publish(topic, msg)
}

//wrap wraps the data with a wnsqproto.NSQMessage and marshals the struct to bytes. It then prepends a prefix to the bytes so the consumer knows that the message must be unwrapped.
func inject(md map[string]string, data []byte) ([]byte, error) {

	var injector func(map[string]string, []byte) ([]byte, error)
	switch {
	case isJSON(data):
		injector = injectorJSON
	case isProto(data):
		injector = injectorProto
	default:
	}

	if injector == nil {
		return data, nil
	}

	// get the data out of the context that we need to inject

	data, err := injector(md, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
