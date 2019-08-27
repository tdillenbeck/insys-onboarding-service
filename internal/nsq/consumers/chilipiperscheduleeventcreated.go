package consumers

import (
	"context"
	"strings"

	"github.com/gogo/protobuf/proto"
	nsq "github.com/nsqio/go-nsq"
	"weavelab.xyz/insys-onboarding-service/internal/app"
	"weavelab.xyz/monorail/shared/protorepo/dist/go/messages/insysproto"
	"weavelab.xyz/monorail/shared/protorepo/dist/go/messages/sharedproto"
	"weavelab.xyz/monorail/shared/protorepo/dist/go/services/insys"
	"weavelab.xyz/monorail/shared/wlib/uuid"
	"weavelab.xyz/monorail/shared/wlib/werror"
)

type ChiliPiperScheduleEventCreatedSubscriber struct {
	onboarderService app.OnboarderService

	onboardersLocationServer insys.OnboardersLocationServer
	onboardingServer         insys.OnboardingServer
}

func NewChiliPiperScheduleEventCreatedSubscriber(onboarderService app.OnboarderService, ols insys.OnboardersLocationServer, onboardingServer insys.OnboardingServer) *ChiliPiperScheduleEventCreatedSubscriber {
	return &ChiliPiperScheduleEventCreatedSubscriber{
		onboarderService: onboarderService,

		onboardersLocationServer: ols,
		onboardingServer:         onboardingServer,
	}
}

func (c ChiliPiperScheduleEventCreatedSubscriber) HandleMessage(ctx context.Context, m *nsq.Message) error {
	var chiliPiperScheduleEventResponse insysproto.CreateChiliPiperScheduleEventResponse

	err := proto.Unmarshal(m.Body, &chiliPiperScheduleEventResponse)
	if err != nil {
		return werror.Wrap(err, "could not unmarshal ChiliPiperScheduleEventCreated message body into proto for insysproto.ChiliPiperScheduleEventResponse struct")
	}

	// Only assign the onboarder on the next steps call event. Assigning the onbaorder will turn on the tracker
	if strings.Contains(chiliPiperScheduleEventResponse.Event.EventType, "next_steps") {
		onboarder, err := c.onboarderService.ReadBySalesforceUserID(ctx, chiliPiperScheduleEventResponse.Event.AssigneeId)
		if err != nil {
			return werror.Wrap(err, "could not read onboarder by salesforce user id").Add("salesforce user id", chiliPiperScheduleEventResponse.Event.AssigneeId)
		}
		locationID, err := uuid.Parse(chiliPiperScheduleEventResponse.Event.LocationId)
		if err != nil {
			return werror.Wrap(err, "could not parse location id from chili piper schedule event create response").Add("LocationId", chiliPiperScheduleEventResponse.Event.LocationId)
		}

		err = c.turnOnOnboardingTracker(ctx, onboarder.ID, locationID)
		if err != nil {
			return werror.Wrap(err, "could not turn on onboarding tracker").Add("locationID", locationID).Add("onboarderID", onboarder.ID)
		}
	}

	return nil
}

func (c ChiliPiperScheduleEventCreatedSubscriber) turnOnOnboardingTracker(ctx context.Context, onboarderID, locationID uuid.UUID) error {
	locationUUID := sharedproto.UUIDToProto(locationID)

	_, err := c.onboardersLocationServer.CreateOrUpdate(ctx, &insysproto.OnboardersLocation{
		OnboarderID: sharedproto.UUIDToProto(onboarderID),
		LocationID:  locationUUID,
	})
	if err != nil {
		return werror.Wrap(err, "could not assign onboarder to location").Add("onboarderID", onboarderID).Add("locationID", locationID)
	}

	// should the CreateTaskInstancesFromTasks not create if the tasks are already created??

	tasks, err := c.onboardingServer.TaskInstances(ctx, &insysproto.TaskInstancesRequest{LocationID: locationUUID})
	if err != nil {
		return werror.Wrap(err, "could not look up task instances for location").Add("locationID", locationID)
	}

	// only create tasks if there are no tasks created
	if len(tasks) == 0 {
		_, err = c.onboardingServer.CreateTaskInstancesFromTasks(
			ctx,
			&insysproto.CreateTaskInstancesFromTasksRequest{LocationID: locationUUID},
		)
		if err != nil {
			return werror.Wrap(err, "could not create tasks instances from tasks").Add("locationID", locationID)
		}
	}

	// TODO: turn on customization flag

	return nil
}
