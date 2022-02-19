package types

import (
	"reflect"
	"testing"
)

const (
	assertEqual     = "AssertEqual(%v): got %v, want %v"
	assertNotEqual  = "AssertNotEqual(%v): got %v, want %v"
	assertDeepEqual = "AssertDeepEqual(%v): got %v, want %v"
	assertSameType  = "AssertSameType(%v): got %v, want %v"
	assertSameKind  = "AssertSameKind(%v): got %v, want %v"
)

func CompareFuncs(t *testing.T, name string, got, want Any) bool {
	g := NewAnyValue(got)
	if g.Kind() == reflect.Ptr {
		return CompareFuncs(t, name, reflect.Indirect(g.ValueOf()), want)
	}

	w := NewAnyValue(want)
	if w.Kind() == reflect.Ptr {
		return CompareFuncs(t, name, got, reflect.Indirect(w.ValueOf()))
	}
	if g.Kind() != reflect.Func {
		t.Errorf("invalid type for function compare(%v) = %v, want %v", name, g.Kind(), reflect.Func)
		return false
	}
	return true
	// w := NewAnyValue(want)
}

func AssertEqual(t *testing.T, name string, got, want Any) bool {
	if got == want {
		return true
	}
	t.Errorf(assertEqual, name, got, want)
	return false
}

func AssertNotEqual(t *testing.T, name string, got, want Any) bool {
	if got != want {
		return true
	}
	t.Errorf(assertNotEqual, name, got, want)
	return false
}

func AssertDeepEqual(t *testing.T, name string, got, want Any) bool {
	if reflect.DeepEqual(got, want) {
		return true
	}
	t.Errorf(assertDeepEqual, name, got, want)
	return false
}

func AssertSameType(t *testing.T, name string, got, want Any) bool {
	g := NewAnyValue(got).TypeOf()
	w := NewAnyValue(want).TypeOf()

	if g == w {
		return true
	}
	t.Errorf(assertSameType, name, g, w)
	return false
}

func AssertSameKind(t *testing.T, name string, got, want Any) bool {
	g := NewAnyValue(got).Kind()
	w := NewAnyValue(want).Kind()

	if g == w {
		return true
	}
	t.Errorf(assertSameKind, name, g, w)
	return false
}
