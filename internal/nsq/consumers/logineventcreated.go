package consumers

import (
	"context"
	"fmt"
	"time"

	"github.com/gogo/protobuf/proto"
	nsq "github.com/nsqio/go-nsq"
	"weavelab.xyz/monorail/shared/grpc-clients/client-grpc-clients/authclient"
	"weavelab.xyz/monorail/shared/grpc-clients/client-grpc-clients/featureflagsclient"
	"weavelab.xyz/monorail/shared/protorepo/dist/go/messages/client/clientproto"
	"weavelab.xyz/monorail/shared/protorepo/dist/go/messages/insysproto"
	"weavelab.xyz/monorail/shared/protorepo/dist/go/services/insys"
	"weavelab.xyz/monorail/shared/wlib/uuid"
	"weavelab.xyz/monorail/shared/wlib/werror"
	"weavelab.xyz/monorail/shared/wlib/wgrpc"
	"weavelab.xyz/monorail/shared/wlib/wlog"
)

type AuthClient interface {
	UserLocations(ctx context.Context, userID uuid.UUID) (*authclient.UserAccess, error)
}

type FeatureFlagsClient interface {
	List(ctx context.Context, locationID uuid.UUID) ([]featureflagsclient.Flag, error)
	Update(ctx context.Context, locationID uuid.UUID, name string, enable bool) error
}

type ZapierClient interface {
	Send(ctx context.Context, username, locationID, salesforceOpportunityID string) error
}

type LogInEventCreatedSubscriber struct {
	authClient         AuthClient
	featureFlagsClient FeatureFlagsClient
	provisioningClient insys.ProvisioningClient
	zapierClient       ZapierClient
}

func NewLogInEventCreatedSubscriber(
	ctx context.Context,
	authclient AuthClient,
	featureFlagsClient FeatureFlagsClient,
	provisioningClient insys.ProvisioningClient,
	zapierClient ZapierClient,
) *LogInEventCreatedSubscriber {
	return &LogInEventCreatedSubscriber{
		authClient:         authclient,
		featureFlagsClient: featureFlagsClient,
		provisioningClient: provisioningClient,
		zapierClient:       zapierClient,
	}
}

func (s LogInEventCreatedSubscriber) HandleMessage(ctx context.Context, m *nsq.Message) error {
	var le clientproto.LoginEvent

	err := proto.Unmarshal(m.Body, &le)
	if err != nil {
		wlog.WErrorC(ctx, werror.Wrap(err, "could not unmarshal LoginEvent message body into proto for clientproto.LoginEvent struct"))
		return nil
	}

	return s.processLoginEventMessage(ctx, le)
}

func (s LogInEventCreatedSubscriber) processLoginEventMessage(ctx context.Context, event clientproto.LoginEvent) error {
	userUUID, err := event.UserID.UUID()
	if err != nil {
		wlog.WErrorC(ctx, werror.Wrap(err, "could not unmarshal LoginEvent User UUID:").Add("userID", event.UserID.String()))
		return nil
	}

	userAccess, err := s.authClient.UserLocations(ctx, userUUID)
	if err != nil {
		return werror.Wrap(err, "could not get userAccess by ID").Add("userID", userUUID)
	}

	// don't capture login for non-practice user
	if userAccess.Type != authclient.UserTypePractice {
		return nil
	}

	var locationIDs []uuid.UUID

	for _, location := range userAccess.Locations {
		locationIDs = append(locationIDs, location.LocationID)
	}

	wlog.InfoC(ctx, fmt.Sprintf("received login event message for user %v. locations: %v", userUUID, locationIDs))

	preprovisions, err := s.getPreprovisionsByLocationIDs(ctx, locationIDs)
	if err != nil {
		wlog.WErrorC(ctx, werror.Wrap(err, fmt.Sprintf("could not fetch preprovisions with location ids: %+v", locationIDs)))
		return nil
	}
	if len(preprovisions) == 0 {
		wlog.WErrorC(ctx, werror.New(fmt.Sprintf("no preprovisions with location ids: %+v", locationIDs)))
		return nil
	}
	preprovisionsWithoutFirstLogin, err := s.filterPreprovisionsToThoseWithoutFirstLogin(ctx, preprovisions)

	if len(preprovisionsWithoutFirstLogin) == 0 {
		wlog.InfoC(ctx, fmt.Sprintf("no preprovisions with without first login for locations ids %v", locationIDs))
		return nil
	}

	for _, preprovision := range preprovisionsWithoutFirstLogin {
		zapierSendErr := s.zapierClient.Send(ctx, userAccess.Username, preprovision.LocationId, preprovision.SalesforceOpportunityId)
		if zapierSendErr != nil {
			return werror.Wrap(err, "failed to fire off zapier call").Add("Username", userAccess.Username).Add("locationId", preprovision.LocationId).Add("SalesforceOpportunityId", preprovision.SalesforceOpportunityId)
		}
		updatePreProvErr := s.setUserFirstLoggedInAtOnPreProvisionRecords(ctx, preprovision.LocationId)
		if updatePreProvErr != nil {
			wlog.InfoC(ctx, fmt.Sprintf("failed to log user first logged in at in preprovisioning with location ID: %s. Error Message: %v", preprovision.LocationId, err))
		}
	}

	return nil
}

func (s LogInEventCreatedSubscriber) setUserFirstLoggedInAtOnPreProvisionRecords(ctx context.Context, locationID string) error {
	preprovisionResponse, err := s.provisioningClient.PreProvisionsByLocationID(ctx, &insysproto.PreProvisionsByLocationIDRequest{LocationId: locationID})
	if err != nil {
		return fmt.Errorf("failed to fetch preprovision for location with ID %s from provisioning service. Error Message: %v", locationID, err)
	}

	if len(preprovisionResponse.PreProvisions) == 0 {
		wlog.WErrorC(ctx, werror.New(fmt.Sprintf("no preprovisions with location id: %+v", locationID)))
		return nil
	}

	for _, preprovision := range preprovisionResponse.PreProvisions {
		preprovision.UserFirstLoggedInAt = time.Now().Format(time.RFC3339)
		_, err = s.provisioningClient.CreateOrUpdatePreProvision(ctx, &insysproto.CreateOrUpdatePreProvisionRequest{PreProvision: preprovision})
		if err != nil {
			return fmt.Errorf("failed to update preprovision user_first_logged_in_at for location with ID %s from provisioning service. Error Message: %v", locationID, err)
		}
	}

	return nil
}

func (s LogInEventCreatedSubscriber) getPreprovisionsByLocationIDs(ctx context.Context, locationIDs []uuid.UUID) ([]insysproto.PreProvision, error) {
	var preprovisions []insysproto.PreProvision

	for _, locationID := range locationIDs {
		preprovisionResponse, err := s.provisioningClient.PreProvisionsByLocationID(ctx, &insysproto.PreProvisionsByLocationIDRequest{LocationId: locationID.String()})
		if err != nil {
			if werror.HasCode(err, wgrpc.CodeNotFound) {
				wlog.InfoC(ctx, fmt.Sprintf("no location with id: %s", locationID))
				continue
			} else {
				return nil, werror.Wrap(err, "could not read location for location by id ").Add("locationID", locationID)
			}
		}
		for _, preprovision := range preprovisionResponse.PreProvisions {
			preprovisions = append(preprovisions, *preprovision)
		}
	}

	return preprovisions, nil
}

func (s LogInEventCreatedSubscriber) filterPreprovisionsToThoseWithoutFirstLogin(ctx context.Context, preprovisions []insysproto.PreProvision) ([]insysproto.PreProvision, error) {
	var preprovisionsWithoutFirstLogin []insysproto.PreProvision

	for _, preprovision := range preprovisions {
		if preprovision.UserFirstLoggedInAt == "" {
			preprovisionsWithoutFirstLogin = append(preprovisionsWithoutFirstLogin, preprovision)
		}
	}

	return preprovisionsWithoutFirstLogin, nil
}
