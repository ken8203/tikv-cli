package cmd

import (
	"context"
	"errors"
	"fmt"
	"io"

	"github.com/spf13/cobra"
	tikverror "github.com/tikv/client-go/v2/error"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a key",
	RunE: func(cmd *cobra.Command, args []string) error {
		return get(cmd.Context(), cmd.OutOrStdout(), args)
	},
}

func get(ctx context.Context, w io.Writer, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("%w 'GET'", ErrInvalidArgs)
	}

	value, err := c.Get(ctx, []byte(args[0]))
	if err != nil {
		if errors.Is(err, tikverror.ErrNotExist) {
			return ErrNotExist
		}
		return err
	}

	if _, err := fmt.Fprintln(w, string(value)); err != nil {
		return err
	}
	return nil
}
