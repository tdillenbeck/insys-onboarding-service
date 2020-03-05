package consumers

import (
	"context"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"google.golang.org/grpc"

	"weavelab.xyz/insys-onboarding-service/internal/mock"
	"weavelab.xyz/insys-onboarding-service/internal/zapier"
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
	locationWithNoLoginsA := uuid.NewV4()
	locationWithNoLoginsB := uuid.NewV4()
	nonExistantLocationID := uuid.NewV4()
	opportunityA := "Opportunity A"
	opportunityB := "Opportunity B"
	opportunityC := "Opportunity C"

	userID := sharedproto.UUIDToProto(uuid.NewV4())

	type fields struct {
		authClient         AuthClient
		featureFlagsClient FeatureFlagsClient
		provisioningClient insys.ProvisioningClient
		zapierClient       *mock.ZapierClient
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
							},
						}, nil
					},
				},
				provisioningClient: &mock.ProvisioningService{
					PreProvisionsByLocationIDFn: func(ctx context.Context, req *insysproto.PreProvisionsByLocationIDRequest, opts []grpc.CallOption) (*insysproto.PreProvisionsByLocationIDResponse, error) {
						switch req.LocationId {
						case locationWithNoLoginsA.String():
							{
								return &insysproto.PreProvisionsByLocationIDResponse{
									PreProvisions: []*insysproto.PreProvision{
										&insysproto.PreProvision{
											LocationId:              locationWithNoLoginsA.String(),
											SalesforceOpportunityId: opportunityA,
										},
									},
								}, nil
							}
						case locationWithPreviousLoginA.String():
							{
								return &insysproto.PreProvisionsByLocationIDResponse{
									PreProvisions: []*insysproto.PreProvision{
										&insysproto.PreProvision{
											LocationId:              locationWithPreviousLoginA.String(),
											SalesforceOpportunityId: opportunityC,
											UserFirstLoggedInAt:     time.Now().String(),
										},
									},
								}, nil
							}
						case locationWithPreviousLoginB.String():
							{
								return &insysproto.PreProvisionsByLocationIDResponse{
									PreProvisions: []*insysproto.PreProvision{
										&insysproto.PreProvision{
											LocationId:              locationWithPreviousLoginB.String(),
											SalesforceOpportunityId: opportunityB,
											UserFirstLoggedInAt:     time.Now().String(),
										},
									},
								}, nil
							}
						case nonExistantLocationID.String():
							{
								return nil, werror.New("Location Not found")
							}
						}
						return nil, nil
					},
					CreateOrUpdatePreProvisionFn: func(ctx context.Context, req *insysproto.CreateOrUpdatePreProvisionRequest, opts []grpc.CallOption) (*insysproto.CreateOrUpdatePreProvisionResponse, error) {
						return nil, nil
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
									LocationID: locationWithNoLoginsA,
								},
								{
									LocationID: locationWithNoLoginsB,
								},
							},
						}, nil
					},
				},
				provisioningClient: &mock.ProvisioningService{
					PreProvisionsByLocationIDFn: func(ctx context.Context, req *insysproto.PreProvisionsByLocationIDRequest, opts []grpc.CallOption) (*insysproto.PreProvisionsByLocationIDResponse, error) {
						switch req.LocationId {
						case locationWithNoLoginsA.String():
							{
								return &insysproto.PreProvisionsByLocationIDResponse{
									PreProvisions: []*insysproto.PreProvision{
										&insysproto.PreProvision{
											LocationId:              locationWithNoLoginsA.String(),
											SalesforceOpportunityId: opportunityA,
										},
									},
								}, nil
							}
						case locationWithPreviousLoginA.String():
							{
								return &insysproto.PreProvisionsByLocationIDResponse{
									PreProvisions: []*insysproto.PreProvision{
										&insysproto.PreProvision{
											LocationId:              locationWithPreviousLoginA.String(),
											SalesforceOpportunityId: opportunityC,
											UserFirstLoggedInAt:     time.Now().String(),
										},
									},
								}, nil
							}
						case locationWithNoLoginsB.String():
							{
								return &insysproto.PreProvisionsByLocationIDResponse{
									PreProvisions: []*insysproto.PreProvision{
										&insysproto.PreProvision{
											LocationId:              locationWithNoLoginsB.String(),
											SalesforceOpportunityId: opportunityB,
										},
									},
								}, nil
							}
						case nonExistantLocationID.String():
							{
								return nil, werror.New("Location Not found")
							}
						}
						return nil, nil
					},
					CreateOrUpdatePreProvisionFn: func(ctx context.Context, req *insysproto.CreateOrUpdatePreProvisionRequest, opts []grpc.CallOption) (*insysproto.CreateOrUpdatePreProvisionResponse, error) {
						return nil, nil
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
					LocationID:              locationWithNoLoginsA.String(),
					SalesforceOpportunityID: opportunityA,
					Username:                "JackFrost@gmail.com",
				},
				zapier.FirstLoginEventPayload{
					LocationID:              locationWithNoLoginsB.String(),
					SalesforceOpportunityID: opportunityB,
					Username:                "JackFrost@gmail.com",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := LogInEventCreatedSubscriber{
				authClient:         tt.fields.authClient,
				featureFlagsClient: tt.fields.featureFlagsClient,
				provisioningClient: tt.fields.provisioningClient,
				zapierClient:       tt.fields.zapierClient,
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
			if !tt.wantZapierCalled && len(tt.fields.zapierClient.PayloadsSent) != 0 {
				t.Errorf("LogInEventCreatedSubscriber.processLoginEventMessage() unexpected zapier payloads. %v", cmp.Diff(tt.zapierPayloadsToBeSent, tt.fields.zapierClient.PayloadsSent))
			}
		})
	}
}
