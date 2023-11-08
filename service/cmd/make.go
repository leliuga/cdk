package cmd

import (
	"github.com/leliuga/cdk/service"
	"github.com/spf13/cobra"
)

// NewMakeCmd returns a new make command.
func NewMakeCmd(options *service.Options) *cobra.Command {
	name := options.Name
	cmd := &cobra.Command{
		Use:     "make",
		Aliases: []string{"m"},
		Short:   "Make for the service " + name,
		Long:    `Make a container file, OCI image or manifests for the service ` + name,
		Args:    cobra.NoArgs,
		RunE:    func(cmd *cobra.Command, args []string) error { return cmd.Usage() },
	}

	cmd.AddCommand(
		NewMakeContainerFileCmd(options),
		NewMakeDeploymentCmd(options),
		NewMakeEnvCmd(options),
		NewMakeOciImageCmd(options),
	)

	return cmd
}
