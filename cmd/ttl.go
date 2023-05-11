package cmd

import (
	"context"
	"fmt"
	"io"

	"github.com/ken8203/tikv-cli/internal/client"
	"github.com/spf13/cobra"
)

var ttlCmd = &cobra.Command{
	Use:   "ttl",
	Short: "Get the TTL of a key",
	RunE: func(cmd *cobra.Command, args []string) error {
		return ttl(cmd.Context(), cmd.OutOrStdout(), args)
	},
}

func ttl(ctx context.Context, w io.Writer, args []string) error {
	if Mode == string(client.ModeTxn) {
		return fmt.Errorf("%w 'TTL'", ErrCommandNotSupported)
	}
	if len(args) < 1 {
		return fmt.Errorf("%w 'TTL'", ErrInvalidArgs)
	}

	ttl, err := c.TTL(ctx, []byte(args[0]))
	if err != nil {
		return err
	}

	if _, err := fmt.Fprintln(w, ttl); err != nil {
		return err
	}
	return nil
}
