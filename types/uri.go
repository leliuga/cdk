package types

import (
	"bytes"
	"database/sql/driver"
	"fmt"
	"net/url"
)

// String outputs the URI as a string.
func (u *URI) String() string {
	return u.URL.String()
}

// Bytes returns the URI as bytes.
func (u *URI) Bytes() []byte {
	return []byte(u.String())
}

// Value outputs the URI as a value.
func (u *URI) Value() (driver.Value, error) {
	return u.String(), nil
}

// MarshalJSON outputs the URI as a json.
func (u *URI) MarshalJSON() ([]byte, error) {
	if !u.Validate() {
		return []byte(`""`), nil
	}

	return []byte(`"` + u.String() + `"`), nil
}

// UnmarshalJSON parses URI from json.
func (u *URI) UnmarshalJSON(data []byte) error {
	v, err := ParseURI(string(bytes.Trim(data, `"`)))
	if err != nil {
		return err
	}

	*u = *v

	return nil
}

// Validate returns true if the URI-Reference is valid.
func (u *URI) Validate() bool {
	return u.IsAbs()
}

// Hash returns the URI as a hash.
func (u *URI) Hash() uint32 {
	return u.hash
}

// BaseUri returns the URI as a base uri.
func (u *URI) BaseUri() String {
	return String(u.URL.Scheme + "://" + u.URL.Host)
}

// Equal checks if two paths are equal.
func (u *URI) Equal(u1 *URI) bool {
	return u.hash == u1.hash
}

// ParseURI parses URI from string.
func ParseURI(value string) (*URI, error) {
	v, err := url.Parse(value)
	if err != nil {
		return nil, fmt.Errorf("failed to parse uri %s: %v", value, err)
	}

	return &URI{
		URL:  v,
		hash: String(value).Hash(),
	}, nil
}
