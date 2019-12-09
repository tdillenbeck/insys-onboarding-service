package app

import (
	"context"

	"weavelab.xyz/monorail/shared/protorepo/dist/go/messages/insysproto"
)

type ChiliPiperScheduleEventPublisher interface {
	PublishCreated(ctx context.Context, response *insysproto.CreateChiliPiperScheduleEventResponse) error
}

type ZapierClient interface {
	Send(ctx context.Context, username, locationID, salesforceOpportunityID string) error
}
