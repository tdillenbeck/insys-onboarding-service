package null

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strconv"
	"time"
	"unicode/utf8"
)

// Custom types must implement the Valuer and Scanner interfaces
// http://golang.org/pkg/database/sql/driver/#Valuer
// http://golang.org/pkg/database/sql/#Scanner

func NewString(s string) String {
	return String{Str: s, Valid: true}
}

type String struct {
	Str   string
	Valid bool
}

func (ns String) String() string {
	if ns.Valid {
		return ns.Str
	}
	return ""
}

func (ns String) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}

	return ns.Str, nil
}

func (ns *String) Scan(value interface{}) error {
	if value == nil {
		ns.Str, ns.Valid = "", false
		return nil
	}
	ns.Valid = true

	return convertAssign(&ns.Str, value)
}

func (ns String) MarshalJSON() ([]byte, error) {
	v, _ := ns.Value()
	return json.Marshal(v)
}

func (ns *String) UnmarshalJSON(rawData []byte) error {

	if string(rawData) == "null" {
		return nil
	}

	err := json.Unmarshal(rawData, &ns.Str)

	if err != nil {
		return err
	}

	ns.Valid = true

	return nil

}

func (s *String) Truncate(max int) {
	s.Str = Truncate(s.Str, max)

}

// Truncate returns a truncated version of
// s with maximum n runes
func Truncate(s string, max int) string {

	c := utf8.RuneCountInString(s)
	if c <= max {
		return s
	}

	var count, i int
	for i, _ = range s {

		if count >= max {
			break
		}

		count++
	}

	return s[:i]

}

func NewBool(v bool) Bool {
	return Bool{Bool: v, Valid: true}
}

type Bool struct {
	Bool  bool
	Valid bool // Valid is true if Bool is not NULL
}

// Scan implements the Scanner interface.
func (n *Bool) Scan(value interface{}) error {
	if value == nil {
		n.Bool, n.Valid = false, false
		return nil
	}
	n.Valid = true
	return convertAssign(&n.Bool, value)
}

// Value implements the driver Valuer interface.
func (n Bool) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.Bool, nil
}

func (ns Bool) String() string {
	return fmt.Sprintf("%t", ns.Bool)
}

func (ns Bool) MarshalJSON() ([]byte, error) {
	v, _ := ns.Value()
	return json.Marshal(v)
}

func (ns *Bool) UnmarshalJSON(rawData []byte) error {

	if string(rawData) == "null" {
		return nil
	}

	err := json.Unmarshal(rawData, &ns.Bool)

	if err != nil {
		return err
	}

	ns.Valid = true

	return nil
}

type Int64 struct {
	Int64 int64
	Valid bool
}

func NewInt64(v int64) Int64 {
	return Int64{Int64: int64(v), Valid: true}
}

func (ns Int64) MarshalJSON() ([]byte, error) {
	v, _ := ns.Value()
	return json.Marshal(v)
}

func (ns *Int64) UnmarshalJSON(rawData []byte) error {

	if string(rawData) == "null" {
		return nil
	}

	err := json.Unmarshal(rawData, &ns.Int64)

	if err != nil {
		return err
	}

	ns.Valid = true

	return nil
}

func (ns *Int64) Scan(value interface{}) error {
	if value == nil {
		ns.Int64, ns.Valid = 0, false
		return nil
	}
	ns.Valid = true
	return convertAssign(&ns.Int64, value)
}

func (ns Int64) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}

	return ns.Int64, nil
}

func NewTimeParse(timeString string) (Time, error) {

	t, err := time.Parse(time.RFC3339, timeString)

	return Time{Time: t, Valid: err == nil}, err

}

func NewTime(time time.Time) Time {

	return Time{Time: time, Valid: !time.IsZero()}

}

type Time struct {
	Time  time.Time
	Valid bool // Valid is true if String is not
}

// Value implements the driver Valuer interface.
func (ns Time) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}

	return ns.Time, nil
}

// Scan implements the Scanner interface.
func (ns *Time) Scan(value interface{}) error {

	if value == nil {
		ns.Time, ns.Valid = time.Time{}, false
		return nil
	}

	ns.Valid = true

	switch str := value.(type) {
	case time.Time:
		ns.Time = str

	case []byte:
		var err error
		ns.Time, err = time.Parse("15:04:05", string(str))
		if err != nil {
			ns.Time = time.Time{}
			return err
		}

	default:
		ns.Valid = false
	}

	return nil

}

func (ns Time) MarshalJSON() ([]byte, error) {

	if ns.Time.IsZero() || ns.Valid == false {
		return json.Marshal(nil)
	}

	return json.Marshal(ns.Time.Format(time.RFC3339))
}

func (ns *Time) UnmarshalJSON(rawData []byte) error {

	if string(rawData) == "null" || len(rawData) == 0 {
		return nil
	}

	err := json.Unmarshal(rawData, &ns.Time)

	if err != nil {
		return err
	}

	ns.Valid = true

	return nil

}

type Date struct {
	Date  time.Time
	Valid bool // Valid is true if String is not NULL
}

func NewTimeDateOnly(dateString string) (Date, error) {

	pt, err := time.Parse(`2006-01-02`, dateString)

	return Date{Date: pt, Valid: err == nil}, err

}

func (ns Date) String() string {
	return fmt.Sprintf("%s", ns.Date)
}

// Value implements the driver Valuer interface.
func (ns Date) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}

	return ns.Date, nil
}

// Scan implements the Scanner interface.
func (ns *Date) Scan(value interface{}) error {
	if value == nil {
		ns.Date, ns.Valid = time.Time{}, false
		return nil
	}
	ns.Valid = true

	switch str := value.(type) {
	case time.Time:
		ns.Date = str

	default:
		ns.Valid = false
	}

	return nil

}

func (ns Date) MarshalJSON() ([]byte, error) {

	if ns.Date.IsZero() {
		return json.Marshal(nil)
	}

	return json.Marshal(ns.Date.Format(`2006-01-02`))
}

func (ns *Date) UnmarshalJSON(rawData []byte) error {

	// Date format is YYYY-MM-DD

	if string(rawData) == "null" || len(rawData) == 0 {
		return nil
	}

	pt, err := time.Parse(`"2006-01-02"`, string(rawData))

	if err == nil {
		ns.Date = pt
		ns.Valid = true

		return nil
	}

	pt, err = time.Parse(`"`+time.RFC3339+`"`, string(rawData))

	if err != nil {
		return err
	}

	ns.Date = pt
	ns.Valid = true

	return nil
}

type ClockTime struct {
	ClockTime time.Time
	Valid     bool // Valid is true if String is not NULL
}

func NewTimeClockTimeOnly(clockString string) (ClockTime, error) {

	pt, err := time.Parse(`15:04:05`, clockString)

	return ClockTime{ClockTime: pt, Valid: err == nil}, err

}

func (ns ClockTime) String() string {
	return fmt.Sprintf("%s", ns.ClockTime)
}

// Value implements the driver Valuer interface.
func (ns ClockTime) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}

	return ns.ClockTime, nil
}

// Scan implements the Scanner interface.
func (ns *ClockTime) Scan(value interface{}) error {
	if value == nil {
		ns.ClockTime, ns.Valid = time.Time{}, false
		return nil
	}
	ns.Valid = true

	switch str := value.(type) {
	case time.Time:
		ns.ClockTime = str

	default:
		ns.Valid = false
	}

	return nil

}

func (ns ClockTime) MarshalJSON() ([]byte, error) {

	if ns.ClockTime.IsZero() {
		return json.Marshal(nil)
	}

	return json.Marshal(ns.ClockTime.Format(`15:04:05`))
}

func (ns *ClockTime) UnmarshalJSON(rawData []byte) error {

	// ClockTime format is YYYY-MM-DD

	if string(rawData) == "null" || len(rawData) == 0 {
		return nil
	}

	pt, err := time.Parse(`"15:04:05"`, string(rawData))

	if err == nil {
		ns.ClockTime = pt
		ns.Valid = true

		return nil
	}

	pt, err = time.Parse(`"`+time.RFC3339+`"`, string(rawData))

	if err != nil {
		return err
	}

	ns.ClockTime = pt
	ns.Valid = true

	return nil
}

type KitchenTime struct {
	KitchenTime time.Time
	Valid       bool
}

func NewKitchenTimeOnly(dateString string) (KitchenTime, error) {

	pt, err := time.Parse(time.Kitchen, dateString)

	return KitchenTime{KitchenTime: pt, Valid: err == nil}, err

}

func (kt KitchenTime) String() string {
	return fmt.Sprintf("%s", kt.KitchenTime.Format(time.Kitchen))
}

func (kt KitchenTime) MarshalJSON() ([]byte, error) {

	if kt.KitchenTime.IsZero() {
		return json.Marshal(nil)
	}

	return json.Marshal(kt.KitchenTime.Format(time.Kitchen))
}

func (kt *KitchenTime) UnmarshalJSON(rawData []byte) error {

	if string(rawData) == "null" || len(rawData) == 0 {
		return nil
	}

	pt, err := time.Parse(time.Kitchen, string(rawData))

	if err == nil {
		kt.KitchenTime = pt
		kt.Valid = true

		return nil
	}

	pt, err = time.Parse(`"`+time.Kitchen+`"`, string(rawData))

	if err != nil {
		return err
	}

	kt.KitchenTime = pt
	kt.Valid = true

	return nil
}

type Float32 struct {
	Float32 float32
	Valid   bool // Valid is true if String is not NULL
}

func NewFloat32(v float32) Float32 {
	return Float32{Float32: v, Valid: true}
}

// Value implements the driver Valuer interface.
func (ns Float32) Value() (driver.Value, error) {
	if !ns.Valid {

		return nil, nil
	}

	return float64(ns.Float32), nil
}

// Scan implements the Scanner interface.
func (ns *Float32) Scan(value interface{}) error {

	if value == nil {
		ns.Valid = false
		return nil
	}
	ns.Valid = true

	switch str := value.(type) {
	case float32:
		ns.Float32 = str
	case float64:
		ns.Float32 = float32(str)

	case []byte:
		f64, err := strconv.ParseFloat(string(str), 32)
		if err != nil {
			ns.Valid = false
			return fmt.Errorf("converting string %q to a %s: %v", string(str), "float32", err)
		}
		ns.Float32 = float32(f64)

	default:

		ns.Valid = false
	}

	return nil

}

func (ns Float32) MarshalJSON() ([]byte, error) {

	if ns.Valid == false {
		return json.Marshal(nil)
	}

	return json.Marshal(ns.Float32)
}

func (ns *Float32) UnmarshalJSON(rawData []byte) error {

	if string(rawData) == "null" || len(rawData) == 0 {
		return nil
	}

	err := json.Unmarshal(rawData, &ns.Float32)

	if err != nil {
		return err
	}

	ns.Valid = true
	return nil
}
