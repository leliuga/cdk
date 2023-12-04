package types

import (
	"bytes"
	"database/sql/driver"
	"fmt"
	"strings"
)

const (
	TypeInvalid Type = iota //
	TypeBoolean
	TypeDateTime
	TypeFloat
	TypeID
	TypeInteger
	TypeString
)

var (
	TypeNames = map[Type]string{
		TypeBoolean:  "Boolean",
		TypeDateTime: "DateTime",
		TypeFloat:    "Float",
		TypeID:       "ID",
		TypeInteger:  "Integer",
		TypeString:   "String",
	}
)

// String outputs the Type as a string.
func (t *Type) String() string {
	if !t.Validate() {
		return ""
	}

	return TypeNames[*t]
}

// Bytes returns the Type as a []byte.
func (t *Type) Bytes() []byte {
	return []byte(t.String())
}

// Value outputs the Type as a value.
func (t *Type) Value() (driver.Value, error) {
	return t.String(), nil
}

// MarshalJSON outputs the Type as a json.
func (t *Type) MarshalJSON() ([]byte, error) {
	return []byte(`"` + t.String() + `"`), nil
}

// UnmarshalJSON parses Type from json.
func (t *Type) UnmarshalJSON(data []byte) error {
	v, err := ParseType(string(bytes.Trim(data, `"`)))
	if err != nil {
		return err
	}

	*t = *v

	return nil
}

// Validate returns true if the Type is valid.
func (t *Type) Validate() bool {
	return *t != TypeInvalid
}

// ParseType parses Type from string.
func ParseType(value string) (*Type, error) {
	value = strings.ToLower(value)
	for k, v := range TypeNames {
		if strings.ToLower(v) == value {
			return &k, nil
		}
	}

	return nil, fmt.Errorf("unsupported type: %s", value)
}
