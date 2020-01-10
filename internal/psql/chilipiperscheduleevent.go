package psql

import (
	"context"
	"database/sql"

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
			canceled_at,
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

func (s *ChiliPiperScheduleEventService) Cancel(ctx context.Context, eventID string) (*app.ChiliPiperScheduleEvent, error) {
	var resultEvent app.ChiliPiperScheduleEvent

	query := `
	  UPDATE insys_onboarding.chili_piper_schedule_events
			SET
				canceled_at = now(),
				updated_at = now()
		 WHERE event_id = $1
		 RETURNING insys_onboarding.chili_piper_schedule_events.*`

	row := s.DB.QueryRowxContext(
		ctx,
		query,
		eventID,
	)

	if row.Err() != nil {
		if row.Err() == sql.ErrNoRows {
			return nil, werror.Wrap(row.Err(), "error returning results from database").SetCode(werror.CodeNotFound)
		}
		return nil, werror.Wrap(row.Err(), "error returning results from database").SetCode(werror.CodeInternal)
	}

	err := row.StructScan(&resultEvent)
	if err != nil {
		return nil, werror.Wrap(err, "error executing chili piper schedule event cancel").SetCode(werror.CodeInternal)
	}

	return &resultEvent, nil
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
		  canceled_at,
		  created_at,
		  updated_at
	  )
	  VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, now(), now())
	  RETURNING id, created_at, updated_at`

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
		scheduleEvent.CanceledAt,
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
	resultEvent.CanceledAt = scheduleEvent.CanceledAt

	return &resultEvent, nil
}

func (s *ChiliPiperScheduleEventService) Update(ctx context.Context, eventID, assigneeID string, startAt, endAt null.Time) (*app.ChiliPiperScheduleEvent, error) {
	var resultEvent app.ChiliPiperScheduleEvent

	query := `
	  UPDATE insys_onboarding.chili_piper_schedule_events
			SET assignee_id = $2,
				start_at = $3,
				end_at = $4,
				updated_at = now()
		 WHERE event_id = $1
		 RETURNING insys_onboarding.chili_piper_schedule_events.*`

	row := s.DB.QueryRowxContext(
		ctx,
		query,
		eventID,
		assigneeID,
		startAt,
		endAt,
	)
	err := row.StructScan(&resultEvent)
	if err != nil {
		return nil, werror.Wrap(err, "error executing chili piper schedule event update")
	}

	return &resultEvent, nil
}

func (s *ChiliPiperScheduleEventService) CanceledCountByLocationIDAndEventType(ctx context.Context, locationID uuid.UUID, eventType string) (int, error) {

	var count int

	query := `
	 SELECT COUNT(*) FROM insys_onboarding.chili_piper_schedule_events
		 WHERE location_id = $1 AND event_type =$2 AND canceled_at IS NOT NULL;`

	row := s.DB.QueryRowxContext(
		ctx,
		query,
		locationID.String(),
		eventType,
	)
	err := row.Scan(&count)
	if err != nil {
		return 0, werror.Wrap(err, "error getting count of canceled events")
	}

	return count, nil

}

func (s *ChiliPiperScheduleEventService) UpdateRescheduleEventCount(ctx context.Context, locationID uuid.UUID, count int, eventType string) error {

	var resultEvent app.RescheduleTracking
	query := `INSERT INTO insys_onboarding.reschedule_tracking
				(id, location_id, event_type, rescheduled_events_count, rescheduled_events_calculated_at, created_at, updated_at)
				VALUES ($1, $2, $3, $4, now(), now(), now())
				ON CONFLICT (location_id, event_type) DO UPDATE SET (rescheduled_events_count,rescheduled_events_calculated_at, updated_at) = (
    				$4,
    				now(),
    				now()
			)`

	row := s.DB.QueryRowxContext(
		ctx,
		query,
		uuid.NewV4(),
		locationID,
		eventType,
		count,
	)
	err := row.StructScan(&resultEvent)
	if err != nil {
		return werror.Wrap(err, "error executing chili piper schedule event update")
	}

	return nil
}
