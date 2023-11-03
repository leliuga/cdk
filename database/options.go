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
func NewOptions(options ...Option) *Options {
	opts := Options{
		SourcesDsn:            []types.URI{},
		ReplicasDsn:           []types.URI{},
		Options:               types.NewMap[string](),
		MaxOpenConnections:    DefaultMaxOpenConnections,
		MaxIdleConnections:    DefaultMaxIdleConnections,
		MaxLifetimeConnection: DefaultMaxLifetimeConnection,
		MaxIdleTimeConnection: DefaultMaxIdleTimeConnection,
	}

	for _, option := range options {
		option(&opts)
	}

	return &opts
}

// WithSourcesDsn sets the sources dsn for the storage.
func WithSourcesDsn(value []types.URI) Option {
	return func(o *Options) {
		o.SourcesDsn = value
	}
}

// WithReplicasDsn sets the replicas dsn for the storage.
func WithReplicasDsn(value []types.URI) Option {
	return func(o *Options) {
		o.ReplicasDsn = value
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
