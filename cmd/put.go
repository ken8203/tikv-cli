package cmd

import (
	"context"
	"fmt"
	"io"

	"github.com/spf13/cobra"
)

var putCmd = &cobra.Command{
	Use:   "put",
	Short: "Put a key",
	RunE: func(cmd *cobra.Command, args []string) error {
		return put(cmd.Context(), cmd.OutOrStdout(), args)
	},
}

func put(ctx context.Context, w io.Writer, args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("%w 'PUT'", ErrInvalidArgs)
	}

	if err := c.Put(ctx, []byte(args[0]), []byte(args[1])); err != nil {
		return err
	}

	if _, err := fmt.Fprintln(w, "OK"); err != nil {
		return err
	}
	return nil
}
