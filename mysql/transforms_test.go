// Author :		Eric<eehsiao@gmail.com>

package mysql

import (
	"database/sql"
	"reflect"
	"testing"
)

type tbTest struct {
	Idx  int64          `TbField:"idx"`
	Name sql.NullString `TbField:"name"`
}

var (
	testTb = tbTest{}
)

func TestStruct4Scan(t *testing.T) {
	type args struct {
		s interface{}
	}

	tests := []struct {
		name  string
		args  args
		wantR []interface{}
	}{
		{
			name: "case 1",
			args: args{
				s: &testTb,
			},
			wantR: []interface{}{
				reflect.ValueOf(&testTb).Elem().Field(0).Addr().Interface(),
				reflect.ValueOf(&testTb).Elem().Field(1).Addr().Interface(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotR := Struct4Scan(tt.args.s); !reflect.DeepEqual(gotR, tt.wantR) {
				t.Errorf("Struct4Scan() = %v, want %v", gotR, tt.wantR)
			}
		})
	}
}

func TestStruce4Query(t *testing.T) {
	type args struct {
		r reflect.Type
	}
	tests := []struct {
		name  string
		args  args
		wantS string
	}{
		{
			name: "case 1",
			args: args{
				r: reflect.TypeOf(tbTest{}),
			},
			wantS: "idx, name",
		},
		{
			name: "case 2",
			args: args{
				r: reflect.TypeOf(args{}),
			},
			wantS: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotS := Struce4Query(tt.args.r); gotS != tt.wantS {
				t.Errorf("Struce4Query() = %v, want %v", gotS, tt.wantS)
			}
		})
	}
}

func TestSerialize(t *testing.T) {
	type args struct {
		i interface{}
	}
	tests := []struct {
		name             string
		args             args
		wantSerialString string
		wantErr          bool
	}{
		{
			name: "case 1",
			args: args{
				i: testTb,
			},
			wantSerialString: "{\"Idx\":0,\"Name\":{\"String\":\"\",\"Valid\":false}}",
			wantErr:          false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSerialString, err := Serialize(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("Serialize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotSerialString != tt.wantSerialString {
				t.Errorf("Serialize() = %v, want %v", gotSerialString, tt.wantSerialString)
			}
		})
	}
}
