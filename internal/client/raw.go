package client

import (
	"context"

	"github.com/pingcap/kvproto/pkg/kvrpcpb"
	"github.com/tikv/client-go/v2/rawkv"
)

type rawClient struct {
	client *rawkv.Client
}

var _ Client = (*rawClient)(nil)

func newRawClient(addrs []string, apiVersion kvrpcpb.APIVersion) (*rawClient, error) {
	client, err := rawkv.NewClientWithOpts(context.Background(), addrs, rawkv.WithAPIVersion(apiVersion))
	if err != nil {
		return nil, err
	}

	return &rawClient{
		client: client,
	}, nil
}

func (c *rawClient) Put(ctx context.Context, key, value []byte) error {
	return c.client.Put(ctx, key, value)
}

func (c *rawClient) Get(ctx context.Context, key []byte) ([]byte, error) {
	return c.client.Get(ctx, key)
}

func (c *rawClient) Delete(ctx context.Context, key []byte) error {
	return c.client.Delete(ctx, key)
}

func (c *rawClient) TTL(ctx context.Context, key []byte) (uint64, error) {
	ttl, err := c.client.GetKeyTTL(ctx, key)
	if err != nil {
		return 0, err
	}

	if ttl != nil {
		return *ttl, nil
	}
	return 0, nil
}

func (c *rawClient) Scan(ctx context.Context, start []byte, limit int) ([]Entry, error) {
	keys, values, err := c.client.Scan(ctx, start, nil, limit)
	if err != nil {
		return nil, err
	}

	entries := make([]Entry, 0, len(keys))
	for i, key := range keys {
		entries[i] = Entry{
			K: key,
			V: values[i],
		}
	}

	return entries, nil
}

func (c *rawClient) Close(ctx context.Context) error {
	return c.client.Close()
}
