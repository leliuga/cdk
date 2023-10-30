package service

import (
	"bytes"
	"strings"
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
func (e Engine) String() string {
	return EngineNames[e]
}

// MarshalJSON outputs the Engine as a json.
func (e Engine) MarshalJSON() ([]byte, error) {
	if !e.Validate() {
		return []byte(`""`), nil
	}

	return []byte(`"` + e.String() + `"`), nil
}

// UnmarshalJSON parses the Engine from json.
func (e *Engine) UnmarshalJSON(data []byte) error {
	str := string(bytes.Trim(data, `"`))
	if engine := ParseEngine(str); engine.Validate() {
		*e = engine
	}

	return nil
}

// Validate returns true if the Engine is valid.
func (e Engine) Validate() bool {
	return e != EngineInvalid
}

// ParseEngine parses the Engine string.
func ParseEngine(value string) Engine {
	value = strings.ToLower(value)
	for k, v := range EngineNames {
		if strings.ToLower(v) == value {
			return k
		}
	}

	return EngineInvalid
}
