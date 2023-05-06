package cmd

import (
	"context"
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
	return delete(cmd.Context(), args)
}

func delete(ctx context.Context, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("%w 'DELETE'", ErrInvalidArgs)
	}

	return c.Delete(ctx, []byte(args[0]))
}
