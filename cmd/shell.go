package cmd

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/ken8203/tikv-cli/internal/terminal"
	"github.com/spf13/cobra"
	tikverror "github.com/tikv/client-go/v2/error"
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
			if len(args) < 2 {
				fmt.Fprintln(os.Stdout, "(error) ERR wrong number of arguments for 'PUT' command")
				break
			}

			if err := client.Put(ctx, []byte(args[0]), []byte(args[1])); err != nil {
				fmt.Fprintln(os.Stdout, err.Error())
			}
		case "get":
			if len(args) < 1 {
				fmt.Fprintln(os.Stdout, "(error) ERR wrong number of arguments for 'GET' command")
				break
			}

			value, err := client.Get(ctx, []byte(args[0]))
			if err != nil {
				if errors.Is(err, tikverror.ErrNotExist) {
					fmt.Fprintln(os.Stdout, "(nil)")
					break
				}
				fmt.Fprintln(os.Stdout, err.Error())
			}

			fmt.Fprintln(os.Stdout, string(value))
		case "delete":
			if len(args) < 1 {
				fmt.Fprintln(os.Stdout, "(error) ERR wrong number of arguments for 'DELETE' command")
				break
			}

			if err := client.Delete(ctx, []byte(args[0])); err != nil {
				fmt.Fprintln(os.Stdout, err.Error())
			}
		default:
			fmt.Fprintf(os.Stdout, "(error) ERR unknown command '%s'\n", command)
		}
	}

	return terminal.New(addr(Host, Port), executeFn).
		Prompt(cmd.Context())
}
