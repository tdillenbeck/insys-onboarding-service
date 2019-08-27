package consumers

import (
	"context"

	"github.com/gogo/protobuf/proto"
	nsq "github.com/nsqio/go-nsq"
	"weavelab.xyz/insys-onboarding-service/internal/app"
	"weavelab.xyz/monorail/shared/protorepo/dist/go/enums/insysenums"
	"weavelab.xyz/monorail/shared/protorepo/dist/go/messages/insysproto"
	"weavelab.xyz/monorail/shared/wlib/uuid"
	"weavelab.xyz/monorail/shared/wlib/werror"
)

const (
	PortingInfoTaskID = "9aec502b-f8b8-4f10-9748-1fe4050eacde"
)

type PortingDataRecordCreatedSubscriber struct {
	taskInstanceService app.TaskInstanceService
}

func NewPortingDataRecordCreatedSubscriber(ctx context.Context, tis app.TaskInstanceService) *PortingDataRecordCreatedSubscriber {
	return &PortingDataRecordCreatedSubscriber{
		taskInstanceService: tis,
	}
}

func (p PortingDataRecordCreatedSubscriber) HandleMessage(ctx context.Context, m *nsq.Message) error {
	var pd insysproto.PortingData
	err := proto.Unmarshal(m.Body, &pd)
	if err != nil {
		return werror.Wrap(err, "could not unmarshal PortingDataCreated message body into proto for insysproto.PortingData struct")
	}

	locationUUID, err := uuid.Parse(pd.LocationId)
	if err != nil {
		return werror.Wrap(err, "could not parse location id  into a uuid").Add("locationID", pd.LocationId)
	}

	taskInstances, err := p.taskInstanceService.ByLocationID(ctx, locationUUID)
	if err != nil {
		return werror.Wrap(err, "no task instances for this location").Add("locationID", pd.LocationId)
	}

	for _, taskInstance := range taskInstances {
		if taskInstance.TaskID.String() == PortingInfoTaskID {
			updatedBy := pd.AuthorizedUserFirstName + " " + pd.AuthorizedUserLastName
			_, err := p.taskInstanceService.Update(ctx, taskInstance.ID, insysenums.OnboardingTaskStatus_Completed, updatedBy)
			if err != nil {
				return werror.Wrap(err, "could not update task to completed").Add("task instance id", taskInstance.ID)
			}
		}
	}

	return nil
}
