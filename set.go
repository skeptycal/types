package types

type (

	// Set implements the Set interface from
	// MIT 6.006 lecture 3: Sets and Sorting
	Set interface {
		Build() error
		Len() int

		// find(key Any) (value Any, found bool)
		// insert(key Any) error
		GetSetter
		Delete(key Any) error

		// IterOrd returns the stored items one by one in key order.
		IterOrd() []Any

		// Min returns the item with the smallest key.
		Min() Any

		// Max returns the item with the largest key.
		Max() Any

		// Prev returns the item with the smallest key larger than k.
		Prev(k Any) Any

		// Next returns the item with the largest key smaller than k.
		Next() Any
	}
)
