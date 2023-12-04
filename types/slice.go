package types

import (
	"reflect"
)

// NewSlice creates a new Slice instance.
func NewSlice[T any]() Slice[T] {
	return make(Slice[T], 0)
}

// Append appends the given value to the slice.
func (s Slice[T]) Append(values ...T) Slice[T] {
	return append(s, values...)
}

// Delete deletes the value at the given index.
func (s Slice[T]) Delete(index int) Slice[T] {
	return append(s[:index], s[index+1:]...)
}

// Clear clears the slice.
func (s Slice[T]) Clear() Slice[T] {
	return s[:0]
}

// Index returns the index of the given value.
func (s Slice[T]) Index(value T) int {
	for k, v := range s {
		if reflect.DeepEqual(v, value) {
			return k
		}
	}

	return -1
}

// Has returns whether the given value exists in the slice.
func (s Slice[T]) Has(value T) bool {
	return s.Index(value) != -1
}

// Len returns the length of the slice.
func (s Slice[T]) Len() int {
	return len(s)
}

// IsEmpty returns whether the slice is empty.
func (s Slice[T]) IsEmpty() bool {
	return len(s) == 0
}

// Equal returns whether the slice is equal to the given slice.
func (s Slice[T]) Equal(other Slice[T]) bool {
	if len(s) != len(other) {
		return false
	}

	for k, v := range s {
		if !reflect.DeepEqual(v, other[k]) {
			return false
		}
	}

	return true
}
