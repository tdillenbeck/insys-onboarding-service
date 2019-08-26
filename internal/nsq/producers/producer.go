package producers

import (
	"context"
	"log"

	"github.com/nsqio/go-nsq"
	"weavelab.xyz/monorail/shared/wlib/werror"
	"weavelab.xyz/monorail/shared/wlib/wnsq"
)

var producer *wnsq.Producer

func Init(nsqdAddress string) {
	var err error

	cfg := nsq.NewConfig()
	producer, err = wnsq.NewProducer(nsqdAddress, cfg)
	if err != nil {
		log.Fatal(err)
	}
}

func sendToNSQ(ctx context.Context, topic string, msg []byte) error {
	err := producer.Publish(ctx, topic, msg)
	if err != nil {
		return werror.Wrap(err)
	}

	return nil
}
