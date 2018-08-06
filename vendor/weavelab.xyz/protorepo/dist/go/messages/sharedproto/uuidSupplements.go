package sharedproto

import (
	"encoding/json"

	"github.com/golang/protobuf/jsonpb"
	"weavelab.xyz/wlib/uuid"
	"weavelab.xyz/wlib/werror"
)

func (u *UUID) MarshalText() (text []byte, err error) {
	uID, err := uuid.New(u.Bytes)
	if err != nil {
		return nil, err
	}

	return []byte(uID.String()), nil
}

func (u UUID) MarshalJSONPB(_ *jsonpb.Marshaler) ([]byte, error) {
	return u.MarshalJSON()
}

func (u UUID) MarshalJSON() (text []byte, err error) {
	uID, err := uuid.New(u.Bytes)
	if err != nil {
		return nil, err
	}

	return json.Marshal(uID.String())
}

func (u *UUID) UnmarshalJSONPB(_ *jsonpb.Unmarshaler, buff []byte) error {
	return u.UnmarshalJSON(buff)
}

func (u *UUID) UnmarshalJSON(buff []byte) error {

	var id uuid.UUID
	if len(buff) == 26 { // base64 encoded byte array
		var dest []byte
		err := json.Unmarshal(buff, &dest)
		if err != nil {
			return werror.Wrap(err, "unable to unmarshal into bytes").Add("json", buff)
		}

		// byte slice representation
		id, err = uuid.New(dest)
		if err != nil {
			return werror.Wrap(err, "unable to convert bytes")
		}

	} else if buff[0] == '{' {
		return werror.New("can't handle object type yet, ask troy to implement")

	} else {
		var dest string
		err := json.Unmarshal(buff, &dest)
		if err != nil {
			return werror.Wrap(err, "unable to unmarshal into string").Add("json", string(buff))
		}

		// ascii hex encoded
		id, err = uuid.Parse(dest)
		if err != nil {
			return werror.Wrap(err).Add("json", string(buff))
		}

	}

	// convert uuid.UUID to protouuid
	u.Bytes = id.Bytes()

	return nil
}

func (u *UUID) UUID() (uuid.UUID, error) {
	if u == nil {
		return uuid.UUID{}, werror.New("cannot convert nil to uuid.UUID")
	}

	result, err := uuid.New(u.Bytes)
	if err != nil {
		return uuid.UUID{}, werror.Wrap(err, "Not valid UUID")
	}

	return result, nil
}

func UUIDToProto(u uuid.UUID) *UUID {
	return &UUID{
		Bytes: u.Bytes(),
	}
}
