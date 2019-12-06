package mock

import (
	"context"
)

type ZapierClient struct {
	SendFn func(ctx context.Context, username, locationID, salesforceOpportunityID string) error
}

func (z *ZapierClient) Send(ctx context.Context, username, locationID, salesforceOpportunityID string) error {
	return z.SendFn(ctx, username, locationID, salesforceOpportunityID)
}
