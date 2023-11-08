// Package types provides common types.
package types

import (
	"net/url"
	"time"
)

type (
	// DateTime represents the date and time using RFC3339 format.
	DateTime struct {
		time.Time
	}

	// Map defines a map of key:value. It implements Map.
	Map[T any] map[string]T

	// Option represents the ui option.
	Option struct {
		Name        string   `json:"name"`
		Description string   `json:"description"`
		Required    bool     `json:"required"`
		Type        Type     `json:"type"`
		Default     any      `json:"default"`
		Min         any      `json:"min,omitempty"`
		Max         any      `json:"max,omitempty"`
		Choices     []string `json:"choices,omitempty"`
	}

	// URI represents the URI.
	URI struct {
		url.URL
	}

	// Options represents the options.
	Options []*Option

	// ContentType represents the content type.
	ContentType uint8

	// Type represents the data type.
	Type uint8
)
