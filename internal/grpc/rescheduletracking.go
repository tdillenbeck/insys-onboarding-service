package grpc

import (
	"context"
	"fmt"

	"github.com/golang/protobuf/ptypes"

	"weavelab.xyz/insys-onboarding-service/internal/app"
	"weavelab.xyz/monorail/shared/protorepo/dist/go/messages/insysproto"
	"weavelab.xyz/monorail/shared/protorepo/dist/go/messages/sharedproto"
	"weavelab.xyz/monorail/shared/protorepo/dist/go/services/insys"
	"weavelab.xyz/monorail/shared/wlib/werror"
	"weavelab.xyz/monorail/shared/wlib/wgrpc"
)

var _ insys.RescheduleTrackingEventServer = &RescheduleTrackingEventServer{}

type RescheduleTrackingEventServer struct {
	rescheduleTrackingEventService app.RescheduleTrackingEventService
}

func NewRescheduleEventServer(rescheduleTrackingService app.RescheduleTrackingEventService) *RescheduleTrackingEventServer {
	return &RescheduleTrackingEventServer{
		rescheduleTrackingEventService: rescheduleTrackingService,
	}
}

func (r *RescheduleTrackingEventServer) ReadRescheduleTracking(ctx context.Context, req *insysproto.RescheduleTrackingRequest) (*insysproto.RescheduleTrackingResponse, error) {

	rescheduleResponse, err := r.rescheduleTrackingEventService.ReadRescheduleTracking(ctx, req)
	if err != nil {
		return nil, wgrpc.Error(wgrpc.CodeInternal, werror.Wrap(err, "error reading reschedule tracking event"))
	}
	fmt.Println(rescheduleResponse)

	return &insysproto.RescheduleTrackingResponse{}, nil
}

func convertRescheduleTrackingToProto(r *app.RescheduleTracking) (*insysproto.RescheduleTrackingResponse, error) {

	if r == nil {
		return nil, werror.New("could not convert nil to reschedule tracking")
	}

	rescheduledAt, err := ptypes.TimestampProto(r.RescheuleEventsCalculatedAt)
	if err != nil {
		return nil, werror.New("could not convert reschedule tracking rescheduled at time")
	}

	createdAt, err := ptypes.TimestampProto(r.CreatedAt)
	if err != nil {
		return nil, werror.New("could not convert reschedule tracking created at time")
	}
	updatedAt, err := ptypes.TimestampProto(r.UpdatedAt)
	if err != nil {
		return nil, werror.New("could not convert reschedule tracking updated at time")
	}

	return &insysproto.RescheduleTrackingResponse{
		LocationId:                    sharedproto.UUIDToProto(r.LocationID),
		EventType:                     r.EventType,
		RescheduledEventsCount:        0, //r.RescheduledEventsCount,
		RescheduledEventsCalculatedAt: rescheduledAt,
		CreatedAt:                     createdAt,
		UpdatedAt:                     updatedAt,
	}, nil

}
