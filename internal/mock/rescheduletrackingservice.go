package mock

import (
	"context"

	"weavelab.xyz/insys-onboarding-service/internal/app"
	"weavelab.xyz/monorail/shared/protorepo/dist/go/messages/insysproto"
	"weavelab.xyz/monorail/shared/wlib/uuid"
)

type RescheduleTrackingEventService struct {
	CreateOrUpdateFn         func(ctx context.Context, locationID uuid.UUID, count int, eventType string) (*app.RescheduleTracking, error)
	ReadRescheduleTrackingFn func(ctx context.Context, in *insysproto.RescheduleTrackingRequest) (*app.RescheduleTracking, error)
}

func (rts *RescheduleTrackingEventService) CreateOrUpdate(ctx context.Context, locationID uuid.UUID, count int, eventType string) (*app.RescheduleTracking, error) {
	return rts.CreateOrUpdateFn(ctx, locationID, count, eventType)
}

func (rts *RescheduleTrackingEventService) ReadRescheduleTracking(ctx context.Context, in *insysproto.RescheduleTrackingRequest) (*app.RescheduleTracking, error) {
	return rts.ReadRescheduleTrackingFn(ctx, in)
}
