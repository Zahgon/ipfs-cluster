// Package single implements a ClusterDAGService that chunks and adds content
// to cluster without sharding, before pinning it.
package single

import (
	"context"
	"sync"

	adder "github.com/ipfs-cluster/ipfs-cluster/adder"
	"github.com/ipfs-cluster/ipfs-cluster/api"

	cid "github.com/ipfs/go-cid"
	ipld "github.com/ipfs/go-ipld-format"
	logging "github.com/ipfs/go-log/v2"
	rpc "github.com/libp2p/go-libp2p-gorpc"
	peer "github.com/libp2p/go-libp2p/core/peer"
)

var logger = logging.Logger("singledags")
var _ = logger // otherwise unused

// DAGService is an implementation of an adder.ClusterDAGService which
// puts the added blocks directly in the peers allocated to them (without
// sharding).
type DAGService struct {
	adder.BaseDAGService

	ctx       context.Context
	rpcClient *rpc.Client

	dests     []peer.ID
	addParams api.AddParams
	local     bool

	bs              *adder.BlockStreamer
	blocks          chan api.NodeWithMeta
	closeBlocksOnce sync.Once
	recentBlocks    *recentBlocks
}

// New returns a new Adder with the given rpc Client. The client is used
// to perform calls to IPFS.BlockStream and Pin content on Cluster.
func New(ctx context.Context, rpc *rpc.Client, opts api.AddParams, local bool) *DAGService {
	_ = "STUB: not implemented"
	// ensure don't Add something and pin it in direct mode.
	return nil
}

// Add puts the given node in the destination peers.
func (dgs *DAGService) Add(ctx context.Context, node ipld.Node) error {
	_ = "STUB: not implemented"
	// Avoid adding the same node multiple times in a row.
	// This is done by the ipfsadd-er, because some nodes are added
	// via dagbuilder, then via MFS, and root nodes once more.
	return nil
}

// FIXME: can't this happen on initialization?  Perhaps the point here
// is the adder only allocates and starts streaming when the first
// block arrives and not on creation.

// ensure our allocs do not carry an empty peer
// mostly an issue with testing mocks

// If this is a local pin, make sure that the local
// peer is among the allocations..
// UNLESS user-allocations are defined!

// replace last allocation with local peer

// Close cleans up the DAGService.
func (dgs *DAGService) Close() error { _ = "STUB: not implemented"; return nil }

// Finalize pins the last Cid added to this DAGService.
func (dgs *DAGService) Finalize(ctx context.Context, root api.Cid) (api.Cid, error) {
	_ = "STUB: not implemented"
	// Close the blocks channel
	return *new(api.Cid), nil
}

// Wait for the BlockStreamer to finish.

// If the streamer failed to put blocks.

// Do not pin, just block put.
// Why? Because some people are uploading CAR files with partial DAGs
// and ideally they should be pinning only when the last partial CAR
// is uploaded. This gives them that option.

// Cluster pin the result

// Allocations returns the add destinations decided by the DAGService.
func (dgs *DAGService) Allocations() []peer.ID {
	_ = "STUB: not implemented"
	// using rpc clients without a host results in an empty peer
	// which cannot be parsed to peer.ID on deserialization.
	return nil
}

// AddMany calls Add for every given node.
func (dgs *DAGService) AddMany(ctx context.Context, nodes []ipld.Node) error {
	_ = "STUB: not implemented"
	return nil
}

type recentBlocks struct {
	blocks [2]cid.Cid
	cur    int
}

func (rc *recentBlocks) Add(n ipld.Node) { _ = "STUB: not implemented"; return }

func (rc *recentBlocks) Has(n ipld.Node) bool { _ = "STUB: not implemented"; return false }
