package mock

import (
	"context"

	"weavelab.xyz/monorail/shared/protorepo/dist/go/messages/insysproto"
)

type ChiliPiperScheduledEventPublisher struct {
	PublishCreatedFn func(ctx context.Context, pd *insysproto.CreateChiliPiperScheduleEventResponse) error
}

func (c *ChiliPiperScheduledEventPublisher) PublishCreated(ctx context.Context, response *insysproto.CreateChiliPiperScheduleEventResponse) error {
	return c.PublishCreatedFn(ctx, response)
}
