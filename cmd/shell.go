package cmd

import (
	"fmt"
	"strings"

	"github.com/ken8203/tikv-cli/internal/terminal"
	"github.com/spf13/cobra"
)

// shellRunE is the entry of shell command.
func shellRunE(cmd *cobra.Command, _ []string) error {
	ctx := cmd.Context()
	term := terminal.New(addr(Host, Port))

	ch, errCh := term.Enter(ctx)
	for command := range ch {
		switch strings.ToLower(command.Cmd) {
		case putCmd.Name():
			if err := put(ctx, command.Args); err != nil {
				fmt.Fprintln(term, err.Error())
				break
			}

			fmt.Fprintln(term, "OK")
			break

		case getCmd.Name():
			value, err := get(ctx, command.Args)
			if err != nil {
				fmt.Fprintln(term, err.Error())
				break
			}

			fmt.Fprintln(term, string(value))
			break

		case deleteCmd.Name():
			if err := delete(ctx, command.Args); err != nil {
				fmt.Fprintln(term, err.Error())
				break
			}

			fmt.Fprintln(term, "OK")
			break

		case ttlCmd.Name():
			ttl, err := ttl(ctx, command.Args)
			if err != nil {
				fmt.Fprintln(term, err.Error())
				break
			}

			fmt.Fprintln(term, ttl)
			break

		default:
			fmt.Fprintf(term, "(error) ERR unknown command '%s'\n", command)
		}
	}

	return <-errCh
}
