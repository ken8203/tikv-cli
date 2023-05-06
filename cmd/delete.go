package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"del"},
	Short:   "Delete a key",
	RunE:    deleteRunE,
}

func deleteRunE(cmd *cobra.Command, args []string) error {
	client, err := newClient()
	if err != nil {
		return fmt.Errorf("new client: %v", err)
	}
	defer client.Close(cmd.Context())

	return client.Delete(cmd.Context(), []byte(args[0]))
}
