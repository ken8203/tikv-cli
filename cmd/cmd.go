package cmd

import (
	"fmt"
	"os"

	"github.com/ken8203/tikv-cli/internal/client"
	pingcaplog "github.com/pingcap/log"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var (
	// Host is the PD host address.
	Host string
	// Port is the PD port.
	Port string
	// Mode is the client mode: raw/txn
	Mode string
	// APIVersion is the API version: v1/v1ttl/v2
	APIVersion string
)

var rootCmd = &cobra.Command{
	Use:   "tikv-cli",
	Short: "Interact with TiKV cluster through PD",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {

	},
	RunE: shellRunE,
}

func init() {
	// Disable logging in tikv/client-go
	pingcaplog.ReplaceGlobals(zap.NewNop(), nil)

	rootCmd.PersistentFlags().StringVarP(&Host, "host", "h", "localhost", "PD host address")
	rootCmd.PersistentFlags().StringVarP(&Port, "port", "p", "2379", "PD port")
	rootCmd.PersistentFlags().StringVarP(&Mode, "mode", "m", "txn", "Client mode")
	rootCmd.PersistentFlags().StringVarP(&APIVersion, "api-version", "a", "v2", "API version")
	rootCmd.PersistentFlags().Bool("help", false, "help for tikv-cli")

	rootCmd.AddCommand(versionCmd, putCmd, getCmd, deleteCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// newClient creates a tikv client.
func newClient() (client.Client, error) {
	var v client.APIVersion
	switch APIVersion {
	case "v1":
		v = client.APIVersion1
	case "v1ttl":
		v = client.APIVersion1TTL
	case "v2":
		v = client.APIVersion2
	default:
		v = client.APIVersion2
	}

	c, err := client.New([]string{addr(Host, Port)}, client.Mode(Mode), v)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func addr(host, port string) string {
	return host + ":" + port
}
