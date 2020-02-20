package mock

import (
	"context"

	"weavelab.xyz/insys-onboarding-service/internal/zapier"
)

type ZapierClient struct {
	SendCalled   bool
	PayloadsSent []zapier.FirstLoginEventPayload
	SendFn       func(ctx context.Context, username, locationID, salesforceOpportunityID string) error
}

func (z *ZapierClient) Send(ctx context.Context, username, locationID, salesforceOpportunityID string) error {
	z.SendCalled = true
	z.PayloadsSent = append(z.PayloadsSent, zapier.FirstLoginEventPayload{Username: username, LocationID: locationID, SalesforceOpportunityID: salesforceOpportunityID})
	return z.SendFn(ctx, username, locationID, salesforceOpportunityID)
}
