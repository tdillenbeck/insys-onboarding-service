package null

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strings"

	"weavelab.xyz/monorail/shared/wlib/uuid"
	"weavelab.xyz/monorail/shared/wlib/werror"
)

// Custom types must implement the Valuer and Scanner interfaces
// http://golang.org/pkg/database/sql/driver/#Valuer
// http://golang.org/pkg/database/sql/#Scanner

type UUID struct {
	UUID  uuid.UUID
	Valid bool // Valid is true if String is not NULL
}

// NewUUIDUUID does no validation on UUID argument for correctness
func NewUUIDUUID(v uuid.UUID) UUID {
	return UUID{UUID: v, Valid: true}
}

func NewUUIDUUIDDefaultAsNull(v uuid.UUID) UUID {
	if v.IsEmpty() {
		return UUID{}
	}

	return NewUUIDUUID(v)
}

func NewUUID(v string) (UUID, error) {

	u, err := uuid.Parse(v)
	if err != nil {
		return UUID{}, err
	}

	return UUID{UUID: u, Valid: true}, err
}

func NewUUIDDefaultAsNull(v string) (UUID, error) {
	if v == "" {
		return UUID{}, nil
	}

	return NewUUID(v)
}

// Value implements the driver Valuer interface.
func (ns UUID) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}

	return ns.String(), nil
}

func (ns UUID) String() string {
	return ns.UUID.String()
}

//AsUUID returns the UUID
func (ns UUID) AsUUID() uuid.UUID {
	return ns.UUID
}

// Scan implements the Scanner interface.
func (ns *UUID) Scan(value interface{}) error {

	if value == nil {
		ns.Valid = false
		return nil
	}

	ns.Valid = true

	switch str := value.(type) {
	// the incoming bytes are a string representation of the UUID
	case []byte:
		var err error
		*ns, err = NewUUID(string(str))
		if err != nil {
			return fmt.Errorf("unable to decode UUID: %s", string(str))
		}

	default:
		ns.Valid = false
		return fmt.Errorf("unknown type %T for UUID", value)
	}

	return nil

}

func (ns UUID) MarshalJSON() ([]byte, error) {
	if ns.Valid == false {
		return json.Marshal(nil)
	}
	u := ns.UUID

	return json.Marshal(u.String())
}

func (ns *UUID) UnmarshalJSON(rawData []byte) error {

	if string(rawData) == "null" || len(rawData) == 0 {
		return nil
	}

	var value string
	err := json.Unmarshal(rawData, &value)
	if err != nil {
		return fmt.Errorf("unable to unmarshal UUID into string: %s", err)
	}

	u, err := uuid.Parse(strings.ToUpper(value))
	if err != nil {
		return fmt.Errorf("unable to parse %s: %s", value, err)
	}

	ns.UUID = u
	ns.Valid = true

	return nil
}

func (ns *UUID) ToUUID() (uuid.UUID, error) {
	if ns.Valid == false {
		return uuid.UUID{}, werror.New("invalid UUID")
	}

	return ns.UUID, nil
}
