// Package resource provides a resource definition.
package resource

import (
	"github.com/leliuga/cdk/types"
)

type (
	// Resource is a struct that represents a resource
	// for naming can be used semantic conventions (https://opentelemetry.io/docs/specs/semconv/)
	Resource struct {
		Name        string         `json:"name"`
		Description string         `json:"description"`
		Kind        Kind           `json:"kind"`
		Quantity    Quantity       `json:"quantity"`
		Attributes  types.Map[any] `json:"attributes"`
	}

	// Quantity represents a quantity
	Quantity struct {
		Unit Unit   `json:"unit"`
		Val  uint64 `json:"val"`
	}

	// Kind represents the kind of a resource
	Kind uint8

	// Unit represents the unit of a resource
	Unit uint8
)
