package mock

import (
	"context"
)

type ZapierClient struct {
	SendCalled                  bool
	UsernameSent                string
	LocationIDSent              string
	SalesforceOpportunityIDSent string
	SendFn                      func(ctx context.Context, username, locationID, salesforceOpportunityID string) error
}

func (z *ZapierClient) Send(ctx context.Context, username, locationID, salesforceOpportunityID string) error {
	z.SendCalled = true
	z.UsernameSent = username
	z.LocationIDSent = locationID
	z.SalesforceOpportunityIDSent = salesforceOpportunityID
	return z.SendFn(ctx, username, locationID, salesforceOpportunityID)
}
