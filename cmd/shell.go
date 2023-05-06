package cmd

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/ken8203/tikv-cli/internal/terminal"
	"github.com/spf13/cobra"
)

// shellRunE is the entry of shell command.
func shellRunE(cmd *cobra.Command, _ []string) error {
	client, err := newClient()
	if err != nil {
		return fmt.Errorf("new client: %v", err)
	}
	defer client.Close(cmd.Context())

	executeFn := func(ctx context.Context, command string, args ...string) {
		switch strings.ToLower(command) {
		case "put":
			if err := put(ctx, args); err != nil {
				fmt.Fprintln(os.Stdout, err.Error())
				break
			}

			fmt.Fprintln(os.Stdout, "OK")
			break

		case "get":
			value, err := get(ctx, args)
			if err != nil {
				fmt.Fprintln(os.Stdout, err.Error())
				break
			}

			fmt.Fprintln(os.Stdout, string(value))
			break

		case "delete":
			if err := delete(ctx, args); err != nil {
				fmt.Fprintln(os.Stdout, err.Error())
				break
			}

			fmt.Fprintln(os.Stdout, "OK")
			break

		case "ttl":
			ttl, err := ttl(ctx, args)
			if err != nil {
				fmt.Fprintln(os.Stdout, err.Error())
				break
			}

			fmt.Fprintln(os.Stdout, ttl)
			break

		default:
			fmt.Fprintf(os.Stdout, "(error) ERR unknown command '%s'\n", command)
		}
	}

	return terminal.New(addr(Host, Port), executeFn).
		Prompt(cmd.Context())
}
