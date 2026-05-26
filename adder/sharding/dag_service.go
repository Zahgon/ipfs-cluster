// Package sharding implements a sharding ClusterDAGService places
// content in different shards while it's being added, creating
// a final Cluster DAG and pinning it.
package sharding

import (
	"context"

	"time"

	"github.com/ipfs-cluster/ipfs-cluster/adder"
	"github.com/ipfs-cluster/ipfs-cluster/api"

	cid "github.com/ipfs/go-cid"
	ipld "github.com/ipfs/go-ipld-format"
	logging "github.com/ipfs/go-log/v2"
	rpc "github.com/libp2p/go-libp2p-gorpc"
	peer "github.com/libp2p/go-libp2p/core/peer"
)

var logger = logging.Logger("shardingdags")

// DAGService is an implementation of a ClusterDAGService which
// shards content while adding among several IPFS Cluster peers,
// creating a Cluster DAG to track and pin that content selectively
// in the IPFS daemons allocated to it.
type DAGService struct {
	adder.BaseDAGService

	ctx       context.Context
	rpcClient *rpc.Client

	addParams api.AddParams
	output    chan<- api.AddedOutput

	addedSet *cid.Set

	// Current shard being built
	currentShard *shard
	// Last flushed shard CID
	previousShard cid.Cid

	// shard tracking
	shards map[string]cid.Cid

	startTime time.Time
	totalSize uint64
}

// New returns a new ClusterDAGService, which uses the given rpc client to perform
// Allocate, IPFSStream and Pin requests to other cluster components.
func New(ctx context.Context, rpc *rpc.Client, opts api.AddParams, out chan<- api.AddedOutput) *DAGService {
	_ = "STUB: not implemented"
	// use a default value for this regardless of what is provided.
	return nil
}

// Add puts the given node in its corresponding shard and sends it to the
// destination peers.
func (dgs *DAGService) Add(ctx context.Context, node ipld.Node) error {
	_ = "STUB: not implemented"
	// FIXME: This will grow in memory
	return nil
}

// Close performs cleanup and should be called when the DAGService is not
// going to be used anymore.
func (dgs *DAGService) Close() error { _ = "STUB: not implemented"; return nil }

// Finalize finishes sharding, creates the cluster DAG and pins it along
// with the meta pin for the root node of the content.
func (dgs *DAGService) Finalize(ctx context.Context, dataRoot api.Cid) (api.Cid, error) {
	_ = "STUB: not implemented"
	return *new(api.Cid), nil
}

// PutDAG to ourselves

//abort

// Stream these blocks and wait until we are done.

// Pin the ClusterDAG

// pin direct

// Update object with response.

// Pin the META pin

// irrelevant. Meta-pins are not pinned

// Log some stats

// Consider doing this? Seems like overkill
//
// // Amend ShardPins to reference clusterDAG root hash as a Parent
// shardParents := cid.NewSet()
// shardParents.Add(clusterDAG)
// for shardN, shard := range dgs.shardNodes {
// 	pin := api.PinWithOpts(shard, dgs.addParams)
// 	pin.Name := fmt.Sprintf("%s-shard-%s", pin.Name, shardN)
// 	pin.Type = api.ShardType
// 	pin.Parents = shardParents
// 	// FIXME: We don't know anymore the shard pin maxDepth
//      // so we'd need to get the pin first.
// 	err := dgs.pin(pin)
// 	if err != nil {
// 		return err
// 	}
// }

// Allocations returns the current allocations for the current shard.
func (dgs *DAGService) Allocations() []peer.ID {
	_ = "STUB: not implemented"
	// FIXME: this is probably not safe in concurrency?  However, there is
	// no concurrent execution of any code in the DAGService I think.
	return nil
}

// ingests a block to the current shard. If it get's full, it
// Flushes the shard and retries with a new one.
func (dgs *DAGService) ingestBlock(ctx context.Context, n ipld.Node) error {
	_ = "STUB: not implemented"
	return nil

	// if we have no currentShard, create one
}

// important: shards use the DAGService context.

// this is not same as n.Size()

// add the block to it if it fits and return

// -------
// Below: block DOES NOT fit in shard
// Flush and retry

// if shard is empty, error

// <-- retry ingest

func (dgs *DAGService) logStats(metaPin, clusterDAGPin api.Cid) { _ = "STUB: not implemented"; return }

func (dgs *DAGService) sendOutput(ao api.AddedOutput) { _ = "STUB: not implemented"; return }

// flushes the dgs.currentShard and returns the LastLink()
func (dgs *DAGService) flushCurrentShard(ctx context.Context) (cid.Cid, error) {
	_ = "STUB: not implemented"
	return *new(cid.Cid), nil
}

// AddMany calls Add for every given node.
func (dgs *DAGService) AddMany(ctx context.Context, nodes []ipld.Node) error {
	_ = "STUB: not implemented"
	return nil
}
