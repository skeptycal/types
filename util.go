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

func Contains(needle Any, haystack []Any) bool {
	for _, x := range haystack {
		if reflect.DeepEqual(needle, x) {
			return true
		}
	}
	return false
}

func Count(needle Any, haystack []Any) int {
	retval := 0
	for _, x := range haystack {
		if reflect.DeepEqual(needle, x) {
			retval += 1
		}
	}
	return retval
}

func ToString(a Any) string { return fmt.Sprintf("%v", a) }
