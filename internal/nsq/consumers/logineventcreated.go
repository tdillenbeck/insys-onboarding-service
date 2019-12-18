package consumers

import (
	"context"
	"fmt"
	"sort"

	"github.com/gogo/protobuf/proto"
	nsq "github.com/nsqio/go-nsq"
	"weavelab.xyz/insys-onboarding-service/internal/app"
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

type LogInEventCreatedSubscriber struct {
	authClient                AuthClient
	featureFlagsClient        FeatureFlagsClient
	onboardersLocationService app.OnboardersLocationService
	provisioningClient        insys.ProvisioningClient
	zapierClient              app.ZapierClient
}

func NewLogInEventCreatedSubscriber(
	ctx context.Context,
	authclient AuthClient,
	featureFlagsClient FeatureFlagsClient,
	onboardersLocationService app.OnboardersLocationService,
	provisioningClient insys.ProvisioningClient,
	zapierClient app.ZapierClient,
) *LogInEventCreatedSubscriber {

	return &LogInEventCreatedSubscriber{
		authClient:                authclient,
		featureFlagsClient:        featureFlagsClient,
		onboardersLocationService: onboardersLocationService,
		provisioningClient:        provisioningClient,
		zapierClient:              zapierClient,
	}
}

func (s LogInEventCreatedSubscriber) HandleMessage(ctx context.Context, m *nsq.Message) error {
	var le clientproto.LoginEvent

	fmt.Println(string(m.Body))

	err := proto.Unmarshal(m.Body, &le)
	if err != nil {
		wlog.ErrorC(ctx, "could not unmarshal LoginEvent message body into proto for clientproto.LoginEvent struct")
		return nil
	}

	return s.processLoginEventMessage(ctx, le)
}

func (s LogInEventCreatedSubscriber) processLoginEventMessage(ctx context.Context, event clientproto.LoginEvent) error {
	userUUID, err := event.UserID.UUID()
	if err != nil {
		return werror.Wrap(err, "could not unmarshal LoginEvent User UUID").Add("UserID", event.UserID)
	}
	userAccess, err := s.authClient.UserLocations(ctx, userUUID)
	if err != nil {
		return werror.Wrap(err, "could not get userAccess by ID").Add("userID", userUUID.String())
	}

	// don't capture login for non-practice user
	if userAccess.Type != authclient.UserTypePractice {
		return nil
	}

	var locationsWithoutFirstLogin []uuid.UUID
	var salesforceOpportunityID string

	for i := 0; i < len(userAccess.Locations); i++ {
		location, err := s.onboardersLocationService.ReadByLocationID(ctx, userAccess.Locations[i].LocationID)
		if err != nil {
			if werror.HasCode(err, wgrpc.CodeNotFound) {
				wlog.InfoC(ctx, fmt.Sprintf("no location with id: %s", userAccess.Locations[i].LocationID.String()))
				continue
			} else {
				return werror.Wrap(err, "could not read location for location by id ").Add("locationID", userAccess.Locations[i].LocationID.String())
			}
		}

		if !location.UserFirstLoggedInAt.Valid {
			locationsWithoutFirstLogin = append(locationsWithoutFirstLogin, location.LocationID)
		}

		// We need an opportunityID to send to Zapier, but we aren't sure which locationID is associated with the Opportunity in Salesforce, so iterate
		// over all, logged in or not, in search for the preprovision, select the most recent preprovision, and pull the opportunity_id.
		provisionResponse, err := s.provisioningClient.PreProvisionsByLocationID(ctx, &insysproto.PreProvisionsByLocationIDRequest{LocationId: location.LocationID.String()})
		if err != nil {
			wlog.InfoC(ctx, fmt.Sprintf("failed to get preprovisions for location with id: %s. error message: %v", location.LocationID.String(), err))
		}

		if provisionResponse != nil && len(provisionResponse.PreProvisions) > 0 {
			pps := sortPreProvisionsByUpdatedDate(provisionResponse.PreProvisions)
			if pps[0].SalesforceOpportunityId != "" {
				salesforceOpportunityID = pps[0].SalesforceOpportunityId
			}
		} else {
			wlog.InfoC(ctx, fmt.Sprintf("no preprovisions for location with id: %s", location.LocationID.String()))
		}

		if salesforceOpportunityID == "" {
			wlog.InfoC(ctx, fmt.Sprintf("no opportunity id for location with id: %s", location.LocationID.String()))
		}
	}

	// exit if there are no locations that have not already been logged in to
	if len(locationsWithoutFirstLogin) == 0 {
		return nil
	}

	for _, locationID := range locationsWithoutFirstLogin {
		features, err := s.featureFlagsClient.List(ctx, locationID)
		if err != nil {
			wlog.InfoC(ctx, fmt.Sprintf("failed to get features for location with id: %s. error message: %v", locationID.String(), err))
			continue
		}

		// we need to ensure that the location is in the onboarding process, so loop through in search of the feature that indicates that it is
		for _, feature := range features {
			if feature.Name == "onboardingBetaEnabled" && feature.Value == true {
				// make call to zapier, and if zapier succeeds, update the database.
				// if not, fail silently as the user is sure to log in again.
				err = s.zapierClient.Send(ctx, userAccess.Username, locationID.String(), salesforceOpportunityID)
				if err != nil {
					wlog.InfoC(ctx, fmt.Sprintf("failed to fire off zapier call to mark Opportunity as `Closed-Won` for location with ID: %s. Error Message: %v", locationID.String(), err))
					continue
				}
				err = s.onboardersLocationService.RecordFirstLogin(ctx, locationID)
				if err != nil {
					wlog.InfoC(ctx, fmt.Sprintf("failed to record first login for location with ID: %s. Error Message: %v", locationID.String(), err))
					continue
				}
			}
		}
	}
	return nil
}

func sortPreProvisionsByUpdatedDate(pps []*insysproto.PreProvision) []*insysproto.PreProvision {
	result := pps
	// only send the most recent one, so sort by updated date
	sort.Slice(result, func(i, j int) bool {
		return result[i].UpdatedAt > result[j].UpdatedAt
	})
	return result
}
