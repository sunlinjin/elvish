// Package types contains basic types for the Elvish runtime.
package types

// Definitions for Value interfaces, some simple Value types and some common
// Value helpers.

// Value is an Elvish value.
type Value interface{}

// Booler wraps the Bool method.
type Booler interface {
	// Bool computes the truth value of the receiver.
	Bool() bool
}

// Stringer wraps the String method.
type Stringer interface {
	// Stringer converts the receiver to a string.
	String() string
}

// ToString converts a Value to string. When the Value type implements
// String(), it is used. Otherwise Repr(NoPretty) is used.
func ToString(v Value) string {
	if s, ok := v.(Stringer); ok {
		return s.String()
	}
	return Repr(v, NoPretty)
}

// Lener wraps the Len method.
type Lener interface {
	// Len computes the length of the receiver.
	Len() int
}

// Iterator wraps the Iterate method.
type Iterator interface {
	// Iterate calls the passed function with each value within the receiver.
	// The iteration is aborted if the function returns false.
	Iterate(func(v Value) bool)
}

// IteratorValue is an iterable Value.
type IteratorValue interface {
	Iterator
	Value
}

func CollectFromIterator(it Iterator) []Value {
	var vs []Value
	if lener, ok := it.(Lener); ok {
		vs = make([]Value, 0, lener.Len())
	}
	it.Iterate(func(v Value) bool {
		vs = append(vs, v)
		return true
	})
	return vs
}

// IterateKeyer wraps the IterateKey method.
type IterateKeyer interface {
	// IterateKey calls the passed function with each value within the receiver.
	// The iteration is aborted if the function returns false.
	IterateKey(func(k Value) bool)
}

// IteratePairer wraps the IteratePair method.
type IteratePairer interface {
	// IteratePair calls the passed function with each key and value within the
	// receiver. The iteration is aborted if the function returns false.
	IteratePair(func(k, v Value) bool)
}

// Indexer wraps the Index method.
type Indexer interface {
	// Index retrieves one value from the receiver at the specified index.
	Index(idx Value) (Value, error)
}

// MustIndex indexes i with k and returns the value. If the operation
// resulted in an error, it panics. It is useful when the caller knows that the
// key must be present.
func MustIndex(i Indexer, k Value) Value {
	v, err := i.Index(k)
	if err != nil {
		panic(err)
	}
	return v
}

// Assocer wraps the Assoc method.
type Assocer interface {
	// Assoc returns a slightly modified version of the receiver with key k
	// associated with value v.
	Assoc(k, v Value) (Value, error)
}

// Dissocer is anything tha can return a slightly modified version of itself with
// the specified key removed, as a new value.
type Dissocer interface {
	// Dissoc returns a slightly modified version of the receiver with key k
	// dissociated with any value.
	Dissoc(k Value) Value
}
