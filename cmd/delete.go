package cmd

import (
	"context"
	"fmt"
	"io"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"del"},
	Short:   "Delete a key",
	RunE: func(cmd *cobra.Command, args []string) error {
		return delete(cmd.Context(), cmd.OutOrStdout(), args)
	},
}

func delete(ctx context.Context, w io.Writer, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("%w 'DELETE'", ErrInvalidArgs)
	}

	if err := c.Delete(ctx, []byte(args[0])); err != nil {
		return err
	}

	if _, err := fmt.Fprintln(w, "OK"); err != nil {
		return err
	}
	return nil
}
