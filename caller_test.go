package types

import (
	"reflect"
	"testing"
)

func TestNewCaller(t *testing.T) {
	type args struct {
		fn callerFunc
	}
	tests := []struct {
		name string
		args args
		want Caller
	}{
		// TODO: Add test cases.
		{"noop", args{fn: CallSetGlobalReturnValue}, &caller{fn: CallSetGlobalReturnValue, fnTrue: CallSetGlobalReturnValue, fnFalse: noop}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewCaller(tt.args.fn)
			if got := c.
			if got := ; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCaller() = %v, want %v", got, tt.want)
			}
		})
	}
}
