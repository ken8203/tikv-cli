package terminal

import (
	"context"
	"io"
	"os"
	"strings"

	"golang.org/x/term"
)

type Terminal struct {
	*term.Terminal
}

func New(prompt string) *Terminal {
	return &Terminal{
		Terminal: term.NewTerminal(os.Stdin, prompt+"> "),
	}
}

func (t *Terminal) Enter(ctx context.Context) (<-chan *Command, <-chan error) {
	errCh := make(chan error, 1)

	state, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		errCh <- err
		return nil, errCh
	}

	ch := make(chan *Command)
	go func() {
		defer term.Restore(int(os.Stdin.Fd()), state)
		defer func() {
			close(ch)
			close(errCh)
		}()

		for {
			select {
			case <-ctx.Done():
				errCh <- ctx.Err()
				return
			default:
				line, err := t.ReadLine()
				if err != nil {
					if err == io.EOF {
						return
					}

					errCh <- err
					return
				}
				if line == "" {
					continue
				}

				ch <- newCommand(line)
			}
		}
	}()

	return ch, errCh
}

func (t *Terminal) SetPrompt(prompt string) {
	t.SetPrompt(prompt + "> ")
}

type Command struct {
	Cmd  string
	Args []string
}

func newCommand(line string) *Command {
	substrings := strings.Fields(line)
	return &Command{
		Cmd:  substrings[0],
		Args: substrings[1:],
	}
}
