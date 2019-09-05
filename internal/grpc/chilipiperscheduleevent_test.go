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

					EventID: "testing event id 1",

					AssigneeID: null.NewString("testing assignee id 1"),
					ContactID:  null.NewString("testing contact id 1"),
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

					EventID: "testing event id 2",

					AssigneeID: null.NewString("testing assignee id 2"),
					ContactID:  null.NewString("testing contact id 2"),
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
		chiliPiperScheduleEventPublisher app.ChiliPiperScheduleEventPublisher
		chiliPiperScheduleEventService   app.ChiliPiperScheduleEventService
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
				chiliPiperScheduleEventPublisher: tt.fields.chiliPiperScheduleEventPublisher,
				chiliPiperScheduleEventService:   tt.fields.chiliPiperScheduleEventService,
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

	noOpChiliPiperScheduleEventPublisher := &mock.ChiliPiperScheduleEventPublisher{
		PublishCreatedFn: func(ctx context.Context, response *insysproto.CreateChiliPiperScheduleEventResponse) error {
			return nil
		},
	}

	successfulChiliPiperScheduleEventService := &mock.ChiliPiperScheduleEventService{
		CreateFn: func(ctx context.Context, scheduleEvent *app.ChiliPiperScheduleEvent) (*app.ChiliPiperScheduleEvent, error) {
			return &app.ChiliPiperScheduleEvent{
				ID:         id,
				LocationID: locationUUID,

				EventID: "testing event id 1",

				AssigneeID: null.NewString("testing assignee id 1"),
				ContactID:  null.NewString("testing contact id 1"),
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
		chiliPiperScheduleEventPublisher app.ChiliPiperScheduleEventPublisher
		chiliPiperScheduleEventService   app.ChiliPiperScheduleEventService
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
		{
			name: "successfully create a chili piper schedule event",
			fields: fields{
				chiliPiperScheduleEventPublisher: noOpChiliPiperScheduleEventPublisher,
				chiliPiperScheduleEventService:   successfulChiliPiperScheduleEventService,
			},
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
				chiliPiperScheduleEventPublisher: tt.fields.chiliPiperScheduleEventPublisher,
				chiliPiperScheduleEventService:   tt.fields.chiliPiperScheduleEventService,
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

func TestChiliPiperScheduleEventServer_Update(t *testing.T) {
	currentTime := time.Now()
	locationUUID := uuid.NewV4()
	existingID := uuid.NewV4()
	eventID := "testing event id"

	successfulChiliPiperScheduleEventService := &mock.ChiliPiperScheduleEventService{
		UpdateFn: func(ctx context.Context, eventID, assigneeID string, startAt, endAt null.Time) (*app.ChiliPiperScheduleEvent, error) {
			return &app.ChiliPiperScheduleEvent{
				ID:         existingID,
				LocationID: locationUUID,

				EventID: eventID,

				AssigneeID: null.NewString(assigneeID),
				ContactID:  null.NewString("testing contact id 1"),
				EventType:  null.NewString("testing event type 1"),
				RouteID:    null.NewString("testing route id 1"),

				StartAt: startAt,
				EndAt:   endAt,

				CreatedAt: currentTime,
				UpdatedAt: currentTime,
			}, nil
		},
	}

	type fields struct {
		chiliPiperScheduleEventPublisher app.ChiliPiperScheduleEventPublisher
		chiliPiperScheduleEventService   app.ChiliPiperScheduleEventService
	}
	type args struct {
		ctx context.Context
		req *insysproto.UpdateChiliPiperScheduleEventRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *insysproto.UpdateChiliPiperScheduleEventResponse
		wantErr bool
	}{
		{
			name:   "successfully update a chili piper schedule event",
			fields: fields{chiliPiperScheduleEventService: successfulChiliPiperScheduleEventService},
			args: args{
				context.Background(),
				&insysproto.UpdateChiliPiperScheduleEventRequest{
					EventId:    eventID,
					AssigneeId: "new assignee id",
					StartAt:    currentTime.Format(time.RFC3339),
					EndAt:      currentTime.Format(time.RFC3339),
				},
			},
			want: &insysproto.UpdateChiliPiperScheduleEventResponse{
				Event: &insysproto.ChiliPiperScheduleEventRecord{
					Id:         existingID.String(),
					LocationId: locationUUID.String(),
					EventId:    eventID,
					EventType:  "testing event type 1",
					RouteId:    "testing route id 1",
					AssigneeId: "new assignee id",
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
				chiliPiperScheduleEventPublisher: tt.fields.chiliPiperScheduleEventPublisher,
				chiliPiperScheduleEventService:   tt.fields.chiliPiperScheduleEventService,
			}
			got, err := s.Update(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("ChiliPiperScheduleEventServer.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !cmp.Equal(got, tt.want) {
				t.Errorf("ChiliPiperScheduleEventServer.Update(). Diff: %v", cmp.Diff(got, tt.want))
			}
		})
	}
}

func TestChiliPiperScheduleEventServer_Cancel(t *testing.T) {
	currentTime := time.Now()
	locationUUID := uuid.NewV4()
	existingID := uuid.NewV4()
	existingStart := null.NewTime(currentTime)
	existingEnd := null.NewTime(currentTime)

	eventID := "testing event id"

	successfulChiliPiperScheduleEventService := &mock.ChiliPiperScheduleEventService{
		CancelFn: func(ctx context.Context, eventID string) (*app.ChiliPiperScheduleEvent, error) {
			return &app.ChiliPiperScheduleEvent{
				ID:         existingID,
				LocationID: locationUUID,

				EventID:    eventID,
				AssigneeID: null.NewString("testing assignee id 1"),
				ContactID:  null.NewString("testing contact id 1"),
				EventType:  null.NewString("testing event type 1"),
				RouteID:    null.NewString("testing route id 1"),

				StartAt:    existingStart,
				EndAt:      existingEnd,
				CanceledAt: null.NewTime(currentTime),

				CreatedAt: currentTime,
				UpdatedAt: currentTime,
			}, nil
		},
	}

	type fields struct {
		chiliPiperScheduleEventPublisher app.ChiliPiperScheduleEventPublisher
		chiliPiperScheduleEventService   app.ChiliPiperScheduleEventService
	}
	type args struct {
		ctx context.Context
		req *insysproto.CancelChiliPiperScheduleEventRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *insysproto.CancelChiliPiperScheduleEventResponse
		wantErr bool
	}{
		{
			name:   "successfully cancel a chili piper schedule event",
			fields: fields{chiliPiperScheduleEventService: successfulChiliPiperScheduleEventService},
			args: args{
				context.Background(),
				&insysproto.CancelChiliPiperScheduleEventRequest{
					EventId: eventID,
				},
			},
			want: &insysproto.CancelChiliPiperScheduleEventResponse{
				Event: &insysproto.ChiliPiperScheduleEventRecord{
					Id:         existingID.String(),
					LocationId: locationUUID.String(),
					EventId:    eventID,
					EventType:  "testing event type 1",
					RouteId:    "testing route id 1",
					AssigneeId: "testing assignee id 1",
					ContactId:  "testing contact id 1",
					StartAt:    currentTime.Format(time.RFC3339),
					EndAt:      currentTime.Format(time.RFC3339),
					CanceledAt: currentTime.Format(time.RFC3339),
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
				chiliPiperScheduleEventPublisher: tt.fields.chiliPiperScheduleEventPublisher,
				chiliPiperScheduleEventService:   tt.fields.chiliPiperScheduleEventService,
			}
			got, err := s.Cancel(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("ChiliPiperScheduleEventServer.Cancel() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !cmp.Equal(got, tt.want) {
				t.Errorf("ChiliPiperScheduleEventServer.Cancel(). Diff: %v", cmp.Diff(got, tt.want))
			}
		})
	}
}
