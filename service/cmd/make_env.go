package cmd

import (
	"fmt"

	"github.com/leliuga/cdk/configurator"
	"github.com/leliuga/cdk/service"
	"github.com/spf13/cobra"
)

func NewMakeEnvCmd(options *service.Options) *cobra.Command {
	var flagFormat string
	cmd := &cobra.Command{
		Use:     "env",
		Aliases: []string{"e"},
		Short:   "Make a environment file",
		Long:    `Make a environment file for the service ` + options.Name,
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			env := configurator.ToEnv(options, options.Name, options.Description, "")
			fmt.Print(env.Marshal(flagFormat))

			return nil
		},
	}
	cmd.Flags().StringVarP(&flagFormat, "format", "f", "init", "Format (json|yaml|dotenv|init)"+"``")

	return cmd
}
