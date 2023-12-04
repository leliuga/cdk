package event

import (
	"bytes"
	"database/sql/driver"
	"fmt"
	"strings"
)

const (
	ActionInvalid Action = iota //
	ActionCreate
	ActionRead
	ActionUpdate
	ActionDelete
	ActionError
)

var (
	ActionNames = map[Action]string{
		ActionCreate: "Create",
		ActionRead:   "Read",
		ActionUpdate: "Update",
		ActionDelete: "Delete",
		ActionError:  "Error",
	}
)

// String outputs the Action as a string.
func (a *Action) String() string {
	if !a.Validate() {
		return ""
	}

	return ActionNames[*a]
}

// Bytes outputs the Action as bytes.
func (a *Action) Bytes() []byte {
	return []byte(a.String())
}

// Value outputs the Action as a value.
func (a *Action) Value() (driver.Value, error) {
	return a.String(), nil
}

// MarshalJSON outputs the Action as a json.
func (a *Action) MarshalJSON() ([]byte, error) {
	return []byte(`"` + a.String() + `"`), nil
}

// UnmarshalJSON parses the Action from json.
func (a *Action) UnmarshalJSON(data []byte) error {
	v, err := ParseAction(string(bytes.Trim(data, `"`)))
	if err != nil {
		return err
	}

	*a = *v

	return nil
}

// Validate returns true if the Action is valid.
func (a *Action) Validate() bool {
	return *a != ActionInvalid
}

// ParseAction parses the Action string.
func ParseAction(value string) (*Action, error) {
	value = strings.ToLower(value)
	for k, v := range ActionNames {
		if strings.ToLower(v) == value {
			return &k, nil
		}
	}

	return nil, fmt.Errorf("unsupported action: %s", value)
}
