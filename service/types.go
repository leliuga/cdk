// Package service provides a service definition.
package service

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/leliuga/cdk/database"
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
		Port                    int32                         `json:"port"`
		Network                 string                        `json:"network"`
		Domain                  string                        `json:"domain"`
		CertificateFile         string                        `json:"certificate_file"`
		CertificateKeyFile      string                        `json:"certificate_key_file"`
		Views                   fiber.Views                   `json:"-"`
		BodyLimit               int                           `json:"body_limit"`
		Concurrency             int                           `json:"concurrency"`
		ReadTimeout             time.Duration                 `json:"read_timeout"`
		WriteTimeout            time.Duration                 `json:"write_timeout"`
		IdleTimeout             time.Duration                 `json:"idle_timeout"`
		ShutdownTimeout         time.Duration                 `json:"shutdown_timeout"`
		ReadBufferSize          int                           `json:"read_buffer_size"`
		WriteBufferSize         int                           `json:"write_buffer_size"`
		EnableTrustedProxyCheck bool                          `json:"enable_trusted_proxy_check"`
		TrustedProxies          []string                      `json:"trusted_proxies"`
		EnablePrintRoutes       bool                          `json:"enable_print_routes"`
		BuildInfo               *BuildInfo                    `json:"build_info"`
		Runtime                 *Runtime                      `json:"runtime"`
		Database                *database.Options             `json:"database"`
		ErrorHandler            func(*fiber.Ctx, error) error `json:"-"`
		Handlers                IHandlers                     `json:"-"`
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
		Provider           Provider              `json:"provider"`
		Region             string                `json:"region"`
		Zone               string                `json:"zone"`
		Namespace          string                `json:"namespace"`
		ServiceAccountName string                `json:"service_account_name"`
		Engine             Engine                `json:"engine"`
		Replicas           int32                 `json:"replicas"`
		Resources          *ResourceRequirements `json:"resources"`
		Probe              *RuntimeProbe         `json:"probe"`
	}

	// ResourceRequirements defines the resource requirements for a Service.
	ResourceRequirements struct {
		Limits   corev1.ResourceList `json:"limits"`
		Requests corev1.ResourceList `json:"requests"`
	}

	// RuntimeProbe defines the runtime probe for a Service.
	RuntimeProbe struct {
		InitialDelaySeconds int32 `json:"initial_delay_seconds"`
		TimeoutSeconds      int32 `json:"timeout_seconds"`
		PeriodSeconds       int32 `json:"period_seconds"`
		SuccessThreshold    int32 `json:"success_threshold"`
		FailureThreshold    int32 `json:"failure_threshold"`
	}

	// Engine defines the engine for a Service runtime.
	Engine uint8

	// Environment defines the environment in which the Service is running.
	Environment uint8

	// Provider defines the cloud provider for a Service runtime.
	Provider uint8

	// Option represents the service option.
	Option func(o *Options)

	// IHandlers represents the service handlers interface.
	IHandlers interface {
		Init(*Service)
	}
)
