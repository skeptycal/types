package types

import "fmt"

type (
	// AnyMap is a map used to store keys and values of any type.
	AnyMap map[Any]Any

	dict struct {
		name      string
		protected bool
		m         map[Any]Any
	}
)

func NewDict(name string, protected bool) *dict {

	return &dict{
		name:      name,
		protected: protected,
		m:         make(AnyMap),
	}
}

func (d dict) Keys() []Any {
	keys := make([]Any, len(d.m))
	for k := range d.m {
		keys = append(keys, k)
	}
	return keys
}

func (d dict) Values() []Any {
	values := make([]Any, len(d.m))
	for v := range d.m {
		values = append(values, v)
	}
	return values
}

func (d dict) Get(key Any) (Any, error) {
	if v, ok := d.m[key]; ok {
		return v, nil
	}
	return nil, fmt.Errorf("key not found: %v", key)
}

func (d dict) Set(key, value Any) error {
	if d.protected {
		return fmt.Errorf("dictionary is write protected")
	}
	if v, ok := d.m[key]; ok && v != value {
		d.m[key] = value
		return nil
	}

	return fmt.Errorf("error setting key: %v[%v]", key, value)
}

func (d dict) Delete(key Any) error {
	if !d.protected {
		delete(d.m, d.m[key])
		return nil
	}
	return fmt.Errorf("error deleting key: %v", key)
}

func (d dict) Protect() {
	d.protected = true
}

func (d dict) Unprotect() {
	d.protected = false
}
