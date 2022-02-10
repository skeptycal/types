package types

import (
	"reflect"

	"github.com/fatih/structs"
)

var ValueOf = reflect.ValueOf

func Kind(a Any) reflect.Kind   { return ValueOf(a).Kind() }
func TypeOf(a Any) reflect.Type { return ValueOf(a).Type() }

func NewStruct(v Any) *structs.Struct {
	if Kind(v) != reflect.Struct {
		return nil
	}
	return structs.New(v)
}
