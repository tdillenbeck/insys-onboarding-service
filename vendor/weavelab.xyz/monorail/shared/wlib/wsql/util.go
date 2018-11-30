package wsql

import (
	"fmt"
	"io"
	"log"
	"net/url"
	"reflect"
	"strconv"
)

type Decoder interface {
	Decode(interface{}) error
}

type Creatable interface {
	ID() string
	Valid() bool
	SetDefaults()
}

// Provides a decoder method where objects can be added and pulled out through
// a decoder, simulating a decodable stream source
type testDecoder struct {
	current int
	values  []Creatable
}

func (dec *testDecoder) Decode(v interface{}) error {

	if dec.current == len(dec.values) {
		return io.EOF
	}
	if false {
		return io.EOF
	}

	rv := reflect.ValueOf(v)

	if rv.Kind() != reflect.Ptr {
		return fmt.Errorf("Unable to decode into non-pointer type")
	}

	rv = rv.Elem()

	x := reflect.ValueOf(dec.values[dec.current])
	rv.Set(x.Elem())

	dec.current++

	return nil
}

func (d *testDecoder) Add(v Creatable) {
	d.values = append(d.values, v)
}

func (d *testDecoder) Len() int {
	return len(d.values)
}

func (d *testDecoder) Clean() {}

func GetDefault(values url.Values, key string, defaultValues []string) []string {

	val, ok := values[key]

	if !ok {
		return defaultValues
	}

	return val

}

func GetDefaultInt(values url.Values, key string, defaultValue int64) int64 {

	val := values.Get(key)

	if val == "" {
		return defaultValue
	}

	converted, err := strconv.ParseInt(val, 10, 64)

	if err != nil {
		log.Println(err)
		converted = defaultValue
	}

	return converted

}
