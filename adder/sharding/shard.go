package sharding

import (
	"context"
	"sync"

	"github.com/ipfs-cluster/ipfs-cluster/adder"
	"github.com/ipfs-cluster/ipfs-cluster/api"
	ipld "github.com/ipfs/go-ipld-format"

	cid "github.com/ipfs/go-cid"
	rpc "github.com/libp2p/go-libp2p-gorpc"
	peer "github.com/libp2p/go-libp2p/core/peer"
)

// a shard represents a set of blocks (or bucket) which have been assigned
// a peer to be block-put and will be part of the same shard in the
// cluster DAG.
type shard struct {
	ctx             context.Context
	rpc             *rpc.Client
	allocations     []peer.ID
	pinOptions      api.PinOptions
	bs              *adder.BlockStreamer
	blocks          chan api.NodeWithMeta
	closeBlocksOnce sync.Once
	// dagNode represents a node with links and will be converted
	// to Cbor.
	dagNode     map[string]cid.Cid
	currentSize uint64
	sizeLimit   uint64
}

func newShard(globalCtx context.Context, ctx context.Context, rpc *rpc.Client, opts api.PinOptions) (*shard, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// This would mean that the empty cid is part of the shared state somehow.

// TODO (hector): get latest metrics for allocations, adjust sizeLimit
// to minimum. This can be done later.

// AddLink tries to add a new block to this shard if it's not full.
// Returns true if the block was added
func (sh *shard) AddLink(ctx context.Context, c cid.Cid, s uint64) {
	_ = "STUB: not implemented"
	return
}

// Allocations returns the peer IDs on which blocks are put for this shard.
func (sh *shard) Allocations() []peer.ID { _ = "STUB: not implemented"; return nil }

func (sh *shard) sendBlock(ctx context.Context, n ipld.Node) error {
	_ = "STUB: not implemented"
	return nil
}

// Close stops any ongoing block streaming.
func (sh *shard) Close() error { _ = "STUB: not implemented"; return nil }

// Flush completes the allocation of this shard by building a CBOR node
// and adding it to IPFS, then pinning it in cluster. It returns the Cid of the
// shard.
func (sh *shard) Flush(ctx context.Context, shardN int, prev cid.Cid) (cid.Cid, error) {
	_ = "STUB: not implemented"
	return *new(cid.Cid), nil
}

// this sets allocations as priority allocation

// use current size, not the limit
// using an indirect graph

// Size returns this shard's current size.
func (sh *shard) Size() uint64 { _ = "STUB: not implemented"; return 0 }

// Limit returns this shard's size limit.
func (sh *shard) Limit() uint64 { _ = "STUB: not implemented"; return 0 }

// LastLink returns the last added link. When finishing sharding,
// the last link of the last shard is the data root for the
// full sharded DAG (the CID that would have resulted from
// adding the content to a single IPFS daemon).
func (sh *shard) LastLink() cid.Cid { _ = "STUB: not implemented"; return *new(cid.Cid) }
