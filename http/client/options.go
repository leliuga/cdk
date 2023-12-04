package client

import (
	"time"

	"github.com/leliuga/cdk/http"
	"github.com/leliuga/cdk/service"
	"github.com/leliuga/cdk/types"
)

const (
	DefaultUserAgent             = service.DefaultApplicationName
	DefaultTimeout               = 20 * time.Second
	DefaultKeepAlive             = 30 * time.Second
	DefaultTLSHandshakeTimeout   = 5 * time.Second
	DefaultExpectContinueTimeout = 1 * time.Second
	DefaultIdleConnectionTimeout = 60 * time.Second
	DefaultResponseHeaderTimeout = 5 * time.Second
	DefaultMaxIdleConnections    = 10
	DefaultMaxConnectionsPerHost = 10
	DefaultWriteBufferSize       = 4 * 1024
	DefaultReadBufferSize        = 4 * 1024
	DefaultQPS                   = 10
	DefaultBurst                 = 100
)

// NewOptions creates a new options.
func NewOptions(options ...Option) *Options {
	opts := Options{
		BaseUri:               &types.URI{},
		Headers:               http.Headers{},
		MaxIdleConnections:    DefaultMaxIdleConnections,
		MaxConnectionsPerHost: DefaultMaxConnectionsPerHost,
		WriteBufferSize:       DefaultWriteBufferSize,
		ReadBufferSize:        DefaultReadBufferSize,
		Timeout:               DefaultTimeout,
		KeepAlive:             DefaultKeepAlive,
		TLSHandshake:          DefaultTLSHandshakeTimeout,
		ExpectContinue:        DefaultExpectContinueTimeout,
		IdleConnection:        DefaultIdleConnectionTimeout,
		ResponseHeader:        DefaultResponseHeaderTimeout,
		QPS:                   DefaultQPS,
		Burst:                 DefaultBurst,
	}

	for _, option := range options {
		option(&opts)
	}

	return &opts
}

// WithBaseUri sets the dsn.
func WithBaseUri(value string) Option {
	return func(o *Options) {
		o.BaseUri, _ = types.ParseURI(value)
	}
}

// WithHeaders sets the headers.
func WithHeaders(value http.Headers) Option {
	return func(o *Options) {
		o.Headers = value
	}
}

// WithProxyURL sets the proxy url.
func WithProxyURL(value string) Option {
	return func(o *Options) {
		o.ProxyURL, _ = types.ParseURI(value)
	}
}

// WithMaxIdleConnections sets the max idle connections.
func WithMaxIdleConnections(value int) Option {
	return func(o *Options) {
		o.MaxIdleConnections = value
	}
}

// WithMaxConnectionsPerHost sets the max connections per host.
func WithMaxConnectionsPerHost(value int) Option {
	return func(o *Options) {
		o.MaxConnectionsPerHost = value
	}
}

// WithWriteBufferSize sets the write buffer size.
func WithWriteBufferSize(value int) Option {
	return func(o *Options) {
		o.WriteBufferSize = value
	}
}

// WithReadBufferSize sets the read buffer size.
func WithReadBufferSize(value int) Option {
	return func(o *Options) {
		o.ReadBufferSize = value
	}
}

// WithTimeout sets the timeout.
func WithTimeout(value time.Duration) Option {
	return func(o *Options) {
		o.Timeout = value
	}
}

// WithKeepAlive sets the keep alive.
func WithKeepAlive(value time.Duration) Option {
	return func(o *Options) {
		o.KeepAlive = value
	}
}

// WithTLSHandshake sets the tls handshake.
func WithTLSHandshake(value time.Duration) Option {
	return func(o *Options) {
		o.TLSHandshake = value
	}
}

// WithExpectContinue sets the expect continue.
func WithExpectContinue(value time.Duration) Option {
	return func(o *Options) {
		o.ExpectContinue = value
	}
}

// WithIdleConnection sets the idle connection.
func WithIdleConnection(value time.Duration) Option {
	return func(o *Options) {
		o.IdleConnection = value
	}
}

// WithResponseHeader sets the response header.
func WithResponseHeader(value time.Duration) Option {
	return func(o *Options) {
		o.ResponseHeader = value
	}
}

// WithQPS sets the qps.
func WithQPS(value float32) Option {
	return func(o *Options) {
		o.QPS = value
	}
}

// WithBurst sets the burst.
func WithBurst(value int) Option {
	return func(o *Options) {
		o.Burst = value
	}
}
