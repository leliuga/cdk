// Package event provides the event definition.
package event

import (
	"github.com/leliuga/cdk/types"
)

type (
	// Event represents the event.
	Event struct {
		*Options `json:",inline"`
	}

	// Options represents the event options.
	Options struct {
		ID         string            `json:"id"`
		Version    string            `json:"version"`
		Schema     types.URI         `json:"schema"`
		Source     types.URI         `json:"source"`
		Kind       Kind              `json:"kind"`
		Action     Action            `json:"action"`
		Attributes types.Map[string] `json:"attributes"`
		Data       []byte            `json:"data"`
		Happen     types.DateTime    `json:"happen"`
	}

	// Action represents the event action.
	Action uint8

	// Kind represents the event topic.
	Kind uint16

	// Option represents the event option.
	Option func(o *Options)
)
