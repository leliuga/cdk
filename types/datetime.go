package types

import (
	"bytes"
	"database/sql/driver"
	"fmt"
	"time"
)

// DateTimeNow returns the current date and time.
func DateTimeNow() *DateTime {
	return &DateTime{Time: time.Now().UTC()}
}

// String outputs the DateTime as a string using RFC3339 format.
func (dt *DateTime) String() string {
	return dt.Time.UTC().Format(time.RFC3339Nano)
}

// Bytes returns the DateTime as bytes.
func (dt *DateTime) Bytes() []byte {
	return []byte(dt.String())
}

// Value outputs the DateTime as a value.
func (dt *DateTime) Value() (driver.Value, error) {
	return dt.String(), nil
}

// MarshalJSON outputs the DateTime as a json.
func (dt *DateTime) MarshalJSON() ([]byte, error) {
	if !dt.Validate() {
		return []byte(`""`), nil
	}

	return []byte(`"` + dt.String() + `"`), nil
}

// UnmarshalJSON parses DateTime from json.
func (dt *DateTime) UnmarshalJSON(data []byte) error {
	v, err := ParseDateTime(string(bytes.Trim(data, `"`)))
	if err != nil {
		return err
	}

	*dt = *v

	return nil
}

// Validate returns true if the DateTime is valid.
func (dt *DateTime) Validate() bool {
	return !dt.IsZero()
}

// ParseDateTime parses DateTime from string.
func ParseDateTime(value string) (*DateTime, error) {
	if v, err := time.Parse(time.RFC3339, value); err == nil {
		return &DateTime{Time: v}, nil
	}

	return nil, fmt.Errorf("unsupported datetime: %s", value)
}
