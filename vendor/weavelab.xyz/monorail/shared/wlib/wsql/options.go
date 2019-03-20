package wsql

import (
	"net/http"

	"weavelab.xyz/monorail/shared/wlib/werror"
)

type Parameters struct {
	Filters      map[string]MapEntry
	Orders       []string
	DefaultOrder []string
	ListMax      int
}

type QueryOptions struct {
	Filters        []Filter
	DefaultFilters []Filter
	GroupBy        []string
	Having         []string
	Query          []string
	Orders         []string
	Limits         Limits
	Distinct       bool
}

type Total struct {
	Total int `db:"total" json:"count"`
}

func NewOptions(r *http.Request) (QueryOptions, error) {

	r.ParseForm()

	filters, err := NewFilter(r.Form["filter"]...)
	if err != nil {
		return QueryOptions{}, err
	}

	return QueryOptions{
		Filters: filters,
		Orders:  r.Form["order_by"],
		Query:   r.Form["q"],
		GroupBy: r.Form["group_by"],
		Limits:  NewLimitsFromURL(r.Form),
	}, nil
}

/*
AddDefaultFilter adds a filter if the query options does not yet contain
 a filter for that field
*/
func (q *QueryOptions) AddDefaultFilter(field string, operator Operator, value string) {
	f := Filter{Field: field, Operator: operator, Value: value}
	q.DefaultFilters = append(q.DefaultFilters, f)
}

func (q *QueryOptions) AddFilter(field string, operator Operator, value string) {
	f := Filter{Field: field, Operator: operator, Value: value}
	q.Filters = append(q.Filters, f)
}

//Filter returns a pointer to the filter with the given field or an nil and an error if it doesn not exist
func (q *QueryOptions) Filter(field string) (*Filter, error) {
	for _, v := range q.Filters {
		if v.Field == field {
			return &v, nil
		}
	}

	return nil, werror.New("Could not find filter for the given field").Add("field", field)
}

func (q *QueryOptions) HasFilter(field string) bool {
	for _, v := range q.Filters {
		if v.Field == field {
			return true
		}
	}

	for _, v := range q.DefaultFilters {
		if v.Field == field {
			return true
		}
	}

	return false
}

func (q *QueryOptions) RemoveFilter(field string) {
	for i, v := range q.Filters {
		if v.Field == field {
			q.Filters = append(q.Filters[:i], q.Filters[i+1:]...)
			return
		}
	}
}

func (q *QueryOptions) FilterSQL(available map[string]MapEntry) (string, map[string]interface{}, error) {

	f, err := NewFilters(q.Filters, available)
	if err != nil {
		return "", nil, err
	}

	sql, parameters := f.ToSQL()
	return sql, parameters, nil
}

func (q *QueryOptions) FilterSlice(available map[string]MapEntry) ([]string, map[string]interface{}, error) {

	f, err := NewFilters(q.Filters, available)
	if err != nil {
		return nil, nil, err
	}

	slice, parameters := f.ToSlice()

	return slice, parameters, nil
}

func (q *QueryOptions) OrderSQL(defaultOrder []string, available []string) (string, error) {
	o, err := NewOrderByFromString(q.Orders, defaultOrder, available)
	if err != nil {
		return "", err
	}
	return o.ToSQL(), nil

}

func (q *QueryOptions) LimitSQL(max int) (string, error) {
	err := q.Limits.SetMaxLimit(max)
	if err != nil {
		return "", err
	}
	return q.Limits.ToSQL(), nil
}
