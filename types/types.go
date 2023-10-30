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

	// URI represents the URI.
	URI struct {
		url.URL
	}

	// ContentType represents the event data mime type.
	ContentType uint8

	// Type represents the event type.
	Type uint8
)
