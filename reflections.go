package types

import (
	"reflect"

	"github.com/fatih/structs"
)

var (
	ValueOf = reflect.ValueOf

	nilStruct = &structs.Struct{}
)

func KindOf(a Any) reflect.Kind { return ValueOf(a).Kind() }
func TypeOf(a Any) reflect.Type { return ValueOf(a).Type() }

func NewStruct(v Any) *structs.Struct {
	if KindOf(v) != reflect.Struct {
		return nilStruct
	}
	return structs.New(v)
}

// A Kind represents the specific kind of type that a Type represents.
// The zero Kind is not a valid kind.
//
// Reference: standard library reflect package
//	 type Kind uint
//
//	 const (
//	 	Invalid Kind = iota
//	 	Bool
//	 	Int
//	 	Int8
//	 	Int16
//	 	Int32
//	 	Int64
//	 	Uint
//	 	Uint8
//	 	Uint16
//	 	Uint32
//	 	Uint64
//	 	Uintptr
//	 	Float32
//	 	Float64
//	 	Complex64
//	 	Complex128
//	 	Array
//	 	Chan
//	 	Func
//	 	Interface
//	 	Map
//	 	Ptr
//	 	Slice
//	 	String
//	 	Struct
//	 	UnsafePointer
//	 )
type Kind = reflect.Kind
