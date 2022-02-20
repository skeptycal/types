package types

import (
	"github.com/fatih/color"
	// . "github.com/fatih/color"
)

// NoColor defines if the output is colorized or not. It's dynamically set to false or true based on the stdout's file descriptor referring to a terminal or not. It's also set to true if the NO_COLOR environment variable is set (regardless of its value). This is a global option and affects all colors. For more control over each color block use the methods DisableColor() individually.
var NoColor = color.NoColor

// NewColor returns a newly created color object.
func NewColor(value ...color.Attribute) *color.Color {
	return color.New(value...)
}