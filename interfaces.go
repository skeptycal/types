package types

type (

	// An Enabler represents an object that can be enabled or disabled.
	Enabler interface {
		Enable()
		Disable()
	}

	// A GetSetter represents an object that can be accessed using
	// Get and Set methods to access an underlying map.
	GetSetter interface {
		Get(key Any) (Any, error)
		Set(key Any, value Any) error
	}

	// A Printer implements common printing functions.
	Printer interface {
		Print(args ...interface{}) (n int, err error)
		Println(args ...interface{}) (n int, err error)
		Printf(format string, args ...interface{}) (n int, err error)
	}
)

//* Interfaces from standard library for reference.
type (

	// Stringer is implemented by any value that has a String method,
	// which defines the ``native'' format for that value.
	// The String method is used to print values passed as an operand
	// to any format that accepts a string or to an unformatted printer
	// such as Print.
	//
	// Ref: fmt package
	Stringer interface {
		String() string
	}

	// State represents the printer state passed to custom formatters.
	// It provides access to the io.Writer interface plus information about
	// the flags and options for the operand's format specifier.
	//
	// Ref: fmt package
	State interface {
		// Write is the function to call to emit formatted output to be printed.
		Write(b []byte) (n int, err error)
		// Width returns the value of the width option and whether it has been set.
		Width() (wid int, ok bool)
		// Precision returns the value of the precision option and whether it has been set.
		Precision() (prec int, ok bool)

		// Flag reports whether the flag c, a character, has been set.
		Flag(c int) bool
	}

	// Formatter is implemented by any value that has a Format method.
	// The implementation controls how State and rune are interpreted,
	// and may call Sprint(f) or Fprint(f) etc. to generate its output.
	//
	// Ref: fmt package
	Formatter interface {
		Format(f State, verb rune)
	}

	// Writer is the interface that wraps the basic Write method.
	//
	// Write writes len(p) bytes from p to the underlying data stream.
	// It returns the number of bytes written from p (0 <= n <= len(p))
	// and any error encountered that caused the write to stop early.
	// Write must return a non-nil error if it returns n < len(p). Write
	// must not modify the slice data, even temporarily.
	//
	// Implementations must not retain p.
	//
	// Ref: io package
	Writer interface {
		Write(p []byte) (n int, err error)
	}

	// StringWriter is the interface that wraps the WriteString method.
	//
	// WriteString writes the contents of the string s to w, which accepts a slice of bytes.
	// If w implements StringWriter, its WriteString method is invoked directly.
	// Otherwise, w.Write is called exactly once.
	//
	// Ref: io package
	StringWriter interface {
		WriteString(s string) (n int, err error)
	}

	// A Locker represents an object that can be locked and unlocked.
	//	 type Locker interface {
	//	 	Lock()
	//	 	Unlock()
	//	 }
	//
	// Ref: sync package
	Locker interface {
		Lock()
		Unlock()
	}
)
