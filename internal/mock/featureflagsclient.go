package mock

import (
	"context"

	"weavelab.xyz/monorail/shared/grpc-clients/client-grpc-clients/featureflagsclient"

	"weavelab.xyz/monorail/shared/wlib/uuid"
)

type FeatureFlagsClient struct {
	ListFn func(ctx context.Context, locationID uuid.UUID) ([]featureflagsclient.Flag, error)
}

func (s *FeatureFlagsClient) List(ctx context.Context, locationID uuid.UUID) ([]featureflagsclient.Flag, error) {
	return s.ListFn(ctx, locationID)
}
