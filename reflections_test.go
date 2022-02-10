package types

import (
	"reflect"
	"testing"
	"unsafe"

	"github.com/fatih/structs"
)

func TestKind(t *testing.T) {
	tests := []struct {
		name string
		a    Any
		want reflect.Kind
	}{
		{"invalid", nil, reflect.Invalid},
		{"bool", true, reflect.Bool},
		{"int", 42, reflect.Int},
		{"uint", uint(42), reflect.Uint},
		{"int", 42, reflect.Int},
		{"Int8", int8(42), reflect.Int8},
		{"Int16", int16(42), reflect.Int16},
		{"Int32", int32(42), reflect.Int32},
		{"Int64", int64(42), reflect.Int64},
		{"Uint", uint(42), reflect.Uint},
		{"Uint8", uint8(42), reflect.Uint8},
		{"Uint16", uint16(42), reflect.Uint16},
		{"Uint32", uint32(42), reflect.Uint32},
		{"Uint64", uint64(42), reflect.Uint64},
		{"Uintptr", uintptr(42), reflect.Uintptr},
		{"Float32", float32(42), reflect.Float32},
		{"Float64", float64(42), reflect.Float64},
		{"Complex64", complex64(42), reflect.Complex64},
		{"Complex128", complex128(42), reflect.Complex128},
		{"Array", [4]int{42, 42, 42, 42}, reflect.Array},
		{"Chan", make(chan int, 1), reflect.Chan},
		{"Func", t.Run, reflect.Func},
		// {"Interface", nil, reflect.Interface},
		{"Map", make(map[string]interface{}), reflect.Map},
		{"Ptr", &ValueOf, reflect.Ptr},
		{"Slice", []int{42}, reflect.Slice},
		{"String", "42", reflect.String},
		{"Struct", ValueOf(42), reflect.Struct},
		{"UnsafePointer", unsafe.Pointer(nil), reflect.UnsafePointer},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Kind(tt.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Kind() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTypeOf(t *testing.T) {

	tests := []struct {
		name string
		a    Any
	}{
		// {"invalid", nil, nil}, // panic
		// {"Interface", nil},	  // panic

		{"bool", true},
		{"int", 42},
		{"uint", uint(42)},
		{"int", 42},
		{"Int8", int8(42)},
		{"Int16", int16(42)},
		{"Int32", int32(42)},
		{"Int64", int64(42)},
		{"Uint", uint(42)},
		{"Uint8", uint8(42)},
		{"Uint16", uint16(42)},
		{"Uint32", uint32(42)},
		{"Uint64", uint64(42)},
		{"Uintptr", uintptr(42)},
		{"Float32", float32(42)},
		{"Float64", float64(42)},
		{"Complex64", complex64(42)},
		{"Complex128", complex128(42)},
		{"Array", [4]int{42, 42, 42, 42}},
		{"Chan", make(chan int, 1)},
		{"Func", t.Run},
		{"Map", make(map[string]interface{})},
		{"Ptr", &ValueOf},
		{"Slice", []int{42}},
		{"String", "42"},
		{"Struct", ValueOf(42)},
		{"UnsafePointer", unsafe.Pointer(nil)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want := reflect.TypeOf(tt.a)
			if got := TypeOf(tt.a); !reflect.DeepEqual(got, want) {
				t.Errorf("TypeOf() = %v, want %v", got, want)
			}
		})
	}
}

func TestNewStruct(t *testing.T) {
	tests := []struct {
		name string
		v    Any
		want *structs.Struct
	}{
		{"42", Cosa{}, structs.New(Cosa{})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewStruct(tt.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewStruct() = %v, want %v", got, tt.want)
			}
		})
	}
}
