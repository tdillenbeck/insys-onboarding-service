package psql

import (
	"context"

	"weavelab.xyz/insys-onboarding-service/internal/app"
	"weavelab.xyz/monorail/shared/wlib/uuid"
	"weavelab.xyz/monorail/shared/wlib/werror"
	"weavelab.xyz/monorail/shared/wlib/wsql"
)

type RescheduleTrackingService struct {
	DB *wsql.PG
}

func (s *RescheduleTrackingService) CreateOrUpdate(ctx context.Context, locationID uuid.UUID, count int, eventType string) error {

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
