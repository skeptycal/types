package types_test

import (
	"reflect"
	"testing"

	. "github.com/skeptycal/types"
)

func TestNewColor(t *testing.T) {
	type args struct {
		value []Attribute
	}
	tests := []struct {
		name string
		args args
		want *Color
	}{
		{"reset", args{value: []Attribute{Reset}}, NewColor(Reset)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewColor(tt.args.value...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewColor() = %v, want %v", got, tt.want)
			}
		})
	}
}
