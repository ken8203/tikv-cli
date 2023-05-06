package terminal

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"
	"unicode"
)

type ExecuteFn func(ctx context.Context, command string)

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
			return err
		}

		t.executeFn(ctx, TrimSpace(s))
	}
}

func (t *Terminal) SetPlaceholder(placeholder string) {
	t.placeholder = placeholder
}

func TrimSpace(s string) string {
	return strings.TrimFunc(s, func(r rune) bool {
		return unicode.IsSpace(r) ||
			r == '\u200B' || // zero-width space
			r == '\u200C' || // zero-width non-joiner
			r == '\u200D' || // zero-width joiner
			r == '\uFEFF' // zero-width no-break space
	})
}
