package consumers

import (
	"context"
	"strings"

	"github.com/gogo/protobuf/proto"
	nsq "github.com/nsqio/go-nsq"
	"weavelab.xyz/insys-onboarding-service/internal/app"
	"weavelab.xyz/monorail/shared/grpc-clients/client-grpc-clients/featureflagsclient"
	"weavelab.xyz/monorail/shared/protorepo/dist/go/messages/insysproto"
	"weavelab.xyz/monorail/shared/protorepo/dist/go/messages/sharedproto"
	"weavelab.xyz/monorail/shared/protorepo/dist/go/services/insys"
	"weavelab.xyz/monorail/shared/wlib/uuid"
	"weavelab.xyz/monorail/shared/wlib/werror"
	"weavelab.xyz/monorail/shared/wlib/wlog"
)

const (
	onboardingFeatureFlagName = "onboardingBetaEnabled"

	installWeaveTaskID    = "720af494-38a4-499f-8633-9c8d5169cd43" // Install weave on other workstations in your office
	installNewPhoneTaskID = "fd4f656c-c9f1-47b8-96ad-3080b999a843" // Install your new phones
	syncPatientDataTaskID = "16a6dc91-ec6b-4b09-b591-a5b0dfa92932" // Sync your patient data to weave
)

type ChiliPiperScheduleEventCreatedSubscriber struct {
	onboarderService app.OnboarderService

	featureFlagsClient *featureflagsclient.Client

	onboardersLocationServer insys.OnboardersLocationServer
	onboardingServer         insys.OnboardingServer
}

func NewChiliPiperScheduleEventCreatedSubscriber(onboarderService app.OnboarderService, ols insys.OnboardersLocationServer, onboardingServer insys.OnboardingServer, ff *featureflagsclient.Client) *ChiliPiperScheduleEventCreatedSubscriber {
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
		wlog.ErrorC(ctx, "could not unmarshal ChiliPiperScheduleEventCreated message body into proto for insysproto.ChiliPiperScheduleEventResponse struct")
		return nil
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
		createdTasks, err := c.onboardingServer.CreateTaskInstancesFromTasks(
			ctx,
			&insysproto.CreateTaskInstancesFromTasksRequest{LocationID: locationUUID},
		)
		if err != nil {
			return werror.Wrap(err, "could not create tasks instances from tasks").Add("locationID", locationID)
		}

		// update the installed schedule tasks to be completed
		for _, taskInstance := range createdTasks.TaskInstances {
			taskInstanceTaskID := taskInstance.TaskID.String()

			if taskInstanceTaskID == installWeaveTaskID ||
				taskInstanceTaskID == installNewPhoneTaskID ||
				taskInstanceTaskID == syncPatientDataTaskID {
				_, err := c.onboardingServer.UpdateTaskInstance(ctx, &insysproto.UpdateTaskInstanceRequest{ID: taskInstance.ID, Status: 2})
				if err != nil {
					return werror.Wrap(err, "could not set status for task instance").Add("taskInstance.ID", taskInstance.ID).Add("locationID", taskInstance.LocationID)
				}
			}
		}

	}

	_, err = c.onboardersLocationServer.CreateOrUpdate(ctx, &insysproto.OnboardersLocation{
		OnboarderID: sharedproto.UUIDToProto(onboarderID),
		LocationID:  locationUUID,
	})
	if err != nil {
		return werror.Wrap(err, "could not assign onboarder to location").Add("onboarderID", onboarderID).Add("locationID", locationID)
	}

	// NOTE: 10/8/2019 - feature flag service appears to not be persisting this value to be true. making the request multiple times is a hack to get the value to persist
	for i := 0; i < 5; i++ {
		err = c.featureFlagsClient.Update(ctx, locationID, onboardingFeatureFlagName, true)
		if err != nil {
			return werror.Wrap(err, "failed to turn on onboarding feature flag")
		}
	}

	return nil
}
