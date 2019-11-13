package grpc

import (
	"context"
	"encoding/json"

	"weavelab.xyz/insys-onboarding-service/internal/app"
	"weavelab.xyz/monorail/shared/protorepo/dist/go/messages/insysproto"
	"weavelab.xyz/monorail/shared/protorepo/dist/go/services/insys"
	"weavelab.xyz/monorail/shared/wlib/werror"
	"weavelab.xyz/monorail/shared/wlib/wgrpc"
)

// ensure interface is implemented correctly
var _ insys.HandoffSnapshotServer = &HandoffSnapshotServer{}

type HandoffSnapshotServer struct {
	handOffSnapshotService app.HandoffSnapshotService
}

func NewHandoffSnapshotServer(handOffSnapshotService app.HandoffSnapshotService) *HandoffSnapshotServer {
	return &HandoffSnapshotServer{
		handOffSnapshotService: handOffSnapshotService,
	}
}

func (s *HandoffSnapshotServer) CreateOrUpdate(ctx context.Context, req *insysproto.HandoffSnapshotCreateOrUpdateRequest) (*insysproto.HandoffSnapshotCreateOrUpdateResponse, error) {
	snapshot, err := convertProtoToHandoffSnapshot(*req)
	if err != nil {
		return nil, wgrpc.Error(wgrpc.CodeInvalidArgument, werror.Wrap(err, "could not convert proto to handoff snapshot").Add("req", req))
	}
	result, err := s.handOffSnapshotService.CreateOrUpdate(ctx, snapshot)
	if err != nil {
		return nil, wgrpc.Error(wgrpc.CodeInternal, werror.Wrap(err, "could not create or update handoff snapshot"))
	}

	proto, err := convertHandoffSnapshotToProto(result)
	if err != nil {
		return nil, wgrpc.Error(wgrpc.CodeInternal, werror.Wrap(err, "could not convert handoff snapshot to proto"))
	}
	return &proto, nil
}

func convertProtoToHandoffSnapshot(proto insysproto.HandoffSnapshotCreateOrUpdateRequest) (app.HandoffSnapshot, error) {
	var result app.HandoffSnapshot

	snapshotJSON, err := json.Marshal(proto.HandoffSnapshot)
	if err != nil {
		return app.HandoffSnapshot{}, werror.Wrap(err, "could not marshal handoff snapshot into json").Add("proto", proto)
	}

	err = json.Unmarshal(snapshotJSON, &result)
	if err != nil {
		return app.HandoffSnapshot{}, werror.Wrap(err, "could not unmarshal handoff snapshot json into struct").Add("snapshotJSON", string(snapshotJSON))
	}

	return result, nil
}

func convertHandoffSnapshotToProto(snapshot app.HandoffSnapshot) (insysproto.HandoffSnapshotCreateOrUpdateResponse, error) {
	var result insysproto.HandoffSnapshotCreateOrUpdateResponse

	snapshotJSON, err := json.Marshal(snapshot)
	if err != nil {
		return insysproto.HandoffSnapshotCreateOrUpdateResponse{}, werror.Wrap(err, "could not marshal handoff snapshot into json").Add("snapshot", snapshot)
	}

	err = json.Unmarshal(snapshotJSON, &result.HandoffSnapshot)
	if err != nil {
		return insysproto.HandoffSnapshotCreateOrUpdateResponse{}, werror.Wrap(err, "could not unmarshal handoff snapshot json into proto").Add("snapshotJSON", string(snapshotJSON))
	}

	return result, nil
}
