package types

type (
	Errer interface {
		Err() error
	}

	Closer interface {
		Close() error
	}

	// An Enabler represents an object that can be enabled or disabled.
	Enabler interface {
		Enable()
		Disable()
	}

	// Protector is used when an object needs to be protected or unprotected (inspired by the "write-protect" tabs of floppy disks)
	Protector interface {
		Protect()
		Unprotect()
	}

	// A GetSetter represents an object that can be accessed using
	// Get and Set methods to access an underlying map.
	//
	// Example methods:
	//  func (p *padding) Get(key Any) (Any, error) {
	//		// TODO get the value that matches the key
	//  	return nil, nil
	//  }
	//  func (p *padding) Set(key, value Any) error {
	//		// TODO set the value that matches the key
	//  	return nil
	//  }
	GetSetter interface {
		Get(key Any) (Any, error)
		Set(key Any, value Any) error
	}

	// Slicer returns the slice of keys and values that are
	// asoociated with the underlying data structure.
	//
	// Example methods:
	// 	func (d *dict) Keys() []Any {
	//		// TODO return a list of keys
	// 		keys := make([]Any, len(d.m))
	//	 	for k := range d.m {
	//	 		keys = append(keys, k)
	//	 	}
	//	 	return keys
	//	 }
	//	 func (d *dict) Values() []Any {
	//		// TODO return a list of values
	//	 	values := make([]Any, len(d.m))
	//	 	for _, v := range d.m {
	//	 		values = append(values, v)
	//	 	}
	//	 	return values
	//	 }
	Slicer interface {
		Keys() []Any
		Values() []Any
	}

	// A Printer implements common printing functions.
	//
	// Example methods:
	// 	func (p *padding) Print(args ...interface{}) (n int, err error) {
	//		// TODO Print unformatted args
	// 		return n, err
	// 	}
	// 	func (p *padding) Println(args ...interface{}) (n int, err error) {
	//		// TODO Print unformatted args with line break (NL)
	// 		return n, err
	// 	}
	// 	func (p *padding) Printf(format string, args ...interface{}) (n int, err error) {
	//		// TODO Print formatted args
	// 		return n, err
	// 	}
	Printer interface {
		Print(args ...interface{}) (n int, err error)
		Println(args ...interface{}) (n int, err error)
		Printf(format string, args ...interface{}) (n int, err error)
	}

	// Dict is a dictionary used to store keys and values of any type.
	// It is not concurrent-safe.
	Dict interface {
		Delete(key Any) error

		GetSetter
		Protector
		Slicer
		Stringer
	}
)

//* Interfaces from standard library for reference.
// Copied here to avoid larger dependencies.
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
