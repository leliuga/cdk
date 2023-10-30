package event

import (
	"bytes"
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
func (a Action) String() string {
	return ActionNames[a]
}

// MarshalJSON outputs the Action as a json.
func (a Action) MarshalJSON() ([]byte, error) {
	if !a.Validate() {
		return []byte(`""`), nil
	}

	return []byte(`"` + a.String() + `"`), nil
}

// UnmarshalJSON parses the Action from json.
func (a *Action) UnmarshalJSON(data []byte) error {
	str := string(bytes.Trim(data, `"`))
	if action := ParseAction(str); action.Validate() {
		*a = action
	}

	return nil
}

// Validate returns true if the Action is valid.
func (a Action) Validate() bool {
	return a != ActionInvalid
}

// ParseAction parses the Action string.
func ParseAction(value string) Action {
	value = strings.ToLower(value)
	for k, v := range ActionNames {
		if strings.ToLower(v) == value {
			return k
		}
	}

	return ActionInvalid
}
