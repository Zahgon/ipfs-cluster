package sharding

import (
	"context"
	"testing"

	"github.com/ipfs-cluster/ipfs-cluster/api"
)

// MockPinStore is used in VerifyShards
type MockPinStore interface {
	// Gets a pin
	PinGet(context.Context, api.Cid) (api.Pin, error)
}

// MockBlockStore is used in VerifyShards
type MockBlockStore interface {
	// Gets a block
	BlockGet(context.Context, api.Cid) ([]byte, error)
}

// VerifyShards checks that a sharded CID has been correctly formed and stored.
// This is a helper function for testing. It returns a map with all the blocks
// from all shards.
func VerifyShards(t *testing.T, rootCid api.Cid, pins MockPinStore, ipfs MockBlockStore, expectedShards int) (map[string]struct{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// traverse shards in order
