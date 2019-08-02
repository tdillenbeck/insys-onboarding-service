package grpc

import (
	"context"
	"database/sql"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/empty"

	"weavelab.xyz/insys-onboarding-service/internal/app"
	"weavelab.xyz/monorail/shared/go-utilities/null"
	"weavelab.xyz/monorail/shared/protorepo/dist/go/messages/insysproto"
	"weavelab.xyz/monorail/shared/protorepo/dist/go/messages/sharedproto"
	"weavelab.xyz/monorail/shared/protorepo/dist/go/services/insys"
	"weavelab.xyz/monorail/shared/wlib/uuid"
	"weavelab.xyz/monorail/shared/wlib/werror"
	"weavelab.xyz/monorail/shared/wlib/wgrpc"
)

var _ insys.OnboarderServer = &OnboarderServer{}

type OnboarderServer struct {
	onboarderService app.OnboarderService
}

func NewOnboarderServer(onbs app.OnboarderService) *OnboarderServer {
	return &OnboarderServer{
		onboarderService: onbs,
	}
}

func (s *OnboarderServer) CreateOrUpdate(ctx context.Context, req *insysproto.Onboarder) (*insysproto.Onboarder, error) {
	onboarder, err := convertProtoToOnboarder(req)
	if err != nil {
		return nil, wgrpc.Error(wgrpc.CodeInvalidArgument, werror.New("could not parse proto request to internal struct"))
	}

	onb, err := s.onboarderService.CreateOrUpdate(ctx, onboarder)
	if err != nil {
		return nil, wgrpc.Error(wgrpc.CodeInternal, werror.New("error inserting or updating data in the database").Add("onboarder", onboarder))
	}

	result, err := convertOnboarderToProto(onb)
	if err != nil {
		return nil, wgrpc.Error(wgrpc.CodeInternal, werror.New("could not parse internal struct to proto"))
	}

	return result, nil
}

func (s *OnboarderServer) Delete(ctx context.Context, req *insysproto.DeleteOnboarderRequest) (*empty.Empty, error) {
	id, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, wgrpc.Error(wgrpc.CodeInvalidArgument, werror.New("could not parse id for delete into uuid.UUID").Add("req.Id", req.Id))
	}

	err = s.onboarderService.Delete(ctx, id)
	if err != nil {
		return nil, wgrpc.Error(wgrpc.CodeInternal, werror.New("could not soft delete onboarder").Add("req.Id", req.Id))
	}

	return &empty.Empty{}, nil
}

func (s *OnboarderServer) ListOnboarders(ctx context.Context, req *insysproto.ListOnboardersRequest) (*insysproto.ListOnboardersResponse, error) {
	onboarders, err := s.onboarderService.List(ctx)
	if err != nil {
		return nil, wgrpc.Error(wgrpc.CodeInternal, werror.New("could not list onboarders from db"))
	}

	result, err := convertOnboardersToListProto(onboarders)
	if err != nil {
		return nil, wgrpc.Error(wgrpc.CodeInternal, werror.New("could not convert internal onboarder into proto struct for list"))
	}

	return result, nil
}

func (s *OnboarderServer) ReadByUserID(ctx context.Context, req *insysproto.Onboarder) (*insysproto.Onboarder, error) {
	userID, err := req.UserID.UUID()
	if err != nil {
		return nil, wgrpc.Error(wgrpc.CodeInvalidArgument, werror.Wrap(err, "could not parse request user id into a uuid").Add("req.UserID", req.UserID))
	}

	onb, err := s.onboarderService.ReadByUserID(ctx, userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, werror.Wrap(err, "no rows returned for user").Add("userID", userID).SetCode(wgrpc.CodeNotFound)
		}
		werr := werror.Cast(err)
		return nil, wgrpc.Error(werr.Code(), werror.Wrap(werr, "error reading onboarder by user id from the database").Add("userID", userID))
	}

	result, err := convertOnboarderToProto(onb)
	if err != nil {
		return nil, wgrpc.Error(wgrpc.CodeInternal, werror.New("could not parse internal struct to proto"))
	}

	return result, nil
}

func convertProtoToOnboarder(proto *insysproto.Onboarder) (*app.Onboarder, error) {
	var err error
	var id uuid.UUID
	var userID uuid.UUID
	var createdAt time.Time
	var updatedAt time.Time

	if proto.CreatedAt != nil {
		createdAt, err = ptypes.Timestamp(proto.CreatedAt)
		if err != nil {
			return nil, werror.Wrap(err, "could not convert proto CreatedAt (*timestamp.Timestamp to time.Time")
		}
	}

	if proto.UpdatedAt != nil {
		updatedAt, err = ptypes.Timestamp(proto.UpdatedAt)
		if err != nil {
			return nil, werror.Wrap(err, "could not convert proto UpdatedAt (*timestamp.Timestamp to time.Time")
		}
	}

	if proto.ID != nil {
		id, err = proto.ID.UUID()
		if err != nil {
			return nil, wgrpc.Error(wgrpc.CodeInternal, werror.Wrap(err, "could not parse proto id into a uuid").Add("proto.ID", proto.ID))
		}
	}

	if proto.UserID != nil {
		userID, err = proto.UserID.UUID()
		if err != nil {
			return nil, wgrpc.Error(wgrpc.CodeInternal, werror.Wrap(err, "could not parse proto userID into a uuid").Add("proto.UserID", proto.UserID))
		}
	}

	return &app.Onboarder{
		ID:                           id,
		UserID:                       userID,
		SalesforceUserID:             null.NewString(proto.SalesforceUserID),
		ScheduleCustomizationLink:    null.NewString(proto.ScheduleCustomizationLink),
		SchedulePortingLink:          null.NewString(proto.SchedulePortingLink),
		ScheduleNetworkLink:          null.NewString(proto.ScheduleNetworkLink),
		ScheduleSoftwareInstallLink:  null.NewString(proto.ScheduleSoftwareInstallLink),
		SchedulePhoneInstallLink:     null.NewString(proto.SchedulePhoneInstallLink),
		ScheduleSoftwareTrainingLink: null.NewString(proto.ScheduleSoftwareTrainingLink),
		SchedulePhoneTrainingLink:    null.NewString(proto.SchedulePhoneTrainingLink),
		CreatedAt:                    createdAt,
		UpdatedAt:                    updatedAt,
	}, nil
}

func convertOnboardersToListProto(onboarders []app.Onboarder) (*insysproto.ListOnboardersResponse, error) {
	var result insysproto.ListOnboardersResponse

	for _, o := range onboarders {
		proto, err := convertOnboarderToProto(&o)
		if err != nil {
			return nil, werror.Wrap(err, "could not convert onboarder to proto").Add("o.ID", o.ID)
		}
		result.Onboarders = append(result.Onboarders, proto)
	}

	return &result, nil
}

func convertOnboarderToProto(onb *app.Onboarder) (*insysproto.Onboarder, error) {
	createdAt, err := ptypes.TimestampProto(onb.CreatedAt)
	if err != nil {
		return nil, werror.Wrap(err, "could not convert onboarder created at time")
	}
	updatedAt, err := ptypes.TimestampProto(onb.UpdatedAt)
	if err != nil {
		return nil, werror.Wrap(err, "could not convert onboarder updated at time")
	}

	deletedAt, err := ptypes.TimestampProto(onb.DeletedAt.Time)
	if err != nil {
		return nil, werror.Wrap(err, "could not convert onboarder deleted at time")
	}

	return &insysproto.Onboarder{
		ID:                           sharedproto.UUIDToProto(onb.ID),
		UserID:                       sharedproto.UUIDToProto(onb.UserID),
		SalesforceUserID:             onb.SalesforceUserID.String(),
		ScheduleCustomizationLink:    onb.ScheduleCustomizationLink.String(),
		SchedulePortingLink:          onb.SchedulePortingLink.String(),
		ScheduleNetworkLink:          onb.ScheduleNetworkLink.String(),
		ScheduleSoftwareInstallLink:  onb.ScheduleSoftwareInstallLink.String(),
		SchedulePhoneInstallLink:     onb.SchedulePhoneInstallLink.String(),
		ScheduleSoftwareTrainingLink: onb.ScheduleSoftwareTrainingLink.String(),
		SchedulePhoneTrainingLink:    onb.SchedulePhoneTrainingLink.String(),
		CreatedAt:                    createdAt,
		UpdatedAt:                    updatedAt,
		DeletedAt:                    deletedAt,
	}, nil
}
