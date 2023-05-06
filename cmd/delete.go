package cmd

import (
	"context"
	"fmt"

	"github.com/ken8203/tikv-cli/internal/client"
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

	return delete(client, cmd.Context(), args)
}

func delete(client client.Client, ctx context.Context, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("%w 'DELETE'", ErrInvalidArgs)
	}

	return client.Delete(ctx, []byte(args[0]))
}
