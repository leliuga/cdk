package types

import (
	"database/sql/driver"
	"fmt"
	"reflect"
	"sort"
	"strings"
)

func NewMap[T any]() Map[T] {
	return make(Map[T])
}

// Set sets the value for the provided key.
func (m Map[T]) Set(key string, value T) {
	m[key] = value
}

// Get returns the value for the provided key.
func (m Map[T]) Get(key string) T {
	return m[key]
}

// Delete deletes the value for the provided key.
func (m Map[T]) Delete(key string) {
	delete(m, key)
}

// Clear clears the map.
func (m Map[T]) Clear() {
	for key := range m {
		delete(m, key)
	}
}

// Has returns whether the provided key exists in the map.
func (m Map[T]) Has(key string) bool {
	_, exists := m[key]
	return exists
}

// Keys returns a slice of keys in the map.
func (m Map[T]) Keys() (keys []string) {
	for key := range m {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	return keys
}

// Values returns a slice of values in the map.
func (m Map[T]) Values() (values []T) {
	for _, value := range m {
		values = append(values, value)
	}

	sort.Slice(values, func(i, j int) bool {
		return Less(values[i], values[j])
	})

	return values
}

// Len returns the length of the map.
func (m Map[T]) Len() int {
	return len(m)
}

// IsEmpty returns whether the map is empty.
func (m Map[T]) IsEmpty() bool {
	return m.Len() == 0
}

// Clone returns a clone of the map.
func (m Map[T]) Clone() Map[T] {
	clone := NewMap[T]()

	for key, value := range m {
		clone[key] = value
	}

	return clone
}

// Merge merges the provided map into the current map.
func (m Map[T]) Merge(maps ...Map[T]) Map[T] {
	for _, mm := range maps {
		for key, value := range mm {
			m[key] = value
		}
	}

	return m
}

// Range iterates over elements in the map.
func (m Map[T]) Range(fn func(key string, value T) bool) {
	for key, value := range m {
		if !fn(key, value) {
			break
		}
	}
}

// String returns a string representation of the map.
func (m Map[T]) String(sep, join string) string {
	parts := make([]string, 0, m.Len())
	for key, value := range m {
		parts = append(parts, fmt.Sprintf("%s%s%s", key, sep, value))
	}

	sort.Strings(parts)

	return strings.Join(parts, join)
}

// Value returns a value representation of the map.
func (m Map[T]) Value() (driver.Value, error) {
	return m.String("=", ";"), nil
}

// AsString returns the value for the provided key as a string.
func (m Map[T]) AsString(key string) string {
	switch reflect.ValueOf(m[key]).Kind() {
	case reflect.Bool:
		return fmt.Sprintf("%t", m[key])
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return fmt.Sprintf("%d", m[key])
	case reflect.Float32, reflect.Float64:
		return fmt.Sprintf("%f", m[key])
	case reflect.Invalid:
		return ""
	}

	return fmt.Sprintf("%s", m[key])
}

// Merge merges the provided maps into a new map.
func Merge[T any](maps ...Map[T]) Map[T] {
	merged := NewMap[T]()
	merged.Merge(maps...)

	return merged
}

// ConflictKeys returns a map of conflicting keys between the two maps.
func ConflictKeys[T any](a, b Map[T]) (conflicts []string) {
	for key := range a {
		if b.Has(key) {
			conflicts = append(conflicts, key)
		}
	}

	return conflicts
}

// DiffKeys returns a map of added and removed keys between the two maps.
func DiffKeys[T any](a, b Map[T]) (added, removed []string) {
	for key := range a {
		if !b.Has(key) {
			added = append(added, key)
		}
	}

	for key := range b {
		if !a.Has(key) {
			removed = append(removed, key)
		}
	}

	return added, removed
}

// ToMap converts a map to a Map.
func ToMap[M ~map[K]T, K comparable, T any](m M) Map[T] {
	newMap := NewMap[T]()
	format := "%s"

	var k K
	switch reflect.ValueOf(k).Kind() {
	case reflect.Bool:
		format = "%t"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		format = "%d"
	case reflect.Float32, reflect.Float64:
		format = "%f"
	}

	for key, value := range m {
		newMap[fmt.Sprintf(format, key)] = value
	}

	return newMap
}
