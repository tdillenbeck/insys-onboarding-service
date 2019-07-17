package grpc

import (
	"context"
	"encoding/json"

	"weavelab.xyz/insys-onboarding-service/internal/app"
	"weavelab.xyz/monorail/shared/protorepo/dist/go/messages/insysproto"
	"weavelab.xyz/monorail/shared/protorepo/dist/go/services/insys"
	"weavelab.xyz/monorail/shared/wlib/uuid"
	"weavelab.xyz/monorail/shared/wlib/werror"
	"weavelab.xyz/monorail/shared/wlib/wgrpc"
)

var _ insys.ChiliPiperScheduleEventServer = &ChiliPiperScheduleEventServer{}

type ChiliPiperScheduleEventServer struct {
	chiliPiperScheduleEventService app.ChiliPiperScheduleEventService
}

func NewChiliPiperScheduleEventServer(s app.ChiliPiperScheduleEventService) *ChiliPiperScheduleEventServer {
	return &ChiliPiperScheduleEventServer{
		chiliPiperScheduleEventService: s,
	}
}

func (s *ChiliPiperScheduleEventServer) ByLocationID(ctx context.Context, req *insysproto.ByLocationIDChiliPiperScheduleEventRequest) (*insysproto.ByLocationIDChiliPiperScheduleEventResponse, error) {
	locationID, err := uuid.Parse(req.LocationId)
	if err != nil {
		return nil, wgrpc.Error(wgrpc.CodeInvalidArgument, werror.Wrap(err, "could not parse request location id into a uuid").Add("req.LocationID", req.LocationId))
	}

	events, err := s.chiliPiperScheduleEventService.ByLocationID(ctx, locationID)
	if err != nil {
		return nil, wgrpc.Error(wgrpc.CodeInternal, werror.Wrap(err, "error querying chili piper schedule events by location id").Add("locationID", locationID))
	}

	result, err := convertChiliPiperScheduleEventsToProto(events)
	if err != nil {
		return nil, wgrpc.Error(wgrpc.CodeInternal, werror.Wrap(err, "error converting database events into proto").Add("events", events))
	}

	return result, nil
}

func (s *ChiliPiperScheduleEventServer) Create(ctx context.Context, req *insysproto.CreateChiliPiperScheduleEventRequest) (*insysproto.CreateChiliPiperScheduleEventResponse, error) {
	return nil, nil
}

func convertChiliPiperScheduleEventsToProto(events []app.ChiliPiperScheduleEvent) (*insysproto.ByLocationIDChiliPiperScheduleEventResponse, error) {
	var result insysproto.ByLocationIDChiliPiperScheduleEventResponse

	eventsJSON, err := json.Marshal(events)
	if err != nil {
		return nil, werror.Wrap(err, "could not marshal chili piper schedule events into json")
	}

	err = json.Unmarshal(eventsJSON, &result.Events)
	if err != nil {
		return nil, werror.Wrap(err, "could not unmarshal chili piper schedule json into proto struct")
	}

	return &result, nil
}
