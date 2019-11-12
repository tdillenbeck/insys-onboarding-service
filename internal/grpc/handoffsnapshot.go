package grpc

import (
	"context"
	"time"

	"weavelab.xyz/insys-onboarding-service/internal/app"
	"weavelab.xyz/monorail/shared/go-utilities/null"
	"weavelab.xyz/monorail/shared/protorepo/dist/go/messages/insysproto"
	"weavelab.xyz/monorail/shared/protorepo/dist/go/services/insys"
	"weavelab.xyz/monorail/shared/wlib/uuid"
	"weavelab.xyz/monorail/shared/wlib/werror"
	"weavelab.xyz/monorail/shared/wlib/wgrpc"
)

var _ insys.HandOffSnapshotServer = &HandOffSnapshotServer{}

type HandOffSnapshotServer struct {
	handOffSnapshotService app.HandOffSnapshotService
}

func NewHandOffSnapshotServer(hoss app.HandOffSnapshotService) *HandOffSnapshotServer {
	return &HandOffSnapshotServer{
		handOffSnapshotService: hoss,
	}
}

func (s *HandOffSnapshotServer) CreateOrUpdate(ctx context.Context, req *insysproto.HandOffSnapshotCreateOrUpdateRequest) (*insysproto.HandOffSnapshotCreateOrUpdateResponse, error) {
	snapshot, err := convertProtoToHandOffSnapshot(*req)
	if err != nil {
		return nil, wgrpc.Error(wgrpc.CodeInvalidArgument, werror.Wrap(err, "could not convert proto to hand-off snapshot").Add("req", req))
	}
	result, err := s.handOffSnapshotService.CreateOrUpdate(ctx, &snapshot)
	if err != nil {
		return nil, wgrpc.Error(wgrpc.CodeInternal, werror.Wrap(err, "could not create or update hand-off snapshot"))
	}

	proto := convertHandOffSnapshotToProto(*result)

	return &proto, nil
}

func convertProtoToHandOffSnapshot(proto insysproto.HandOffSnapshotCreateOrUpdateRequest) (app.HandOffSnapshot, error) {
	var onboardersLocationID uuid.UUID
	var csatRecipientUserID null.UUID
	var csatSentAt null.Time
	var err error

	onboardersLocationID, err = uuid.Parse(proto.HandoffSnapshot.OnboardersLocationId)
	if err != nil {
		return app.HandOffSnapshot{}, err
	}

	if proto.HandoffSnapshot.CsatRecipientUserId != "" {
		csatRecipientUserID, err = null.NewUUID(proto.HandoffSnapshot.CsatRecipientUserId)
		if err != nil {
			return app.HandOffSnapshot{}, err
		}
	}

	if proto.HandoffSnapshot.CsatSentAt != "" {
		parsedCsatSentAt, err := time.Parse(time.RFC3339, proto.HandoffSnapshot.CsatSentAt)
		if err != nil {
			return app.HandOffSnapshot{}, err
		}
		csatSentAt = null.NewTime(parsedCsatSentAt.UTC())
	}

	return app.HandOffSnapshot{
		OnboardersLocationID:                      onboardersLocationID,
		CustomerSatisfactionSurveyRecipientUserID: csatRecipientUserID,
		CustomerSatisfactionSurveySentAt:          csatSentAt,
	}, nil
}

func convertHandOffSnapshotToProto(snapshot app.HandOffSnapshot) insysproto.HandOffSnapshotCreateOrUpdateResponse {
	csatSentAt := ""
	if snapshot.CustomerSatisfactionSurveySentAt.Valid {
		csatSentAt = snapshot.CustomerSatisfactionSurveySentAt.Time.Format(time.RFC3339)
	}

	return insysproto.HandOffSnapshotCreateOrUpdateResponse{
		HandoffSnapshot: &insysproto.HandOffSnapshotRecord{
			Id:                   snapshot.ID.String(),
			OnboardersLocationId: snapshot.OnboardersLocationID.String(),
			CsatRecipientUserId:  snapshot.CustomerSatisfactionSurveyRecipientUserID.String(),
			CsatSentAt:           csatSentAt,
			CreatedAt:            snapshot.CreatedAt.Format(time.RFC3339),
			UpdatedAt:            snapshot.UpdatedAt.Format(time.RFC3339),
		},
	}
}
