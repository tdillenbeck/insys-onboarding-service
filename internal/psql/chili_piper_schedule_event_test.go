package psql

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"weavelab.xyz/insys-onboarding-service/internal/app"
	"weavelab.xyz/monorail/shared/go-utilities/null"
	"weavelab.xyz/monorail/shared/wlib/uuid"
	"weavelab.xyz/monorail/shared/wlib/wsql"
)

func TestChiliPiperScheduleService_ByLocationID(t *testing.T) {
	type fields struct {
		DB *wsql.PG
	}
	type args struct {
		ctx        context.Context
		locationID uuid.UUID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []app.ChiliPiperScheduleEvent
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &ChiliPiperScheduleEventService{
				DB: tt.fields.DB,
			}
			got, err := s.ByLocationID(tt.args.ctx, tt.args.locationID)
			if (err != nil) != tt.wantErr {
				t.Errorf("ChiliPiperScheduleService.ByLocationID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ChiliPiperScheduleService.ByLocationID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChiliPiperScheduleService_Create(t *testing.T) {
	db := initDBConnection(t, psqlConnString)
	clearExistingData(db)

	locationID := uuid.NewV4()
	startAt, endAt := time.Now(), time.Now()

	type fields struct {
		DB *wsql.PG
	}
	type args struct {
		ctx           context.Context
		scheduleEvent *app.ChiliPiperScheduleEvent
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *app.ChiliPiperScheduleEvent
		wantErr bool
	}{
		{
			name:   "successfully create chili piper schedule event",
			fields: fields{DB: db},
			args: args{
				context.Background(),
				&app.ChiliPiperScheduleEvent{
					LocationID: locationID,
					EventID:    null.NewString("testing event id"),
					RouteID:    null.NewString("testing route id"),
					AssigneeID: null.NewString("testing assignee id"),

					StartAt: null.NewTime(startAt),
					EndAt:   null.NewTime(endAt),
				},
			},
			want: &app.ChiliPiperScheduleEvent{
				LocationID: locationID,
				EventID:    null.NewString("testing event id"),
				RouteID:    null.NewString("testing route id"),
				AssigneeID: null.NewString("testing assignee id"),

				StartAt: null.NewTime(startAt),
				EndAt:   null.NewTime(endAt),
			},
			wantErr: false,
		},
	}

	// custom functions to ignore fields in cmp.Equal comparison
	opts := []cmp.Option{
		cmpopts.IgnoreFields(app.ChiliPiperScheduleEvent{}, "ID", "CreatedAt", "UpdatedAt"),
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &ChiliPiperScheduleEventService{
				DB: tt.fields.DB,
			}
			got, err := s.Create(tt.args.ctx, tt.args.scheduleEvent)
			if (err != nil) != tt.wantErr {
				t.Errorf("ChiliPiperScheduleService.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("ChiliPiperScheduleService.Create(). Diff :%v", cmp.Diff(got, tt.want, opts...))
			}
		})
	}
}
