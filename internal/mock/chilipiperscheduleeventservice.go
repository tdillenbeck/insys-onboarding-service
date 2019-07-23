package mock

import (
	"context"

	"weavelab.xyz/insys-onboarding-service/internal/app"
	"weavelab.xyz/monorail/shared/go-utilities/null"
	"weavelab.xyz/monorail/shared/wlib/uuid"
)

type ChiliPiperScheduleEventService struct {
	ByLocationIDFn func(ctx context.Context, locationID uuid.UUID) ([]app.ChiliPiperScheduleEvent, error)
	CreateFn       func(ctx context.Context, scheduleEvent *app.ChiliPiperScheduleEvent) (*app.ChiliPiperScheduleEvent, error)
	UpdateFn       func(ctx context.Context, eventID, assigneeID string, startAt, endAt null.Time) (*app.ChiliPiperScheduleEvent, error)
}

func (cpse *ChiliPiperScheduleEventService) ByLocationID(ctx context.Context, locationID uuid.UUID) ([]app.ChiliPiperScheduleEvent, error) {
	return cpse.ByLocationIDFn(ctx, locationID)
}

func (cpse *ChiliPiperScheduleEventService) Create(ctx context.Context, scheduleEvent *app.ChiliPiperScheduleEvent) (*app.ChiliPiperScheduleEvent, error) {
	return cpse.CreateFn(ctx, scheduleEvent)
}

func (cpse *ChiliPiperScheduleEventService) Update(ctx context.Context, eventID, assigneeID string, startAt, endAt null.Time) (*app.ChiliPiperScheduleEvent, error) {
	return cpse.UpdateFn(ctx, eventID, assigneeID, startAt, endAt)
}
