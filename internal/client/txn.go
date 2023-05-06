package client

import (
	"context"

	"github.com/pingcap/kvproto/pkg/kvrpcpb"
	"github.com/tikv/client-go/v2/txnkv"
)

type txnClient struct {
	client *txnkv.Client
}

var _ Client = (*txnClient)(nil)

func newTxnClient(addrs []string, apiVersion kvrpcpb.APIVersion) (*txnClient, error) {
	client, err := txnkv.NewClient(addrs, txnkv.WithAPIVersion(apiVersion))
	if err != nil {
		return nil, err
	}

	return &txnClient{
		client: client,
	}, nil
}

func (c *txnClient) Put(ctx context.Context, key, value []byte) error {
	tx, err := c.client.Begin()
	if err != nil {
		return err
	}

	if err := tx.Set(key, value); err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func (c *txnClient) Get(ctx context.Context, key []byte) ([]byte, error) {
	tx, err := c.client.Begin()
	if err != nil {
		return nil, err
	}

	return tx.Get(ctx, key)
}

func (c *txnClient) Delete(ctx context.Context, key []byte) error {
	tx, err := c.client.Begin()
	if err != nil {
		return err
	}

	if err := tx.Delete(key); err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func (c *txnClient) Close(ctx context.Context) error {
	return c.client.Close()
}
