package wnsq

import (
	"github.com/golang/protobuf/proto"
	"weavelab.xyz/wlib/wnsq/wnsqproto"
)

func isProto(b []byte) bool {

	w := wnsqproto.OpenTracing{}
	err := proto.Unmarshal(b, &w)
	if err != nil {
		return false
	}
	return true
}

func injectorProto(md map[string]string, b []byte) ([]byte, error) {

	o := wnsqproto.OpenTracing{
		OpenTracing: md,
	}

	out, err := proto.Marshal(&o)
	if err != nil {
		return nil, err
	}

	// append the data to the existing proto message
	d := append(b, out...)
	return d, nil

}

func extractProto(b []byte) (map[string]string, error) {

	m := wnsqproto.OpenTracing{}
	err := proto.Unmarshal(b, &m)
	if err != nil {
		return nil, err
	}

	return m.OpenTracing, nil
}
