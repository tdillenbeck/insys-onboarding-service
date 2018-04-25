package sharedproto

import (
	"encoding/json"

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

func (u UUID) MarshalJSON() (text []byte, err error) {
	uID, err := uuid.New(u.Bytes)
	if err != nil {
		return nil, err
	}

	return json.Marshal(uID.String())
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
