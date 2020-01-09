package psql

import (
	"context"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"weavelab.xyz/insys-onboarding-service/internal/app"
	"weavelab.xyz/monorail/shared/go-utilities/null"
	"weavelab.xyz/monorail/shared/wlib/uuid"
	"weavelab.xyz/monorail/shared/wlib/wsql"
)

func TestHandoffSnapshotService_CreateOrUpdate(t *testing.T) {
	db := initDBConnection(t)
	clearExistingData(db)

	// SETUP.  Snapshot needs an OnboarderLocation
	onboarderService := OnboarderService{DB: db}
	onboarder, err := onboarderService.CreateOrUpdate(context.Background(), &app.Onboarder{UserID: uuid.NewV4()})
	if err != nil {
		t.Fatal(err)
	}

	_ = onboarder

	onboarderLocationService := OnboardersLocationService{DB: db}
	onboarderLocation, err := onboarderLocationService.CreateOrUpdate(
		context.Background(),
		&app.OnboardersLocation{
			OnboarderID: onboarder.ID,
			LocationID:  uuid.NewV4(),
		},
	)
	if err != nil {
		t.Fatal(err)
	}

	onboardersLocationID := onboarderLocation.ID

	updatedSurveySentAt := null.NewTime(time.Now().Add(5 * time.Hour))

	pointOfContact := null.NewString("client@example.com")
	reasonForPurchase := null.NewString("reason")
	customizations := null.NewBool(true)
	customizationSetup := null.NewString("notes about customizations")
	faxPortSubmitted := null.NewString("yes")
	routerType := null.NewString("Red")
	disclaimerTypeSent := null.NewString("email")
	routerMakeAndModel := null.NewString("make and model")
	networkDecision := null.NewString("notes about network")
	billingNotes := null.NewString("notes about billing")
	notes := null.NewString("general notes")
	notes1 := null.NewString("general notes and more")
	notes2 := null.NewString("general notes and even more")

	type fields struct {
		DB *wsql.PG
	}
	type args struct {
		ctx      context.Context
		snapshot app.HandoffSnapshot
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    app.HandoffSnapshot
		wantErr bool
	}{
		{
			name: "insert handoff snapshot",
			fields: fields{
				DB: db,
			},
			args: args{
				ctx: context.Background(),
				snapshot: app.HandoffSnapshot{
					OnboardersLocationID: onboardersLocationID,
				},
			},
			want: app.HandoffSnapshot{
				OnboardersLocationID: onboardersLocationID,
			},
		},
		{
			name: "Add fields",
			fields: fields{
				DB: db,
			},
			args: args{
				ctx: context.Background(),
				snapshot: app.HandoffSnapshot{
					OnboardersLocationID: onboardersLocationID,
					Customizations:       customizations,
					FaxPortSubmitted:     faxPortSubmitted,

				},
			},
			want: app.HandoffSnapshot{
				OnboardersLocationID: onboardersLocationID,
				Customizations:       customizations,
				FaxPortSubmitted:     faxPortSubmitted,
			},
		},
		{
			name: "Add notes",
			fields: fields{
				DB: db,
			},
			args: args{
				ctx: context.Background(),
				snapshot: app.HandoffSnapshot{
					OnboardersLocationID: onboardersLocationID,
					Notes:                notes,

				},
			},
			want: app.HandoffSnapshot{
				OnboardersLocationID: onboardersLocationID,
				Notes:                notes,
			},
		},
		{
			name: "update notes",
			fields: fields{
				DB: db,
			},
			args: args{
				ctx: context.Background(),
				snapshot: app.HandoffSnapshot{
					OnboardersLocationID: onboardersLocationID,
					Notes:                notes1,
				},
			},
			want: app.HandoffSnapshot{
				OnboardersLocationID: onboardersLocationID,
				Notes:                notes1,
			},
		},
		{
			name: "update notes 2",
			fields: fields{
				DB: db,
			},
			args: args{
				ctx: context.Background(),
				snapshot: app.HandoffSnapshot{
					OnboardersLocationID: onboardersLocationID,
					CSATSentAt:           updatedSurveySentAt,
					Notes:                notes2,
				},
			},
			want: app.HandoffSnapshot{
				OnboardersLocationID: onboardersLocationID,
				CSATSentAt:           updatedSurveySentAt,
				Notes:                notes2,
			},
		},
		{
			name: "Fill out all fields",
			fields: fields{
				DB: db,
			},
			args: args{
				ctx: context.Background(),
				snapshot: app.HandoffSnapshot{
					OnboardersLocationID: onboardersLocationID,
					PointOfContactEmail:  pointOfContact,
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
			want: app.HandoffSnapshot{
				OnboardersLocationID: onboardersLocationID,
				PointOfContactEmail:  pointOfContact,
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
	}

	// custom functions to ignore fields in cmp.Equal comparison
	opts := []cmp.Option{
		cmpopts.IgnoreFields(app.HandoffSnapshot{}, "ID", "CreatedAt", "UpdatedAt", "CSATSentAt"),
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hos := HandoffSnapshotService{
				DB: tt.fields.DB,
			}
			got, err := hos.CreateOrUpdate(tt.args.ctx, tt.args.snapshot)
			if (err != nil) != tt.wantErr {
				t.Errorf("HandoffSnapshotService.CreateOrUpdate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("HandoffSnapshotService.CreateOrUpdate() = %v", cmp.Diff(got, tt.want, opts...))
			}
			read, err := hos.ReadByOnboardersLocationID(tt.args.ctx, got.OnboardersLocationID)
			if !cmp.Equal(read, tt.want, opts...) {
				t.Errorf("HandoffSnapshotService.ReadByOnboardersLocationID() = %v", cmp.Diff(read, tt.want, opts...))
			}
		})
	}
}

func TestHandoffSnapshotService_Submit(t *testing.T) {
	db := initDBConnection(t)
	clearExistingData(db)

	// SETUP.  Snapshot needs an OnboarderLocation
	onboarderService := OnboarderService{DB: db}
	onboarder, err := onboarderService.CreateOrUpdate(context.Background(), &app.Onboarder{UserID: uuid.NewV4()})
	if err != nil {
		t.Fatal(err)
	}

	_ = onboarder

	onboarderLocationService := OnboardersLocationService{DB: db}
	onboarderLocation, err := onboarderLocationService.CreateOrUpdate(
		context.Background(),
		&app.OnboardersLocation{
			OnboarderID: onboarder.ID,
			LocationID:  uuid.NewV4(),
		},
	)
	if err != nil {
		t.Fatal(err)
	}

	onboardersLocationID := onboarderLocation.ID

	userID := null.NewString("client@example.com")
	pointOfContactEmail := null.NewString("client2@example.com")
	reasonForPurchase := null.NewString("reason")
	customizations := null.NewBool(true)
	faxPortSubmitted := null.NewString("yes")
	routerType := null.NewString("Red")
	disclaimerTypeSent := null.NewString("email")
	routerMakeAndModel := null.NewString("make and model")
	networkDecision := null.NewString("notes about network")
	billingNotes := null.NewString("notes about billing")
	notes := null.NewString("general notes")

	type fields struct {
		DB *wsql.PG
	}
	type args struct {
		ctx      context.Context
		snapshot app.HandoffSnapshot
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    app.HandoffSnapshot
		wantErr bool
	}{
		{
			name: "Fill out all fields",
			fields: fields{
				DB: db,
			},
			args: args{
				ctx: context.Background(),
				snapshot: app.HandoffSnapshot{
					OnboardersLocationID: onboardersLocationID,
					PointOfContactEmail:  pointOfContactEmail,
					ReasonForPurchase:    reasonForPurchase,
					Customizations:       customizations,
					FaxPortSubmitted:     faxPortSubmitted,
					RouterType:           routerType,
					DisclaimerTypeSent:   disclaimerTypeSent,
					RouterMakeAndModel:   routerMakeAndModel,
					NetworkDecision:      networkDecision,
					BillingNotes:         billingNotes,
					Notes:                notes,
				},
			},
			want: app.HandoffSnapshot{
				OnboardersLocationID: onboardersLocationID,
				PointOfContactEmail:  pointOfContactEmail,
				ReasonForPurchase:    reasonForPurchase,
				Customizations:       customizations,
				FaxPortSubmitted:     faxPortSubmitted,
				RouterType:           routerType,
				DisclaimerTypeSent:   disclaimerTypeSent,
				RouterMakeAndModel:   routerMakeAndModel,
				NetworkDecision:      networkDecision,
				BillingNotes:         billingNotes,
				Notes:                notes,
			},
		},
	}

	// custom functions to ignore fields in cmp.Equal comparison
	opts := []cmp.Option{
		cmpopts.IgnoreFields(app.HandoffSnapshot{}, "ID", "CreatedAt", "UpdatedAt", "CSATSentAt"),
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hos := HandoffSnapshotService{
				DB: tt.fields.DB,
			}

			// Create snapshot
			got, err := hos.CreateOrUpdate(tt.args.ctx, tt.args.snapshot)
			if (err != nil) != tt.wantErr {
				t.Errorf("HandoffSnapshotService.CreateOrUpdate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("HandoffSnapshotService.CreateOrUpdate() = %v", cmp.Diff(got, tt.want, opts...))
			}

			// Read to verify it wrote correctly
			read, err := hos.ReadByOnboardersLocationID(tt.args.ctx, got.OnboardersLocationID)
			if !cmp.Equal(read, tt.want, opts...) {
				t.Errorf("HandoffSnapshotService.ReadByOnboardersLocationID() = %v", cmp.Diff(read, tt.want, opts...))
			}

			// Submit CSAT
			csat, err := hos.SubmitCSAT(tt.args.ctx, got.OnboardersLocationID, userID.String())
			if (err != nil) != tt.wantErr {
				t.Errorf("HandoffSnapshotService.SubmitCSAT() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if csat.CSATSentAt == got.CSATSentAt {
				t.Errorf("HandoffSnapshotService.SubmitCSAT(), CSAT submission not recorded")
			}

			// Submit CSAT again and it should pass
			csat2, err := hos.SubmitCSAT(tt.args.ctx, got.OnboardersLocationID, userID.String())
			if (err != nil) != tt.wantErr {
				t.Errorf("HandoffSnapshotService.SubmitCSAT() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if csat.CSATSentAt == csat2.CSATSentAt {
				t.Errorf("HandoffSnapshotService.SubmitCSAT(), 2nd CSAT submission not recorded")
			}

			// Submit Handoff
			submitted, err := hos.SubmitHandoff(tt.args.ctx, got.OnboardersLocationID)
			if (err != nil) != tt.wantErr {
				t.Errorf("HandoffSnapshotService.SubmitHandoff() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if submitted.HandedOffAt == read.HandedOffAt {
				t.Errorf("HandoffSnapshotService.SubmitCSAT(), handoff submission not recorded")
			}

			// Submit Handoff again and it should fail
			submitted2, err := hos.SubmitHandoff(tt.args.ctx, got.OnboardersLocationID)
			if err == nil {
				t.Errorf("HandoffSnapshotService.SubmitHandoff() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if submitted.HandedOffAt == submitted2.HandedOffAt {
				t.Errorf("HandoffSnapshotService.SubmitCSAT(), handoff submission allowed twice")
			}
		})
	}
}
