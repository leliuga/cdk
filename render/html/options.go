package html

import (
	"github.com/flosch/pongo2/v6"
)

// NewOptions creates a new options.
func NewOptions(options ...Option) *Options {
	opts := Options{}

	for _, option := range options {
		option(&opts)
	}

	return &opts
}

// WithDirectory sets the directory for the options.
func WithDirectory(value string) Option {
	return func(o *Options) {
		o.Directory = value
	}
}

// WithDebug sets the debug for the options.
func WithDebug(value bool) Option {
	return func(o *Options) {
		o.Debug = value
	}
}

// WithMinify sets the minify for the options.
func WithMinify(value bool) Option {
	return func(o *Options) {
		o.Minify = value
	}
}

// WithCache sets the cache for the options.
func WithCache(value bool) Option {
	return func(o *Options) {
		o.Cache = value
	}
}

// WithVariables sets the variables for the options.
func WithVariables(value pongo2.Context) Option {
	return func(o *Options) {
		o.Variables = value
	}
}

// WithFilters sets the filters for the options.
func WithFilters(value map[string]pongo2.FilterFunction) Option {
	return func(o *Options) {
		o.filters = value
	}
}
