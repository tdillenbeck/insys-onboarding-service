package psql

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"weavelab.xyz/insys-onboarding-service/internal/app"
	"weavelab.xyz/monorail/shared/wlib/uuid"
	"weavelab.xyz/monorail/shared/wlib/wsql"
)

func TestRescheduleTrackingService_CreateOrUpdate(t *testing.T) {
	db := initDBConnection(t)
	clearExistingData(db)

	locationID := uuid.NewV4()

	type fields struct {
		DB *wsql.PG
	}
	type args struct {
		ctx        context.Context
		locationID uuid.UUID
		count      int
		eventType  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *app.RescheduleTracking
		wantErr bool
	}{
		{
			name:   "successfully create reschedule tracking record",
			fields: fields{DB: db},
			args: args{
				context.Background(),
				locationID,
				4,
				"software_install_call",
			},
			want: &app.RescheduleTracking{
				LocationID:             locationID,
				EventType:              "software_install_call",
				RescheduledEventsCount: 4,
			},
			wantErr: false,
		},
	}

	opts := []cmp.Option{
		cmpopts.IgnoreFields(app.RescheduleTracking{}, "ID", "CreatedAt", "UpdatedAt", "RescheuleEventsCalculatedAt"),
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &RescheduleTrackingEventService{
				DB: tt.fields.DB,
			}
			got, err := s.CreateOrUpdate(tt.args.ctx, tt.args.locationID, tt.args.count, tt.args.eventType)
			if (err != nil) != tt.wantErr {
				t.Errorf("RescheduleTrackingService.CreateOrUpdate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("RescheduleTrackingService.CreateOrUpdate(). Diff: %v", cmp.Diff(got, tt.want, opts...))
			}
		})
	}
}
