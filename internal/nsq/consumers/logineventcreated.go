package consumers

import (
	"context"
	"fmt"

	"github.com/kr/pretty"

	"github.com/gogo/protobuf/proto"
	"weavelab.xyz/monorail/shared/protorepo/dist/go/messages/client/clientproto"
	"weavelab.xyz/monorail/shared/wlib/werror"

	nsq "github.com/nsqio/go-nsq"
)

type LogInEventCreatedSubscriber struct {
}

func NewLogInEventCreatedSubscriber(ctx context.Context) *LogInEventCreatedSubscriber {
	return &LogInEventCreatedSubscriber{}
}

func (p LogInEventCreatedSubscriber) HandleMessage(ctx context.Context, m *nsq.Message) error {
	var le clientproto.LoginEvent

	fmt.Println(string(m.Body))

	err := proto.Unmarshal(m.Body, &le)
	if err != nil {
		return werror.Wrap(err, "could not unmarshal LoginEvent message body into proto for clientproto.LoginEvent struct")
	}

	pretty.Println(le)

	return nil
}
