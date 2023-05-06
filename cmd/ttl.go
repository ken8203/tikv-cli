package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var ttlCmd = &cobra.Command{
	Use:   "ttl",
	Short: "Get the TTL of a key",
	RunE:  ttlRunE,
}

func ttlRunE(cmd *cobra.Command, args []string) error {
	client, err := newClient()
	if err != nil {
		return fmt.Errorf("new client: %v", err)
	}
	defer client.Close(cmd.Context())

	value, err := ttl(cmd.Context(), args)
	if err != nil {
		return err
	}

	if _, err := fmt.Fprintln(os.Stdout, value); err != nil {
		return err
	}
	return nil
}

func ttl(ctx context.Context, args []string) (uint64, error) {
	if Mode == "txn" {
		return 0, fmt.Errorf("%w 'TTL'", ErrCommandNotSupported)
	}
	if len(args) < 1 {
		return 0, fmt.Errorf("%w 'TTL'", ErrInvalidArgs)
	}

	ttl, err := c.TTL(ctx, []byte(args[0]))
	if err != nil {
		return 0, err
	}

	return ttl, nil
}
