package grpc

import (
	"context"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"weavelab.xyz/insys-onboarding-service/internal/app"
	"weavelab.xyz/insys-onboarding-service/internal/mock"
	"weavelab.xyz/monorail/shared/go-utilities/null"
	"weavelab.xyz/monorail/shared/protorepo/dist/go/messages/insysproto"
	"weavelab.xyz/monorail/shared/wlib/uuid"
)

func TestChiliPiperScheduleEventServer_ByLocationID(t *testing.T) {
	id := uuid.NewV4()
	locationUUID := uuid.NewV4()
	currentTime := time.Now()

	successfulChiliPiperScheduleEventService := &mock.ChiliPiperScheduleEventService{
		ByLocationIDFn: func(ctx context.Context, locationID uuid.UUID) ([]app.ChiliPiperScheduleEvent, error) {
			return []app.ChiliPiperScheduleEvent{
				{
					ID:         id,
					LocationID: locationUUID,

					AssigneeID: null.NewString("testing assignee id 1"),
					ContactID:  null.NewString("testing contact id 1"),
					EventID:    null.NewString("testing event id 1"),
					EventType:  null.NewString("testing event type 1"),
					RouteID:    null.NewString("testing route id 1"),

					StartAt: null.NewTime(currentTime),
					EndAt:   null.NewTime(currentTime),

					CreatedAt: currentTime,
					UpdatedAt: currentTime,
				},
				{
					ID:         id,
					LocationID: locationUUID,

					AssigneeID: null.NewString("testing assignee id 2"),
					ContactID:  null.NewString("testing contact id 2"),
					EventID:    null.NewString("testing event id 2"),
					EventType:  null.NewString("testing event type 2"),
					RouteID:    null.NewString("testing route id 2"),

					StartAt: null.NewTime(currentTime),
					EndAt:   null.NewTime(currentTime),

					CreatedAt: currentTime,
					UpdatedAt: currentTime,
				},
			}, nil
		},
	}

	type fields struct {
		chiliPiperScheduleEventService app.ChiliPiperScheduleEventService
	}
	type args struct {
		ctx context.Context
		req *insysproto.ByLocationIDChiliPiperScheduleEventRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *insysproto.ByLocationIDChiliPiperScheduleEventResponse
		wantErr bool
	}{
		{
			name:   "successfully retrieves chili piper schedule events by location id",
			fields: fields{chiliPiperScheduleEventService: successfulChiliPiperScheduleEventService},
			args: args{
				context.Background(),
				&insysproto.ByLocationIDChiliPiperScheduleEventRequest{
					LocationId: locationUUID.String(),
				},
			},
			want: &insysproto.ByLocationIDChiliPiperScheduleEventResponse{
				Events: []*insysproto.ChiliPiperScheduleEventRecord{
					{
						Id:         id.String(),
						LocationId: locationUUID.String(),
						EventId:    "testing event id 1",
						EventType:  "testing event type 1",
						RouteId:    "testing route id 1",
						AssigneeId: "testing assignee id 1",
						ContactId:  "testing contact id 1",
						StartAt:    currentTime.Format(time.RFC3339),
						EndAt:      currentTime.Format(time.RFC3339),
						CreatedAt:  currentTime.Format(time.RFC3339Nano),
						UpdatedAt:  currentTime.Format(time.RFC3339Nano),
					},
					{
						Id:         id.String(),
						LocationId: locationUUID.String(),
						EventId:    "testing event id 2",
						EventType:  "testing event type 2",
						RouteId:    "testing route id 2",
						AssigneeId: "testing assignee id 2",
						ContactId:  "testing contact id 2",
						StartAt:    currentTime.Format(time.RFC3339),
						EndAt:      currentTime.Format(time.RFC3339),
						CreatedAt:  currentTime.Format(time.RFC3339Nano),
						UpdatedAt:  currentTime.Format(time.RFC3339Nano),
					},
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &ChiliPiperScheduleEventServer{
				chiliPiperScheduleEventService: tt.fields.chiliPiperScheduleEventService,
			}
			got, err := s.ByLocationID(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("ChiliPiperScheduleEventServer.ByLocationID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !cmp.Equal(got, tt.want) {
				t.Errorf("ChiliPiperScheduleEventServer.ByLocationID(). Diff: %v", cmp.Diff(got, tt.want))
			}
		})
	}
}

func TestChiliPiperScheduleEventServer_Create(t *testing.T) {
	id := uuid.NewV4()
	locationUUID := uuid.NewV4()
	currentTime := time.Now()

	successfulChiliPiperScheduleEventService := &mock.ChiliPiperScheduleEventService{
		CreateFn: func(ctx context.Context, scheduleEvent *app.ChiliPiperScheduleEvent) (*app.ChiliPiperScheduleEvent, error) {
			return &app.ChiliPiperScheduleEvent{
				ID:         id,
				LocationID: locationUUID,

				AssigneeID: null.NewString("testing assignee id 1"),
				ContactID:  null.NewString("testing contact id 1"),
				EventID:    null.NewString("testing event id 1"),
				EventType:  null.NewString("testing event type 1"),
				RouteID:    null.NewString("testing route id 1"),

				StartAt: null.NewTime(currentTime),
				EndAt:   null.NewTime(currentTime),

				CreatedAt: currentTime,
				UpdatedAt: currentTime,
			}, nil
		},
	}

	type fields struct {
		chiliPiperScheduleEventService app.ChiliPiperScheduleEventService
	}
	type args struct {
		ctx context.Context
		req *insysproto.CreateChiliPiperScheduleEventRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *insysproto.CreateChiliPiperScheduleEventResponse
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:   "successfully create a chili piper schedule event",
			fields: fields{chiliPiperScheduleEventService: successfulChiliPiperScheduleEventService},
			args: args{
				context.Background(),
				&insysproto.CreateChiliPiperScheduleEventRequest{
					Event: &insysproto.ChiliPiperScheduleEventRecord{
						LocationId: locationUUID.String(),
						EventId:    "testing event id 1",
						EventType:  "testing event type 1",
						RouteId:    "testing route id 1",
						AssigneeId: "testing assignee id 1",
						ContactId:  "testing contact id 1",
						StartAt:    currentTime.Format(time.RFC3339),
						EndAt:      currentTime.Format(time.RFC3339),
					},
				},
			},
			want: &insysproto.CreateChiliPiperScheduleEventResponse{
				Event: &insysproto.ChiliPiperScheduleEventRecord{
					Id:         id.String(),
					LocationId: locationUUID.String(),
					EventId:    "testing event id 1",
					EventType:  "testing event type 1",
					RouteId:    "testing route id 1",
					AssigneeId: "testing assignee id 1",
					ContactId:  "testing contact id 1",
					StartAt:    currentTime.Format(time.RFC3339),
					EndAt:      currentTime.Format(time.RFC3339),
					CreatedAt:  currentTime.Format(time.RFC3339Nano),
					UpdatedAt:  currentTime.Format(time.RFC3339Nano),
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &ChiliPiperScheduleEventServer{
				chiliPiperScheduleEventService: tt.fields.chiliPiperScheduleEventService,
			}
			got, err := s.Create(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("ChiliPiperScheduleEventServer.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !cmp.Equal(got, tt.want) {
				t.Errorf("ChiliPiperScheduleEventServer.Create(). Diff:  %v", cmp.Diff(got, tt.want))
			}
		})
	}
}
