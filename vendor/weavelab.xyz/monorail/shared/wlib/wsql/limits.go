package wsql

import (
	"errors"
	"fmt"
	"net/url"
)

const GENERIC_LIST_DEFAULT_LIMIT = 50

var errLimitOverMax = errors.New("limit: over maximum permitted value")

type Limits struct {
	limit int
	skip  int
	max   int
}

func NewLimits(limit, skip int, max int) (Limits, error) {
	l := Limits{limit: limit, skip: skip, max: max}
	if limit > max {
		return l, errLimitOverMax
	}

	return l, nil
}

func NewLimitsFromURL(values url.Values) Limits {

	limit := GetDefaultInt(values, "limit", 0)
	if limit < 0 {
		limit = 0
	}

	skip := GetDefaultInt(values, "skip", 0)
	if skip < 0 {
		skip = 0
	}

	return Limits{limit: int(limit), skip: int(skip)}

}

// If limit <= 0 do not limit the returned results?
func (l *Limits) ToSQL() string {

	if l.limit > l.max && l.max > 0 {
		l.limit = l.max
	}

	if l.limit == 0 && l.max > 0 {
		l.limit = GENERIC_LIST_DEFAULT_LIMIT
	}

	if l.limit <= 0 && l.skip <= 0 && l.max <= 0 {
		return " "
	}

	if l.limit <= 0 {
		return fmt.Sprintf("OFFSET %d", l.skip)
	}

	if l.limit >= 0 && l.skip <= 0 {
		return fmt.Sprintf("LIMIT %d", l.limit)
	}

	return fmt.Sprintf("LIMIT %d OFFSET %d", l.limit, l.skip)

}

// 0 - no maximum limit
// > 0 - limit
func (l *Limits) SetMaxLimit(max int) error {

	if max >= 0 {
		l.max = max
	}

	if l.max > 0 && l.limit > l.max {
		l.limit = l.max
		return errLimitOverMax
	}

	if l.limit == 0 {
		l.limit = l.max
	}

	return nil
}
