package server

import (
	"time"

	"github.com/leliuga/cdk/render"
)

// Default values for the HTTP server
const (
	DefaultName              = "le:cdk"
	DefaultPort              = 3000
	DefaultReadTimeout       = 10 * time.Second
	DefaultReadHeaderTimeout = 5 * time.Second
	DefaultWriteTimeout      = 10 * time.Second
	DefaultIdleTimeout       = 65 * time.Second
	DefaultKeepAliveTimeout  = 65 * time.Second
	DefaultShutdownTimeout   = 10 * time.Second
)

// NewOptions creates a new Options instance.
func NewOptions(options ...Option) *Options {
	opts := &Options{
		Name:              DefaultName,
		Port:              DefaultPort,
		ReadTimeout:       DefaultReadTimeout,
		ReadHeaderTimeout: DefaultReadHeaderTimeout,
		WriteTimeout:      DefaultWriteTimeout,
		IdleTimeout:       DefaultIdleTimeout,
		KeepAliveTimeout:  DefaultKeepAliveTimeout,
		ShutdownTimeout:   DefaultShutdownTimeout,
	}

	for _, option := range options {
		option(opts)
	}

	return opts
}

// WithName sets the name for the options.
func WithName(value string) Option {
	return func(o *Options) {
		o.Name = value
	}
}

// WithPort sets the port for the options.
func WithPort(value uint32) Option {
	return func(o *Options) {
		o.Port = value
	}
}

// WithCertificateFile sets the certificate file for the options.
func WithCertificateFile(value string) Option {
	return func(o *Options) {
		o.CertificateFile = value
	}
}

// WithCertificateKeyFile sets the certificate key file for the options.
func WithCertificateKeyFile(value string) Option {
	return func(o *Options) {
		o.CertificateKeyFile = value
	}
}

// WithRenderer sets the renderer for the options.
func WithRenderer(value render.IRenderer) Option {
	return func(o *Options) {
		o.Renderer = value
	}
}

// WithReadTimeout sets the read timeout for the options.
func WithReadTimeout(value time.Duration) Option {
	return func(o *Options) {
		o.ReadTimeout = value
	}
}

// WithReadHeaderTimeout sets the read header timeout for the options.
func WithReadHeaderTimeout(value time.Duration) Option {
	return func(o *Options) {
		o.ReadHeaderTimeout = value
	}
}

// WithWriteTimeout sets the write timeout for the options.
func WithWriteTimeout(value time.Duration) Option {
	return func(o *Options) {
		o.WriteTimeout = value
	}
}

// WithIdleTimeout sets the idle timeout for the options.
func WithIdleTimeout(value time.Duration) Option {
	return func(o *Options) {
		o.IdleTimeout = value
	}
}

// WithKeepAliveTimeout sets the keep alive timeout for the options.
func WithKeepAliveTimeout(value time.Duration) Option {
	return func(o *Options) {
		o.KeepAliveTimeout = value
	}
}

// WithShutdownTimeout sets the shutdown timeout for the options.
func WithShutdownTimeout(value time.Duration) Option {
	return func(o *Options) {
		o.ShutdownTimeout = value
	}
}
