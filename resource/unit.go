package resource

import (
	"bytes"
	"strconv"
	"strings"
)

const (
	UnitInvalid Unit = iota
	UnitBit
	UnitByte
	UnitHertz
)

var (
	UnitNames = map[Unit]string{
		UnitBit:   "bit",
		UnitByte:  "byte",
		UnitHertz: "hertz",
	}
)

// String outputs the Unit as a string.
func (u Unit) String() string {
	return UnitNames[u]
}

// Bytes returns the Unit as a []byte.
func (u Unit) Bytes() []byte {
	return []byte(strconv.Itoa(int(u)))
}

// MarshalJSON outputs the Unit as a json.
func (u Unit) MarshalJSON() ([]byte, error) {
	if !u.Validate() {
		return []byte(`""`), nil
	}

	return []byte(`"` + u.String() + `"`), nil
}

// UnmarshalJSON parses Unit from json.
func (u *Unit) UnmarshalJSON(data []byte) error {
	s := string(bytes.Trim(data, `"`))
	if r := ParseUnit(s); r.Validate() {
		*u = r
	}

	return nil
}

// Validate returns true if the Unit is valid.
func (u Unit) Validate() bool {
	return u != UnitInvalid
}

// ParseUnit parses Unit from string.
func ParseUnit(value string) Unit {
	value = strings.ToLower(value)
	for k, v := range UnitNames {
		if v == value {
			return k
		}
	}

	return UnitInvalid
}
