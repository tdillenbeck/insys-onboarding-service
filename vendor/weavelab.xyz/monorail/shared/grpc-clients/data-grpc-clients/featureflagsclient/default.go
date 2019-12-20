package featureflagsclient

import (
	"context"

	"weavelab.xyz/monorail/shared/wlib/uuid"
)

//FeatureFlag is unique in that it gets called all over the place and we want to be able to call vendor/featureflag.Enabled
//Rather than setting up a feature flag client, making it public and calling that internal/featureflag Enabled in every repo.
type featureflagclient interface {
	Enabled(ctx context.Context, locationID uuid.UUID, flag string) bool
}

var defaultClient featureflagclient

func Enabled(ctx context.Context, locationID uuid.UUID, flag string) bool {
	if defaultClient == nil {
		return false
	}
	return defaultClient.Enabled(ctx, locationID, flag)
}

func InitDefault(f featureflagclient) {
	defaultClient = f
}
