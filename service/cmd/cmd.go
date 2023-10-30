package cmd

import (
	"strings"

	"github.com/leliuga/cdk/service"
	"github.com/spf13/cobra"
)

// Execute executes the service.
func Execute(svc *service.Service, commands ...*cobra.Command) error {
	name := svc.Options.Name
	cmd := &cobra.Command{
		Use:   strings.ToLower(name),
		Short: "The " + name + " is a service.",
		Long: `The ` + name + ` is a service within a distributed application (` + service.DefaultApplicationName + `).

All of service features can be driven through the various commands below.
For help with any of those, simply call them with --help.`,
		SilenceUsage:  true,
		SilenceErrors: true,
	}

	cmd.AddCommand(
		NewInspectCmd(svc.Options),
		NewMakeCmd(svc.Options),
		NewServeCmd(svc),
	)
	cmd.AddCommand(commands...)
	cmd.InitDefaultHelpCmd()

	return cmd.Execute()
}
