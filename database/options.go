package database

import (
	"time"

	"github.com/leliuga/cdk/types"
)

const (
	// DefaultMaxOpenConnections represents the default max open connections.
	DefaultMaxOpenConnections    = 5
	DefaultMaxIdleConnections    = 1
	DefaultMaxLifetimeConnection = 65 * time.Second
	DefaultMaxIdleTimeConnection = 15 * time.Second
)

// NewOptions returns new options.
func NewOptions(option ...Option) *Options {
	options := Options{
		Dsn:                   types.URI{},
		Options:               types.NewMap[string](),
		MaxOpenConnections:    DefaultMaxOpenConnections,
		MaxIdleConnections:    DefaultMaxIdleConnections,
		MaxLifetimeConnection: DefaultMaxLifetimeConnection,
		MaxIdleTimeConnection: DefaultMaxIdleTimeConnection,
	}

	for _, o := range option {
		o(&options)
	}

	return &options
}

// WithDSN sets the dsn for the storage.
func WithDSN(value string) Option {
	return func(o *Options) {
		o.Dsn = types.ParseURI(value)
	}
}

// WithOptions sets the options for the storage.
func WithOptions(value types.Map[string]) Option {
	return func(o *Options) {
		o.Options = value
	}
}

// WithMaxOpenConnections sets the max open connections for the storage.
func WithMaxOpenConnections(value int) Option {
	return func(o *Options) {
		o.MaxOpenConnections = value
	}
}

// WithMaxIdleConnections sets the max idle connections for the storage.
func WithMaxIdleConnections(value int) Option {
	return func(o *Options) {
		o.MaxIdleConnections = value
	}
}

// WithMaxLifetimeConnection sets the max lifetime connection for the storage.
func WithMaxLifetimeConnection(value time.Duration) Option {
	return func(o *Options) {
		o.MaxLifetimeConnection = value
	}
}

// WithMaxIdleTimeConnection sets the max idle time connection for the storage.
func WithMaxIdleTimeConnection(value time.Duration) Option {
	return func(o *Options) {
		o.MaxIdleTimeConnection = value
	}
}
