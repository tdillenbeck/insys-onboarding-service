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
	event, err := convertProtoToChiliPiperScheduleEvent(req)
	if err != nil {
		return nil, wgrpc.Error(wgrpc.CodeInvalidArgument, werror.Wrap(err, "could not convert request into internal chili piper schedule events truct").Add("req", req))
	}

	createResponse, err := s.chiliPiperScheduleEventService.Create(ctx, event)
	if err != nil {
		return nil, wgrpc.Error(wgrpc.CodeInternal, werror.Wrap(err, "error creating chili piper schedule event").Add("event", event))
	}

	result, err := convertChiliPiperScheduleEventToProto(createResponse)
	if err != nil {
		return nil, wgrpc.Error(wgrpc.CodeInternal, werror.Wrap(err, "error converting chili piper schedule event into proto").Add("createResponse", createResponse))
	}

	return result, nil
}

func (s *ChiliPiperScheduleEventServer) Update(ctx context.Context, req *insysproto.UpdateChiliPiperScheduleEventRequest) (*insysproto.UpdateChiliPiperScheduleEventResponse, error) {
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

func convertChiliPiperScheduleEventToProto(event *app.ChiliPiperScheduleEvent) (*insysproto.CreateChiliPiperScheduleEventResponse, error) {
	var result insysproto.CreateChiliPiperScheduleEventResponse

	eventJSON, err := json.Marshal(event)
	if err != nil {
		return nil, werror.Wrap(err, "could not marshal chili piper schedule event into json").Add("event", event)
	}

	err = json.Unmarshal(eventJSON, &result.Event)
	if err != nil {
		return nil, werror.Wrap(err, "could not unmarshal chili piper schedule json into proto struct").Add("eventJSON", string(eventJSON))
	}

	return &result, nil
}

func convertProtoToChiliPiperScheduleEvent(in *insysproto.CreateChiliPiperScheduleEventRequest) (*app.ChiliPiperScheduleEvent, error) {
	var result app.ChiliPiperScheduleEvent

	inJSON, err := json.Marshal(in.Event)
	if err != nil {
		return nil, werror.Wrap(err, "could not marshal input proto to json").Add("in", in)
	}

	err = json.Unmarshal(inJSON, &result)
	if err != nil {
		return nil, werror.Wrap(err, "could not unmarshal json into iternal chili piper schedule event").Add("inJSON", string(inJSON))
	}

	return &result, nil
}
