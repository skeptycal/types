package types

import (
	"reflect"
	"testing"
	"unsafe"

	"github.com/fatih/structs"
)

var testSample int = 42
var ptrSample = &testSample
var ReflectTests []struct {
	name string
	a    Any
	want reflect.Kind
}

func init() {
	copy(ReflectTests, reflectTests)
}

var reflectTests = []struct {
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
	{"Func", IsComparable, reflect.Func},
	{"Map", make(map[string]interface{}), reflect.Map},
	{"Ptr", ptrSample, reflect.Ptr},
	{"Slice", []int{42}, reflect.Slice},
	{"String", "42", reflect.String},
	{"UnsafePointer", unsafe.Pointer(nil), reflect.UnsafePointer},
	// {"Interface", nil, reflect.Interface},
	// {"ValueOf(42)", ValueOf(42), ValueOf(ValueOf(42)).Kind()},
}

func TestKind(t *testing.T) {
	t.Parallel()
	for _, tt := range reflectTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := KindOf(tt.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Kind() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTypeOf(t *testing.T) {
	t.Parallel()
	for _, tt := range reflectTests {
		t.Run(tt.name, func(t *testing.T) {
			want := reflect.TypeOf(tt.a)
			if got := TypeOf(tt.a); !reflect.DeepEqual(got, want) {
				t.Errorf("TypeOf() = %v, want %v", got, want)
			}
		})
	}
}

func TestNewStruct(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		v    Any
		want *structs.Struct
	}{
		{"Cosa", Cosa{}, structs.New(Cosa{})},
		{"int 42", nil, nilStruct},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewStruct(tt.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewStruct() = %v, want %v", got, tt.want)
			}
		})
	}
}
