package wsql

import (
	"reflect"

	"weavelab.xyz/wlib/werror"
)

const (
	dbTagName = "db"
)

func ValidColumns(a interface{}, other ...string) ([]string, error) {

	// extract the column names and pointers to the values for later use
	t := reflect.TypeOf(a)

	// if it is a pointer then get the underlying struct
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	numFields := t.NumField()

	var columns = make([]string, 0, numFields)
	// loop through the fields and get the tags
	for i := 0; i < numFields; i++ {
		f := t.Field(i)
		if f.PkgPath != "" {
			continue
		}

		col := f.Name

		tag := f.Tag
		dbTag := tag.Get(dbTagName)
		if dbTag != "" {
			col = dbTag
		}

		columns = append(columns, col)
	}

	return columns, nil
}

func ColumnValues(a interface{}) ([]interface{}, error) {

	v := reflect.ValueOf(a)
	k := v.Kind()

	if k == reflect.Ptr {
		v = reflect.Indirect(v)
		k = v.Kind()
	}

	if k != reflect.Struct {
		return nil, werror.New("unknown type").Add("type", a)
	}

	numFields := v.NumField()

	var values = make([]interface{}, 0, numFields)

	// loop through the fields and get the tags
	for i := 0; i < numFields; i++ {

		fv := v.Field(i)
		if fv.CanSet() {
			values = append(values, fv.Interface())
		}
	}

	return values, nil
}
