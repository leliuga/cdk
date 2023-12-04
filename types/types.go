// Package types provides common types.
package types

import (
	"database/sql/driver"
	"net/url"
	"time"
)

type (
	// DateTime represents the date and time using RFC3339 format.
	DateTime struct {
		time.Time
	}

	// Map defines a map of key:value. It implements Map.
	Map[T any] map[String]T

	// Slice defines a slice of T. It implements Slice.
	Slice[T any] []T

	// Option represents the ui option.
	Option struct {
		Name        String   `json:"name"`
		Description String   `json:"description"`
		Required    bool     `json:"required"`
		Type        *Type    `json:"type"`
		Default     any      `json:"default"`
		Min         any      `json:"min,omitempty"`
		Max         any      `json:"max,omitempty"`
		Choices     []String `json:"choices,omitempty"`
	}

	// URI represents the URI.
	URI struct {
		*url.URL
		hash uint32
	}

	// Path represents the Path.
	Path struct {
		p    String
		hash uint32
	}

	// Options represents the options.
	Options []*Option

	// ContentType represents the content type.
	ContentType uint8

	// Type represents the data type.
	Type uint8

	// String represents the string.
	String string

	// IType represents the stringer interface.
	IType interface {
		String() string
		Bytes() []byte
		Value() (driver.Value, error)
		Validate() bool
	}
)
