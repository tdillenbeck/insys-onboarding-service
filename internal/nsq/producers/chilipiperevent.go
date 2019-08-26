package producers

import (
	"context"

	"github.com/gogo/protobuf/proto"
	"weavelab.xyz/monorail/shared/protorepo/dist/go/messages/insysproto"
	"weavelab.xyz/monorail/shared/wlib/werror"
)

type ChiliPiperScheduleEventPublisher struct {
	createdTopic string
}

func NewChiliPiperScheduleEventCreatedPublisher(createdTopic string) *ChiliPiperScheduleEventPublisher {
	return &ChiliPiperScheduleEventPublisher{
		createdTopic: createdTopic,
	}
}

func (p ChiliPiperScheduleEventPublisher) PublishCreated(ctx context.Context, response *insysproto.CreateChiliPiperScheduleEventResponse) error {
	msg, err := proto.Marshal(response)
	if err != nil {
		return werror.Wrap(err, "unable to marshal CreateChiliPiperScheduleEventResponse to proto for publishing on nsq")
	}

	err = sendToNSQ(ctx, p.createdTopic, msg)
	if err != nil {
		return werror.Wrap(err, "unable to publish message to NSQ created topic")
	}

	return nil
}
