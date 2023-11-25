package types

import (
	"bytes"
	"database/sql/driver"
	"time"
)

// String outputs the DateTime as a string using RFC3339 format.
func (dt DateTime) String() string {
	return dt.Time.UTC().Format(time.RFC3339Nano)
}

// Value outputs the DateTime as a value.
func (dt DateTime) Value() (driver.Value, error) {
	return dt.String(), nil
}

// MarshalJSON outputs the DateTime as a json.
func (dt DateTime) MarshalJSON() ([]byte, error) {
	if !dt.Validate() {
		return []byte(`""`), nil
	}

	return []byte(`"` + dt.String() + `"`), nil
}

// UnmarshalJSON parses DateTime from json.
func (dt *DateTime) UnmarshalJSON(data []byte) error {
	s := string(bytes.Trim(data, `"`))
	if v := ParseDateTime(s); v.Validate() {
		*dt = v
	}

	return nil
}

// Validate returns true if the DateTime is valid.
func (dt DateTime) Validate() bool {
	return !dt.IsZero()
}

// DateTimeNow returns the current date and time.
func DateTimeNow() DateTime {
	return DateTime{Time: time.Now().UTC()}
}

// ParseDateTime parses DateTime from string.
func ParseDateTime(value string) DateTime {
	v, err := time.Parse(time.RFC3339, value)
	if err != nil {
		return DateTime{}
	}

	return DateTime{Time: v}
}
