package featureflagsclient

import (
	"context"
	"time"

	"github.com/patrickmn/go-cache"
	"weavelab.xyz/monorail/shared/wlib/uuid"
)

type CacheClient struct {
	featureflagclient

	cacheDuration time.Duration
	cache         *cache.Cache
}

func NewCacheClient(ctx context.Context, ffClient featureflagclient, cacheDuration time.Duration, cleanupInterval time.Duration) (*CacheClient, error) {

	cache := cache.New(cacheDuration, cleanupInterval)

	client := CacheClient{
		featureflagclient: ffClient,

		cacheDuration: cacheDuration,
		cache:         cache,
	}

	return &client, nil

}

//Check if a featureFlag is enabled for a locationID
func (c *CacheClient) Enabled(ctx context.Context, locationID uuid.UUID, flag string) bool {
	if c.cacheDuration == 0 {
		return c.Enabled(ctx, locationID, flag)
	}

	key := key(locationID, flag)

	cached, ok := c.cache.Get(key)
	if ok {
		switch result := cached.(type) {
		case bool:
			return result
		}
	}

	result := c.featureflagclient.Enabled(ctx, locationID, flag)

	c.cache.Set(key, result, cache.DefaultExpiration)

	return result
}

func key(locationID uuid.UUID, flag string) string {
	return locationID.String() + flag
}
