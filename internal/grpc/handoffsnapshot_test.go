package grpc

import (
	"context"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"testing"
	"time"
	"weavelab.xyz/insys-onboarding-service/internal/app"
	"weavelab.xyz/insys-onboarding-service/internal/mock"
	"weavelab.xyz/monorail/shared/go-utilities/null"
	"weavelab.xyz/monorail/shared/protorepo/dist/go/messages/insysproto"
	"weavelab.xyz/monorail/shared/wlib/uuid"
)

// TestHandoffSnapshotServer_SubmitCSAT, tests create, update, submit full cycle.  Note: order and timing matters in this test.
func TestHandoffSnapshotServer_HandoffCycle(t *testing.T) {

	userID := uuid.NewV4()
	onboardersLocationID := uuid.NewV4()

	pointOfContactEmail := null.NewString("client@example.com")
	reasonForPurchase := "reason"
	customizations := true
	customizationSetup := "notes about setup"
	faxPortSubmitted := "yes"
	routerType := "Red"
	disclaimerTypeSent := "email"
	routerMakeAndModel := "make and model"
	networkDecision := "notes about network"
	billingNotes := "notes about billing"
	notes := "general notes"

	testSnapshot := app.HandoffSnapshot{}

	type fields struct {
		handoffSnapshotService app.HandoffSnapshotService
	}
	type args struct {
		ctx                        context.Context
		req                        *insysproto.HandoffSnapshotCreateOrUpdateRequest
		handoffSnapshotReadRequest *insysproto.HandoffSnapshotReadRequest
		submitCsatParams           *insysproto.SubmitCSATRequest
		submitHandoffRequest       *insysproto.SubmitHandoffRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *insysproto.HandoffSnapshotResponse
		wantErr bool
	}{
		{
			name: "create and submit handoff successfully",
			fields: fields{
				handoffSnapshotService: &mock.HandoffSnapshotService{
					CreateOrUpdateFn: func(ctx context.Context, snapshot app.HandoffSnapshot) (app.HandoffSnapshot, error) {
						testSnapshot.ID = uuid.NewV4()
						testSnapshot.CreatedAt = time.Now()
						testSnapshot.UpdatedAt = time.Now()

						testSnapshot.OnboardersLocationID = snapshot.OnboardersLocationID
						testSnapshot.PointOfContactEmail = snapshot.PointOfContactEmail
						testSnapshot.ReasonForPurchase = snapshot.ReasonForPurchase
						testSnapshot.Customizations = snapshot.Customizations
						testSnapshot.CustomizationSetup = snapshot.CustomizationSetup
						testSnapshot.FaxPortSubmitted = snapshot.FaxPortSubmitted
						testSnapshot.RouterType = snapshot.RouterType
						testSnapshot.DisclaimerTypeSent = snapshot.DisclaimerTypeSent
						testSnapshot.RouterMakeAndModel = snapshot.RouterMakeAndModel
						testSnapshot.NetworkDecision = snapshot.NetworkDecision
						testSnapshot.BillingNotes = snapshot.BillingNotes
						testSnapshot.Notes = snapshot.Notes
						return testSnapshot, nil
					},
					ReadByOnboardersLocationIDFn: func(ctx context.Context, onboardersLocationId uuid.UUID) (app.HandoffSnapshot, error) {
						return testSnapshot, nil
					},
					SubmitCSATFn: func(ctx context.Context, onboardersLocationId uuid.UUID, csatRecipientUserEmail string) (app.HandoffSnapshot, error) {
						testSnapshot.CSATSentAt = null.NewTime(time.Now())
						testSnapshot.CsatRecipientUserEmail = null.NewString(csatRecipientUserEmail)
						return testSnapshot, nil
					},
					SubmitHandoffFn: func(ctx context.Context, onboardersLocationId uuid.UUID) (app.HandoffSnapshot, error) {
						testSnapshot.HandedOffAt = null.NewTime(time.Now())
						return testSnapshot, nil
					},
				},
			},
			args: args{
				ctx: context.Background(),
				req: &insysproto.HandoffSnapshotCreateOrUpdateRequest{
					HandoffSnapshot: &insysproto.HandoffSnapshotRecord{
						OnboardersLocationId: onboardersLocationID.String(),
						PointOfContactEmail:  pointOfContactEmail.String(),
						ReasonForPurchase:    reasonForPurchase,
						Customizations:       customizations,
						CustomizationSetup:   customizationSetup,
						FaxPortSubmitted:     faxPortSubmitted,
						RouterType:           routerType,
						DisclaimerTypeSent:   disclaimerTypeSent,
						RouterMakeAndModel:   routerMakeAndModel,
						NetworkDecision:      networkDecision,
						BillingNotes:         billingNotes,
						Notes:                notes,
					},
				},
				submitCsatParams: &insysproto.SubmitCSATRequest{
					OnboardersLocationId:   onboardersLocationID.String(),
					CsatRecipientUserEmail: userID.String(),
				},
				handoffSnapshotReadRequest: &insysproto.HandoffSnapshotReadRequest{
					OnboardersLocationId: onboardersLocationID.String(),
				},
				submitHandoffRequest: &insysproto.SubmitHandoffRequest{
					OnboardersLocationId: onboardersLocationID.String(),
				},
			},
			want: &insysproto.HandoffSnapshotResponse{
				HandoffSnapshot: &insysproto.HandoffSnapshotRecord{
					OnboardersLocationId: onboardersLocationID.String(),
					PointOfContactEmail:  pointOfContactEmail.String(),
					ReasonForPurchase:    reasonForPurchase,
					Customizations:       customizations,
					CustomizationSetup:   customizationSetup,
					FaxPortSubmitted:     faxPortSubmitted,
					RouterType:           routerType,
					DisclaimerTypeSent:   disclaimerTypeSent,
					RouterMakeAndModel:   routerMakeAndModel,
					NetworkDecision:      networkDecision,
					BillingNotes:         billingNotes,
					Notes:                notes,
				},
			},
			wantErr: false,
		},
		{
			name: "creates partial snapshot and fail to submit",
			fields: fields{
				handoffSnapshotService: &mock.HandoffSnapshotService{
					CreateOrUpdateFn: func(ctx context.Context, snapshot app.HandoffSnapshot) (app.HandoffSnapshot, error) {
						testSnapshot.ID = uuid.NewV4()
						testSnapshot.CreatedAt = time.Now()
						testSnapshot.UpdatedAt = time.Now()

						testSnapshot.OnboardersLocationID = snapshot.OnboardersLocationID
						testSnapshot.PointOfContactEmail = snapshot.PointOfContactEmail
						testSnapshot.ReasonForPurchase = snapshot.ReasonForPurchase
						testSnapshot.Customizations = snapshot.Customizations
						testSnapshot.CustomizationSetup = snapshot.CustomizationSetup
						return testSnapshot, nil
					},
					ReadByOnboardersLocationIDFn: func(ctx context.Context, onboardersLocationId uuid.UUID) (app.HandoffSnapshot, error) {
						return testSnapshot, nil
					},
					SubmitCSATFn: func(ctx context.Context, onboardersLocationId uuid.UUID, csatRecipientUserEmail string) (app.HandoffSnapshot, error) {
						testSnapshot.CSATSentAt = null.NewTime(time.Now())
						testSnapshot.CsatRecipientUserEmail = null.NewString(csatRecipientUserEmail)
						return testSnapshot, nil
					},
					SubmitHandoffFn: func(ctx context.Context, onboardersLocationId uuid.UUID) (app.HandoffSnapshot, error) {
						testSnapshot.HandedOffAt = null.NewTime(time.Now())
						return testSnapshot, nil
					},
				},
			},
			args: args{
				ctx: context.Background(),
				req: &insysproto.HandoffSnapshotCreateOrUpdateRequest{
					HandoffSnapshot: &insysproto.HandoffSnapshotRecord{
						OnboardersLocationId: onboardersLocationID.String(),
						PointOfContactEmail:  pointOfContactEmail.String(),
						ReasonForPurchase:    reasonForPurchase,
						Customizations:       customizations,
						CustomizationSetup:   customizationSetup,
					},
				},
				submitCsatParams: &insysproto.SubmitCSATRequest{
					OnboardersLocationId:   onboardersLocationID.String(),
					CsatRecipientUserEmail: userID.String(),
				},
				handoffSnapshotReadRequest: &insysproto.HandoffSnapshotReadRequest{
					OnboardersLocationId: onboardersLocationID.String(),
				},
				submitHandoffRequest: &insysproto.SubmitHandoffRequest{
					OnboardersLocationId: onboardersLocationID.String(),
				},
			},
			want: &insysproto.HandoffSnapshotResponse{
				HandoffSnapshot: &insysproto.HandoffSnapshotRecord{
					OnboardersLocationId: onboardersLocationID.String(),
					PointOfContactEmail:  pointOfContactEmail.String(),
					ReasonForPurchase:    reasonForPurchase,
					Customizations:       customizations,
					CustomizationSetup:   customizationSetup,
				},
			},
			wantErr: true,
		},
	}

	// custom functions to ignore fields in cmp.Equal comparison
	opts := []cmp.Option{
		cmpopts.IgnoreFields(insysproto.HandoffSnapshotRecord{}, "Id", "CreatedAt", "UpdatedAt", "CsatSentAt"),
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testSnapshot = app.HandoffSnapshot{}

			s := &HandoffSnapshotServer{
				handoffSnapshotService: tt.fields.handoffSnapshotService,
			}
			got, err := s.CreateOrUpdate(tt.args.ctx, tt.args.req)
			if err != nil {
				t.Errorf("HandoffSnapshotServer.CreateOrUpdate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("HandoffSnapshotServer.CreateOrUpdate() = %v", cmp.Diff(got, tt.want, opts...))
			}

			// Test the read function
			read, err := s.ReadByOnboardersLocationID(tt.args.ctx, tt.args.handoffSnapshotReadRequest)
			if !cmp.Equal(read, tt.want, opts...) {
				t.Errorf("HandoffSnapshotService.ReadByOnboardersLocationID() = %v", cmp.Diff(read, tt.want, opts...))
			}

			// Submit CSAT
			csat, err := s.SubmitCSAT(tt.args.ctx, tt.args.submitCsatParams)
			if err != nil {
				t.Errorf("HandoffSnapshotServer.SubmitCSAT() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if csat.HandoffSnapshot.CsatSentAt == got.HandoffSnapshot.CsatSentAt {
				t.Errorf("HandoffSnapshotServer.SubmitCSAT(), CSAT submission not recorded")
			}

			// This allows us to check that the 2nd CSAT timestamp changed.  We only have second granularity in now()
			time.Sleep(1 * time.Second)

			// Submit CSAT again and it should pass
			csat2, err := s.SubmitCSAT(tt.args.ctx, tt.args.submitCsatParams)
			if err != nil {
				t.Errorf("HandoffSnapshotServer.SubmitCSAT() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if csat2.HandoffSnapshot.CsatSentAt == csat.HandoffSnapshot.CsatSentAt {
				t.Errorf("HandoffSnapshotServer.SubmitCSAT(), 2nd csat submission not recorded")
			}

			// Submit Handoff
			submitted, err := s.SubmitHandoff(tt.args.ctx, tt.args.submitHandoffRequest)
			if (err != nil) != tt.wantErr {
				t.Errorf("HandoffSnapshotService.SubmitHandoff() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err != nil {
				t.Logf("expected error submitting handoff due to missing fields %v", err)
				return
			}
			if submitted.HandoffSnapshot.HandedOffAt == read.HandoffSnapshot.HandedOffAt {
				t.Errorf("HandoffSnapshotService.SubmitCSAT(), handoff submission not recorded")
			}

			// Submit Handoff again and it should fail
			submitted2, err := s.SubmitHandoff(tt.args.ctx, tt.args.submitHandoffRequest)
			if err == nil {
				t.Errorf("HandoffSnapshotService.SubmitHandoff() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if submitted2 != nil && submitted.HandoffSnapshot.HandedOffAt == submitted2.HandoffSnapshot.HandedOffAt {
				t.Errorf("HandoffSnapshotService.SubmitCSAT(), handoff submission allowed twice")
			}

			// Attempt to update handoff and it should fail
			updated, err := s.CreateOrUpdate(tt.args.ctx, tt.args.req)
			if err == nil {
				t.Errorf("HandoffSnapshotService.SubmitHandoff() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if updated != nil && submitted.HandoffSnapshot.HandedOffAt == submitted2.HandoffSnapshot.HandedOffAt {
				t.Errorf("HandoffSnapshotService.SubmitCSAT(), handoff update allowed after submission")
			}
		})
	}
}
