package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	cmd.AddCommand(
		dataCmd,
		outputsCmd,
		variablesCmd,
	)
}

var cmd = &cobra.Command{
	Use:   "tfsort",
	Short: "",
}

func Execute() error {
	return cmd.Execute()
}
