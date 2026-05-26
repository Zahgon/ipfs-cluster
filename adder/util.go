package adder

import (
	"context"
	"errors"
	"sync"

	"github.com/ipfs-cluster/ipfs-cluster/api"

	cid "github.com/ipfs/go-cid"
	ipld "github.com/ipfs/go-ipld-format"
	rpc "github.com/libp2p/go-libp2p-gorpc"
	peer "github.com/libp2p/go-libp2p/core/peer"
)

// ErrBlockAdder is returned when adding a to multiple destinations
// block fails on all of them.
var ErrBlockAdder = errors.New("failed to put block on all destinations")

// BlockStreamer helps streaming nodes to multiple destinations, as long as
// one of them is still working.
type BlockStreamer struct {
	dests     []peer.ID
	rpcClient *rpc.Client
	blocks    <-chan api.NodeWithMeta

	ctx    context.Context
	cancel context.CancelFunc
	errMu  sync.Mutex
	err    error
}

// NewBlockStreamer creates a BlockStreamer given an rpc client, allocated
// peers and a channel on which the blocks to stream are received.
func NewBlockStreamer(ctx context.Context, rpcClient *rpc.Client, dests []peer.ID, blocks <-chan api.NodeWithMeta) *BlockStreamer {
	_ = "STUB: not implemented"
	return nil
}

// Done returns a channel which gets closed when the BlockStreamer has
// finished.
func (bs *BlockStreamer) Done() <-chan struct{} { _ = "STUB: not implemented"; return nil }

func (bs *BlockStreamer) setErr(err error) { _ = "STUB: not implemented"; return }

// Err returns any errors that happened after the operation of the
// BlockStreamer, for example when blocks could not be put to all nodes.
func (bs *BlockStreamer) Err() error { _ = "STUB: not implemented"; return nil }

func (bs *BlockStreamer) streamBlocks() {
	_ = "STUB: not implemented"

	// Nothing should be sent on out.
	// We drain though
	return
}

// FIXME: replicate everywhere.

// IpldNodeToNodeWithMeta converts an ipld.Node to api.NodeWithMeta.
func IpldNodeToNodeWithMeta(n ipld.Node) api.NodeWithMeta {
	_ = "STUB: not implemented"
	return *new(api.NodeWithMeta)
}

// BlockAllocate helps allocating blocks to peers.
func BlockAllocate(ctx context.Context, rpc *rpc.Client, pinOpts api.PinOptions) ([]peer.ID, error) {
	_ = "STUB: not implemented"
	// Find where to allocate this file
	return nil, nil
}

// Pin helps sending local RPC pin requests.
func Pin(ctx context.Context, rpc *rpc.Client, pin api.Pin) error {
	_ = "STUB: not implemented"
	return nil
}

// use ourself to pin

// ErrDAGNotFound is returned whenever we try to get a block from the DAGService.
var ErrDAGNotFound = errors.New("dagservice: a Get operation was attempted while cluster-adding (this is likely a bug)")

// BaseDAGService partially implements an ipld.DAGService.
// It provides the methods which are not needed by ClusterDAGServices
// (Get*, Remove*) so that they can save adding this code.
type BaseDAGService struct {
}

// Get always returns errNotFound
func (dag BaseDAGService) Get(ctx context.Context, key cid.Cid) (ipld.Node, error) {
	_ = "STUB: not implemented"
	return *new(ipld.Node), nil
}

// GetMany returns an output channel that always emits an error
func (dag BaseDAGService) GetMany(ctx context.Context, keys []cid.Cid) <-chan *ipld.NodeOption {
	_ = "STUB: not implemented"
	return nil
}

// Remove is a nop
func (dag BaseDAGService) Remove(ctx context.Context, key cid.Cid) error {
	_ = "STUB: not implemented"

	// RemoveMany is a nop
	return nil
}

func (dag BaseDAGService) RemoveMany(ctx context.Context, keys []cid.Cid) error {
	_ = "STUB: not implemented"
	return nil
}
