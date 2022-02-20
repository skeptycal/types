package types

import (
	"fmt"
	"os"
	"reflect"

	"github.com/mattn/go-isatty"
)

var (

	// NoColor defines if the output is colorized or not. It's dynamically set to
	// false or true based on the stdout's file descriptor referring to a terminal
	// or not. It's also set to true if the NO_COLOR environment variable is
	// set (regardless of its value). This is a global option and affects all
	// colors. For more control over each color block use the methods
	// DisableColor() individually.
	//
	// Reference: color.NoColor from https://github.com/fatih/color
	noColor bool = noColorEnvExists() || os.Getenv("TERM") == "dumb" ||
		(!isatty.IsTerminal(os.Stdout.Fd()) && !isatty.IsCygwinTerminal(os.Stdout.Fd()))
	// NoColor = color.NoColor
)

// noColorEnvExists returns true if the environment variable NO_COLOR exists.
//
// Reference: color.noColorExists from https://github.com/fatih/color
func noColorEnvExists() bool {
	_, exists := os.LookupEnv("NO_COLOR")
	return exists
}

// WithLock runs fn while holding lk.
func WithLock(lk Locker, fn func()) {
	defer lk.Unlock() // in case fn panics
	lk.Lock()
	fn()
}

// IsComparable returns true if the underlying value
// is of a type that is capable of comparisions, e.g.
// equal, not equal
//
// Bools, strings and most numeric values are comparable.
//
// Next, types that have a Len() method are considered
// comparable by this function based on their length and
// item type alone. This is different from the standard
// library approach.
func IsComparable(v Any) bool { return new_any(v).IsComparable() }

// IsOrdered returns true if the underlying value is ordered.
// This means that it is capable of order based comparisons, e.g.
// less than, greater than
//
// Strings and most numeric values are ordered.
//
//
func IsOrdered(v Any) bool { return new_any(v).IsOrdered() }

// IsDeepComparable returns true if the underlying value
// is of a type that is capable of DeepEqual, the Go
// standard library approach to rigorous comparisons.
func IsDeepComparable(v Any) bool { return new_any(v).IsDeepComparable() }

// IsIterable returns true if the underlying value is
// made up of smaller units that can be read out one by
// one.
//
// Maps, strings, and slices naturally come to mind, but this
// package also adds functionality to iterate over most numeric
// values and structs.
func IsIterable(v Any) bool { return new_any(v).IsIterable() }

// HasAlternate returns true if the underlying value has
// alternate methods in addition to the Go standard library
// operations.
func HasAlternate(v Any) bool { return new_any(v).HasAlternate() }

func isComparable1(v Any) bool {

	if _, ok := ValueOf(v).Type().MethodByName("Len"); ok {
		return true
	}

	switch v.(type) {
	case bool:
		return true
	case int, uint, float64, float32, string, []byte:
		return true
	default:
		return false
	}
}

func isComparable2(v reflect.Value) bool {
	if IsOrdered(v) {
		return true
	}
	if v.Kind() == reflect.Bool {
		return true
	}
	return false
}

func isOrdered2(v reflect.Value) bool {
	return v.Kind() > 1 && v.Kind() < 17

}

// Contains returns true if the underlying iterable
// sequence (haystack) contains the search term
// (needle) in at least one position.
func Contains(needle Any, haystack []Any) bool {
	for _, x := range haystack {
		if reflect.DeepEqual(needle, x) {
			return true
		}
	}
	return false
}

// Count returns the number of times the search term
// (needle) occurs in the underlying iterable
// sequence (haystack).
func Count(needle Any, haystack []Any) int {
	retval := 0
	for _, x := range haystack {
		if reflect.DeepEqual(needle, x) {
			retval += 1
		}
	}
	return retval
}

// ToString converts the given argument to the
// standard string representation. If a implements
// fmt.Stringer, it is used, otherwise the slower
// fmt.Sprintf is used as a backup.
func ToString(a Any) string {
	if v, ok := a.(fmt.Stringer); ok {
		return v.String()
	}
	return fmt.Sprintf("%v", a)
}