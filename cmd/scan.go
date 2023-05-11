package cmd

import (
	"context"
	"fmt"
	"io"

	"github.com/spf13/cobra"
)

const defaultLimit = 100

var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Scan keys",
	RunE: func(cmd *cobra.Command, args []string) error {
		return scan(cmd.Context(), cmd.OutOrStdout(), args)
	},
}

func scan(ctx context.Context, w io.Writer, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("%w 'SCAN'", ErrInvalidArgs)
	}

	entries, err := c.Scan(ctx, []byte(args[0]), defaultLimit)
	if err != nil {
		return err
	}

	for i, entry := range entries {
		fmt.Fprintf(w, "%*d) key=%s value=%s\n", padding(len(entries)), i, entry.Key(), entry.Value())
	}
	return nil
}

func padding(n int) int {
	var count int
	for n > 0 {
		n = n / 10
		count++
	}
	return count
}
