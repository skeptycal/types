package types

import (
	"fmt"
	"reflect"
)

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
