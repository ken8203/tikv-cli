package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var putCmd = &cobra.Command{
	Use:   "put",
	Short: "Put a key",
	RunE:  putRunE,
}

func putRunE(cmd *cobra.Command, args []string) error {
	client, err := newClient()
	if err != nil {
		return fmt.Errorf("new client: %v", err)
	}
	defer client.Close(cmd.Context())

	return client.Put(cmd.Context(), []byte(args[0]), []byte(args[1]))
}
