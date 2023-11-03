package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/leliuga/cdk/service"
	"github.com/spf13/cobra"
)

// NewMakeOciImageCmd returns a new make OCI image command.
func NewMakeOciImageCmd(options *service.Options) *cobra.Command {
	imageTag := imageName(options, options.BuildInfo.Commit)
	imageTagLatest := imageName(options, "latest")

	cmd := &cobra.Command{
		Use:     "image",
		Aliases: []string{"i"},
		Short:   "Make a OCI image",
		Long:    `Make a OCI image for the service ` + options.Name,
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			p := exec.Command("docker", "build", "--platform", options.BuildInfo.Platform, "-t", imageTag, "-t", imageTagLatest, "--push", "-f", "-", ".")
			p.Stdin = strings.NewReader(containerFile(options))
			p.Stdout = os.Stdout
			p.Stderr = os.Stderr

			return p.Run()
		},
	}

	return cmd
}

func imageName(options *service.Options, version string) string {
	name := strings.ToLower(options.Name)

	return fmt.Sprintf("%s-%s:%s", service.DefaultImagePrefix, name, version)
}
