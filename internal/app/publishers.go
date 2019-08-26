package app

import (
	"context"

	"weavelab.xyz/monorail/shared/protorepo/dist/go/messages/insysproto"
)

type ChiliPiperScheduleEventPublisher interface {
	PublishCreated(ctx context.Context, pd *insysproto.CreateChiliPiperScheduleEventResponse) error
}
