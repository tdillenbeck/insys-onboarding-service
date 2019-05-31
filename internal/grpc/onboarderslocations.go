package grpc

import (
	"context"
	"time"

	"github.com/golang/protobuf/ptypes"

	"weavelab.xyz/insys-onboarding-service/internal/app"

	"weavelab.xyz/monorail/shared/protorepo/dist/go/messages/insysproto"
	"weavelab.xyz/monorail/shared/protorepo/dist/go/messages/sharedproto"
	"weavelab.xyz/monorail/shared/protorepo/dist/go/services/insys"
	"weavelab.xyz/monorail/shared/wlib/uuid"
	"weavelab.xyz/monorail/shared/wlib/werror"
	"weavelab.xyz/monorail/shared/wlib/wgrpc"
)

var _ insys.OnboardersLocationServer = &OnboardersLocationServer{}

type OnboardersLocationServer struct {
	onboardersLocationService app.OnboardersLocationService
	taskInstanceService       app.TaskInstanceService
}

func NewOnboardersLocationServer(onbls app.OnboardersLocationService, tis app.TaskInstanceService) *OnboardersLocationServer {
	return &OnboardersLocationServer{
		onboardersLocationService: onbls,
		taskInstanceService:       tis,
	}
}

// CreateOrUpdate is responsible or assigning an onboarder to a location. This will:
//  1. create or update the record in the onboarders location table
//  2. check if the location that the onboarder is being assigned has already been setup with tasks
//  3. if there are tasks, update their links to the new onboarder
func (s *OnboardersLocationServer) CreateOrUpdate(ctx context.Context, req *insysproto.OnboardersLocation) (*insysproto.OnboardersLocation, error) {
	onboardersLocation, err := convertProtoToOnboardersLocation(req)
	if err != nil {
		return nil, wgrpc.Error(wgrpc.CodeInternal, werror.New("could not parse proto request to internal struct"))
	}

	onbl, err := s.onboardersLocationService.CreateOrUpdate(ctx, onboardersLocation)
	if err != nil {
		return nil, wgrpc.Error(wgrpc.CodeInternal, werror.New("error inserting or updating data in the database"))
	}

	taskInstances, err := s.taskInstanceService.ByLocationID(ctx, onbl.LocationID)
	if err != nil {
		return nil, wgrpc.Error(wgrpc.CodeInternal, werror.New("error looking up task instances for location"))
	}

	if len(taskInstances) > 0 {
		err = s.taskInstanceService.SyncTaskInstanceLinksFromOnboarderLinks(ctx, onbl.LocationID)
		if err != nil {
			return nil, wgrpc.Error(wgrpc.CodeInternal, werror.New("error updating the existing tasks"))
		}
	}

	result, err := convertOnboardersLocationToProto(onbl)
	if err != nil {
		return nil, wgrpc.Error(wgrpc.CodeInternal, werror.New("could not parse internal struct to proto"))
	}

	return result, nil
}

func (s *OnboardersLocationServer) ReadByLocationID(ctx context.Context, req *insysproto.OnboardersLocation) (*insysproto.OnboardersLocation, error) {
	locationID, err := req.LocationID.UUID()
	if err != nil {
		return nil, wgrpc.Error(wgrpc.CodeInternal, werror.Wrap(err, "could not parse request location id into a uuid").Add("req.LocationID", req.LocationID))
	}

	onbl, err := s.onboardersLocationService.ReadByLocationID(ctx, locationID)
	if err != nil {
		return nil, wgrpc.Error(wgrpc.CodeInternal, werror.New("error reading onboarders locaiton by location id from the database"))
	}

	result, err := convertOnboardersLocationToProto(onbl)
	if err != nil {
		return nil, wgrpc.Error(wgrpc.CodeInternal, werror.New("could not parse internal struct to proto"))
	}

	return result, nil
}

func convertProtoToOnboardersLocation(proto *insysproto.OnboardersLocation) (*app.OnboardersLocation, error) {
	var err error
	var id uuid.UUID
	var onboarderID uuid.UUID
	var locationID uuid.UUID
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

	if proto.OnboarderID != nil {
		onboarderID, err = proto.OnboarderID.UUID()
		if err != nil {
			return nil, wgrpc.Error(wgrpc.CodeInternal, werror.Wrap(err, "could not parse proto onboarderID into a uuid").Add("proto.OnboarderID", proto.OnboarderID))
		}
	}

	if proto.LocationID != nil {
		locationID, err = proto.LocationID.UUID()
		if err != nil {
			return nil, wgrpc.Error(wgrpc.CodeInternal, werror.Wrap(err, "could not parse proto locationID into a uuid").Add("proto.LocationID", proto.LocationID))
		}
	}

	return &app.OnboardersLocation{
		ID:          id,
		OnboarderID: onboarderID,
		LocationID:  locationID,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
	}, nil
}

func convertOnboardersLocationToProto(onbl *app.OnboardersLocation) (*insysproto.OnboardersLocation, error) {
	if onbl == nil {
		return nil, werror.New("could not convert nil onboarders location")
	}

	createdAt, err := ptypes.TimestampProto(onbl.CreatedAt)
	if err != nil {
		return nil, werror.Wrap(err, "could not convert onboarders location created at time")
	}
	updatedAt, err := ptypes.TimestampProto(onbl.UpdatedAt)
	if err != nil {
		return nil, werror.Wrap(err, "could not convert onboarders location updated at time")
	}

	return &insysproto.OnboardersLocation{
		ID:          sharedproto.UUIDToProto(onbl.ID),
		OnboarderID: sharedproto.UUIDToProto(onbl.OnboarderID),
		LocationID:  sharedproto.UUIDToProto(onbl.LocationID),
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
	}, nil
}
