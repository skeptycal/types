package types

import (
	"reflect"

	"github.com/fatih/structs"
)

var nilStruct = &structs.Struct{}

func ValueOf(i interface{}) reflect.Value { return reflect.ValueOf(i) }
func KindOf(a Any) reflect.Kind           { return ValueOf(a).Kind() }
func TypeOf(a Any) reflect.Type {

	if a == nil {
		return reflect.Type(nil)
	}

	return ValueOf(a).Type()
}

func NewStruct(v Any) *structs.Struct {
	if KindOf(v) != reflect.Struct {
		return nilStruct
	}
	return structs.New(v)
}
