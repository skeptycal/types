package types

import (
	"reflect"

	"github.com/fatih/structs"
)

var nilStruct = &structs.Struct{}

var ValueOf = reflect.ValueOf

func KindOf(a Any) reflect.Kind { return ValueOf(a).Kind() }
func TypeOf(a Any) reflect.Type { return ValueOf(a).Type() }

func NewStruct(v Any) *structs.Struct {
	if KindOf(v) != reflect.Struct {
		return nilStruct
	}
	return structs.New(v)
}
