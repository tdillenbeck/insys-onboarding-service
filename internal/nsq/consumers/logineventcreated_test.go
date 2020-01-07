package consumers

import (
	"context"
	"reflect"
	"testing"
	"time"

	"google.golang.org/grpc"
	"weavelab.xyz/insys-onboarding-service/internal/app"
	"weavelab.xyz/insys-onboarding-service/internal/mock"
	"weavelab.xyz/monorail/shared/go-utilities/null"
	"weavelab.xyz/monorail/shared/grpc-clients/client-grpc-clients/authclient"
	"weavelab.xyz/monorail/shared/grpc-clients/client-grpc-clients/featureflagsclient"
	"weavelab.xyz/monorail/shared/protorepo/dist/go/messages/client/clientproto"
	"weavelab.xyz/monorail/shared/protorepo/dist/go/messages/insysproto"
	"weavelab.xyz/monorail/shared/protorepo/dist/go/messages/sharedproto"
	"weavelab.xyz/monorail/shared/protorepo/dist/go/services/insys"
	"weavelab.xyz/monorail/shared/wlib/uuid"
	"weavelab.xyz/monorail/shared/wlib/werror"
)

/*
func TestLogInEventCreatedSubscriber_processLoginEventMessage(t *testing.T) {

	locationWithPreviousLoginA := uuid.NewV4()
	locationWithPreviousLoginB := uuid.NewV4()
	locationWithNoLoginsA := uuid.NewV4()
	nonExistantLocationID := uuid.NewV4()

	onboarderA := uuid.NewV4()
	onboarderB := uuid.NewV4()
	onboarderC := uuid.NewV4()

	userID := sharedproto.UUIDToProto(uuid.NewV4())

	mockOnboardersLocationService := mock.OnboarderLocationService{
		ReadByLocationIDFn: func(ctx context.Context, locationID uuid.UUID) (*app.OnboardersLocation, error) {
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
		},
		RecordFirstLoginFn: func(ctx context.Context, locationID uuid.UUID) error {
			return nil
		},
	}

	mockAuthClient := mock.Auth{
		UserLocationsFn: func(ctx context.Context, userID uuid.UUID) (*authclient.UserAccess, error) {
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
		},
	}

	mockFeatureFlagClient := mock.FeatureFlagsClient{
		ListFn: func(ctx context.Context, locationID uuid.UUID) ([]featureflagsclient.Flag, error) {
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
		},
	}

	mockProvisioningService := mock.ProvisioningService{
		// provide two preprovisions with varying dates to ensure that the function only uses the most recent one
		PreProvisionsByLocationIDFn: func(ctx context.Context, req *insysproto.PreProvisionsByLocationIDRequest, opts []grpc.CallOption) (*insysproto.PreProvisionsByLocationIDResponse, error) {
			return &insysproto.PreProvisionsByLocationIDResponse{
				PreProvisions: []*insysproto.PreProvision{
					&insysproto.PreProvision{
						SalesforceOpportunityId: "older opportunityID",
						UpdatedAt:               time.Now().Add(time.Hour * -10).String(),
					},
					&insysproto.PreProvision{
						SalesforceOpportunityId: "opportunityID",
						UpdatedAt:               time.Now().String(),
					},
				},
			}, nil
		},
	}

	mockZapierClient := mock.ZapierClient{
		SendFn: func(ctx context.Context, username, locationID, salesforceOpportunityID string) error {
			return nil
		},
	}

	type fields struct {
		authClient                AuthClient
		featureFlagsClient        FeatureFlagsClient
		onboardersLocationService app.OnboardersLocationService
		provisioningService       insys.ProvisioningClient
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
				provisioningService:       &mockProvisioningService,
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
				authClient:                tt.fields.authClient,
				featureFlagsClient:        tt.fields.featureFlagsClient,
				onboardersLocationService: tt.fields.onboardersLocationService,
				provisioningClient:        tt.fields.provisioningService,
				zapierClient:              tt.fields.zapierClient,
			}
			if err := p.processLoginEventMessage(tt.args.ctx, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("LogInEventCreatedSubscriber.processLoginEventMessage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
*/

func TestLogInEventCreatedSubscriber_processLoginEventMessage(t *testing.T) {

	locationWithPreviousLoginA := uuid.NewV4()
	locationWithPreviousLoginB := uuid.NewV4()
	locationWithNoLoginsA := uuid.NewV4()
	nonExistantLocationID := uuid.NewV4()

	onboarderA := uuid.NewV4()
	onboarderB := uuid.NewV4()
	onboarderC := uuid.NewV4()

	userID := sharedproto.UUIDToProto(uuid.NewV4())

	mockOnboardersLocationService := &mock.OnboarderLocationService{
		ReadByLocationIDFn: func(ctx context.Context, locationID uuid.UUID) (*app.OnboardersLocation, error) {
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
		},
		RecordFirstLoginFn: func(ctx context.Context, locationID uuid.UUID) error {
			return nil
		},
	}

	mockAuthClient := &mock.Auth{
		UserLocationsFn: func(ctx context.Context, userID uuid.UUID) (*authclient.UserAccess, error) {
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
		},
	}

	mockFeatureFlagClient := &mock.FeatureFlagsClient{
		ListFn: func(ctx context.Context, locationID uuid.UUID) ([]featureflagsclient.Flag, error) {
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
		},
	}

	mockProvisioningService := &mock.ProvisioningService{
		// provide two preprovisions with varying dates to ensure that the function only uses the most recent one
		PreProvisionsByLocationIDFn: func(ctx context.Context, req *insysproto.PreProvisionsByLocationIDRequest, opts []grpc.CallOption) (*insysproto.PreProvisionsByLocationIDResponse, error) {
			return &insysproto.PreProvisionsByLocationIDResponse{
				PreProvisions: []*insysproto.PreProvision{
					&insysproto.PreProvision{
						SalesforceOpportunityId: "older opportunityID",
						UpdatedAt:               time.Now().Add(time.Hour * -10).String(),
					},
					&insysproto.PreProvision{
						SalesforceOpportunityId: "opportunityID",
						UpdatedAt:               time.Now().String(),
					},
				},
			}, nil
		},
	}

	mockZapierClient := &mock.ZapierClient{
		SendFn: func(ctx context.Context, username, locationID, salesforceOpportunityID string) error {
			return nil
		},
	}

	type fields struct {
		authClient                AuthClient
		featureFlagsClient        FeatureFlagsClient
		onboardersLocationService app.OnboardersLocationService
		provisioningClient        insys.ProvisioningClient
		zapierClient              ZapierClient
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
			name: "run successfully",
			fields: fields{
				authClient:                mockAuthClient,
				featureFlagsClient:        mockFeatureFlagClient,
				onboardersLocationService: mockOnboardersLocationService,
				provisioningClient:        mockProvisioningService,
				zapierClient:              mockZapierClient,
			},
			args: args{
				ctx: context.Background(),
				event: clientproto.LoginEvent{
					UserID: userID,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := LogInEventCreatedSubscriber{
				authClient:                tt.fields.authClient,
				featureFlagsClient:        tt.fields.featureFlagsClient,
				onboardersLocationService: tt.fields.onboardersLocationService,
				provisioningClient:        tt.fields.provisioningClient,
				zapierClient:              tt.fields.zapierClient,
			}
			if err := s.processLoginEventMessage(tt.args.ctx, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("LogInEventCreatedSubscriber.processLoginEventMessage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLogInEventCreatedSubscriber_filterLocationsToThoseWithoutFirstLoginForUser(t *testing.T) {
	type fields struct {
		authClient                *authclient.Auth
		featureFlagsClient        *featureflagsclient.Client
		onboardersLocationService app.OnboardersLocationService
		provisioningClient        insys.ProvisioningClient
		zapierClient              ZapierClient
	}
	type args struct {
		ctx         context.Context
		locationIDs []uuid.UUID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []uuid.UUID
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := LogInEventCreatedSubscriber{
				authClient:                tt.fields.authClient,
				featureFlagsClient:        tt.fields.featureFlagsClient,
				onboardersLocationService: tt.fields.onboardersLocationService,
				provisioningClient:        tt.fields.provisioningClient,
				zapierClient:              tt.fields.zapierClient,
			}
			got, err := s.filterLocationsToThoseWithoutFirstLoginForUser(tt.args.ctx, tt.args.locationIDs)
			if (err != nil) != tt.wantErr {
				t.Errorf("LogInEventCreatedSubscriber.filterLocationsToThoseWithoutFirstLoginForUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LogInEventCreatedSubscriber.filterLocationsToThoseWithoutFirstLoginForUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLogInEventCreatedSubscriber_getMostRecentOpportunityIDForLocations(t *testing.T) {
	type fields struct {
		authClient                *authclient.Auth
		featureFlagsClient        *featureflagsclient.Client
		onboardersLocationService app.OnboardersLocationService
		provisioningClient        insys.ProvisioningClient
		zapierClient              ZapierClient
	}
	type args struct {
		ctx         context.Context
		locationIDs []uuid.UUID
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := LogInEventCreatedSubscriber{
				authClient:                tt.fields.authClient,
				featureFlagsClient:        tt.fields.featureFlagsClient,
				onboardersLocationService: tt.fields.onboardersLocationService,
				provisioningClient:        tt.fields.provisioningClient,
				zapierClient:              tt.fields.zapierClient,
			}
			if got := s.getMostRecentOpportunityIDForLocations(tt.args.ctx, tt.args.locationIDs); got != tt.want {
				t.Errorf("LogInEventCreatedSubscriber.getMostRecentOpportunityIDForLocations() = %v, want %v", got, tt.want)
			}
		})
	}
}
