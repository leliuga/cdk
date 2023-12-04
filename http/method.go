package http

import (
	"bytes"
	"database/sql/driver"
	"strings"

	"github.com/pkg/errors"
)

// Common HTTP methods, these are defined in RFC 7231 section 4.3.
const (
	MethodInvalid Method = iota //
	MethodGet                   // RFC 7231, 4.3.1
	MethodHead                  // RFC 7231, 4.3.2
	MethodPost                  // RFC 7231, 4.3.3
	MethodPut                   // RFC 7231, 4.3.4
	MethodDelete                // RFC 7231, 4.3.5
	MethodConnect               // RFC 7231, 4.3.6
	MethodOptions               // RFC 7231, 4.3.7
	MethodTrace                 // RFC 7231, 4.3.8
	MethodPatch                 // RFC 5789
)

var (
	// MethodNames is a map of Method to string.
	MethodNames = map[Method]string{
		MethodGet:     "GET",
		MethodHead:    "HEAD",
		MethodPost:    "POST",
		MethodPut:     "PUT",
		MethodDelete:  "DELETE",
		MethodConnect: "CONNECT",
		MethodOptions: "OPTIONS",
		MethodTrace:   "TRACE",
		MethodPatch:   "PATCH",
	}
)

// String outputs the Method as a string.
func (m *Method) String() string {
	if !m.Validate() {
		return ""
	}

	return MethodNames[*m]
}

// Bytes returns the Method as a []byte.
func (m *Method) Bytes() []byte {
	return []byte(m.String())
}

// Value outputs the Method as a value.
func (m *Method) Value() (driver.Value, error) {
	return m.String(), nil
}

// MarshalJSON outputs the Method as a json.
func (m *Method) MarshalJSON() ([]byte, error) {
	return []byte(`"` + m.String() + `"`), nil
}

// UnmarshalJSON parses the Method from json.
func (m *Method) UnmarshalJSON(data []byte) error {
	v, err := ParseMethod(string(bytes.Trim(data, `"`)))
	if err != nil {
		return err
	}

	*m = *v

	return nil
}

// Validate returns true if the Method is valid.
func (m *Method) Validate() bool {
	return *m != MethodInvalid
}

// ParseMethod parses the Method from string.
func ParseMethod(value string) (*Method, error) {
	value = strings.ToUpper(value)
	for k, v := range MethodNames {
		if v == value {
			return &k, nil
		}
	}

	return nil, errors.Errorf("unsupported method: %s", value)
}
