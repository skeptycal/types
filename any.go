package types

import "reflect"

// Any represents a object that may contain any
// valid type.
type Any interface{}

type any struct {
	v reflect.Value
	i interface{}
}

func (a *any) Value() Any         { return a.v }
func (a *any) Type() Any          { return a.v.Type() }
func (a *any) Interface() Any     { return a.i }
func (a *any) Kind() reflect.Kind { return a.v.Kind() }
func (a *any) IsComparable() bool { return a.Kind() > 1 && a.Kind() < 17 }
