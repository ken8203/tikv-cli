package cmd

import "errors"

var (
	ErrNotExist            = errors.New("(nil)")
	ErrInvalidArgs         = errors.New("(error) ERR wrong number of arguments for command")
	ErrCommandNotSupported = errors.New("(error) ERR txn mode doesn't support command")
)
