package grpc

import (
	"context"
	"database/sql"
	"encoding/json"
	"strings"
	"weavelab.xyz/monorail/shared/wlib/wlog"
	"weavelab.xyz/monorail/shared/wlib/wlog/tag"

	"weavelab.xyz/monorail/shared/wlib/uuid"

	"weavelab.xyz/insys-onboarding-service/internal/app"
	"weavelab.xyz/monorail/shared/protorepo/dist/go/messages/insysproto"
	"weavelab.xyz/monorail/shared/protorepo/dist/go/services/insys"
	"weavelab.xyz/monorail/shared/wlib/werror"
	"weavelab.xyz/monorail/shared/wlib/wgrpc"
)

// ensure interface is implemented correctly
var _ insys.HandoffSnapshotServer = &HandoffSnapshotServer{}

type HandoffSnapshotServer struct {
	handoffSnapshotService app.HandoffSnapshotService
}

func NewHandoffSnapshotServer(handoffSnapshotService app.HandoffSnapshotService) *HandoffSnapshotServer {
	return &HandoffSnapshotServer{
		handoffSnapshotService: handoffSnapshotService,
	}
}

func (s *HandoffSnapshotServer) CreateOrUpdate(ctx context.Context, req *insysproto.HandoffSnapshotCreateOrUpdateRequest) (*insysproto.HandoffSnapshotResponse, error) {
	snapshot, err := convertProtoToHandoffSnapshot(*req)
	if err != nil {
		return nil, wgrpc.Error(wgrpc.CodeInvalidArgument, werror.Wrap(err, "could not convert proto to handoff snapshot").Add("req", req))
	}

	onboardersLocationID := snapshot.OnboardersLocationID
	result, err := s.handoffSnapshotService.ReadByOnboardersLocationID(ctx, onboardersLocationID)
	if err != nil && err != sql.ErrNoRows {
		return nil, wgrpc.Error(wgrpc.CodeInternal, werror.Wrap(err, "could not read onboardersLocationID"))
	}
	if result.HandedOffAt.Valid {
		return nil, wgrpc.Error(wgrpc.CodePermissionDenied, werror.New("handoff has already been submitted"))
	}

	result, err = s.handoffSnapshotService.CreateOrUpdate(ctx, snapshot)
	if err != nil {
		return nil, wgrpc.Error(wgrpc.CodeInternal, werror.Wrap(err, "could not create or update handoff snapshot"))
	}

	proto, err := convertHandoffSnapshotToProto(result)
	if err != nil {
		return nil, wgrpc.Error(wgrpc.CodeInternal, werror.Wrap(err, "could not convert handoff snapshot to proto"))
	}
	return &proto, nil
}

func (s *HandoffSnapshotServer) ReadByOnboardersLocationID(ctx context.Context, req *insysproto.HandoffSnapshotReadRequest) (*insysproto.HandoffSnapshotResponse, error) {
	onboardersLocationID, err := uuid.Parse(req.OnboardersLocationId)
	if err != nil {
		return nil, wgrpc.Error(wgrpc.CodeInvalidArgument, werror.Wrap(err, "error parsing: ").Add("req.OnboardersLocationId", req.OnboardersLocationId))
	}

	result, err := s.handoffSnapshotService.ReadByOnboardersLocationID(ctx, onboardersLocationID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, wgrpc.Error(wgrpc.CodeNotFound, werror.Wrap(err, "no handoff snapshot found for onboarders location id").Add("onboardersLocationID", onboardersLocationID))
		}
		return nil, werror.Wrap(err, "failed to get porting data")
	}

	proto, err := convertHandoffSnapshotToProto(result)
	if err != nil {
		return nil, wgrpc.Error(wgrpc.CodeInternal, werror.Wrap(err, "could not convert handoff snapshot to proto"))
	}
	return &proto, nil
}

func (s *HandoffSnapshotServer) SubmitCSAT(ctx context.Context, req *insysproto.SubmitCSATRequest) (*insysproto.HandoffSnapshotResponse, error) {
	onboardersLocationID, err := uuid.Parse(req.OnboardersLocationId)
	if err != nil {
		return nil, wgrpc.Error(wgrpc.CodeInvalidArgument, werror.Wrap(err, "error parsing: ").Add("req.OnboardersLocationId", req.OnboardersLocationId))
	}

	wlog.InfoC(ctx, "submitting CSAT", tag.String("OnboardersLocationId", req.OnboardersLocationId), tag.String("CsatRecipientUserEmail", req.CsatRecipientUserEmail))
	result, err := s.handoffSnapshotService.ReadByOnboardersLocationID(ctx, onboardersLocationID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, wgrpc.Error(wgrpc.CodeNotFound, werror.Wrap(err, "no handoff snapshot found for onboarders location id").Add("onboardersLocationID", onboardersLocationID))
		}
		return nil, werror.Wrap(err, "failed to get porting data")
		wlog.InfoC(ctx, "submitting CSAT", tag.String("OnboardersLocationId", req.OnboardersLocationId), tag.String("CsatRecipientUserEmail", req.CsatRecipientUserEmail))

	}

	missingFields := validateCsatSubmit(result)
	if missingFields != "" {
		return nil, wgrpc.Error(wgrpc.CodeInternal, werror.New("missing csat fields").Add("missing_fields", missingFields))
	}

	result, err = s.handoffSnapshotService.SubmitCSAT(ctx, onboardersLocationID, req.CsatRecipientUserEmail)
	if err != nil {
		return nil, wgrpc.Error(wgrpc.CodeInternal, werror.Wrap(err, "error submitting csat"))
	}

	proto, err := convertHandoffSnapshotToProto(result)
	if err != nil {
		return nil, wgrpc.Error(wgrpc.CodeInvalidArgument, werror.Wrap(err, "could not convert handoff snapshot to proto"))
	}
	return &proto, nil
}

func (s *HandoffSnapshotServer) SubmitHandoff(ctx context.Context, req *insysproto.SubmitHandoffRequest) (*insysproto.HandoffSnapshotResponse, error) {
	onboardersLocationID, err := uuid.Parse(req.OnboardersLocationId)
	if err != nil {
		return nil, wgrpc.Error(wgrpc.CodeInvalidArgument, werror.Wrap(err, "error parsing: ").Add("req.OnboardersLocationId", req.OnboardersLocationId))
	}

	result, err := s.handoffSnapshotService.ReadByOnboardersLocationID(ctx, onboardersLocationID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, wgrpc.Error(wgrpc.CodeNotFound, werror.Wrap(err, "no handoff snapshot found for onboarders location id").Add("onboardersLocationID", onboardersLocationID))
		}
		return nil, werror.Wrap(err, "failed to get handoff snapshot")
	}

	if result.HandedOffAt.Valid {
		return nil, wgrpc.Error(wgrpc.CodeInternal, werror.New("this snapshot has already been handed off and cannot be updated"))
	}

	missingFields := validateHandoffSubmit(result)
	if missingFields != "" {
		return nil, wgrpc.Error(wgrpc.CodeInternal, werror.New("missing handoff fields: "+missingFields))
	}

	result, err = s.handoffSnapshotService.SubmitHandoff(ctx, onboardersLocationID)
	if err != nil {
		return nil, wgrpc.Error(wgrpc.CodeInternal, werror.Wrap(err, "error submitting handoff"))
	}

	proto, err := convertHandoffSnapshotToProto(result)
	if err != nil {
		return nil, wgrpc.Error(wgrpc.CodeInvalidArgument, werror.Wrap(err, "could not convert handoff snapshot to proto"))
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

func convertHandoffSnapshotToProto(snapshot app.HandoffSnapshot) (insysproto.HandoffSnapshotResponse, error) {
	var result insysproto.HandoffSnapshotResponse

	snapshotJSON, err := json.Marshal(snapshot)
	if err != nil {
		return insysproto.HandoffSnapshotResponse{}, werror.Wrap(err, "could not marshal handoff snapshot into json").Add("snapshot", snapshot)
	}

	err = json.Unmarshal(snapshotJSON, &result.HandoffSnapshot)
	if err != nil {
		return insysproto.HandoffSnapshotResponse{}, werror.Wrap(err, "could not unmarshal handoff snapshot json into proto").Add("snapshotJSON", string(snapshotJSON))
	}

	return result, nil
}

func validateCsatSubmit(snapshot app.HandoffSnapshot) string {
	var missingFields []string

	if !snapshot.PointOfContactEmail.Valid || snapshot.PointOfContactEmail.String() == "" {
		missingFields = append(missingFields, "point_of_contact_email")
	}

	if missingFields == nil {
		return ""
	}

	return strings.Join(missingFields, ", ")
}

func validateHandoffSubmit(snapshot app.HandoffSnapshot) string {
	var missingFields []string

	if !snapshot.CSATSentAt.Valid {
		missingFields = append(missingFields, "csat_sent_at")
	}
	if !snapshot.PointOfContactEmail.Valid || snapshot.PointOfContactEmail.String() == "" {
		missingFields = append(missingFields, "point_of_contact_email")
	}
	if !snapshot.ReasonForPurchase.Valid || snapshot.ReasonForPurchase.String() == "" {
		missingFields = append(missingFields, "reason_for_purchase")
	}
	if !snapshot.BillingNotes.Valid || snapshot.BillingNotes.String() == "" {
		missingFields = append(missingFields, "billing_notes")
	}
	if !snapshot.Notes.Valid || snapshot.Notes.String() == "" {
		missingFields = append(missingFields, "notes")
	}

	if missingFields == nil {
		return ""
	}

	return strings.Join(missingFields, ", ")
}
