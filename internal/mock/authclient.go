package mock

import (
	"context"

	"weavelab.xyz/monorail/shared/grpc-clients/client-grpc-clients/authclient"
	"weavelab.xyz/monorail/shared/wlib/uuid"
)

type Auth struct {
	UserLocationsFn func(ctx context.Context, userID uuid.UUID) (*authclient.UserAccess, error)
}

func (a *Auth) UserLocations(ctx context.Context, userID uuid.UUID) (*authclient.UserAccess, error) {
	return a.UserLocationsFn(ctx, userID)
}
