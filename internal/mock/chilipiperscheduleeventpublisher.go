package mock

import (
	"context"

	"weavelab.xyz/monorail/shared/protorepo/dist/go/messages/insysproto"
)

type ChiliPiperScheduleEventPublisher struct {
	PublishCreatedFn func(ctx context.Context, pd *insysproto.CreateChiliPiperScheduleEventResponse) error
}

func (c *ChiliPiperScheduleEventPublisher) PublishCreated(ctx context.Context, response *insysproto.CreateChiliPiperScheduleEventResponse) error {
	return c.PublishCreatedFn(ctx, response)
}
