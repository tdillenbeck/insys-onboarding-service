package psql

import (
	"context"

	"weavelab.xyz/insys-onboarding-service/internal/app"
	"weavelab.xyz/monorail/shared/wlib/uuid"
	"weavelab.xyz/monorail/shared/wlib/wsql"
)

type ChiliPiperScheduleEventService struct {
	DB *wsql.PG
}

func (s *ChiliPiperScheduleEventService) ByLocationID(ctx context.Context, locationID uuid.UUID) ([]app.ChiliPiperScheduleEvent, error) {
	return nil, nil
}

func (s *ChiliPiperScheduleEventService) Create(ctx context.Context, scheduleEvent *app.ChiliPiperScheduleEvent) (*app.ChiliPiperScheduleEvent, error) {
	return nil, nil
}
