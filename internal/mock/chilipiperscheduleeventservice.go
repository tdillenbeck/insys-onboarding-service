package mock

import (
	"context"

	"weavelab.xyz/insys-onboarding-service/internal/app"
	"weavelab.xyz/monorail/shared/go-utilities/null"
	"weavelab.xyz/monorail/shared/wlib/uuid"
)

type ChiliPiperScheduleEventService struct {
	ByLocationIDFn               func(ctx context.Context, locationID uuid.UUID) ([]app.ChiliPiperScheduleEvent, error)
	CancelFn                     func(ctx context.Context, eventID string) (*app.ChiliPiperScheduleEvent, error)
	CreateFn                     func(ctx context.Context, scheduleEvent *app.ChiliPiperScheduleEvent) (*app.ChiliPiperScheduleEvent, error)
	UpdateFn                     func(ctx context.Context, eventID, assigneeID string, startAt, endAt null.Time) (*app.ChiliPiperScheduleEvent, error)
	CancelCountByLocationFn      func(ctx context.Context, locationID uuid.UUID, eventType string) (int, error)
	UpdateRescheduleEventCountFn func(ctx context.Context, locationID uuid.UUID, count int, eventType string) error
}

func (cpse *ChiliPiperScheduleEventService) ByLocationID(ctx context.Context, locationID uuid.UUID) ([]app.ChiliPiperScheduleEvent, error) {
	return cpse.ByLocationIDFn(ctx, locationID)
}

func (cpse *ChiliPiperScheduleEventService) Cancel(ctx context.Context, eventID string) (*app.ChiliPiperScheduleEvent, error) {
	return cpse.CancelFn(ctx, eventID)
}

func (cpse *ChiliPiperScheduleEventService) Create(ctx context.Context, scheduleEvent *app.ChiliPiperScheduleEvent) (*app.ChiliPiperScheduleEvent, error) {
	return cpse.CreateFn(ctx, scheduleEvent)
}

func (cpse *ChiliPiperScheduleEventService) Update(ctx context.Context, eventID, assigneeID string, startAt, endAt null.Time) (*app.ChiliPiperScheduleEvent, error) {
	return cpse.UpdateFn(ctx, eventID, assigneeID, startAt, endAt)
}

func (cpse *ChiliPiperScheduleEventService) CanceledCountByLocationIDAndEventType(ctx context.Context, locationID uuid.UUID, eventType string) (int, error) {
	return cpse.CancelCountByLocationFn(ctx, locationID, eventType)
}
func (cpse *ChiliPiperScheduleEventService) UpdateRescheduleEventCount(ctx context.Context, locationID uuid.UUID, count int, eventType string) error {
	return nil
}
