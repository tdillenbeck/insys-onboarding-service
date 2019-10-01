package mock

import (
	"context"
)

type ZapierClient struct {
	SendFn func(ctx context.Context, username, locationID string) error
}

func (z *ZapierClient) Send(ctx context.Context, username, locationID string) error {
	return z.SendFn(ctx, username, locationID)
}
