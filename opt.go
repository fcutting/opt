// Package opt provides a generic Option type that holds a value of a provided
// type and a boolean flag indicating whether the value was provided.
package opt

import (
	"encoding/json"
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

// Exists reports whether the value was provided.
// If o is nil, Exists returns false.
func (o *Option[T]) Exists() (exists bool) {
	if o == nil {
		return false
	}

	return o.exists
}

// Get returns the value.
// if o is nil, Get always returns the zero value of the type.
// If the value is not provided, Get always returns the zero value of the type.
func (o *Option[T]) Get() (value T) {
	if o == nil || !o.exists {
		return value
	}

	return o.value
}

// GetWithDefault returns the value, or returns the defaultValue if the value is not provided.
// If o is nil, GetWithDefault always returns the defaultValue.
// If the value is not provided, GetWithDefault always returns the defaultValue.
func (o *Option[T]) GetWithDefault(defaultValue T) (value T) {
	if o == nil || !o.exists {
		return defaultValue
	}

	return o.value
}
