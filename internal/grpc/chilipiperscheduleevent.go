package grpc

import (
	"context"

	"weavelab.xyz/monorail/shared/protorepo/dist/go/messages/insysproto"
	"weavelab.xyz/monorail/shared/protorepo/dist/go/services/insys"
)

var _ insys.ChiliPiperScheduleEventServer = &ChiliPiperScheduleEventServer{}

type ChiliPiperScheduleEventServer struct{}

func NewChiliPiperScheduleEventServer() *ChiliPiperScheduleEventServer {
	return &ChiliPiperScheduleEventServer{}
}

func (s *ChiliPiperScheduleEventServer) ByLocationID(ctx context.Context, req *insysproto.ByLocationIDChiliPiperScheduleEventRequest) (*insysproto.ByLocationIDChiliPiperScheduleEventResponse, error) {
	return nil, nil
}

func (s *ChiliPiperScheduleEventServer) Create(ctx context.Context, req *insysproto.CreateChiliPiperScheduleEventRequest) (*insysproto.CreateChiliPiperScheduleEventResponse, error) {
	return nil, nil
}
