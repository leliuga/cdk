package cmd

import (
	"fmt"
	"strings"

	compose "github.com/compose-spec/compose-go/types"
	"github.com/goccy/go-yaml"
	"github.com/leliuga/cdk/service"
	"github.com/leliuga/cdk/types"
	"github.com/spf13/cobra"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

// NewMakeDeploymentCmd returns a new make deployment command.
func NewMakeDeploymentCmd(options *service.Options) *cobra.Command {
	var flagFormat string
	serviceName := strings.ToLower(options.Name)
	cmd := &cobra.Command{
		Use:     "deployment",
		Aliases: []string{"d"},
		Short:   "Make a deployment manifest",
		Long:    `Make a deployment manifest (Kubernetes or Docker Swarm) for the service ` + options.Name,
		Args:    cobra.NoArgs,
		Example: serviceName + ` make deployment | kubectl apply -f -
  Make a deployment manifest for the service ` + options.Name + ` and apply it to the Kubernetes cluster

` + options.Name + ` make deployment | docker stack deploy -c - ` + serviceName + `
  Make a deployment manifest for the service ` + options.Name + ` and deploy it to the Docker Swarm cluster
`,
		RunE: func(cmd *cobra.Command, args []string) error {
			switch options.Runtime.Engine {
			case service.EngineKubernetes:
				if flagFormat == "terraform" {
					fmt.Print(kubernetesDeploymentTerraform(options))
					return nil
				}

				fmt.Print(kubernetesDeploymentNative(options))
			case service.EngineDockerSwarm:
				if flagFormat == "terraform" {
					fmt.Print(dockerSwarmDeploymentTerraform(options))
					return nil
				}

				fmt.Print(dockerSwarmDeploymentNative(options))
			}

			return nil
		},
	}
	cmd.Flags().StringVarP(&flagFormat, "format", "f", "native", "Format (native|terraform)"+"``")

	return cmd
}

func kubernetesDeploymentNative(options *service.Options) string {
	instanceName := fmt.Sprintf("service-%s", strings.ToLower(options.Name))
	servicePortName := "http"
	labelPrefix := strings.ToLower("service." + service.DefaultDomain + "/")
	labels := types.Map[string]{
		labelPrefix + "application": service.DefaultApplicationName,
		labelPrefix + "name":        options.Name,
		labelPrefix + "domain":      options.Domain,
		labelPrefix + "vendor":      service.DefaultVendor,
		labelPrefix + "repository":  options.BuildInfo.Repository,
		labelPrefix + "version":     options.BuildInfo.Commit,
		labelPrefix + "go":          options.BuildInfo.GoVersion,
		"kubernetes.io/arch":        options.BuildInfo.Architecture,
		"kubernetes.io/os":          options.BuildInfo.OS,
	}

	selectorLabels := types.Map[string]{
		labelPrefix + "application": labels[labelPrefix+"application"],
		labelPrefix + "name":        labels[labelPrefix+"name"],
	}

	deploy, _ := yaml.MarshalWithOptions(appsv1.Deployment{
		TypeMeta:   metav1.TypeMeta{APIVersion: "apps/v1", Kind: "Deployment"},
		ObjectMeta: metav1.ObjectMeta{Name: instanceName, Namespace: options.Runtime.Namespace, Labels: labels},
		Spec: appsv1.DeploymentSpec{
			Replicas: &options.Runtime.Replicas,
			Selector: &metav1.LabelSelector{MatchLabels: selectorLabels},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{Name: instanceName, Namespace: options.Runtime.Namespace, Labels: labels},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  instanceName,
							Image: imageName(options, "latest"),
							Ports: []corev1.ContainerPort{
								{
									Name:          servicePortName,
									ContainerPort: options.Port,
									Protocol:      corev1.ProtocolTCP,
								},
							},
							Resources: options.Runtime.ToResourceRequirements(),
							LivenessProbe: &corev1.Probe{
								ProbeHandler:        corev1.ProbeHandler{HTTPGet: &corev1.HTTPGetAction{Path: service.DefaultPathMonitoring, Port: intstr.FromInt32(options.Port), Scheme: corev1.URISchemeHTTP}},
								InitialDelaySeconds: options.Runtime.Probe.InitialDelaySeconds,
								TimeoutSeconds:      options.Runtime.Probe.TimeoutSeconds,
								PeriodSeconds:       options.Runtime.Probe.PeriodSeconds,
								SuccessThreshold:    options.Runtime.Probe.SuccessThreshold,
								FailureThreshold:    options.Runtime.Probe.FailureThreshold,
							},
							ImagePullPolicy: corev1.PullAlways,
						},
					},
					RestartPolicy:                 corev1.RestartPolicyAlways,
					TerminationGracePeriodSeconds: (*int64)(&options.ShutdownTimeout),
					ServiceAccountName:            options.Runtime.ServiceAccountName,
					Hostname:                      instanceName,
				},
			},
			Strategy: appsv1.DeploymentStrategy{Type: appsv1.RollingUpdateDeploymentStrategyType, RollingUpdate: &appsv1.RollingUpdateDeployment{MaxSurge: &intstr.IntOrString{Type: intstr.Int, IntVal: 1}, MaxUnavailable: &intstr.IntOrString{Type: intstr.Int, IntVal: 0}}},
		},
	}, yaml.UseJSONMarshaler())

	svc, _ := yaml.MarshalWithOptions(corev1.Service{
		TypeMeta:   metav1.TypeMeta{APIVersion: "v1", Kind: "Service"},
		ObjectMeta: metav1.ObjectMeta{Name: instanceName, Namespace: options.Runtime.Namespace, Labels: labels},
		Spec: corev1.ServiceSpec{
			Ports: []corev1.ServicePort{{
				Name:       servicePortName,
				Protocol:   corev1.ProtocolTCP,
				Port:       80,
				TargetPort: intstr.FromString(servicePortName),
			}},
			Selector:        selectorLabels,
			Type:            corev1.ServiceTypeClusterIP,
			SessionAffinity: corev1.ServiceAffinityClientIP,
		},
	}, yaml.UseJSONMarshaler())

	return fmt.Sprintf("%s---\n%s", deploy, svc)
}

func kubernetesDeploymentTerraform(options *service.Options) string {
	return "TODO: kubernetesDeploymentTerraform"
}

func dockerSwarmDeploymentNative(options *service.Options) string {
	instanceName := fmt.Sprintf("service-%s", strings.ToLower(options.Name))
	replicas := uint64(options.Runtime.Replicas)
	maxAttempts := uint64(options.Runtime.Probe.FailureThreshold)
	limitMemoryBytes, _ := options.Runtime.Resources.Limits.Memory().AsInt64()
	reservedMemoryBytes, _ := options.Runtime.Resources.Requests.Memory().AsInt64()

	labelPrefix := strings.ToLower("service." + service.DefaultDomain + "/")
	labels := compose.Labels{
		labelPrefix + "application": service.DefaultApplicationName,
		labelPrefix + "name":        options.Name,
		labelPrefix + "domain":      options.Domain,
		labelPrefix + "vendor":      service.DefaultVendor,
		labelPrefix + "repository":  options.BuildInfo.Repository,
		labelPrefix + "version":     options.BuildInfo.Commit,
		labelPrefix + "go":          options.BuildInfo.GoVersion,
	}
	svc, _ := yaml.MarshalWithOptions(compose.Project{
		Name: instanceName,
		Services: compose.Services{
			compose.ServiceConfig{
				Name: instanceName,
				Deploy: &compose.DeployConfig{
					Mode:     "replicated",
					Replicas: &replicas,
					Resources: compose.Resources{
						Limits: &compose.Resource{
							NanoCPUs:    options.Runtime.Resources.Limits.Cpu().String(),
							MemoryBytes: compose.UnitBytes(limitMemoryBytes),
						},
						Reservations: &compose.Resource{
							NanoCPUs:    options.Runtime.Resources.Limits.Cpu().String(),
							MemoryBytes: compose.UnitBytes(reservedMemoryBytes),
						},
					},
					RestartPolicy: &compose.RestartPolicy{
						Condition:   "on-failure",
						MaxAttempts: &maxAttempts,
					},
				},
				Hostname: instanceName,
				Image:    imageName(options, "latest"),
				Labels:   labels,
				Logging: &compose.LoggingConfig{
					Driver: "json-file",
					Options: map[string]string{
						"max-size": "20m",
						"max-file": "3",
						"tag":      "{{.ImageName}}|{{.Name}}|{{.ImageFullID}}|{{.FullID}}",
					},
				},
			},
		},
		Networks: compose.Networks{
			"default": compose.NetworkConfig{
				Name: strings.ToLower(service.DefaultApplicationName),
			},
		},
	}, yaml.UseJSONMarshaler())

	return fmt.Sprintf("%s", svc)
}

func dockerSwarmDeploymentTerraform(options *service.Options) string {
	return "TODO: dockerSwarmDeploymentTerraform"
}
