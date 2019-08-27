package featureflagsclient

import (
	"context"

	"weavelab.xyz/monorail/shared/protorepo/dist/go/messages/platformproto"
	"weavelab.xyz/monorail/shared/protorepo/dist/go/messages/sharedproto"
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

//List all featureFlags
func (c *Client) List(ctx context.Context, locationID uuid.UUID) ([]Flag, error) {
	if c == nil {
		return nil, werror.New("FeatureFlagClient is nil")
	}

	f, err := c.c.ListLocationFlags(ctx, &platformproto.LocationFlagsRequest{LocationID: locationID.String()})
	if err != nil {
		return nil, werror.Wrap(err).Add("locationID", locationID.String())
	}

	//map to flag struct
	flags := make([]Flag, 0, len(f.Flags))
	for _, v := range f.Flags {

		flags = append(flags, Flag{
			Name:        v.Name,
			Value:       v.Value,
			DisplayName: v.DisplayName,
		})
	}

	return flags, nil
}

func (c *Client) Update(ctx context.Context, locationID uuid.UUID, name string, enable bool) error {
	req := &platformproto.LocationFlagUpdateRequest{
		LocationID: locationID.String(),
		Name:       name,
		Value:      enable,
	}

	_, err := c.c.UpdateLocationFlag(ctx, req)
	if err != nil {
		return werror.Wrap(err, "error updating feature flag")
	}

	return nil
}

// ListLocationsByFlag lists all locations that have the specified feature flag set with the value provided. Only
// locations with records in the lisa.location_featureflags table will be returned
func (c *Client) ListLocationsByFlag(ctx context.Context, flag string, value bool) ([]uuid.UUID, error) {
	locationIDs := make([]uuid.UUID, 0)

	req := &platformproto.ListLocationsByFlagRequest{
		Name:  flag,
		Value: value,
	}

	resp, err := c.c.ListLocationsByFlag(ctx, req)
	if err != nil {
		return locationIDs, werror.Wrap(err, "failed to fetch locations by flag and value")
	}

	for _, locationID := range resp.LocationIDs {
		lID, err := locationID.UUID()
		if err != nil {
			return locationIDs, werror.Wrap(err, "invalid location id").Add("locationID", locationID.String())
		}

		locationIDs = append(locationIDs, lID)
	}

	return locationIDs, nil
}

// UpdatePercentage updates the feature flag value of the
// requested percentage of locations to the specified value and returns
// the count of locations changed.
func (c *Client) UpdatePercentage(ctx context.Context, flag string, value bool, percent float32) (int, error) {
	req := &platformproto.MassFlagUpdateRequest{
		Name:    flag,
		Value:   value,
		Percent: percent,
	}

	resp, err := c.c.MassLocationFlagUpdate(ctx, req)
	if err != nil {
		return 0, werror.Wrap(err, "failed to mass update percentage of locations")
	}
	return int(resp.Count), nil
}

// UpdateLocations updates the feature flag value of the
// requested list of locations to the specified value and returns
// the count of locations changed.
func (c *Client) UpdateLocations(ctx context.Context, flag string, value bool, locationIDs ...uuid.UUID) (int, error) {

	protoIDs := make([]*sharedproto.UUID, len(locationIDs))
	for i := range locationIDs {
		protoIDs[i] = sharedproto.UUIDToProto(locationIDs[i])
	}
	req := &platformproto.MassFlagUpdateRequest{
		Name:        flag,
		Value:       value,
		LocationIDs: protoIDs,
	}

	resp, err := c.c.MassLocationFlagUpdate(ctx, req)
	if err != nil {
		return 0, werror.Wrap(err, "failed to mass update locations")
	}
	return int(resp.Count), nil
}
