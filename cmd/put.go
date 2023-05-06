package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
)

var putCmd = &cobra.Command{
	Use:   "put",
	Short: "Put a key",
	RunE:  putRunE,
}

func putRunE(cmd *cobra.Command, args []string) error {
	return put(cmd.Context(), args)
}

func put(ctx context.Context, args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("%w 'PUT'", ErrInvalidArgs)
	}

	if err := c.Put(ctx, []byte(args[0]), []byte(args[1])); err != nil {
		return err
	}
	return nil
}
