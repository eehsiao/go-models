// Author :		Eric<eehsiao@gmail.com>

package model

import (
	"encoding/json"
	"reflect"
	"strings"
)

// Struct4Scan : transfer struct to slice for scan
func Struct4Scan(s interface{}) (r []interface{}) {
	if s != nil {
		vals := reflect.ValueOf(s).Elem()
		for i := 0; i < vals.NumField(); i++ {
			r = append(r, vals.Field(i).Addr().Interface())
		}
	}

	return
}

// Struce4Query : transfer struct to string for query
func Struce4Query(r reflect.Type) (s string) {
	if r != nil && r.NumField() > 0 {
		var f []string
		for i := 0; i < r.NumField(); i++ {
			f = append(f, r.Field(i).Tag.Get(TableFieldTag))
		}

		s = strings.Join(f, ", ")
	}

	return
}

// Struce4Query : transfer struct to string for query
func Struce4QuerySlice(r reflect.Type) (s []string) {
	if r != nil && r.NumField() > 0 {
		for i := 0; i < r.NumField(); i++ {
			s = append(s, r.Field(i).Tag.Get(TableFieldTag))
		}
	}

	return
}

// Serialize : transfer object to string, the object's members must be public
func Serialize(i interface{}) (serialString string, err error) {
	bytes, err := json.Marshal(i)
	serialString = string(bytes)

	return
}
