// Package opt provides a generic Option type that holds a value of a provided
// type and a boolean flag indicating whether the value was provided.
package opt

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// Represents the JSON null string in bytes.
// This isn't a constant because Go doesn't allow for converted types to be
// constants.
var nullBytes = []byte("null")

// Option represents a generic type that holds a value of any type T and a
// boolean flag indication whether the value was provided.
type Option[T any] struct {
	// value holds the value of type T.
	value T

	// exists indicates whether the value was provided.
	exists bool
}

// MarshalJSON marshals the Option to JSON.
// If the value is provided, MarshalJSON marshals the value.
// If the value is not provided, MarshalJSON returns "null".
func (o Option[T]) MarshalJSON() (data []byte, err error) {
	if o.exists {
		return json.Marshal(o.value)
	}

	return nullBytes, nil
}

// UnmarshalJSON unmarshals the Option from JSON.
// If the data is not "null", UnmarshalJSON unmarshals the value and sets
// exists to true.
// If the data is "null", the value is not set and UnmarshalJSON returns nil.
func (o *Option[T]) UnmarshalJSON(data []byte) (err error) {
	if reflect.DeepEqual(data, nullBytes) {
		tZeroValue := *(new(T))
		tKind := reflect.ValueOf(tZeroValue).Kind()

		switch tKind {
		case reflect.Ptr, reflect.Map, reflect.Slice:
			o.exists = true
		}
		return nil
	}

	// I check if the Unmarshal works first before setting exists to true because
	// if the Unmarshal fails and the caller continues despite the error then
	// exists being true is incorrect
	if err = json.Unmarshal(data, &o.value); err != nil {
		return
	}

	o.exists = true
	return nil
}

// String returns a string representation of the value.
// If the value is not provided, String returns "<empty>".
func (o Option[T]) String() (str string) {
	if !o.exists {
		return "<empty>"
	}

	return fmt.Sprint(o.value)
}

// Exists reports whether the value was provided.
func (o Option[T]) Exists() (exists bool) {
	return o.exists
}

// Unwrap returns the value.
// If the value is not provided, Unwrap returns the zero value of the type.
func (o Option[T]) Unwrap() (value T) {
	if !o.exists {
		return value
	}

	return o.value
}

// MustUnwrap returns the value.
// If the value is not provided, MustUnwrap may panic.
func (o Option[T]) MustUnwrap() (value T) {
	return o.value
}

// UnwrapDefault returns the value, or returns the defaultValue if the value
// is not provided.
// If the value is not provided, UnwrapDefault returns the defaultValue.
func (o Option[T]) UnwrapDefault(defaultValue T) (value T) {
	if !o.exists {
		return defaultValue
	}

	return o.value
}
