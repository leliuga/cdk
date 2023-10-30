package cmd

import (
	"fmt"

	"github.com/goccy/go-json"
	"github.com/goccy/go-yaml"
	"github.com/leliuga/cdk/service"
	"github.com/spf13/cobra"
)

// NewInspectCmd returns a new inspect command.
func NewInspectCmd(options *service.Options) *cobra.Command {
	name := options.Name
	var flagFormat string
	cmd := &cobra.Command{
		Use:     "inspect",
		Aliases: []string{"i"},
		Short:   "Inspect a service " + name,
		Long:    `Inspect a service ` + name,
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			switch flagFormat {
			case "json":
				if marshal, err := json.Marshal(&options); err == nil {
					fmt.Print(string(marshal) + "\n")
				}
			case "yaml":
				if marshal, err := yaml.MarshalWithOptions(&options, yaml.UseJSONMarshaler()); err == nil {
					fmt.Print(string(marshal))
				}
			}

			return nil
		},
	}
	cmd.Flags().StringVarP(&flagFormat, "format", "f", "yaml", "Format (json|yaml)"+"``")

	return cmd
}
