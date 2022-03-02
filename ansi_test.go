package types

import (
	"os"
	"testing"
)

func TestAttn(t *testing.T) {
	tests := []struct {
		name   string
		format string
		args   []interface{}
	}{
		// TODO: Add test cases.
		{"ATTN example ", "(yellow bold text on red background)", []interface{}{FgYellow, BgHiRed, Bold}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Attn(tt.format, tt.args...)
		})
	}
}

func ExampleNewColor() {
	in := "(yellow bold log INFO text on red background)"
	NewColor(FgYellow, BgHiRed, Bold).Fprintln(os.Stdout, in)
	// Attn(in)

	// output:
	// (yellow bold log INFO text on red background)
}
