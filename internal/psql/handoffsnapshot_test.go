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

	surveySentAt := null.NewTime(time.Now())
	updatedSurveySentAt := null.NewTime(time.Now().Add(5 * time.Hour))

	userID := null.NewUUIDUUID(uuid.NewV4())
	updatedUserID := null.NewUUIDUUID(uuid.NewV4())

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
					CSATRecipientUserID:  userID,
					CSATSentAt:           surveySentAt,
				},
			},
			want: app.HandoffSnapshot{
				OnboardersLocationID: onboardersLocationID,
				CSATRecipientUserID:  userID,
				CSATSentAt:           surveySentAt,
			},
		},
		{
			name: "update handoff snapshot",
			fields: fields{
				DB: db,
			},
			args: args{
				ctx: context.Background(),
				snapshot: app.HandoffSnapshot{
					OnboardersLocationID: onboardersLocationID,
					CSATRecipientUserID:  updatedUserID,
					CSATSentAt:           updatedSurveySentAt,
				},
			},
			want: app.HandoffSnapshot{
				OnboardersLocationID: onboardersLocationID,
				CSATRecipientUserID:  updatedUserID,
				CSATSentAt:           updatedSurveySentAt,
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
		})
	}
}
