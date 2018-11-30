package wsql

import (
	"log"
	"regexp"
	"strconv"
	"strings"

	"weavelab.xyz/monorail/shared/go-utilities/dev"
	"weavelab.xyz/monorail/shared/wlib/werror"
)

const (
	ILike            = Operator("")
	Equal            = Operator("bool")
	DateInterval     = Operator("date_interval")
	NumberComparison = Operator("number")
	IsNull           = Operator("is_null")
	NotNull          = Operator("not_null")
	In               = Operator("in")
	NotIn            = Operator("not_in")
)

type filterSet struct {
	filters []filter
}

// Supplied from the user
type Filter struct {
	Field    string
	Operator Operator
	Value    string
}

type Operator string

func (o Operator) String() string {
	s := string(o)
	if v, ok := filterOperatorsToString[s]; ok {
		return v
	}
	return s
}

var filterOperators = []string{"!~", ">~", "<~", "~", ">", "<"}
var filterOperatorsToString = map[string]string{
	string(Equal): "=",
	"!~":          "!=",
	">~":          ">=",
	"<~":          "<=",
	"~":           "=",
}

var urlFilterRegex = regexp.MustCompile(`(` + strings.Join(filterOperators, "|") + `)`)

func NewFilter(values ...string) ([]Filter, error) {
	var filters []Filter
	for _, v := range values {
		f, err := ParseURLFilter(v)
		if err != nil {
			return nil, err
		}
		filters = append(filters, f)
	}
	return filters, nil
}

func ParseURLFilter(filter string) (Filter, error) {
	operator := urlFilterRegex.FindString(filter)
	if operator == "" {
		return Filter{}, werror.New("no operator found in filter").Add("filter", filter)
	}

	split := strings.SplitN(filter, operator, 2)
	if len(split) != 2 {
		return Filter{}, werror.New("unable to parse filter").Add("filter", filter)
	}

	return Filter{
		Field:    split[0],
		Operator: Operator(operator),
		Value:    split[1],
	}, nil
}

type filter struct {
	field     string
	fieldMap  []string
	fieldType Operator
	operator  Operator
	value     string
	values    []string
}

type FilterMap map[string]MapEntry

type MapEntry struct {
	Name   string
	Type   Operator
	Fields []string
}

func NewMapEntry(name string, operator Operator, fields ...string) MapEntry {
	name = strings.ToLower(name)

	return MapEntry{
		Name:   name,
		Type:   operator,
		Fields: fields,
	}
}

func NewFilterMap(maps ...MapEntry) FilterMap {
	f := make(map[string]MapEntry)
	for _, v := range maps {
		f[v.Name] = v
	}
	return f
}

func NewFilters(requested []Filter, allowed FilterMap) (filterSet, error) {

	var f filterSet

	if len(requested) == 0 {
		return filterSet{}, nil
	}

	for _, v := range requested {

		field := strings.ToLower(v.Field)
		// check to see if the requested filter exists in
		// the allowed filter map, if so add it
		var ok bool
		var fieldMap MapEntry
		if fieldMap, ok = allowed[field]; !ok {
			if dev.IsDev() {
				// burn the world
				return filterSet{}, werror.New("unknown filter").Add("filter", v.Field).Add("allowed filters", allowed)
			}
			continue
		}

		f.Add(filter{
			fieldMap:  fieldMap.Fields,
			fieldType: fieldMap.Type,

			field:    field,
			operator: v.Operator,
			value:    v.Value},
		)
	}

	return f, nil
}

func (f *filterSet) HasFilter(field string) bool {
	for _, v := range f.filters {
		if v.field == field {
			return true
		}
	}

	return false
}

func (f *filterSet) AddQuery(query []string, fieldMap FilterMap, mapper func(string) string) {

	terms := query

	// decide which fields to map where

	for _, v := range terms {

		if len(v) == 0 {
			continue
		}

		fieldToUse := mapper(v)

		mappedField, ok := fieldMap[fieldToUse]

		if !ok {
			log.Println("Unknown query field ", fieldToUse, " from mapper")
			continue
		}

		filter := filter{
			field:     fieldToUse,
			fieldMap:  mappedField.Fields,
			fieldType: mappedField.Type,
			operator:  "ILIKE",
			value:     v + "%",
		}

		f.Add(filter)
	}
}

func (f *filterSet) Add(newFilter filter) {
	if newFilter.operator == In || newFilter.operator == NotIn {
		for i, v := range f.filters {
			if newFilter.field == v.field && newFilter.operator == v.operator {
				v.values = append(v.values, newFilter.value)
				f.filters[i] = v
				return
			}
		}
	}

	newFilter.values = []string{newFilter.value}
	f.filters = append(f.filters, newFilter)

}

func (f *filterSet) ToSQL() (string, map[string]interface{}) {
	slice, parameters := f.ToSlice()
	sql := strings.Join(slice, " AND ")
	return sql, parameters
}

func (f *filterSet) ToSlice() ([]string, map[string]interface{}) {
	parameters := make(map[string]interface{})

	// pre-allocate space with a little extra
	var slice = make([]string, 0, len(f.filters)+1)

	// named parameter: :f{field}

	countAnd := 0
	for i, v := range f.filters {

		parameter := "f" + strconv.Itoa(i) + v.field
		if v.fieldType != "bool" && v.fieldType != "date_interval" && v.fieldType != "lowerEqual" {

			// this shouldn't be applied to all where operators
			switch v.operator {

			case "EQUAL":
				parameters[parameter] = v.value
			case "LIKE":
				parameters[parameter] = "%" + v.value
			case "~":
				v.operator = "ILIKE"
				if v.fieldType == "contains" {
					parameters[parameter] = "%" + v.value + "%"
				} else {
					parameters[parameter] = v.value
				}
			case "!~":
				v.operator = "NOT ILIKE"
				parameters[parameter] = v.value

			default:

				// do an exact match for filter searches, otherwise
				// patientId=9 will match 9, 9555, 5559, etc
				// status=active will match active, inactive, etc

				// match should still be case insensitive
				if v.fieldType == "contains" {
					parameters[parameter] = "%" + v.value + "%"
				} else {
					parameters[parameter] = v.value
				}
			}
		} else if v.fieldType == "lowerEqual" {
			parameters[parameter] = strings.ToLower(v.value)
		} else if v.operator == In || v.operator == NotIn {
			// do nothing here, parameters are handled separately
		} else {

			parameters[parameter] = v.value
		}

		sql := "("

		operator := v.operator.String()

		switch v.fieldType {
		case "", "contains", "bool", "number", "lowerEqual":

			countOr := 0
			for _, field := range v.fieldMap {
				if countOr > 0 {
					sql += " OR "
				}

				switch v.operator {
				case IsNull:
					sql += field + " IS NULL"
				case NotNull:
					sql += field + " IS NOT NULL"

				case In, NotIn:

					params := make([]string, len(v.values))

					for i, v := range v.values {
						p := parameter + "In" + strconv.Itoa(i)
						params[i] = ":" + p

						parameters[p] = v
					}

					op := " IN "
					if v.operator == NotIn {
						op = " NOT IN "
					}

					sql += field + op + "(" + strings.Join(params, ", ") + ")"

				default:
					sql += field + " " + operator + " :" + parameter
				}

				countOr++
			}

		case "date_interval":

			field := v.fieldMap[0]

			sql += field + ` ` + operator + ` current_date + :` + parameter + ` * interval '1 day'`

		case "birthday_range":
			// TODO: instead of current_date use date in practice's timezone and not weave's timezone

			// age at beginning of the interval is not equal to the age at the end of the interval
			// find the birthday end date range

			parameter_end_value := "31"
			for _, v := range f.filters {
				if v.fieldType == "birthday_range_end" {
					parameter_end_value = v.value
				}
			}
			parameter_end := parameter + "_end"
			parameters[parameter_end] = parameter_end_value

			sql += `(extract(year from ` + v.fieldMap[0] + `) > 1875) AND (extract(year from age(current_date + (:` + parameter + ` - 1) * INTERVAL '1 day', ` + v.fieldMap[0] + `)) != extract(year from age(current_date + (:` + parameter_end + `- 1)  * INTERVAL '1 day', ` + v.fieldMap[0] + `)))`

		case "date_range_interval":

			// TODO: breaks when the field is a leap year
			field := `date_part('doy',` + v.fieldMap[0] + `)`

			sql += field + ` ` + operator + ` date_part('doy', current_date) + :` + parameter

		case "date_range":
			field := v.fieldMap[0]

			sql += `date_part('doy',` + field + `) >= date_part('doy', current_date)
							AND date_part('doy',` + field + `) <= date_part('doy', current_date + :` + parameter + ` * interval '1 day')`

		default:
			sql += " 1=1 "
		}

		countAnd++

		sql += ")"

		slice = append(slice, sql)
	}

	return slice, parameters

}
