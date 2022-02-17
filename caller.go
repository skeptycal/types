package types

var (
	globalReturnValue interface{}
	RetVal            = NewCaller(CallSetGlobalReturnValue)
)

func CallSetGlobalReturnValue(any interface{}) { globalReturnValue = any }

type (
	Caller interface {
		Call(interface{})
		Enabler
	}

	caller struct {
		fn      func(v interface{})
		fnFalse func(v interface{})
		fnTrue  func(v interface{})
	}
)

// NewCaller returns a new instance of Caller
// with the function enabled.
func NewCaller(fn func(v interface{})) Caller {
	d := caller{
		fn:      fn,
		fnTrue:  fn,
		fnFalse: func(v interface{}) {},
	}
	d.Enable()
	return &d
}

func (d *caller) Call(v interface{}) { d.fn(v) }
func (d *caller) Enable()            { d.fn = d.fnTrue }
func (d *caller) Disable()           { d.fn = d.fnFalse }

func noop(any interface{}) {} // noop function
