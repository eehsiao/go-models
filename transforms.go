// Author :		Eric<eehsiao@gmail.com>

package model

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	sb "github.com/eehsiao/sqlbuilder"
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

// Struce4Query : transfer struct to string for query
func Inst2Fields(t interface{}) (s []string) {
	r := reflect.TypeOf(t)
	if r != nil && r.NumField() > 0 {
		for i := 0; i < r.NumField(); i++ {
			s = append(s, r.Field(i).Tag.Get(TableFieldTag))
		}
	}

	return
}

// Struce4Query : transfer struct to string for query
func Inst2FieldWithoutID(t interface{}) (s []string) {
	r := reflect.TypeOf(t)
	if r != nil && r.NumField() > 0 {
		for i := 0; i < r.NumField(); i++ {
			f := r.Field(i).Tag.Get(TableFieldTag)
			if f != "id" && f != "idx" {
				s = append(s, f)
			}
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

func Inst2Set(t interface{}, wout ...string) (s []sb.Set) {
	r := reflect.TypeOf(t)
	rv := reflect.ValueOf(t)
	if r != nil && r.NumField() > 0 {
		for i := 0; i < r.NumField(); i++ {
			var (
				f   string
				v   interface{}
				brk bool
			)
			f = r.Field(i).Tag.Get(TableFieldTag)
			v = rv.Field(i).Interface()

			fmt.Println(reflect.TypeOf(v).Name())
			switch strings.ReplaceAll(reflect.TypeOf(v).Name(), "sql.", "") {
			case "NullBool":
				v = Iif(v.(sql.NullBool).Valid, v.(sql.NullBool).Bool, nil)
			case "NullString":
				v = Iif(v.(sql.NullString).Valid, v.(sql.NullString).String, nil)
			case "NullFloat64":
				v = Iif(v.(sql.NullFloat64).Valid, v.(sql.NullFloat64).Float64, nil)
			case "NullInt32":
				v = Iif(v.(sql.NullInt32).Valid, v.(sql.NullInt32).Int32, nil)
			case "NullInt64":
				v = Iif(v.(sql.NullInt64).Valid, v.(sql.NullInt64).Int64, nil)
			case "NullTime":
				v = Iif(v.(sql.NullTime).Valid, v.(sql.NullTime).Time, nil)
			}

			brk = true
			for _, w := range wout {
				if f == w {
					brk = true
					continue
				}
			}
			if brk {
				s = append(s, sb.Set{f, v})
			}
		}
	}

	return
}
