package service

import (
	"bytes"
	"database/sql/driver"
	"strings"

	"github.com/pkg/errors"
)

const (
	EnvironmentInvalid Environment = iota //
	EnvironmentDevelopment
	EnvironmentStaging
	EnvironmentProduction
)

var (
	EnvironmentNames = map[Environment]string{
		EnvironmentDevelopment: "development",
		EnvironmentStaging:     "staging",
		EnvironmentProduction:  "production",
	}
)

// String outputs the Environment as a string.
func (e *Environment) String() string {
	if !e.Validate() {
		return ""
	}

	return EnvironmentNames[*e]
}

// Bytes returns the Environment as a []byte.
func (e *Environment) Bytes() []byte {
	return []byte(e.String())
}

// Value outputs the Environment as a value.
func (e *Environment) Value() (driver.Value, error) {
	return e.String(), nil
}

// MarshalJSON outputs the Environment as a json.
func (e *Environment) MarshalJSON() ([]byte, error) {
	return []byte(`"` + e.String() + `"`), nil
}

// UnmarshalJSON parses the Environment from json.
func (e *Environment) UnmarshalJSON(data []byte) error {
	v, err := ParseEnvironment(string(bytes.Trim(data, `"`)))
	if err != nil {
		return err
	}

	*e = *v

	return nil
}

// Validate returns true if the Environment is valid.
func (e *Environment) Validate() bool {
	return *e != EnvironmentInvalid
}

// ParseEnvironment parses the Environment from string.
func ParseEnvironment(value string) (*Environment, error) {
	value = strings.ToLower(value)
	for k, v := range EnvironmentNames {
		if v == value {
			return &k, nil
		}
	}

	return nil, errors.Errorf("unsupported environment: %s", value)
}
