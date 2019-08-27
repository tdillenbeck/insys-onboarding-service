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

const (
	onboardingFeatureFlagName = "onboardingBetaEnabled"
)

type FeatureFlagsClient interface {
	Update(ctx context.Context, locationID uuid.UUID, name string, enable bool) error
}

type ChiliPiperScheduleEventCreatedSubscriber struct {
	onboarderService app.OnboarderService

	featureFlagsClient FeatureFlagsClient

	onboardersLocationServer insys.OnboardersLocationServer
	onboardingServer         insys.OnboardingServer
}

func NewChiliPiperScheduleEventCreatedSubscriber(onboarderService app.OnboarderService, ols insys.OnboardersLocationServer, onboardingServer insys.OnboardingServer, ff FeatureFlagsClient) *ChiliPiperScheduleEventCreatedSubscriber {
	return &ChiliPiperScheduleEventCreatedSubscriber{
		onboarderService: onboarderService,

		featureFlagsClient: ff,

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

	// only turn on the onboarding tracker if the chili piper event is a next_steps call
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

// turnOnOnboardingTracker will assign an onboarder to a location, create the task instances, and turn on the feature flag for a location.
func (c ChiliPiperScheduleEventCreatedSubscriber) turnOnOnboardingTracker(ctx context.Context, onboarderID, locationID uuid.UUID) error {
	locationUUID := sharedproto.UUIDToProto(locationID)

	// if the onboarding tracker has already been setup, don't do anything
	tasks, err := c.onboardingServer.TaskInstances(ctx, &insysproto.TaskInstancesRequest{LocationID: locationUUID})
	if err != nil {
		return werror.Wrap(err, "could not look up task instances for location").Add("locationID", locationID)
	}

	if len(tasks.TaskInstances) == 0 {
		_, err = c.onboardingServer.CreateTaskInstancesFromTasks(
			ctx,
			&insysproto.CreateTaskInstancesFromTasksRequest{LocationID: locationUUID},
		)
		if err != nil {
			return werror.Wrap(err, "could not create tasks instances from tasks").Add("locationID", locationID)
		}
	}

	_, err = c.onboardersLocationServer.CreateOrUpdate(ctx, &insysproto.OnboardersLocation{
		OnboarderID: sharedproto.UUIDToProto(onboarderID),
		LocationID:  locationUUID,
	})
	if err != nil {
		return werror.Wrap(err, "could not assign onboarder to location").Add("onboarderID", onboarderID).Add("locationID", locationID)
	}

	err = c.featureFlagsClient.Update(ctx, locationID, onboardingFeatureFlagName, true)
	if err != nil {
		return werror.Wrap(err, "failed to turn on onboarding feature flag")
	}

	return nil
}
