package wsql

import (
	"reflect"
)

func sliceLength(a interface{}) int {
	if a == nil {
		return 0
	}

	v := reflect.ValueOf(a)

	if v.Type().Kind() != reflect.Slice {
		return 0
	}

	return v.Len()

}
