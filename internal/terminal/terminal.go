package terminal

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

type ExecuteFn func(ctx context.Context, command string, args ...string)

type Terminal struct {
	placeholder string
	executeFn   ExecuteFn
}

func New(placeholder string, executeFn ExecuteFn) *Terminal {
	return &Terminal{
		placeholder: placeholder,
		executeFn:   executeFn,
	}
}

func (t *Terminal) Prompt(ctx context.Context) error {
	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Fprint(os.Stdout, t.placeholder+"> ")
		s, err := r.ReadString('\n')
		if err != nil {
			if errors.Is(err, io.EOF) {
				fmt.Fprintln(os.Stdout, "\nbye bye")
				return nil
			}

			return err
		}

		substrings := strings.Fields(s)
		if len(substrings) == 0 {
			continue
		}

		t.executeFn(ctx, substrings[0], substrings[1:]...)
	}
}

func (t *Terminal) SetPlaceholder(placeholder string) {
	t.placeholder = placeholder
}
