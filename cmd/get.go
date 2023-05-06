package cmd

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	tikverror "github.com/tikv/client-go/v2/error"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a key",
	RunE:  getRunE,
}

func getRunE(cmd *cobra.Command, args []string) error {
	value, err := get(cmd.Context(), args)
	if err != nil {
		return err
	}

	if _, err := fmt.Fprintln(os.Stdout, string(value)); err != nil {
		return err
	}
	return nil
}

func get(ctx context.Context, args []string) ([]byte, error) {
	if len(args) < 1 {
		return nil, fmt.Errorf("%w 'GET'", ErrInvalidArgs)
	}

	value, err := c.Get(ctx, []byte(args[0]))
	if err != nil {
		if errors.Is(err, tikverror.ErrNotExist) {
			return nil, ErrNotExist
		}
		return nil, err
	}

	return value, nil
}
