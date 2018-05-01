package wvault

import (
	"context"
	"time"
)

type Secret interface {
	Name() string
	Parent() Secret
	Expiration() time.Time
	Refresh(ctx context.Context, recreate bool) (bool, error) // renews the lease
	LeaseID() string                                          // returns the leaseID associated with the secret

	Token() string // returns the token used for the secret
}
