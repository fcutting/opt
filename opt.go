// Package opt provides a generic Option type that holds a value of a provided
// type and a boolean flag indicating whether the value was provided.
package opt

// Option represents a generic type that holds a value of any type T and a
// boolean flag indication whether the value was provided.
type Option[T any] struct {
	// value holds the value of type T.
	value T

	// exists indicates whether the value was provided.
	exists bool
}
