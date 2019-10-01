package mock

import (
	"context"

	"weavelab.xyz/monorail/shared/grpc-clients/client-grpc-clients/featureflagsclient"

	"weavelab.xyz/monorail/shared/wlib/uuid"
)

type FeatureFlagsClient struct {
	ListFn   func(ctx context.Context, locationID uuid.UUID) ([]featureflagsclient.Flag, error)
	UpdateFn func(ctx context.Context, locationID uuid.UUID, name string, enable bool) error
}

func (s *FeatureFlagsClient) List(ctx context.Context, locationID uuid.UUID) ([]featureflagsclient.Flag, error) {
	return s.ListFn(ctx, locationID)
}

func (s *FeatureFlagsClient) Update(ctx context.Context, locationID uuid.UUID, name string, enable bool) error {
	return s.UpdateFn(ctx, locationID, name, enable)
}
