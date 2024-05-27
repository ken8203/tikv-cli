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
	TTL(ctx context.Context, key []byte) (uint64, error)
	Scan(ctx context.Context, start []byte, limit int) ([]Entry, error)
	Close(ctx context.Context) error
}

func New(addrs []string, mode Mode, apiVersion APIVersion, keySpace string) (Client, error) {
	switch mode {
	case ModeTxn:
		return newTxnClient(addrs, kvrpcpb.APIVersion(apiVersion), keySpace)
	case ModeRaw:
		return newRawClient(addrs, kvrpcpb.APIVersion(apiVersion), keySpace)
	default:
		return nil, fmt.Errorf("invalid mode: %s", mode)
	}
}
