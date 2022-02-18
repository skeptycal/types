package types

import (
	"reflect"
	"testing"

	"github.com/fatih/color"
)

func TestNewColor(t *testing.T) {
	type args struct {
		value []color.Attribute
	}
	tests := []struct {
		name string
		args args
		want *color.Color
	}{
		{"reset", args{value: []color.Attribute{color.Reset}}, color.New(color.Reset)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewColor(tt.args.value...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewColor() = %v, want %v", got, tt.want)
			}
		})
	}
}
