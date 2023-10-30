package cmd

import (
	"github.com/leliuga/cdk/service"
	"github.com/spf13/cobra"
)

// NewServeCmd returns a new serve command.
func NewServeCmd(svc *service.Service) *cobra.Command {
	name := svc.Options.Name
	cmd := &cobra.Command{
		Use:     "serve",
		Aliases: []string{"s"},
		Short:   "Serve a service " + name,
		Long:    `Serve a service ` + name,
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return svc.Serve()
		},
	}

	return cmd
}
