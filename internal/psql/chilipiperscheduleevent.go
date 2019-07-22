package psql

import (
	"context"
	"time"

	"weavelab.xyz/insys-onboarding-service/internal/app"
	"weavelab.xyz/monorail/shared/go-utilities/null"
	"weavelab.xyz/monorail/shared/wlib/uuid"
	"weavelab.xyz/monorail/shared/wlib/werror"
	"weavelab.xyz/monorail/shared/wlib/wsql"
)

type ChiliPiperScheduleEventService struct {
	DB *wsql.PG
}

func (s *ChiliPiperScheduleEventService) ByLocationID(ctx context.Context, locationID uuid.UUID) ([]app.ChiliPiperScheduleEvent, error) {
	var resultEvents []app.ChiliPiperScheduleEvent

	query := `
	  SELECT
		id,
		location_id,
		event_id,
		event_type,
		route_id,
		assignee_id,
		contact_id,
		start_at,
		end_at,
		created_at,
		updated_at
	  FROM insys_onboarding.chili_piper_schedule_events
	  WHERE location_id = $1 `

	rows, err := s.DB.QueryxContext(ctx, query, locationID.String())
	if err != nil {
		return nil, werror.Wrap(err, "error executing ByLocationID query")
	}
	defer rows.Close()

	for rows.Next() {
		var event app.ChiliPiperScheduleEvent
		err = rows.StructScan(&event)
		if err != nil {
			return nil, werror.Wrap(err, "error scanning result from ByLocationID query into chili piper schedule event struct")
		}

		resultEvents = append(resultEvents, event)
	}

	return resultEvents, nil
}

func (s *ChiliPiperScheduleEventService) Create(ctx context.Context, scheduleEvent *app.ChiliPiperScheduleEvent) (*app.ChiliPiperScheduleEvent, error) {
	var resultEvent app.ChiliPiperScheduleEvent

	query := `
	  INSERT INTO insys_onboarding.chili_piper_schedule_events
	  (
		  id,
		  event_id,
		  event_type,
		  route_id,
		  assignee_id,
		  start_at,
		  end_at,
		  contact_id,
		  location_id,
		  created_at,
		  updated_at
	  )
	  VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
	  RETURNING id, created_at, updated_at`

	currentTime := time.Now()
	row := s.DB.QueryRowContext(
		ctx,
		query,
		uuid.NewV4().String(),
		scheduleEvent.EventID,
		scheduleEvent.EventType,
		scheduleEvent.RouteID,
		scheduleEvent.AssigneeID,
		scheduleEvent.StartAt,
		scheduleEvent.EndAt,
		scheduleEvent.ContactID,
		scheduleEvent.LocationID.String(),
		currentTime,
		currentTime,
	)
	err := row.Scan(
		&resultEvent.ID,
		&resultEvent.CreatedAt,
		&resultEvent.UpdatedAt,
	)
	if err != nil {
		return nil, werror.Wrap(err, "error executing chili piper schedule event create")
	}

	resultEvent.EventID = scheduleEvent.EventID
	resultEvent.EventType = scheduleEvent.EventType
	resultEvent.RouteID = scheduleEvent.RouteID
	resultEvent.AssigneeID = scheduleEvent.AssigneeID
	resultEvent.StartAt = scheduleEvent.StartAt
	resultEvent.EndAt = scheduleEvent.EndAt
	resultEvent.ContactID = scheduleEvent.ContactID
	resultEvent.LocationID = scheduleEvent.LocationID

	return &resultEvent, nil
}

func (s *ChiliPiperScheduleEventService) Update(ctx context.Context, id uuid.UUID, assigneeID string, startAt, endAt null.Time) (*app.ChiliPiperScheduleEvent, error) {
	return nil, nil
}
