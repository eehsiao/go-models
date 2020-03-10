// Author :		Eric<eehsiao@gmail.com>

package model

import (
	"database/sql"
	"reflect"
	"testing"

	sb "github.com/eehsiao/sqlbuilder"
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

func TestInst2Fields(t *testing.T) {
	type args struct {
		r interface{}
	}
	tests := []struct {
		name  string
		args  args
		wantR []string
	}{
		{
			name: "case 1",
			args: args{
				r: tbTest{},
			},
			wantR: []string{"idx", "name"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotR := Inst2Fields(tt.args.r); !reflect.DeepEqual(gotR, tt.wantR) {
				t.Errorf("Inst2Fields() = %v, want %v", gotR, tt.wantR)
			}
		})
	}
}

func TestInst2FieldWithoutID(t *testing.T) {
	type args struct {
		r interface{}
	}
	tests := []struct {
		name  string
		args  args
		wantR []string
	}{
		{
			name: "case 1",
			args: args{
				r: tbTest{},
			},
			wantR: []string{"name"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotR := Inst2FieldWithoutID(tt.args.r); !reflect.DeepEqual(gotR, tt.wantR) {
				t.Errorf("Inst2FieldWithoutID() = %v, want %v", gotR, tt.wantR)
			}
		})
	}
}

func TestInst2Set(t *testing.T) {
	type args struct {
		r    interface{}
		wout []string
	}
	tests := []struct {
		name  string
		args  args
		wantS []sb.Set
	}{
		{
			name: "case 1",
			args: args{
				r:    tbTest{Idx: 1, Name: sql.NullString{"test", true}},
				wout: []string{},
			},
			wantS: []sb.Set{{"idx", int64(1)}, {"name", "test"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotS := Inst2Set(tt.args.r, tt.args.wout...); !reflect.DeepEqual(gotS, tt.wantS) {
				t.Errorf("Inst2Set() = %v, want %v", gotS, tt.wantS)
			}
		})
	}
}

func TestInst2FieldWout(t *testing.T) {
	type args struct {
		r    interface{}
		wout []string
	}
	tests := []struct {
		name  string
		args  args
		wantR []string
	}{
		{
			name: "case 1",
			args: args{
				r:    tbTest{},
				wout: []string{"idx"},
			},
			wantR: []string{"name"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotR := Inst2FieldWout(tt.args.r, tt.args.wout...); !reflect.DeepEqual(gotR, tt.wantR) {
				t.Errorf("Inst2FieldWout() = %v, want %v", gotR, tt.wantR)
			}
		})
	}
}

func TestInst2Values(t *testing.T) {
	type args struct {
		r    interface{}
		wout []string
	}
	tests := []struct {
		name  string
		args  args
		wantS []interface{}
	}{
		{
			name: "case 1",
			args: args{
				r:    tbTest{Idx: 1, Name: sql.NullString{"test", true}},
				wout: []string{},
			},
			wantS: []interface{}{int64(1), "test"},
		},
		{
			name: "case 2",
			args: args{
				r:    tbTest{Idx: 1, Name: sql.NullString{"test", true}},
				wout: []string{"idx"},
			},
			wantS: []interface{}{"test"},
		},
		{
			name: "case 3",
			args: args{
				r:    tbTest{Idx: 1, Name: sql.NullString{"test", true}},
				wout: []string{"name"},
			},
			wantS: []interface{}{int64(1)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotS := Inst2Values(tt.args.r, tt.args.wout...); !reflect.DeepEqual(gotS, tt.wantS) {
				t.Errorf("Inst2Values() = %v, want %v", gotS, tt.wantS)
			}
		})
	}
}
