package types

type (
	// Cosa is a generic 'thing.' It is an empty stuct used for 'preallocating' zero resource objects
	Cosa struct{}

	// Any represents a object that may contain any
	// valid type.
	Any interface{}
)
