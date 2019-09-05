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

func TestChiliPiperScheduleService_ByLocationID(t *testing.T) {
	db := initDBConnection(t)
	clearExistingData(db)

	locationID := uuid.NewV4()
	currentTime := time.Now()

	eventService := ChiliPiperScheduleEventService{DB: db}

	// create records for us to retrieve
	_, err := eventService.Create(
		context.Background(),
		&app.ChiliPiperScheduleEvent{
			LocationID: locationID,
			EventID:    "testing event id 1",

			AssigneeID: null.NewString("testing assignee id 1"),
			ContactID:  null.NewString("testing contact id 1"),
			EventType:  null.NewString("testing event type 1"),
			RouteID:    null.NewString("testing route id 1"),

			StartAt: null.NewTime(currentTime),
			EndAt:   null.NewTime(currentTime),
		},
	)
	if err != nil {
		t.Fatal("could not create ChiliPiperScheduleEvent for setup in ByLocationID")
	}
	_, err = eventService.Create(
		context.Background(),
		&app.ChiliPiperScheduleEvent{
			LocationID: locationID,
			EventID:    "testing event id 2",

			AssigneeID: null.NewString("testing assignee id 2"),
			ContactID:  null.NewString("testing contact id 2"),
			EventType:  null.NewString("testing event type 2"),
			RouteID:    null.NewString("testing route id 2"),

			StartAt: null.NewTime(currentTime),
			EndAt:   null.NewTime(currentTime),
		},
	)
	if err != nil {
		t.Fatal("could not create ChiliPiperScheduleEvent for setup in ByLocationID")
	}

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
		{
			name:   "successfully retrieve records by location id",
			fields: fields{DB: db},
			args: args{
				context.Background(),
				locationID,
			},
			want: []app.ChiliPiperScheduleEvent{
				{
					LocationID: locationID,
					EventID:    "testing event id 1",

					AssigneeID: null.NewString("testing assignee id 1"),
					ContactID:  null.NewString("testing contact id 1"),
					EventType:  null.NewString("testing event type 1"),
					RouteID:    null.NewString("testing route id 1"),

					StartAt:    null.NewTime(currentTime),
					EndAt:      null.NewTime(currentTime),
					CanceledAt: null.Time{},
				},
				{
					LocationID: locationID,
					EventID:    "testing event id 2",

					AssigneeID: null.NewString("testing assignee id 2"),
					ContactID:  null.NewString("testing contact id 2"),
					EventType:  null.NewString("testing event type 2"),
					RouteID:    null.NewString("testing route id 2"),

					StartAt:    null.NewTime(currentTime),
					EndAt:      null.NewTime(currentTime),
					CanceledAt: null.Time{},
				},
			},
			wantErr: false,
		},
		{
			name:   "returns nil when there are no scheduled events for a location",
			fields: fields{DB: db},
			args: args{
				context.Background(),
				uuid.NewV4(),
			},
			want:    nil,
			wantErr: false,
		},
	}

	// custom functions to ignore fields in cmp.Equal comparison
	opts := []cmp.Option{
		cmpopts.IgnoreFields(app.ChiliPiperScheduleEvent{}, "ID", "CreatedAt", "UpdatedAt"),
		cmp.Comparer(func(x, y null.Time) bool {
			// xor the valid fields to handle empty null.Time struct comparison
			if x.Valid != y.Valid {
				return true
			}
			diff := x.Time.Sub(y.Time)
			return diff < (1 * time.Millisecond)
		}),
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

			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("ChiliPiperScheduleService.ByLocationID(). Diff :%v", cmp.Diff(got, tt.want, opts...))
			}
		})
	}
}

func TestChiliPiperScheduleService_Create(t *testing.T) {
	db := initDBConnection(t)
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
					EventID:    "testing event id",
					EventType:  null.NewString("testing event type"),
					RouteID:    null.NewString("testing route id"),
					AssigneeID: null.NewString("testing assignee id"),

					StartAt: null.NewTime(startAt),
					EndAt:   null.NewTime(endAt),
				},
			},
			want: &app.ChiliPiperScheduleEvent{
				LocationID: locationID,
				EventID:    "testing event id",
				EventType:  null.NewString("testing event type"),
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

func TestChiliPiperScheduleEventService_Update(t *testing.T) {
	db := initDBConnection(t)
	clearExistingData(db)

	locationID := uuid.NewV4()
	eventService := ChiliPiperScheduleEventService{DB: db}
	currentTime := time.Now()
	reassignedAssignee := uuid.NewV4().String()
	rescheduleStartAt := currentTime.Add(24 * time.Hour)
	rescheduleEndAt := currentTime.Add(25 * time.Hour)

	// create record to update
	existingEventForReassignment, err := eventService.Create(
		context.Background(),
		&app.ChiliPiperScheduleEvent{
			LocationID: locationID,
			EventID:    "testing event id 1",

			AssigneeID: null.NewString("testing assignee id 1"),
			ContactID:  null.NewString("testing contact id 1"),
			EventType:  null.NewString("testing event type 1"),
			RouteID:    null.NewString("testing route id 1"),

			StartAt: null.NewTime(currentTime),
			EndAt:   null.NewTime(currentTime),
		},
	)
	if err != nil {
		t.Fatal("could not create ChiliPiperScheduleEvent for reassignment in setup for update")
	}
	existingEventForReschedule, err := eventService.Create(
		context.Background(),
		&app.ChiliPiperScheduleEvent{
			LocationID: locationID,
			EventID:    "testing event id 2",

			AssigneeID: null.NewString("testing assignee id 1"),
			ContactID:  null.NewString("testing contact id 1"),
			EventType:  null.NewString("testing event type 1"),
			RouteID:    null.NewString("testing route id 1"),

			StartAt: null.NewTime(currentTime),
			EndAt:   null.NewTime(currentTime),
		},
	)
	if err != nil {
		t.Fatal("could not create ChiliPiperScheduleEvent for reschedule in setup for update")
	}

	type fields struct {
		DB *wsql.PG
	}
	type args struct {
		ctx        context.Context
		eventID    string
		assigneeID string
		startAt    null.Time
		endAt      null.Time
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *app.ChiliPiperScheduleEvent
		wantErr bool
	}{
		{
			name:   "successfully reassign an onboarder",
			fields: fields{DB: db},
			args: args{
				context.Background(),
				existingEventForReassignment.EventID,
				reassignedAssignee,
				existingEventForReassignment.StartAt,
				existingEventForReassignment.EndAt,
			},
			want: &app.ChiliPiperScheduleEvent{
				ID:         existingEventForReassignment.ID,
				LocationID: existingEventForReassignment.LocationID,

				AssigneeID: null.NewString(reassignedAssignee),
				ContactID:  existingEventForReassignment.ContactID,
				EventID:    existingEventForReassignment.EventID,
				EventType:  existingEventForReassignment.EventType,
				RouteID:    existingEventForReassignment.RouteID,

				StartAt: existingEventForReassignment.StartAt,
				EndAt:   existingEventForReassignment.EndAt,

				CreatedAt: existingEventForReassignment.CreatedAt,
				UpdatedAt: existingEventForReassignment.UpdatedAt,
			},
			wantErr: false,
		},
		{
			name:   "successfully reschedule a meeting to another time",
			fields: fields{DB: db},
			args: args{
				context.Background(),
				existingEventForReschedule.EventID,
				existingEventForReschedule.AssigneeID.String(),
				null.NewTime(rescheduleStartAt),
				null.NewTime(rescheduleEndAt),
			},
			want: &app.ChiliPiperScheduleEvent{
				ID:         existingEventForReschedule.ID,
				LocationID: existingEventForReschedule.LocationID,

				AssigneeID: existingEventForReschedule.AssigneeID,
				ContactID:  existingEventForReschedule.ContactID,
				EventID:    existingEventForReschedule.EventID,
				EventType:  existingEventForReschedule.EventType,
				RouteID:    existingEventForReschedule.RouteID,

				StartAt: null.NewTime(rescheduleStartAt),
				EndAt:   null.NewTime(rescheduleEndAt),

				CreatedAt: existingEventForReschedule.CreatedAt,
				UpdatedAt: existingEventForReschedule.UpdatedAt,
			},
			wantErr: false,
		},
	}

	// custom functions to ignore fields in cmp.Equal comparison
	opts := []cmp.Option{
		cmpopts.IgnoreFields(app.ChiliPiperScheduleEvent{}, "UpdatedAt"),
		cmp.Comparer(func(x, y null.Time) bool {
			diff := x.Time.Sub(y.Time)
			return diff < (1 * time.Millisecond)
		}),
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &ChiliPiperScheduleEventService{
				DB: tt.fields.DB,
			}
			got, err := s.Update(tt.args.ctx, tt.args.eventID, tt.args.assigneeID, tt.args.startAt, tt.args.endAt)
			if (err != nil) != tt.wantErr {
				t.Errorf("ChiliPiperScheduleEventService.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("ChiliPiperScheduleEventService.Update(). Diff: %v", cmp.Diff(got, tt.want, opts...))
			}
		})
	}
}

func TestChiliPiperScheduleEventService_Cancel(t *testing.T) {
	db := initDBConnection(t)
	clearExistingData(db)

	currentTime := time.Now()
	locationID := uuid.NewV4()
	eventService := ChiliPiperScheduleEventService{DB: db}

	existingEventForCancelation, err := eventService.Create(
		context.Background(),
		&app.ChiliPiperScheduleEvent{
			LocationID: locationID,
			EventID:    "testing event id 2",
			CanceledAt: null.NewTime(currentTime),
		},
	)
	if err != nil {
		t.Fatal("could not create ChiliPiperScheduleEvent for reassignment in setup for cancel -> id = 2")
	}

	type fields struct {
		DB *wsql.PG
	}
	type args struct {
		ctx     context.Context
		eventID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *app.ChiliPiperScheduleEvent
		wantErr bool
	}{
		{
			name:   "successfully cancel an appointment",
			fields: fields{DB: db},
			args: args{
				context.Background(),
				existingEventForCancelation.EventID,
			},
			want: &app.ChiliPiperScheduleEvent{
				ID:         existingEventForCancelation.ID,
				LocationID: existingEventForCancelation.LocationID,
				CreatedAt:  existingEventForCancelation.CreatedAt,
				EventID:    existingEventForCancelation.EventID,
			},
			wantErr: false,
		},
	}

	opts := []cmp.Option{
		cmpopts.IgnoreFields(app.ChiliPiperScheduleEvent{}, "UpdatedAt", "CanceledAt"),
		cmp.Comparer(func(x, y null.Time) bool {
			diff := x.Time.Sub(y.Time)
			return diff < (1 * time.Millisecond)
		}),
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &ChiliPiperScheduleEventService{
				DB: tt.fields.DB,
			}
			got, err := s.Cancel(tt.args.ctx, tt.args.eventID)
			if (err != nil) != tt.wantErr {
				t.Errorf("ChiliPiperScheduleEventService.Cancel() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("ChiliPiperScheduleEventService.Cancel(). Diff: %v", cmp.Diff(got, tt.want, opts...))
			}

			//a canceled event will have its UpdatedAt and CanceledAt fields set to the current time
			updatedDiff := time.Now().Sub(got.UpdatedAt)
			if updatedDiff > (1 * time.Millisecond) {
				t.Errorf("Updated at is not within the range. Diff: %v", updatedDiff)
			}

			canceledDiff := time.Now().Sub(got.CanceledAt.Time)
			if canceledDiff > (1 * time.Millisecond) {
				t.Errorf("Canceled at is not within the range. Diff: %v", canceledDiff)
			}
		})
	}
}
