package consumers

import (
	"context"
	"fmt"

	nsq "github.com/nsqio/go-nsq"
)

type PortingDataRecordCreatedSubscriber struct {
}

func NewPortingDataRecordCreatedSubscriber(ctx context.Context) *PortingDataRecordCreatedSubscriber {
	return &PortingDataRecordCreatedSubscriber{}
}

func (p PortingDataRecordCreatedSubscriber) HandleMessage(ctx context.Context, m *nsq.Message) error {
	fmt.Println("#################")
	fmt.Println("received message")
	fmt.Println(m.Body)
	fmt.Println("##################")

	// 	var pd insysproto.PortingData
	// 	err := proto.Unmarshal(m.Body, &pd)
	// 	if err != nil {
	// 		return werror.Wrap(err, "could not unmarshall PortingDataCreated message body into proto for insysproto.PortingData struct")
	// 	}

	return nil
}
