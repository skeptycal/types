package types

import "reflect"

// Any represents a object that may contain any
// valid type.
type Any interface{}

type any struct {
	v       reflect.Value
	k       reflect.Kind
	t       reflect.Type
	kindmap kindInfo
	i       interface{}
}

// NewAnyValue returns a new AnyValue interface, e.g.
// v := NewAnyValue(uint(42))
//
// AnyValue is a wrapper around the Any interface,
// or interface{}, which may contain any value.
// The original interface{} value is returned by
// v.Interface()
//
// The extra features of this wrapper allow
// value, type, and kind information, as well
// as whether the type is comparable, ordered,
// and/or iterable.
func NewAnyValue(a Any) AnyValue { return new_any(a) }

func new_any(a Any) *any {
	return &any{
		v:       ValueOf(a),
		kindmap: NewKindInfo(a),
		i:       a,
	}
}

func (a *any) ValueOf() reflect.Value { return ValueOf(a.i) }

// Kind returns the type of the underlying variable.
// This is cached in a jit struct field upon request.
func (a *any) TypeOf() reflect.Type {

	if a.t == nil {
		a.t = TypeOf(a.i)
	}

	return a.t
}

func (a *any) KindInfo() KindInfo {
	return a.kindmap
}

// Kind returns the kind of type of the underlying variable.
// This is cached in a jit struct field upon request.
func (a *any) Kind() reflect.Kind {
	if a.k == 0 {
		a.k = a.ValueOf().Kind()
	}
	return a.k
}
func (a *any) Interface() Any { return a.i }
func (a *any) IsComparable() bool {
	// return KindMaps[a.Kind()].IsComparable()

	return a.KindInfo().IsComparable()
}

// func (a *any) IsOrdered() bool        { return KindMaps.IsOrdered(a.Kind()) }
func (a *any) IsOrdered() bool        { return a.KindInfo().IsOrdered() }
func (a *any) IsDeepComparable() bool { return a.KindInfo().IsDeepComparable() }
func (a *any) IsIterable() bool       { return a.KindInfo().IsIterable() }
func (a *any) HasAlternate() bool     { return a.KindInfo().HasAlternate() }

// AnyValue is a wrapper around the Any interface,
// or interface{}, which may contain any value.
// The extra features of this wrapper allow
// value, type, and kind information, as well
// as whether the type is comparable, ordered,
// and/or iterable.
type AnyValue interface {

	// ValueOf returns a new Value initialized to the
	// concrete value stored in the interface i.
	// ValueOf(nil) returns the zero Value.
	ValueOf() reflect.Value

	// TypeOf returns the object's type.
	TypeOf() reflect.Type

	// Kind returns v's Kind. If v is the zero Value
	// (IsValid returns false), Kind returns Invalid.
	Kind() reflect.Kind

	// Interface returns the original underlying interface.
	Interface() Any

	KindInfo
}
