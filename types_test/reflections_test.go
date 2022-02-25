package types_test

import (
	"io"
	"reflect"
	"testing"
	"unsafe"

	"github.com/fatih/structs"
	. "github.com/skeptycal/types"
)

var testSample int = 42
var ptrSample = &testSample

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
	{"io.ReadCloser", io.NopCloser(nil), reflect.Interface},
	{"Interface", nil, reflect.Interface},
	{"byte slice element", []byte("fake")[0], reflect.Uint8},
}

func Test_ValueOf(t *testing.T) {
	t.Parallel()
	for _, tt := range reflectTests {
		tRun(t, tt.name, NewAnyValue(tt.a).ValueOf(), ValueOf(tt.a))
	}
}

func Test_KindOf(t *testing.T) {
	t.Parallel()
	for _, tt := range reflectTests {
		tRun(t, tt.name, NewAnyValue(tt.a).Kind(), KindOf(tt.a))
	}
}

func Test_TypeOf(t *testing.T) {
	t.Parallel()
	for _, tt := range reflectTests {
		tRun(t, tt.name, NewAnyValue(tt.a).TypeOf(), TypeOf(tt.a))
	}
}

func Test_Indirect(t *testing.T) {
	t.Parallel()
	// LimitResult = true
	for _, tt := range reflectTests {
		want := ValueOf(tt.a)
		got := NewAnyValue(tt.a).Indirect().ValueOf()
		name := tName(tt.name, tt.name, tt.a)
		tRun(t, name, got, want)
	}
	// LimitResult = false
}

func Test_Addr(t *testing.T) {
	t.Parallel()
	for _, tt := range reflectTests {
		want := ValueOf(tt.a)
		got := Addr(want)

		if !want.CanAddr() {
			continue
		}

		name := tName(tt.name, tt.name, tt.a)
		tRun(t, name, got, want.Addr())
	}
}

func Test_Interface(t *testing.T) {
	t.Parallel()
	for _, tt := range reflectTests {
		v := ValueOf(tt.a)

		got := Interface(v)

		if !v.IsValid() {
			continue
		}

		if v.IsZero() {
			continue
		}

		if !v.CanInterface() {
			continue
		}

		name := tName(tt.name, tt.name, tt.a)
		tRun(t, name, got, v.Interface())
	}
}

func Test_Elem(t *testing.T) {
	t.Parallel()
	for _, tt := range reflectTests {
		want := ValueOf(tt.a)
		got := Elem(want)

		if want.Kind() != reflect.Ptr || want.Kind() != reflect.Interface {
			continue
		}
		name := tName(tt.name, tt.name, tt.a)
		tRun(t, name, got, want.Elem())
	}
}

func Test_Convert(t *testing.T) {
	t.Parallel()

	wantType := ValueOf(int(42)).Type()

	for _, tt := range reflectTests {
		v := ValueOf(tt.a)
		_ = Interface(Convert(v, wantType))

		if !v.IsValid() {
			continue
		}

		want := tt.a
		got := tt.a

		if v.CanConvert(wantType) {
			got = Convert(v, wantType).Interface()
			want = v.Convert(wantType).Interface()
		}

		name := tName(tt.name, tt.name, tt.a)
		tRun(t, name, got, want)
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

func guardReflectType(v reflect.Value) reflect.Type {
	if v.Kind() == reflect.Invalid {
		return reflect.Type(nil)
	}
	return v.Type()
}
