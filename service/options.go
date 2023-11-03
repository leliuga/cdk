package service

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/goccy/go-json"
	"github.com/goccy/go-yaml"
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

	DefaultConfigDirectory = "/etc/leliuga/"
)

// Default paths for the service
const (
	DefaultPathMonitoring = "/monitoring"
	DefaultPathDiscovery  = "/discovery"
)

// NewOptions creates a new options.
func NewOptions(options ...Option) *Options {
	opts := Options{
		Name:                    DefaultName,
		Port:                    DefaultPort,
		Network:                 DefaultNetwork,
		Domain:                  strings.ToLower(DefaultName + "." + DefaultDomain),
		BodyLimit:               DefaultBodyLimit,
		Concurrency:             DefaultConcurrency,
		ReadTimeout:             DefaultReadTimeout,
		WriteTimeout:            DefaultWriteTimeout,
		IdleTimeout:             DefaultIdleTimeout,
		ShutdownTimeout:         DefaultShutdownTimeout,
		ReadBufferSize:          DefaultReadBufferSize,
		WriteBufferSize:         DefaultWriteBufferSize,
		EnableTrustedProxyCheck: DefaultEnableTrustedProxyCheck,
		DisableStartupMessage:   false,
		EnablePrintRoutes:       true,
		TrustedProxies:          []string{},
		BuildInfo:               NewBuildInfo("", "", ""),
		Runtime:                 NewRuntime(),
		ErrorHandler:            fiber.DefaultErrorHandler,
		Kernel:                  NewKernel(),
		Database:                database.NewOptions(),
	}

	for _, option := range options {
		option(&opts)
	}

	return &opts
}

// NewOptionsFromConfig creates a new options from config.
func NewOptionsFromConfig(cfgName string, options ...Option) (*Options, error) {
	opts := NewOptions(
		options...,
	)

	filename := strings.ToLower(path.Join(DefaultConfigDirectory, opts.Name, cfgName))
	ext := filepath.Ext(filename)

	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	switch ext {
	case ".yaml", ".yml":
		if err = yaml.UnmarshalWithOptions(content, opts, yaml.UseJSONUnmarshaler()); err != nil {
			return nil, err
		}
	case ".json":
		if err = json.Unmarshal(content, opts); err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("unsupported config file extension: %s", ext)
	}

	return opts, nil
}

// WithName sets the name for the service.
func WithName(value string) Option {
	return func(o *Options) {
		o.Name = value
		o.Domain = strings.ToLower(value + "." + DefaultDomain)
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
		o.Domain = strings.ToLower(value)
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
func WithTrustedProxies(values []string) Option {
	return func(o *Options) {
		o.TrustedProxies = values
	}
}

// WithDisableStartupMessage sets the disable startup message for the service.
func WithDisableStartupMessage(value bool) Option {
	return func(o *Options) {
		o.DisableStartupMessage = value
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

// WithKernel sets the kernel for the service.
func WithKernel(value IKernel) Option {
	return func(o *Options) {
		o.Kernel = value
	}
}
