package consumers

import (
	"context"
	"fmt"

	"github.com/gogo/protobuf/proto"
	nsq "github.com/nsqio/go-nsq"
	"weavelab.xyz/insys-onboarding-service/internal/app"
	"weavelab.xyz/monorail/shared/grpc-clients/client-grpc-clients/authclient"
	"weavelab.xyz/monorail/shared/protorepo/dist/go/messages/client/clientproto"
	"weavelab.xyz/monorail/shared/wlib/uuid"
	"weavelab.xyz/monorail/shared/wlib/werror"
	"weavelab.xyz/monorail/shared/wlib/wlog"
)

type LogInEventCreatedSubscriber struct {
	onboardersLocationService app.OnboardersLocationService
	authClient                app.AuthClient
	featureFlagsClient        app.FeatureFlagsClient
	zapierClient              app.ZapierClient
}

func NewLogInEventCreatedSubscriber(ctx context.Context, onboardersLocationService app.OnboardersLocationService, authclient app.AuthClient, featureFlagsClient app.FeatureFlagsClient, zapierClient app.ZapierClient) *LogInEventCreatedSubscriber {
	return &LogInEventCreatedSubscriber{
		onboardersLocationService: onboardersLocationService,
		authClient:                authclient,
		featureFlagsClient:        featureFlagsClient,
		zapierClient:              zapierClient,
	}
}

func (s LogInEventCreatedSubscriber) HandleMessage(ctx context.Context, m *nsq.Message) error {
	var le *clientproto.LoginEvent

	err := proto.Unmarshal(m.Body, le)
	if err != nil {
		return werror.Wrap(err, "could not unmarshal LoginEvent message body into proto for clientproto.LoginEvent struct")
	}

	return s.processLoginEventMessage(ctx, le)
}

func (s LogInEventCreatedSubscriber) processLoginEventMessage(ctx context.Context, event *clientproto.LoginEvent) error {
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

	var locations []uuid.UUID
	var locationsWithoutFirstLogin []uuid.UUID

	for _, location := range userAccess.Locations {
		locations = append(locations, location.LocationID)
	}

	for i := 0; i < len(locations); i++ {
		location, err := s.onboardersLocationService.ReadByLocationID(ctx, locations[i])
		if err != nil {
			return werror.Wrap(err, "could not read location for location by id ").Add("locationID", locations[i].String())
		}

		if !location.UserFirstLoggedInAt.Valid {
			locationsWithoutFirstLogin = append(locationsWithoutFirstLogin, location.LocationID)
		}
	}

	// exit if there are no locations that have not already been logged in to
	if len(locationsWithoutFirstLogin) == 0 {
		return nil
	}

	for _, locationID := range locationsWithoutFirstLogin {
		features, err := s.featureFlagsClient.List(ctx, locationID)
		if err != nil {
			wlog.InfoC(ctx, fmt.Sprintf("failed to get features for location with ID: %s. Error Message: %v", locationID.String(), err))
			continue
		}

		for _, feature := range features {
			if feature.Name == "onboardingBetaEnabled" && feature.Value == true {
				// make call to zapier, and if zapier succeeds, update the database.
				// if not, fail silently as the user is sure to log in again.
				err = s.zapierClient.Send(ctx, userAccess.Username, locationID.String())
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
