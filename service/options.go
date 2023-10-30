package service

import (
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/leliuga/cdk/database"
)

// Default values for the HTTP server
const (
	DefaultName                    = "service"
	DefaultPort                    = 3000
	DefaultNetwork                 = "tcp4"
	DefaultBodyLimit               = 4 * 1024 * 1024
	DefaultConcurrency             = 256 * 1024
	DefaultReadTimeout             = 5 * time.Second
	DefaultWriteTimeout            = 5 * time.Second
	DefaultIdleTimeout             = 65 * time.Second
	DefaultShutdownTimeout         = 10 * time.Second
	DefaultReadBufferSize          = 4 * 1024
	DefaultWriteBufferSize         = 4 * 1024
	DefaultEnableTrustedProxyCheck = false
	DefaultCompressedFileSuffix    = ".gz"
)

// Default paths for the service
const (
	DefaultPathMonitoring = "/monitoring"
	DefaultPathDiscovery  = "/discovery"
)

// NewOptions creates a new options.
func NewOptions(option ...Option) *Options {
	options := Options{
		Name:                    DefaultName,
		Port:                    DefaultPort,
		Network:                 DefaultNetwork,
		Domain:                  DefaultDomain,
		BodyLimit:               DefaultBodyLimit,
		Concurrency:             DefaultConcurrency,
		ReadTimeout:             DefaultReadTimeout,
		WriteTimeout:            DefaultWriteTimeout,
		IdleTimeout:             DefaultIdleTimeout,
		ShutdownTimeout:         DefaultShutdownTimeout,
		ReadBufferSize:          DefaultReadBufferSize,
		WriteBufferSize:         DefaultWriteBufferSize,
		EnableTrustedProxyCheck: DefaultEnableTrustedProxyCheck,
		TrustedProxies:          []string{},
		BuildInfo:               NewBuildInfo("", "", ""),
		Runtime:                 NewRuntime(),
		ErrorHandler:            fiber.DefaultErrorHandler,
		Handlers:                nil,
	}

	for _, o := range option {
		o(&options)
	}

	return &options
}

// WithName sets the name for the service.
func WithName(value string) Option {
	return func(o *Options) {
		o.Name = value
		o.Domain = strings.ToLower(value) + "." + DefaultDomain
	}
}

// WithPort sets the port for the service.
func WithPort(value int32) Option {
	return func(o *Options) {
		o.Port = value
	}
}

// WithNetwork sets the network for the service.
func WithNetwork(value string) Option {
	return func(o *Options) {
		o.Network = value
	}
}

// WithDomain sets the domain for the service.
func WithDomain(value string) Option {
	return func(o *Options) {
		o.Domain = value
	}
}

// WithCertificateFile sets the certificate file for the service.
func WithCertificateFile(value string) Option {
	return func(o *Options) {
		o.CertificateFile = value
	}
}

// WithCertificateKeyFile sets the certificate key file for the service.
func WithCertificateKeyFile(value string) Option {
	return func(o *Options) {
		o.CertificateKeyFile = value
	}
}

// WithViews sets the views for the service.
func WithViews(value fiber.Views) Option {
	return func(o *Options) {
		o.Views = value
	}
}

// WithBodyLimit sets the body limit for the service.
func WithBodyLimit(value int) Option {
	return func(o *Options) {
		o.BodyLimit = value
	}
}

// WithConcurrency sets the concurrency for the service.
func WithConcurrency(value int) Option {
	return func(o *Options) {
		o.Concurrency = value
	}
}

// WithReadTimeout sets the read timeout for the service.
func WithReadTimeout(value time.Duration) Option {
	return func(o *Options) {
		o.ReadTimeout = value
	}
}

// WithWriteTimeout sets the write timeout for the service.
func WithWriteTimeout(value time.Duration) Option {
	return func(o *Options) {
		o.WriteTimeout = value
	}
}

// WithIdleTimeout sets the idle timeout for the service.
func WithIdleTimeout(value time.Duration) Option {
	return func(o *Options) {
		o.IdleTimeout = value
	}
}

// WithShutdownTimeout sets the shutdown timeout for the service.
func WithShutdownTimeout(value time.Duration) Option {
	return func(o *Options) {
		o.ShutdownTimeout = value
	}
}

// WithReadBufferSize sets the read buffer size for the service.
func WithReadBufferSize(value int) Option {
	return func(o *Options) {
		o.ReadBufferSize = value
	}
}

// WithWriteBufferSize sets the write buffer size for the service.
func WithWriteBufferSize(value int) Option {
	return func(o *Options) {
		o.WriteBufferSize = value
	}
}

// WithEnableTrustedProxyCheck sets the enable trusted proxy check for the service.
func WithEnableTrustedProxyCheck(value bool) Option {
	return func(o *Options) {
		o.EnableTrustedProxyCheck = value
	}
}

// WithTrustedProxies sets the trusted proxies for the service.
func WithTrustedProxies(value []string) Option {
	return func(o *Options) {
		o.TrustedProxies = value
	}
}

// WithEnablePrintRoutes sets the enable print routes for the service.
func WithEnablePrintRoutes(value bool) Option {
	return func(o *Options) {
		o.EnablePrintRoutes = value
	}
}

// WithBuildInfo sets the build info for the service.
func WithBuildInfo(repository, commit, when string) Option {
	return func(o *Options) {
		o.BuildInfo = NewBuildInfo(repository, commit, when)
	}
}

// WithRuntime sets the runtime for the service.
func WithRuntime(value *Runtime) Option {
	return func(o *Options) {
		o.Runtime = value
	}
}

// WithDatabase sets the database for the service.
func WithDatabase(value *database.Options) Option {
	return func(o *Options) {
		o.Database = value
	}
}

// WithErrorHandler sets the error handler for the service.
func WithErrorHandler(value func(*fiber.Ctx, error) error) Option {
	return func(o *Options) {
		o.ErrorHandler = value
	}
}

// WithHandlers sets the handlers for the service.
func WithHandlers(value IHandlers) Option {
	return func(o *Options) {
		o.Handlers = value
	}
}
