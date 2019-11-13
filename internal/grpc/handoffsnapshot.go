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
var _ insys.HandOffSnapshotServer = &HandOffSnapshotServer{}

type HandOffSnapshotServer struct {
	handOffSnapshotService app.HandOffSnapshotService
}

func NewHandOffSnapshotServer(handOffSnapshotService app.HandOffSnapshotService) *HandOffSnapshotServer {
	return &HandOffSnapshotServer{
		handOffSnapshotService: handOffSnapshotService,
	}
}

func (s *HandOffSnapshotServer) CreateOrUpdate(ctx context.Context, req *insysproto.HandOffSnapshotCreateOrUpdateRequest) (*insysproto.HandOffSnapshotCreateOrUpdateResponse, error) {
	snapshot, err := convertProtoToHandOffSnapshot(*req)
	if err != nil {
		return nil, wgrpc.Error(wgrpc.CodeInvalidArgument, werror.Wrap(err, "could not convert proto to hand-off snapshot").Add("req", req))
	}
	result, err := s.handOffSnapshotService.CreateOrUpdate(ctx, snapshot)
	if err != nil {
		return nil, wgrpc.Error(wgrpc.CodeInternal, werror.Wrap(err, "could not create or update hand-off snapshot"))
	}

	proto, err := convertHandOffSnapshotToProto(result)
	if err != nil {
		return nil, wgrpc.Error(wgrpc.CodeInternal, werror.Wrap(err, "could not convert hand-off snapshot to proto"))
	}
	return &proto, nil
}

func convertProtoToHandOffSnapshot(proto insysproto.HandOffSnapshotCreateOrUpdateRequest) (app.HandOffSnapshot, error) {
	var result app.HandOffSnapshot

	snapshotJSON, err := json.Marshal(proto.HandoffSnapshot)
	if err != nil {
		return app.HandOffSnapshot{}, werror.Wrap(err, "could not marshal hand-off snapshot into json").Add("proto", proto)
	}

	err = json.Unmarshal(snapshotJSON, &result)
	if err != nil {
		return app.HandOffSnapshot{}, werror.Wrap(err, "could not unmarshal hand-off snapshot json into struct").Add("snapshotJSON", string(snapshotJSON))
	}

	return result, nil
}

func convertHandOffSnapshotToProto(snapshot app.HandOffSnapshot) (insysproto.HandOffSnapshotCreateOrUpdateResponse, error) {
	var result insysproto.HandOffSnapshotCreateOrUpdateResponse

	snapshotJSON, err := json.Marshal(snapshot)
	if err != nil {
		return insysproto.HandOffSnapshotCreateOrUpdateResponse{}, werror.Wrap(err, "could not marshal hand-off snapshot into json").Add("snapshot", snapshot)
	}

	err = json.Unmarshal(snapshotJSON, &result.HandoffSnapshot)
	if err != nil {
		return insysproto.HandOffSnapshotCreateOrUpdateResponse{}, werror.Wrap(err, "could not unmarshal hand-off snapshot json into proto").Add("snapshotJSON", string(snapshotJSON))
	}

	return result, nil
}
