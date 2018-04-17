package server

import (
	"context"
	"reflect"
	"testing"
	"time"

	"weavelab.xyz/insys-onboarding/exampleproto"
	"weavelab.xyz/wlib/uuid"
	"weavelab.xyz/wlib/wgrpc/wgrpcproto"
)

func TestServerImpl_ExampleRequest(t *testing.T) {
	type fields struct {
		Delay time.Duration
	}
	type args struct {
		ctx context.Context
		in  *exampleproto.ExampleRequestMessage
	}
	id := uuid.NewV4()

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *exampleproto.ExampleResponseMessage
		wantErr bool
	}{
		{
			"hey",
			fields{
				Delay: time.Second,
			},
			args{
				ctx: context.TODO(),
				in: &exampleproto.ExampleRequestMessage{
					SomeID: wgrpcproto.UUIDProto(id),
				},
			},
			&exampleproto.ExampleResponseMessage{
				Message: "hey",
			},
			false,
		},
		{
			"pass no ID get error",
			fields{
				Delay: time.Second,
			},
			args{
				ctx: context.TODO(),
				in:  &exampleproto.ExampleRequestMessage{},
			},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &ServerImpl{
				Delay: tt.fields.Delay,
			}
			got, err := e.ExampleRequest(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("ServerImpl.ExampleRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ServerImpl.ExampleRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}
