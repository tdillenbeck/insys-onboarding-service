// Copyright 2012 Kamil Kisiel. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// MODIFIED BY WEAVE 2014

package null

/*
Package sqlstruct provides some convenience functions for using structs with
the Go standard library's database/sql package.

The package matches struct field names to SQL query column names. A field can
also specify a matching column with "sql" tag, if it's different from field
name.  Unexported fields or fields marked with `sql:"-"` are ignored, just like
with "encoding/json" package.

*/

import (
	"database/sql/driver"
	"fmt"
	"reflect"
	"strings"
	"time"

	"weavelab.xyz/monorail/shared/wlib/werror"
)

// tagName is the name of the tag to use on struct fields
// MODIFIED
const tagName = "db"

// fieldInfo is a mapping of field tag values to their indices
type fieldInfo map[string]interface{}

// getFieldInfo creates a fieldInfo for the provided type. Fields that are not tagged
// with the {tagName} tag and unexported fields ARE NOT included.
func getFieldInfo(s interface{}, useMagic bool) (fieldInfo, error) {
	if s == nil {
		return nil, werror.New("Cannot get field info for nil")
	}

	v := reflect.ValueOf(s)

	//Check if s is a pointer, if so get the value it points to.
	if v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return nil, werror.New("Cannot get field info for nil")
		}
		v = v.Elem()
	}

	typ := v.Type()

	finfo := make(fieldInfo)

	n := typ.NumField()
	for i := 0; i < n; i++ {
		f := typ.Field(i)
		tag := f.Tag.Get(tagName)

		// Skip unexported fields or fields marked with "-"
		if f.PkgPath != "" || tag == "-" {
			continue
		}

		// Handle embedded structs or structs tagged with a '+'
		if (f.Anonymous || tag == "+") && f.Type.Kind() == reflect.Struct {
			r, err := getFieldInfo(v.Field(i).Interface(), useMagic)
			if err != nil {
				return nil, err
			}
			for k, v := range r {
				addToFieldInfo(finfo, k, v, true, useMagic)
			}
			continue
		}

		// Only add if the user supplied the value

		vfi := v.Field(i).Interface()
		var value interface{}

		var valid bool
		switch t := vfi.(type) {
		case driver.Valuer:
			var err error
			value, err = t.Value()
			valid = true
			if err != nil || value == nil {
				valid = false
			}

		case time.Time:
			valid = !t.IsZero()
			value = t

		case string:
			valid = len(t) != 0
			value = t

		case bool, float32, float64, int, int32, int64:
			valid = true
			value = t

		default:
			return nil, fmt.Errorf("SQLStruct - Unhandled type %s %T", f.Name, t)

		}

		// Use field name for untagged fields
		if tag == "" {
			tag = strings.ToLower(f.Name)
		}

		addToFieldInfo(finfo, tag, value, valid, useMagic)

	}

	return finfo, nil
}

func addToFieldInfo(m map[string]interface{}, field string, value interface{}, valid bool, useMagic bool) {

	if useMagic {
		isMagic, include, newValue := magicTimestamp(field, value)
		if isMagic {
			if include == false {
				return
			} else {
				value = newValue
				valid = true
			}
		}
	}

	if valid == false {
		return
	}

	m[field] = value

}

// Columns returns a string containing comma-separated list of column names as
// defined by the type s. s must be a struct that has exported fields tagged with the "db" tag.
func Columns(s interface{}) string {
	// handle nulls....
	// Coalesce
	columns, _, _ := cols(s, false)
	return strings.Join(columns, ", ")
}

func ColumnsValues(s interface{}) ([]string, map[string]interface{}, error) {
	columns, values, err := cols(s, false)
	if err != nil {
		return nil, nil, err
	}

	return columns, values, nil
}

// ColumnsValuesTimestamps is the same as ColumnsValues, but treats
// timestamps (created_?at and modified_?at) magically
func ColumnsValuesTimestamps(s interface{}) ([]string, map[string]interface{}, error) {
	columns, values, err := cols(s, true)
	if err != nil {
		return nil, nil, err
	}

	return columns, values, nil
}

func cols(s interface{}, magicTimestamps bool) ([]string, map[string]interface{}, error) {

	fields, err := getFieldInfo(s, magicTimestamps)
	if err != nil {
		return nil, nil, err
	}

	values := make(map[string]interface{})
	names := make([]string, 0, len(fields))
	for f, value := range fields {
		names = append(names, f+" = :"+f)
		values[f] = value
	}

	return names, values, nil
}

// magicTimestamp
// Returns:
//		bool: is magic timestamp or not
// 		bool: if magic, whether or not to include in query
//		bool: if magic, value to use, in place of given value
//
// Handles created_at, createdat, modified_at, modifiedat
// assumes incoming name is always lowercase
func magicTimestamp(name string, value interface{}) (bool, bool, interface{}) {

	switch name {
	case "created_at", "createdat":
		return true, false, nil
	case "modified_at", "modifiedat":
		return true, true, "NOW()"
	}

	return false, false, value
}
