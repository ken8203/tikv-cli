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
			if err := put(ctx, term, command.Args); err != nil {
				fmt.Fprintln(term, err.Error())
				break
			}

			break

		case getCmd.Name():
			if err := get(ctx, term, command.Args); err != nil {
				fmt.Fprintln(term, err.Error())
				break
			}

			break

		case deleteCmd.Name():
			if err := delete(ctx, term, command.Args); err != nil {
				fmt.Fprintln(term, err.Error())
				break
			}

			break

		case ttlCmd.Name():
			if err := ttl(ctx, term, command.Args); err != nil {
				fmt.Fprintln(term, err.Error())
				break
			}

			break

		case versionCmd.Name():
			if err := version(term); err != nil {
				fmt.Fprintln(term, err.Error())
				break
			}

			break

		case scanCmd.Name():
			if err := scan(ctx, term, command.Args); err != nil {
				fmt.Fprintln(term, err.Error())
				break
			}

			break

		default:
			fmt.Fprintf(term, "(error) ERR unknown command '%s'\n", command)
		}
	}

	return <-errCh
}
