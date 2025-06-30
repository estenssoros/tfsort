package cmd

import (
	"fmt"

	"github.com/estenssoros/tfsort/pkg/tfsort"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var outputsCmd = &cobra.Command{
	Use:     "outputs",
	Short:   "",
	Args:    cobra.ExactArgs(1),
	PreRunE: func(cmd *cobra.Command, args []string) error { return nil },
	RunE: func(cmd *cobra.Command, args []string) error {
		output, err := tfsort.Outputs(args[0])
		if err != nil {
			return errors.Wrap(err, "tfsort.Run")
		}
		fmt.Println(output)
		return nil
	},
}
