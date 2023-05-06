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
func shellRunE(cmd *cobra.Command, args []string) error {
	client, err := newClient()
	if err != nil {
		return fmt.Errorf("new client: %v", err)
	}
	defer client.Close(cmd.Context())

	executeFn := func(ctx context.Context, command string) {
		fields := strings.Fields(command)
		if len(fields) == 0 {
			return
		}

		switch strings.ToLower(fields[0]) {
		case "put":
			if len(fields) != 3 {
				fmt.Fprintln(os.Stdout, "(error) ERR wrong number of arguments for 'put' command")
				break
			}

			if err := client.Put(ctx, []byte(fields[1]), []byte(fields[2])); err != nil {
				fmt.Fprintln(os.Stdout, err.Error())
			}
		case "get":
			if len(fields) != 2 {
				fmt.Fprintln(os.Stdout, "(error) ERR wrong number of arguments for 'get' command")
				break
			}

			value, err := client.Get(ctx, []byte(fields[1]))
			if err != nil {
				if errors.Is(err, tikverror.ErrNotExist) {
					fmt.Fprintln(os.Stdout, "(nil)")
					break
				}
				fmt.Fprintln(os.Stdout, err.Error())
			}

			fmt.Fprintln(os.Stdout, string(value))
		case "delete":
			if len(fields) != 2 {
				fmt.Fprintln(os.Stdout, "(error) ERR wrong number of arguments for 'delete' command")
				break
			}

			if err := client.Delete(ctx, []byte(fields[1])); err != nil {
				fmt.Fprintln(os.Stdout, err.Error())
			}
		}
	}

	return terminal.New(addr(Host, Port), executeFn).
		Prompt(cmd.Context())
}
