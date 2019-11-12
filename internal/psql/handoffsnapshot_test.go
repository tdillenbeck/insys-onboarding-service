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

func TestHandOffSnapshotService_CreateOrUpdate(t *testing.T) {
	db := initDBConnection(t)
	clearExistingData(db)

	onboardersLocationID := uuid.NewV4()

	surveySentAt := null.NewTime(time.Now())
	updatedSurveySentAt := null.NewTime(time.Now().Add(5 * time.Hour))

	userID := null.NewUUIDUUID(uuid.NewV4())
	updatedUserID := null.NewUUIDUUID(uuid.NewV4())

	type fields struct {
		DB *wsql.PG
	}
	type args struct {
		ctx      context.Context
		snapshot *app.HandOffSnapshot
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *app.HandOffSnapshot
		wantErr bool
	}{
		{
			name: "insert hand-off snapshot",
			fields: fields{
				DB: db,
			},
			args: args{
				ctx: context.Background(),
				snapshot: &app.HandOffSnapshot{
					OnboardersLocationID:                      onboardersLocationID,
					CustomerSatisfactionSurveyRecipientUserID: userID,
					CustomerSatisfactionSurveySentAt:          surveySentAt,
				},
			},
			want: &app.HandOffSnapshot{
				OnboardersLocationID:                      onboardersLocationID,
				CustomerSatisfactionSurveyRecipientUserID: userID,
				CustomerSatisfactionSurveySentAt:          surveySentAt,
			},
		},
		{
			name: "update hand-off snapshot",
			fields: fields{
				DB: db,
			},
			args: args{
				ctx: context.Background(),
				snapshot: &app.HandOffSnapshot{
					OnboardersLocationID:                      onboardersLocationID,
					CustomerSatisfactionSurveyRecipientUserID: updatedUserID,
					CustomerSatisfactionSurveySentAt:          updatedSurveySentAt,
				},
			},
			want: &app.HandOffSnapshot{
				OnboardersLocationID:                      onboardersLocationID,
				CustomerSatisfactionSurveyRecipientUserID: updatedUserID,
				CustomerSatisfactionSurveySentAt:          updatedSurveySentAt,
			},
		},
	}

	// custom functions to ignore fields in cmp.Equal comparison
	opts := []cmp.Option{
		cmpopts.IgnoreFields(app.HandOffSnapshot{}, "ID", "CreatedAt", "UpdatedAt", "CustomerSatisfactionSurveySentAt"),
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hos := HandOffSnapshotService{
				DB: tt.fields.DB,
			}
			got, err := hos.CreateOrUpdate(tt.args.ctx, tt.args.snapshot)
			if (err != nil) != tt.wantErr {
				t.Errorf("HandOffSnapshotService.CreateOrUpdate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("HandOffSnapshotService.CreateOrUpdate() = %v", cmp.Diff(got, tt.want, opts...))
			}
		})
	}
}
