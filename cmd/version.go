package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// Version can be set the version info from the -ldflags when building
var Version string

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of tikv-cli",
	RunE: func(_ *cobra.Command, _ []string) error {
		_, err := os.Stdout.Write([]byte("tikv-cli " + Version))
		return err
	},
}
