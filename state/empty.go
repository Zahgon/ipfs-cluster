package state

import (
	"context"

	"github.com/ipfs-cluster/ipfs-cluster/api"
)

type empty struct{}

func (e *empty) List(ctx context.Context, out chan<- api.Pin) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *empty) Has(ctx context.Context, c api.Cid) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *empty) Get(ctx context.Context, c api.Cid) (api.Pin, error) {
	_ = "STUB: not implemented"
	return *new(api.Pin), nil
}

// Empty returns an empty read-only state.
func Empty() ReadOnly { _ = "STUB: not implemented"; return *new(ReadOnly) }
