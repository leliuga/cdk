package service

import (
	"bytes"
	"strings"
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
func (e Environment) String() string {
	return EnvironmentNames[e]
}

// MarshalJSON outputs the Environment as a json.
func (e Environment) MarshalJSON() ([]byte, error) {
	if !e.Validate() {
		return []byte(`""`), nil
	}

	return []byte(`"` + e.String() + `"`), nil
}

// UnmarshalJSON parses the Environment from json.
func (e *Environment) UnmarshalJSON(data []byte) error {
	str := string(bytes.Trim(data, `"`))
	if environment := ParseEnvironment(str); environment.Validate() {
		*e = environment
	}

	return nil
}

// Validate returns true if the Environment is valid.
func (e Environment) Validate() bool {
	return e != EnvironmentInvalid
}

// ParseEnvironment parses the Environment from string.
func ParseEnvironment(value string) Environment {
	value = strings.ToLower(value)
	for k, v := range EnvironmentNames {
		if v == value {
			return k
		}
	}

	return EnvironmentInvalid
}
