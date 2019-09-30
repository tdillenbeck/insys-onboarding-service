package consumers

import (
	"context"
	"fmt"

	"github.com/gogo/protobuf/proto"
	nsq "github.com/nsqio/go-nsq"
	"weavelab.xyz/insys-onboarding-service/internal/app"
	"weavelab.xyz/insys-onboarding-service/internal/zapier"
	"weavelab.xyz/monorail/shared/grpc-clients/client-grpc-clients/authclient"
	"weavelab.xyz/monorail/shared/grpc-clients/client-grpc-clients/featureflagsclient"
	"weavelab.xyz/monorail/shared/protorepo/dist/go/messages/client/clientproto"
	"weavelab.xyz/monorail/shared/wlib/uuid"
	"weavelab.xyz/monorail/shared/wlib/werror"
	"weavelab.xyz/monorail/shared/wlib/wlog"
)

type LogInEventCreatedSubscriber struct {
	onboardersLocationService app.OnboardersLocationService
	authClient                authclient.Auth
	featureFlagsClient        featureflagsclient.Client
	zapierClient              *zapier.ZapierClient
}

func NewLogInEventCreatedSubscriber(ctx context.Context, onboardersLocationService app.OnboardersLocationService, authclient authclient.Auth, featureFlagsClient featureflagsclient.Client, zapierClient *zapier.ZapierClient) *LogInEventCreatedSubscriber {
	return &LogInEventCreatedSubscriber{
		onboardersLocationService: onboardersLocationService,
		authClient:                authclient,
		featureFlagsClient:        featureFlagsClient,
		zapierClient:              zapierClient,
	}
}

func (p LogInEventCreatedSubscriber) HandleMessage(ctx context.Context, m *nsq.Message) error {
	var le clientproto.LoginEvent

	err := proto.Unmarshal(m.Body, &le)
	if err != nil {
		return werror.Wrap(err, "could not unmarshal LoginEvent message body into proto for clientproto.LoginEvent struct")
	}

	userUUID, err := le.UserID.UUID()
	if err != nil {
		return werror.Wrap(err, "could not unmarshal LoginEvent User UUID")
	}

	userAccess, err := p.authClient.UserLocations(ctx, userUUID)
	if err != nil {
		return werror.Wrap(err, "could not get userAccess for user with id: "+userUUID.String())
	}

	// don't capture login for non-practice user
	if userAccess.Type != authclient.UserTypePractice {
		return nil
	}

	var locations []uuid.UUID

	for _, location := range userAccess.Locations {
		locations = append(locations, location.LocationID)
	}

	hasLocationsWithoutLoginRecorded, err := p.onboardersLocationService.HasLocationsWithoutLoginRecorded(ctx, locations)

	if err != nil {
		return werror.Wrap(err, "could not get hasLocationsWithoutLoginRecorded for user with id: "+userUUID.String())
	}

	// exit if there are no locations that have not already been logged in to
	if !hasLocationsWithoutLoginRecorded {
		return nil
	}

	for _, locationID := range locations {
		features, err := p.featureFlagsClient.List(ctx, locationID)
		if err != nil {
			wlog.InfoC(ctx, fmt.Sprintf("failed to get features for location with ID: %s. Error Message: %v", locationID.String(), err))
		}

		// hasOnboardingBetaEnabled is active only for those locations being onboarded, and so it's the indicator that we use.
		hasOnboardingBetaEnabled := false
		for _, feature := range features {
			if feature.Name == "onboardingBetaEnabled" && feature.Value == true {
				hasOnboardingBetaEnabled = true
				break
			}
		}

		if !hasOnboardingBetaEnabled {
			return nil
		}

		// make call to zapier, and if zapier succeeds, update the database.
		// if not, fail silently as the user is sure to log in again.
		err = p.zapierClient.Send(ctx, userAccess.Username, locationID.String())
		if err != nil {
			wlog.InfoC(ctx, fmt.Sprintf("failed to fire off zapier call to mark Opportunity as `Closed-Won` for location with ID: %s. Error Message: %v", locationID.String(), err))
			continue
		}
		err = p.onboardersLocationService.RecordFirstLogin(ctx, locationID)
		if err != nil {
			wlog.InfoC(ctx, fmt.Sprintf("failed to record first login for location with ID: %s. Error Message: %v", locationID.String(), err))
		}
	}

	return nil
}
