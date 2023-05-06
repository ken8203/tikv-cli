package client

import (
	"context"
	"fmt"

	"github.com/pingcap/kvproto/pkg/kvrpcpb"
)

type Mode string

const (
	ModeTxn Mode = "txn"
	ModeRaw Mode = "raw"
)

type APIVersion int32

const (
	APIVersion1 APIVersion = iota
	APIVersion1TTL
	APIVersion2
)

type Client interface {
	Put(ctx context.Context, key, value []byte) error
	Get(ctx context.Context, key []byte) ([]byte, error)
	Delete(ctx context.Context, key []byte) error
	Close(ctx context.Context) error
}

func New(addrs []string, mode Mode, apiVersion APIVersion) (Client, error) {
	switch mode {
	case ModeTxn:
		return newTxnClient(addrs, kvrpcpb.APIVersion(apiVersion))
	case ModeRaw:
		return newRawClient(addrs, kvrpcpb.APIVersion(apiVersion))
	default:
		return nil, fmt.Errorf("invalid mode: %s", mode)
	}
}
