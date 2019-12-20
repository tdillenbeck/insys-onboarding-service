package featureflagsclient

import (
	"context"

	"weavelab.xyz/monorail/shared/protorepo/dist/go/messages/platformproto"
	"weavelab.xyz/monorail/shared/protorepo/dist/go/services/platform"
	"weavelab.xyz/monorail/shared/wlib/uuid"
	"weavelab.xyz/monorail/shared/wlib/wdns"
	"weavelab.xyz/monorail/shared/wlib/werror"
	"weavelab.xyz/monorail/shared/wlib/wgrpc/wgrpcclient"
	"weavelab.xyz/monorail/shared/wlib/wlog"
)

const (
	defaultAddress = "client-feature-flags.client.svc.cluster.local:grpc"
)

type Client struct {
	c platform.FeatureFlagsClient
}
type Flag struct {
	Name        string
	Value       bool
	DisplayName string
}

func New(ctx context.Context, addr string) (*Client, error) {

	if addr == "" {
		var err error
		addr, err = wdns.ResolveAddress(defaultAddress)
		if err != nil {
			return nil, werror.Wrap(err, "unable to use default address")
		}
	}

	g, err := wgrpcclient.NewDefault(ctx, addr)
	if err != nil {
		return nil, werror.Wrap(err, "unable to setup grpc client")
	}

	c := platform.NewFeatureFlagsClient(g)

	client := Client{
		c: c,
	}

	return &client, nil

}

//Check if a featureFlag is enabled for a locationID
func (c *Client) Enabled(ctx context.Context, locationID uuid.UUID, flag string) bool {
	if c == nil {
		return false
	}

	ff, err := c.c.LocationFlag(ctx, &platformproto.LocationFlagsRequest{FeatureFlag: flag, LocationID: locationID.String()})
	if err != nil {
		wlog.WError(werror.Wrap(err).Add("locationID", locationID).Add("flagname", flag))
		return false
	}

	return ff.Value
}
