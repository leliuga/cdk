// Package service provides a service definition.
package service

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/leliuga/cdk/database"
	"github.com/leliuga/cdk/types"
	corev1 "k8s.io/api/core/v1"
)

type (
	// Service represents the service.
	Service struct {
		*Options
		*fiber.App
	}

	// Options represents the service options.
	Options struct {
		Name                    string                        `json:"name"`
		Description             string                        `json:"description"`
		Port                    int32                         `json:"port"                       env:"PORT"`
		Network                 string                        `json:"network"`
		Domain                  string                        `json:"domain"                     env:"DOMAIN"`
		CertificateFile         string                        `json:"certificate_file"           env:"CERTIFICATE_FILE"`
		CertificateKeyFile      string                        `json:"certificate_key_file"       env:"CERTIFICATE_KEY_FILE"`
		Views                   fiber.Views                   `json:"-"`
		BodyLimit               int                           `json:"body_limit"                 env:"BODY_LIMIT"`
		Concurrency             int                           `json:"concurrency"                env:"CONCURRENCY"`
		ReadTimeout             time.Duration                 `json:"read_timeout"               env:"READ_TIMEOUT"`
		WriteTimeout            time.Duration                 `json:"write_timeout"              env:"WRITE_TIMEOUT"`
		IdleTimeout             time.Duration                 `json:"idle_timeout"               env:"IDLE_TIMEOUT"`
		ShutdownTimeout         time.Duration                 `json:"shutdown_timeout"           env:"SHUTDOWN_TIMEOUT"`
		ReadBufferSize          int                           `json:"read_buffer_size"           env:"READ_BUFFER_SIZE"`
		WriteBufferSize         int                           `json:"write_buffer_size"          env:"WRITE_BUFFER_SIZE"`
		EnableTrustedProxyCheck bool                          `json:"enable_trusted_proxy_check" env:"ENABLE_TRUSTED_PROXY_CHECK"`
		TrustedProxies          []string                      `json:"trusted_proxies"            env:"TRUSTED_PROXIES"`
		DisableStartupMessage   bool                          `json:"disable_startup_message"    env:"DISABLE_STARTUP_MESSAGE"`
		EnablePrintRoutes       bool                          `json:"enable_print_routes"        env:"ENABLE_PRINT_ROUTES"`
		BuildInfo               *BuildInfo                    `json:"build_info"`
		Runtime                 *Runtime                      `json:"runtime"                    env:"RUNTIME"`
		Database                *database.Options             `json:"database"                   env:"DATABASE"`
		ErrorHandler            func(*fiber.Ctx, error) error `json:"-"`
		Kernel                  IKernel                       `json:"-"`
	}

	// BuildInfo defines the build information for a Service.
	BuildInfo struct {
		Repository   string `json:"repository"`
		Commit       string `json:"commit"`
		When         string `json:"when"`
		GoVersion    string `json:"go_version"`
		Platform     string `json:"platform"`
		OS           string `json:"os"`
		Architecture string `json:"architecture"`
	}

	// Runtime defines the runtime for a Service.
	Runtime struct {
		Provider           Provider              `json:"provider"             env:"PROVIDER"`
		Region             string                `json:"region"               env:"REGION"`
		Zone               string                `json:"zone"                 env:"ZONE"`
		Namespace          string                `json:"namespace"            env:"NAMESPACE"`
		ServiceAccountName string                `json:"service_account_name" env:"SERVICE_ACCOUNT_NAME"`
		Engine             Engine                `json:"engine"               env:"ENGINE"`
		Replicas           int32                 `json:"replicas"             env:"REPLICAS"`
		Resources          *ResourceRequirements `json:"resources"            env:"RESOURCES"`
		Probe              *RuntimeProbe         `json:"probe"                env:"PROBE"`
	}

	// ResourceRequirements defines the resource requirements for a Service.
	ResourceRequirements struct {
		Limits   corev1.ResourceList `json:"limits"   env:"LIMITS"`
		Requests corev1.ResourceList `json:"requests" env:"REQUESTS"`
	}

	// RuntimeProbe defines the runtime probe for a Service.
	RuntimeProbe struct {
		InitialDelaySeconds int32 `json:"initial_delay_seconds" env:"INITIAL_DELAY_SECONDS"`
		TimeoutSeconds      int32 `json:"timeout_seconds"       env:"TIMEOUT_SECONDS"`
		PeriodSeconds       int32 `json:"period_seconds"        env:"PERIOD_SECONDS"`
		SuccessThreshold    int32 `json:"success_threshold"     env:"SUCCESS_THRESHOLD"`
		FailureThreshold    int32 `json:"failure_threshold"     env:"FAILURE_THRESHOLD"`
	}

	// Kernel represents the service kernel.
	Kernel struct {
		IKernel

		instances types.Map[any]
	}

	// Engine defines the engine for a Service runtime.
	Engine uint8

	// Environment defines the environment in which the Service is running.
	Environment uint8

	// Provider defines the cloud provider for a Service runtime.
	Provider uint8

	// Option represents the service option.
	Option func(o *Options)

	// IKernel represents the service kernel interface.
	IKernel interface {
		// Boot the kernel.
		Boot(*Service) error

		// Shutdown the kernel.
		Shutdown(context.Context) error

		// Set an instance to the kernel.
		Set(key string, instance any)

		// Get an instance from the kernel.
		Get(key string) any

		// Has an instance from the kernel.
		Has(key string) bool

		// Instances returns all instances from the kernel.
		Instances() types.Map[any]
	}
)
