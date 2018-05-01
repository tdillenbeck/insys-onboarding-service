package wsql

import (
	"database/sql/driver"
	"encoding/json"
	"strconv"
	"strings"

	"weavelab.xyz/wlib/uuid"
	"weavelab.xyz/wlib/werror"
)

func isNull(arr []byte) bool {
	// if it is a null array it will be "{NULL}"
	return len(arr) == 6 && arr[1] == 'N'
}

type Array []int

func (a Array) Value() (driver.Value, error) {
	if len(a) == 0 {
		return nil, nil
	}

	joined := "{"
	for i, v := range []int(a) {
		if i > 0 {
			joined += ","
		}
		joined += strconv.Itoa(v) + " "
	}
	joined += "}"

	return joined, nil
}

func (a *Array) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	switch str := value.(type) {
	case []byte:
		// check for a null response
		if isNull(str) {
			return nil
		}

		arr := strings.Trim(string(str), "{}")

		arrSplit := strings.Split(arr, ",")

		for _, v := range arrSplit {
			// if one of the values in the int array is NULL, skip it

			if len(v) > 0 && v[0] == 'N' {
				continue
			}

			i, err := strconv.Atoi(v)
			if err != nil {
				return err
			}

			*a = append(*a, i)
		}

	default:
		return werror.New("Unsupported type").Add("type", value)
	}

	return nil

}

type ArrayUUID struct {
	UUIDs []uuid.UUID
}

func (a *ArrayUUID) Value() (driver.Value, error) {
	if len(a.UUIDs) == 0 {
		return nil, nil
	}

	joined := "{"
	for i, v := range a.UUIDs {
		if i > 0 {
			joined += ","
		}
		joined += v.String() + " "
	}
	joined += "}"

	return joined, nil
}

func (a *ArrayUUID) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	switch str := value.(type) {
	case []byte:

		// check for a null response
		if isNull(str) {
			return nil
		}

		arr := strings.Trim(string(str), "{}")

		arrSplit := strings.Split(arr, ",")

		for _, v := range arrSplit {
			// if one of the values in the int array is NULL, skip it
			if len(v) > 0 && v[0] == 'N' {
				continue
			}

			i, err := uuid.Parse(v)
			if err != nil {
				return werror.Wrap(err, "unable to parse").Add("value", string(v))
			}

			a.UUIDs = append(a.UUIDs, i)
		}

	default:
		return werror.New("unsupported type").Add("type", value)
	}

	return nil

}

func (a ArrayUUID) MarshalJSON() ([]byte, error) {
	return json.Marshal(a.UUIDs)
}
