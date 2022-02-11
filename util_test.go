package types

import (
	"fmt"
	"os"
	"sync"
	"testing"
	"time"
)

var setupDone bool = false

func SetupEnv(t *testing.T) {
	t.Helper()

	if !setupDone {
		t.Setenv("NO_COLOR", "true")
		setupDone = true
	}
}

func CleanupEnv(t *testing.T) {
	t.Helper()

	if setupDone {
		t.Cleanup(func() {
			os.Unsetenv("NO_COLOR")
		})
		setupDone = false
	}
}

func TestWithLock(t *testing.T) {
	lockfunc := func() {
		fmt.Sprintf("lock time: %v", time.Now())
	}
	type args struct {
		lk Locker
		fn func()
	}
	tests := []struct {
		name string
		args args
	}{
		{"sync.Mutex", args{&sync.Mutex{}, lockfunc}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			WithLock(tt.args.lk, tt.args.fn)
		})
	}
}

func TestContains(t *testing.T) {
	type args struct {
		needle   Any
		haystack []Any
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"true", args{"v", []Any{"v", "v"}}, true},
		{"wrong type", args{3, []Any{"v", "v"}}, false},
		{"wrong value", args{"d", []Any{"v", "v"}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Contains(tt.args.needle, tt.args.haystack); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCount(t *testing.T) {
	type args struct {
		needle   Any
		haystack []Any
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"true", args{"v", []Any{"v", "v"}}, 2},
		{"true", args{"f", []Any{"v", "v", "f"}}, 1},
		{"wrong type", args{3, []Any{"v", "v"}}, 0},
		{"wrong value", args{"d", []Any{"v", "v"}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Count(tt.args.needle, tt.args.haystack); got != tt.want {
				t.Errorf("Count() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToString(t *testing.T) {
	// p := uint(42)
	tests := []struct {
		name string
		a    Any
		want string
	}{
		{"int 42", 42, "42"},
		{"uint 42", uint(42), "42"},
		// {"pointer uint 42", &p, "address"},
		{"nil", nil, "<nil>"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToString(tt.a); got != tt.want {
				t.Errorf("ToString() = %v, want %v", got, tt.want)
			}
		})
	}
}
