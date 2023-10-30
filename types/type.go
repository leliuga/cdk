package types

import (
	"bytes"
	"strconv"
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
func (t Type) String() string {
	return TypeNames[t]
}

// Bytes returns the Type as a []byte.
func (t Type) Bytes() []byte {
	return []byte(strconv.Itoa(int(t)))
}

// MarshalJSON outputs the Type as a json.
func (t Type) MarshalJSON() ([]byte, error) {
	if !t.Validate() {
		return []byte(`""`), nil
	}

	return []byte(`"` + t.String() + `"`), nil
}

// UnmarshalJSON parses Type from json.
func (t *Type) UnmarshalJSON(data []byte) error {
	s := string(bytes.Trim(data, `"`))
	if r := ParseType(s); r.Validate() {
		*t = r
	}

	return nil
}

// Validate returns true if the Type is valid.
func (t Type) Validate() bool {
	return t != TypeInvalid
}

// ParseType parses Type from string.
func ParseType(value string) Type {
	value = strings.ToLower(value)
	for k, v := range TypeNames {
		if strings.ToLower(v) == value {
			return k
		}
	}

	return TypeInvalid
}
