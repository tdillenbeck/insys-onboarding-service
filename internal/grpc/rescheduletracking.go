package grpc

import (
	"context"
	"fmt"

	"weavelab.xyz/insys-onboarding-service/internal/app"
	"weavelab.xyz/monorail/shared/protorepo/dist/go/messages/insysproto"
	"weavelab.xyz/monorail/shared/protorepo/dist/go/services/insys"
	"weavelab.xyz/monorail/shared/wlib/uuid"
)

var _ insys.RescheduleTrackingEventServer = &RescheduleTrackingEventServer{}

type RescheduleTrackingEventServer struct {
	rescheduleTrackingService app.RescheduleTrackingEventService
}

func NewRescheduleEventServer(rescheduleTrackingService app.RescheduleTrackingEventService) *RescheduleTrackingEventServer {
	return &RescheduleTrackingEventServer{
		rescheduleTrackingService: rescheduleTrackingService,
	}
}

func (r *RescheduleTrackingEventServer) ReadRescheduleTracking(ctx context.Context, in *insysproto.RescheduleTrackingRequest) (*insysproto.RescheduleTrackingResponse, error) {

	locationID, err := uuid.Parse(in.LocationId)
	if err != nil {
		return &insysproto.RescheduleTrackingResponse{}, nil
	}
	fmt.Println(locationID)

	return &insysproto.RescheduleTrackingResponse{}, nil
}
