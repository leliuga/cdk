package cmd

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/leliuga/cdk/service"
	"github.com/leliuga/cdk/types"
	"github.com/spf13/cobra"
)

// NewMakeContainerFileCmd returns a new make container file command.
func NewMakeContainerFileCmd(options *service.Options) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "containerfile",
		Aliases: []string{"c"},
		Short:   "Make a container file",
		Long:    `Make a container file for the service ` + options.Name,
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Print(containerFile(options))

			return nil
		},
	}

	return cmd
}

func containerFile(options *service.Options) string {
	var buf bytes.Buffer
	serviceName := strings.ToLower(options.Name)
	labelPrefix := "org.opencontainers.image."
	labels := types.Map[string]{
		labelPrefix + "title":         options.Name,
		labelPrefix + "description":   "A service " + options.Name + " for " + service.DefaultApplicationName,
		labelPrefix + "licenses":      "MPL-2.0",
		labelPrefix + "authors":       service.DefaultApplicationName + " Authors",
		labelPrefix + "documentation": options.BuildInfo.Repository + "/blob/" + options.BuildInfo.Commit + "/README.md",
		labelPrefix + "source":        options.BuildInfo.Repository,
		labelPrefix + "version":       options.BuildInfo.Commit,
	}
	probe := options.Runtime.Probe

	buf.WriteString("## Build\n")
	buf.WriteString(fmt.Sprintf("FROM ghcr.io/leliuga/:%v-alpine AS build\n\n", options.BuildInfo.GoVersion))
	buf.WriteString("ARG TARGETOS TARGETARCH\n")
	buf.WriteString("RUN apk add --update --no-cache git build-base\n")
	buf.WriteString("WORKDIR /src\n")
	buf.WriteString("COPY go.mod go.sum .\n")
	buf.WriteString("RUN --mount=type=cache,target=/go/src go mod download\n")
	buf.WriteString("RUN --mount=type=cache,target=/go/src go mod verify\n")
	buf.WriteString("COPY . .\n")
	buf.WriteString("RUN --mount=type=cache,target=/go/src go list -mod=readonly all\n")
	buf.WriteString("RUN --mount=type=cache,target=/go/src --mount=type=cache,target=/root/.cache/go-build OS=$TARGETOS ARCH=$TARGETARCH make " + serviceName + "\n\n")

	buf.WriteString("## Deployment\n")
	buf.WriteString(fmt.Sprintf("FROM %s AS final\n\n", service.DefaultBaseImage))

	for _, key := range labels.Keys() {
		buf.WriteString("LABEL " + key + "=\"" + labels[key] + "\"\n")
	}

	buf.WriteString(fmt.Sprintf("COPY --from=build /src/bin/%s /usr/bin/%s\n", serviceName, serviceName))
	buf.WriteString(fmt.Sprintf("EXPOSE %v/tcp\n", options.Port))
	buf.WriteString(fmt.Sprintf("HEALTHCHECK --start-period=%vs --interval=%vs --timeout=%vs --retries=%v CMD wget --no-verbose --tries=1 --spider 'http://localhost:%v%s' || exit 1\n", probe.InitialDelaySeconds, probe.PeriodSeconds, probe.TimeoutSeconds, probe.FailureThreshold, options.Port, service.DefaultPathMonitoring))
	buf.WriteString(fmt.Sprintf(`CMD ["%s", "serve"]`, serviceName))
	buf.WriteString("\nSTOPSIGNAL SIGTERM\n")

	return buf.String()
}
