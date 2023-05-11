package cmd

import (
	"fmt"
	"io"

	"github.com/spf13/cobra"
)

// Version can be set the version info from the -ldflags when building
var Version string

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of tikv-cli",
	RunE: func(cmd *cobra.Command, _ []string) error {
		return version(cmd.OutOrStdout())
	},
}

func version(w io.Writer) error {
	_, err := fmt.Fprintln(w, "tikv-cli", Version)
	return err
}
