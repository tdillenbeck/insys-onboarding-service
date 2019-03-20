package null

import (
	"database/sql/driver"
	"encoding/json"

	"weavelab.xyz/monorail/shared/wlib/uuid"
)

// Custom types must implement the Valuer and Scanner interfaces
// http://golang.org/pkg/database/sql/driver/#Valuer
// http://golang.org/pkg/database/sql/#Scanner

type MSGUID struct {
	UUID  []byte
	Valid bool // Valid is true if String is not NULL
}

func NewMSGuid(v string) (MSGUID, error) {

	u, err := uuid.NewHex(v)

	return MSGUID{UUID: u.Bytes(), Valid: true}, err
}

// Value implements the driver Valuer interface.
func (ns MSGUID) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}

	// convert back to MS format
	uuid := msConvertFromUUIDOrder(ns.UUID)

	return uuid, nil

}

func (ns MSGUID) String() string {
	uuid, err := uuid.New(ns.UUID)
	if err != nil {
		return ""
	}

	return uuid.String()
}

// Scan implements the Scanner interface.
func (ns *MSGUID) Scan(value interface{}) error {

	if value == nil {
		ns.Valid = false
		return nil
	}

	ns.Valid = true

	switch str := value.(type) {
	case []byte:

		ns.UUID = msConvertToUUIDOrder(str)
		// reformat the uuid bytes into network order

	default:
		ns.Valid = false
	}

	return nil

}

var msConvertFromUUIDOrder = msConvertToUUIDOrder

func msConvertToUUIDOrder(in []byte) []byte {

	if len(in) != 16 {
		return in
	}

	out := append([]byte{in[3]}, in[2], in[1], in[0],
		in[5], in[4],
		in[7], in[6])
	out = append(out, in[8:]...)

	return out
}

func (ns MSGUID) MarshalJSON() ([]byte, error) {

	if ns.Valid == false {
		return json.Marshal(nil)
	}

	u, err := uuid.New(ns.UUID)
	if err != nil {
		return nil, err
	}

	return json.Marshal(u.String())
}

func (ns *MSGUID) UnmarshalJSON(rawData []byte) error {

	if len(rawData) == 0 || string(rawData) == "null" {
		return nil
	}

	var value string
	err := json.Unmarshal(rawData, &value)

	if err != nil {
		return err
	}

	u, err := uuid.NewHex(value)

	if err != nil {
		return err
	}

	ns.UUID = u.Bytes()
	ns.Valid = true

	return nil
}
