package cmd

import (
	"context"
	"fmt"

	"github.com/ken8203/tikv-cli/internal/client"
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

	return put(client, cmd.Context(), args)
}

func put(client client.Client, ctx context.Context, args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("%w 'PUT'", ErrInvalidArgs)
	}

	if err := client.Put(ctx, []byte(args[0]), []byte(args[1])); err != nil {
		return err
	}
	return nil
}
