package consumers

import (
	"context"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"

	"google.golang.org/grpc"
	"weavelab.xyz/insys-onboarding-service/internal/app"
	"weavelab.xyz/insys-onboarding-service/internal/mock"
	"weavelab.xyz/insys-onboarding-service/internal/zapier"
	"weavelab.xyz/monorail/shared/go-utilities/null"
	"weavelab.xyz/monorail/shared/grpc-clients/client-grpc-clients/authclient"
	"weavelab.xyz/monorail/shared/protorepo/dist/go/messages/client/clientproto"
	"weavelab.xyz/monorail/shared/protorepo/dist/go/messages/insysproto"
	"weavelab.xyz/monorail/shared/protorepo/dist/go/messages/sharedproto"
	"weavelab.xyz/monorail/shared/protorepo/dist/go/services/insys"
	"weavelab.xyz/monorail/shared/wlib/uuid"
	"weavelab.xyz/monorail/shared/wlib/werror"
)

func TestLogInEventCreatedSubscriber_processLoginEventMessage(t *testing.T) {
	locationWithPreviousLoginA := uuid.NewV4()
	locationWithPreviousLoginB := uuid.NewV4()
	locationWithPreviousLoginC := uuid.NewV4()
	locationWithNoLoginsA := uuid.NewV4()
	locationWithNoLoginsB := uuid.NewV4()
	nonExistantLocationID := uuid.NewV4()

	onboarderA := uuid.NewV4()
	onboarderB := uuid.NewV4()
	onboarderC := uuid.NewV4()

	userID := sharedproto.UUIDToProto(uuid.NewV4())

	type fields struct {
		authClient                AuthClient
		featureFlagsClient        FeatureFlagsClient
		onboardersLocationService app.OnboardersLocationService
		provisioningClient        insys.ProvisioningClient
		zapierClient              *mock.ZapierClient
	}
	type args struct {
		ctx   context.Context
		event clientproto.LoginEvent
	}
	tests := []struct {
		name                   string
		fields                 fields
		args                   args
		wantErr                bool
		wantZapierCalled       bool
		zapierPayloadsToBeSent []zapier.FirstLoginEventPayload
	}{
		{
			name: "when all the locations for a user have a login event already recorded don't send a message to Zapier",
			fields: fields{
				authClient: &mock.Auth{
					UserLocationsFn: func(ctx context.Context, userID uuid.UUID) (*authclient.UserAccess, error) {
						return &authclient.UserAccess{
							FirstName: "Jack",
							LastName:  "Frost",
							Username:  "JackFrost@gmail.com",
							Type:      authclient.UserTypePractice,
							Locations: []authclient.Location{
								{
									LocationID: locationWithPreviousLoginA,
								},
								{
									LocationID: locationWithPreviousLoginB,
								},
								{
									LocationID: locationWithPreviousLoginC,
								},
							},
						}, nil
					},
				},
				onboardersLocationService: &mock.OnboarderLocationService{
					ReadByLocationIDFn: func(ctx context.Context, locationID uuid.UUID) (*app.OnboardersLocation, error) {
						switch locationID {
						case locationWithPreviousLoginC:
							{
								return &app.OnboardersLocation{
									LocationID:          locationWithPreviousLoginC,
									OnboarderID:         onboarderA,
									UserFirstLoggedInAt: null.NewTime(time.Now()),
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
				},
				provisioningClient: &mock.ProvisioningService{
					PreProvisionsByLocationIDFn: func(ctx context.Context, req *insysproto.PreProvisionsByLocationIDRequest, opts []grpc.CallOption) (*insysproto.PreProvisionsByLocationIDResponse, error) {
						return &insysproto.PreProvisionsByLocationIDResponse{
							PreProvisions: []*insysproto.PreProvision{
								{
									SalesforceOpportunityId: "older opportunityID",
									UpdatedAt:               time.Now().Add(time.Hour * -10).String(),
								},
								{
									SalesforceOpportunityId: "opportunityID",
									UpdatedAt:               time.Now().String(),
								},
							},
						}, nil
					},
				},
				zapierClient: nil,
			},
			args: args{
				ctx: context.Background(),
				event: clientproto.LoginEvent{
					UserID: userID,
				},
			},
			wantZapierCalled: false,
		},
		{
			name: "when there are multiple locations without first login event associated with user, call zapier",
			fields: fields{
				authClient: &mock.Auth{
					UserLocationsFn: func(ctx context.Context, userID uuid.UUID) (*authclient.UserAccess, error) {
						return &authclient.UserAccess{
							FirstName: "Jack",
							LastName:  "Frost",
							Username:  "JackFrost@gmail.com",
							Type:      authclient.UserTypePractice,
							Locations: []authclient.Location{
								{
									LocationID: locationWithPreviousLoginA,
								},
								{
									LocationID: locationWithNoLoginsB,
								},
								{
									LocationID: locationWithNoLoginsA,
								},
							},
						}, nil
					},
				},
				onboardersLocationService: &mock.OnboarderLocationService{
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
						case locationWithNoLoginsB:
							{
								return &app.OnboardersLocation{
									LocationID:          locationWithNoLoginsB,
									OnboarderID:         onboarderB,
									UserFirstLoggedInAt: null.Time{},
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
				},
				provisioningClient: &mock.ProvisioningService{
					PreProvisionsByLocationIDFn: func(ctx context.Context, req *insysproto.PreProvisionsByLocationIDRequest, opts []grpc.CallOption) (*insysproto.PreProvisionsByLocationIDResponse, error) {
						return &insysproto.PreProvisionsByLocationIDResponse{
							PreProvisions: []*insysproto.PreProvision{
								{
									SalesforceOpportunityId: "older opportunityID",
									UpdatedAt:               time.Now().Add(time.Hour * -10).String(),
								},
								{
									SalesforceOpportunityId: "opportunityID",
									UpdatedAt:               time.Now().String(),
								},
							},
						}, nil
					},
				},
				zapierClient: &mock.ZapierClient{
					SendFn: func(ctx context.Context, username, locationID, salesforceOpportunityID string) error {
						return nil
					},
				},
			},
			args: args{
				ctx: context.Background(),
				event: clientproto.LoginEvent{
					UserID: userID,
				},
			},
			wantZapierCalled: true,
			zapierPayloadsToBeSent: []zapier.FirstLoginEventPayload{
				zapier.FirstLoginEventPayload{
					LocationID:              locationWithNoLoginsB.String(),
					SalesforceOpportunityID: "opportunityID",
					Username:                "JackFrost@gmail.com",
				},
				zapier.FirstLoginEventPayload{
					LocationID:              locationWithNoLoginsA.String(),
					SalesforceOpportunityID: "opportunityID",
					Username:                "JackFrost@gmail.com",
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
			if tt.wantZapierCalled && tt.fields.zapierClient.SendCalled != tt.wantZapierCalled {
				t.Error("LogInEventCreatedSubscriber.processLoginEventMessage() Zapier called erroneously")
			}
			if tt.wantZapierCalled && !cmp.Equal(tt.zapierPayloadsToBeSent, tt.fields.zapierClient.PayloadsSent) {
				t.Errorf("LogInEventCreatedSubscriber.processLoginEventMessage() unexpected zapier payloads. %v", cmp.Diff(tt.zapierPayloadsToBeSent, tt.fields.zapierClient.PayloadsSent))
			}
		})
	}
}
