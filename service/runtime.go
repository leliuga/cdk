package service

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/leliuga/cdk/types"
	"github.com/leliuga/cdk/validation"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
)

var (
	NamespaceRegex = regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`)
)

// Default values for the Service runtime
const (
	DefaultServiceNamespace                        = "leliuga"
	DefaultServiceAccountName                      = "leliuga"
	DefaultServiceReplicas                         = 1
	DefaultServiceResourcesLimitCPU                = "2000m"
	DefaultServiceResourcesLimitMemory             = "1Gi"
	DefaultServiceResourcesLimitEphemeralStorage   = "1Gi"
	DefaultServiceResourcesRequestCPU              = "100m"
	DefaultServiceResourcesRequestMemory           = "32Mi"
	DefaultServiceResourcesRequestEphemeralStorage = "100Mi"
	DefaultServiceProbeInitialDelaySeconds         = 3
	DefaultServiceProbeTimeoutSeconds              = 1
	DefaultServiceProbePeriodSeconds               = 10
	DefaultServiceProbeSuccessThreshold            = 1
	DefaultServiceProbeFailureThreshold            = 3
)

const (
	InvalidNamespace = "A namespace must consist of lower case alphanumeric characters or '-', and must start and end with an alphanumeric character."
)

// Default values for the image
const (
	// DefaultImageVendor is the default image vendor for the service container
	DefaultImageVendor = "ghcr.io/leliuga"

	// DefaultImagePrefix is the default image prefix for the service container
	DefaultImagePrefix = DefaultImageVendor + "/service"

	// DefaultBaseImage is the default base image for the service container
	DefaultBaseImage = DefaultImageVendor + "/base"

	// DefaultGolangImage is the default golang image for the service container
	DefaultGolangImage = DefaultImageVendor + "/golang"
)

// NewRuntime creates a new Runtime.
func NewRuntime() *Runtime {
	return &Runtime{
		Provider:           ProviderBareMetal,
		Namespace:          DefaultServiceNamespace,
		ServiceAccountName: DefaultServiceAccountName,
		Engine:             EngineKubernetes,
		Replicas:           DefaultServiceReplicas,
		Resources: &ResourceRequirements{
			Limits: corev1.ResourceList{
				corev1.ResourceCPU:              resource.MustParse(DefaultServiceResourcesLimitCPU),
				corev1.ResourceMemory:           resource.MustParse(DefaultServiceResourcesLimitMemory),
				corev1.ResourceEphemeralStorage: resource.MustParse(DefaultServiceResourcesLimitEphemeralStorage),
			},
			Requests: corev1.ResourceList{
				corev1.ResourceCPU:              resource.MustParse(DefaultServiceResourcesRequestCPU),
				corev1.ResourceMemory:           resource.MustParse(DefaultServiceResourcesRequestMemory),
				corev1.ResourceEphemeralStorage: resource.MustParse(DefaultServiceResourcesRequestEphemeralStorage),
			},
		},

		Probe: &RuntimeProbe{
			InitialDelaySeconds: DefaultServiceProbeInitialDelaySeconds,
			TimeoutSeconds:      DefaultServiceProbeTimeoutSeconds,
			PeriodSeconds:       DefaultServiceProbePeriodSeconds,
			SuccessThreshold:    DefaultServiceProbeSuccessThreshold,
			FailureThreshold:    DefaultServiceProbeFailureThreshold,
		},
	}
}

// Validate makes Runtime validatable by implementing [validation.Validatable] interface.
func (r *Runtime) Validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.Provider, validation.In(validation.ToAnySliceFromMapKeys(ProviderNames)...).Error(fmt.Sprintf("A provider value must be one of: %s", strings.Join(types.ToMap(ProviderNames).Values(), ", ")))),
		validation.Field(&r.Namespace, validation.Required, validation.Length(1, 63), validation.Match(NamespaceRegex).Error(InvalidNamespace)),
		validation.Field(&r.Engine, validation.Required, validation.In(validation.ToAnySliceFromMapKeys(EngineNames)...).Error(fmt.Sprintf("A engine value must be one of: %s", strings.Join(types.ToMap(EngineNames).Values(), ", ")))),
	)
}

// ToResourceRequirements converts the Runtime to a ResourceRequirements.
func (r *Runtime) ToResourceRequirements() corev1.ResourceRequirements {
	return corev1.ResourceRequirements{
		Limits:   r.Resources.Limits,
		Requests: r.Resources.Requests,
	}
}
