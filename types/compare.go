package types

import (
	"reflect"
)

// Equal returns whether the provided values are equal.
func Equal(a, b any) bool {
	valueA := reflect.ValueOf(a)
	valueB := reflect.ValueOf(b)
	kindA := valueA.Kind()
	kindB := valueB.Kind()

	if kindA != kindB {
		return false
	}

	switch kindA {
	case reflect.Bool:
		return valueA.Bool() == valueB.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return valueA.Int() == valueB.Int()
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return valueA.Uint() == valueB.Uint()
	case reflect.Float32, reflect.Float64:
		return valueA.Float() == valueB.Float()
	case reflect.Complex64, reflect.Complex128:
		return valueA.Complex() == valueB.Complex()
	case reflect.String:
		return valueA.String() == valueB.String()
	}

	return false
}

// NotEqual returns whether the provided values are not equal.
func NotEqual(a, b any) bool {
	return !Equal(a, b)
}

// Less returns whether the first value is less than the second.
func Less(a, b any) bool {
	valueA := reflect.ValueOf(a)
	valueB := reflect.ValueOf(b)
	kindA := valueA.Kind()
	kindB := valueB.Kind()

	if kindA != kindB {
		return false
	}

	switch kindA {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return valueA.Int() < valueB.Int()
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return valueA.Uint() < valueB.Uint()
	case reflect.Float32, reflect.Float64:
		return valueA.Float() < valueB.Float()
	case reflect.String:
		return valueA.String() < valueB.String()
	}

	return false
}

// More returns whether the first value is more than the second.
func More(a, b any) bool {
	return !Less(b, a)
}
