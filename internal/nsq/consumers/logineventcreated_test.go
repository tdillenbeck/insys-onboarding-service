package consumers

import (
	"context"
	"testing"
	"time"

	"weavelab.xyz/insys-onboarding-service/internal/app"
	"weavelab.xyz/insys-onboarding-service/internal/mock"
	"weavelab.xyz/monorail/shared/go-utilities/null"
	"weavelab.xyz/monorail/shared/grpc-clients/client-grpc-clients/authclient"
	"weavelab.xyz/monorail/shared/grpc-clients/client-grpc-clients/featureflagsclient"
	"weavelab.xyz/monorail/shared/protorepo/dist/go/messages/client/clientproto"
	"weavelab.xyz/monorail/shared/protorepo/dist/go/messages/sharedproto"
	"weavelab.xyz/monorail/shared/wlib/uuid"
	"weavelab.xyz/monorail/shared/wlib/werror"
)

func TestLogInEventCreatedSubscriber_processLoginEventMessage(t *testing.T) {

	locationWithPreviousLoginA := uuid.NewV4()
	locationWithPreviousLoginB := uuid.NewV4()
	locationWithNoLoginsA := uuid.NewV4()
	nonExistantLocationID := uuid.NewV4()

	onboarderA := uuid.NewV4()
	onboarderB := uuid.NewV4()
	onboarderC := uuid.NewV4()

	userID := sharedproto.UUIDToProto(uuid.NewV4())

	mockOnboardersLocationService := mock.OnboarderLocationService{}

	mockOnboardersLocationService.ReadByLocationIDFn = func(ctx context.Context, locationID uuid.UUID) (*app.OnboardersLocation, error) {

		switch locationID {
		case locationWithNoLoginsA:
			{
				return &app.OnboardersLocation{
					LocationID:          locationWithNoLoginsA,
					OnboarderID:         onboarderA,
					UserFirstLoggedInAt: null.Time{},
				}, nil
			}
		case locationWithPreviousLoginB:
			{
				return &app.OnboardersLocation{
					LocationID:          locationWithPreviousLoginB,
					OnboarderID:         onboarderB,
					UserFirstLoggedInAt: null.NewTime(time.Now()),
				}, nil
			}
		case locationWithPreviousLoginA:
			{
				return &app.OnboardersLocation{
					LocationID:          locationWithPreviousLoginA,
					OnboarderID:         onboarderC,
					UserFirstLoggedInAt: null.NewTime(time.Now()),
				}, nil
			}

		case nonExistantLocationID:
			{
				return nil, werror.New("Location Not found")
			}
		}

		return nil, nil
	}

	mockOnboardersLocationService.RecordFirstLoginFn = func(ctx context.Context, locationID uuid.UUID) error {
		return nil
	}

	mockAuthClient := mock.Auth{}

	mockAuthClient.UserLocationsFn = func(ctx context.Context, userID uuid.UUID) (*authclient.UserAccess, error) {
		return &authclient.UserAccess{
			FirstName: "Jack",
			LastName:  "Frost",
			Type:      authclient.UserTypePractice,
			Locations: []authclient.Location{
				authclient.Location{
					LocationID: locationWithPreviousLoginA,
				},
				authclient.Location{
					LocationID: locationWithPreviousLoginB,
				},
				authclient.Location{
					LocationID: locationWithNoLoginsA,
				},
			},
		}, nil
	}

	mockFeatureFlagClient := mock.FeatureFlagsClient{}
	mockFeatureFlagClient.ListFn = func(ctx context.Context, locationID uuid.UUID) ([]featureflagsclient.Flag, error) {
		return []featureflagsclient.Flag{
			featureflagsclient.Flag{
				Name:  "onboardingBetaEnabled",
				Value: true,
			},
			featureflagsclient.Flag{
				Name:  "otherflag",
				Value: false,
			},
			featureflagsclient.Flag{
				Name:  "anotherflag",
				Value: true,
			},
		}, nil
	}

	mockZapierClient := mock.ZapierClient{}
	mockZapierClient.SendFn = func(ctx context.Context, username, locationID string) error {
		return nil
	}

	type fields struct {
		onboardersLocationService app.OnboardersLocationService
		authClient                app.AuthClient
		featureFlagsClient        app.FeatureFlagsClient
		zapierClient              app.ZapierClient
	}
	type args struct {
		ctx   context.Context
		event clientproto.LoginEvent
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Ensure login recorded for location",
			fields: fields{
				authClient:                &mockAuthClient,
				featureFlagsClient:        &mockFeatureFlagClient,
				onboardersLocationService: &mockOnboardersLocationService,
				zapierClient:              &mockZapierClient,
			},
			args: args{
				ctx: context.Background(),
				event: clientproto.LoginEvent{
					UserID: userID,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := LogInEventCreatedSubscriber{
				onboardersLocationService: tt.fields.onboardersLocationService,
				authClient:                tt.fields.authClient,
				featureFlagsClient:        tt.fields.featureFlagsClient,
				zapierClient:              tt.fields.zapierClient,
			}
			if err := p.processLoginEventMessage(tt.args.ctx, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("LogInEventCreatedSubscriber.processLoginEventMessage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
