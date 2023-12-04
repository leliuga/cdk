package service

import (
	"bytes"
	"database/sql/driver"
	"strings"

	"github.com/pkg/errors"
)

const (
	EngineInvalid Engine = iota //
	EngineKubernetes
	EngineDockerSwarm
)

var (
	EngineNames = map[Engine]string{
		EngineKubernetes:  "Kubernetes",
		EngineDockerSwarm: "Docker Swarm",
	}
)

// String outputs the Engine as a string.
func (e *Engine) String() string {
	if !e.Validate() {
		return ""
	}

	return EngineNames[*e]
}

// Bytes returns the Engine as a []byte.
func (e *Engine) Bytes() []byte {
	return []byte(e.String())
}

// Value outputs the Engine as a value.
func (e *Engine) Value() (driver.Value, error) {
	return e.String(), nil
}

// MarshalJSON outputs the Engine as a json.
func (e *Engine) MarshalJSON() ([]byte, error) {
	return []byte(`"` + e.String() + `"`), nil
}

// UnmarshalJSON parses the Engine from json.
func (e *Engine) UnmarshalJSON(data []byte) error {
	v, err := ParseEngine(string(bytes.Trim(data, `"`)))
	if err != nil {
		return err
	}

	*e = *v

	return nil
}

// Validate returns true if the Engine is valid.
func (e *Engine) Validate() bool {
	return *e != EngineInvalid
}

// ParseEngine parses the Engine string.
func ParseEngine(value string) (*Engine, error) {
	value = strings.ToLower(value)
	for k, v := range EngineNames {
		if strings.ToLower(v) == value {
			return &k, nil
		}
	}

	return nil, errors.Errorf("unsupported engine: %s", value)
}
