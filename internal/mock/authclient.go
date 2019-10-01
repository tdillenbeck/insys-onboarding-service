package mock

import (
	"context"

	"weavelab.xyz/monorail/shared/grpc-clients/client-grpc-clients/authclient"
	"weavelab.xyz/monorail/shared/wlib/uuid"
)

type AuthClient struct {
	UserLocationsFn func(ctx context.Context, userID uuid.UUID) (*authclient.UserAccess, error)
}

func (s *AuthClient) UserLocations(ctx context.Context, userID uuid.UUID) (*authclient.UserAccess, error) {
	return s.UserLocationsFn(ctx, userID)
}
