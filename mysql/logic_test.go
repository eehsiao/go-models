// Author :		Eric<eehsiao@gmail.com>

package mysql

import (
	"reflect"
	"testing"
)

func TestIif(t *testing.T) {
	type args struct {
		l bool
		a interface{}
		b interface{}
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			name: "case 1",
			args: args{
				l: true,
				a: "a",
				b: "b",
			},
			want: "a",
		},
		{
			name: "case 1",
			args: args{
				l: false,
				a: "a",
				b: "b",
			},
			want: "b",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Iif(tt.args.l, tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Iif() = %v, want %v", got, tt.want)
			}
		})
	}
}
