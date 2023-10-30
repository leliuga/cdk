package types

import (
	"bytes"
	"database/sql/driver"
	"net/url"
)

// String outputs the URI as a string.
func (u URI) String() string {
	return u.URL.String()
}

// Value outputs the URI as a value.
func (U URI) Value() (driver.Value, error) {
	return U.String(), nil
}

// MarshalJSON outputs the URI as a json.
func (u URI) MarshalJSON() ([]byte, error) {
	if !u.Validate() {
		return []byte(`""`), nil
	}

	return []byte(`"` + u.String() + `"`), nil
}

// UnmarshalJSON parses URI from json.
func (u *URI) UnmarshalJSON(data []byte) error {
	s := string(bytes.Trim(data, `"`))
	if v := ParseURI(s); v.Validate() {
		*u = v
	}

	return nil
}

// Validate returns true if the URI-Reference is valid.
func (u URI) Validate() bool {
	return u.IsAbs()
}

// ParseURI parses URI from string.
func ParseURI(value string) URI {
	v, err := url.Parse(value)
	if err != nil {
		return URI{}
	}

	return URI{URL: *v}
}
