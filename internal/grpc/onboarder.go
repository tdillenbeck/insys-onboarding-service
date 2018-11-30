package grpc

import (
	"context"
	"time"

	"github.com/golang/protobuf/ptypes"

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
		return nil, wgrpc.Error(wgrpc.CodeInternal, werror.New("could not parse proto request to internal struct"))
	}

	onb, err := s.onboarderService.CreateOrUpdate(ctx, onboarder)
	if err != nil {
		return nil, wgrpc.Error(wgrpc.CodeInternal, werror.New("error inserting or updating data in the database"))
	}

	result, err := convertOnboarderToProto(onb)
	if err != nil {
		return nil, wgrpc.Error(wgrpc.CodeInternal, werror.New("could not parse internal struct to proto"))
	}

	return result, nil
}

func (s *OnboarderServer) ReadByUserID(ctx context.Context, req *insysproto.Onboarder) (*insysproto.Onboarder, error) {
	userID, err := req.UserID.UUID()
	if err != nil {
		return nil, wgrpc.Error(wgrpc.CodeInternal, werror.Wrap(err, "could not parse request user id into a uuid").Add("req.UserID", req.UserID))
	}

	onb, err := s.onboarderService.ReadByUserID(ctx, userID)
	if err != nil {
		return nil, wgrpc.Error(wgrpc.CodeInternal, werror.New("error reading onboarder by user id from the database"))
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

func convertOnboarderToProto(onb *app.Onboarder) (*insysproto.Onboarder, error) {
	createdAt, err := ptypes.TimestampProto(onb.CreatedAt)
	if err != nil {
		return nil, werror.Wrap(err, "could not convert onboarder created at time")
	}
	updatedAt, err := ptypes.TimestampProto(onb.UpdatedAt)
	if err != nil {
		return nil, werror.Wrap(err, "could not convert onboarder updated at time")
	}

	return &insysproto.Onboarder{
		ID:                           sharedproto.UUIDToProto(onb.ID),
		UserID:                       sharedproto.UUIDToProto(onb.UserID),
		ScheduleCustomizationLink:    onb.ScheduleCustomizationLink.String(),
		SchedulePortingLink:          onb.SchedulePortingLink.String(),
		ScheduleNetworkLink:          onb.ScheduleNetworkLink.String(),
		ScheduleSoftwareInstallLink:  onb.ScheduleSoftwareInstallLink.String(),
		SchedulePhoneInstallLink:     onb.SchedulePhoneInstallLink.String(),
		ScheduleSoftwareTrainingLink: onb.ScheduleSoftwareTrainingLink.String(),
		SchedulePhoneTrainingLink:    onb.SchedulePhoneTrainingLink.String(),
		CreatedAt:                    createdAt,
		UpdatedAt:                    updatedAt,
	}, nil
}
