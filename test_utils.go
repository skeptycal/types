package types

import (
	"fmt"
	"reflect"
	"testing"
)

var (
	LimitResult            bool
	DefaultTestResultLimit = 15
)

func limitTestResultLength(v Any) string {
	s := fmt.Sprintf("%v", v)

	if len(s) > DefaultTestResultLimit && LimitResult {
		return s[:DefaultTestResultLimit-3] + "..."
	}

	return s
}

func tName(testname, funcname, argname Any) string {
	if argname == "" {
		return fmt.Sprintf("%v: %v()", testname, funcname)
	}
	return fmt.Sprintf("%v: %v(%v)", testname, funcname, argname)
}

func typeGuardExclude(needle Any, notAllowed []Any) bool {
	return !Contains(needle, notAllowed)
}

func tTypeError(t *testing.T, name string, got, want Any) {
	t.Errorf("%v = %v(%T), want %v(%T)", name, limitTestResultLength(got), got, limitTestResultLength(want), want)
}

func tError(t *testing.T, name string, got, want Any) {
	t.Errorf("%v = %v, want %v", name, limitTestResultLength(got), limitTestResultLength(want))
}

func tRun(t *testing.T, name string, got, want Any) {

	gotV := ValueOf(got)
	wantV := ValueOf(want)
	if NewAnyValue(got).IsComparable() && NewAnyValue(want).IsComparable() {
		t.Run(name, func(t *testing.T) {
			if gotV != wantV {
				if !reflect.DeepEqual(got, want) {
					tTypeError(t, name, got, want)
				}
			}
		})
	}
}
