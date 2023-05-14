package cmd

import (
	"context"
	"fmt"
	"io"
	"strings"
)

func use(ctx context.Context, w io.Writer, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("%w 'USE'", ErrInvalidArgs)
	}

	for _, arg := range args {
		v := strings.SplitN(arg, "=", 2)
		if len(v) < 2 {
			continue
		}

		switch v[0] {
		case "mode":
			if _, err := fmt.Sscanf(arg, "mode=%s", &Mode); err != nil {
				return err
			}
			break

		case "api-version":
			if _, err := fmt.Sscanf(arg, "api-version=%s", &APIVersion); err != nil {
				return err
			}
			break
		}
	}

	var err error
	if c, err = newClient(); err != nil {
		return err
	}

	if _, err := fmt.Fprintf(w, "Switch to mode [%s] and api version [%s]\n", Mode, APIVersion); err != nil {
		return err
	}
	return nil
}
