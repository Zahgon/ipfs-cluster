// Package balanced implements an allocator that can sort allocations
// based on multiple metrics, where metrics may be an arbitrary way to
// partition a set of peers.
//
// For example, allocating by ["tag:region", "disk"] the resulting peer
// candidate order will balanced between regions and ordered by the value of
// the weight of the disk metric.
package balanced

import (
	"context"

	api "github.com/ipfs-cluster/ipfs-cluster/api"
	logging "github.com/ipfs/go-log/v2"
	rpc "github.com/libp2p/go-libp2p-gorpc"
	peer "github.com/libp2p/go-libp2p/core/peer"
)

var logger = logging.Logger("allocator")

// Allocator is an allocator that partitions metrics and orders
// the final list of allocation by selecting for each partition.
type Allocator struct {
	config    *Config
	rpcClient *rpc.Client
}

// New returns an initialized Allocator.
func New(cfg *Config) (*Allocator, error) { _ = "STUB: not implemented"; return nil, nil }

// SetClient provides us with an rpc.Client which allows
// contacting other components in the cluster.
func (a *Allocator) SetClient(c *rpc.Client) {
	_ = "STUB: not implemented"

	// Shutdown is called on cluster shutdown. We just invalidate
	// any metrics from this point.
	return
}

func (a *Allocator) Shutdown(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

type partitionedMetric struct {
	metricName       string
	curChoosingIndex int
	noMore           bool
	partitions       []*partition // they are in order of their values
}

type partition struct {
	value            string
	weight           int64
	aggregatedWeight int64
	peers            map[peer.ID]bool   // the bool tracks whether the peer has been picked already out of the partition when doing the final sort.
	sub              *partitionedMetric // all peers in sub-partitions will have the same value for this metric
}

// Returns a partitionedMetric which has partitions and subpartitions based
// on the metrics and values given by the "by" slice. The partitions
// are ordered based on the cumulative weight.
func partitionMetrics(set api.MetricsSet, by []string) *partitionedMetric {
	_ = "STUB: not implemented"
	return nil
}

// For sorting based on weight (more to less)

// if weight is equal, sort by aggregated weight of
// all sub-partitions.

// If subpartitions weight the same, do strict order
// based on value string

// Descending!

// we are done

// process sub-partitions

// not needed anymore

// only leave metrics for peers in current partition

// Add the aggregated weight of the subpartitions

func partitionValues(metrics []api.Metric) []*partition { _ = "STUB: not implemented"; return nil }

// We group peers with the same value in the same partition.

// Sometimes two metrics have the same value / weight, but we
// still want to put them in different partitions. Otherwise
// their weights get added and they form a bucket and
// therefore not they are not selected in order: 3 peers with
// freespace=100 and one peer with freespace=200 would result
// in one of the peers with freespace 100 being chosen first
// because the partition's weight is 300.
//
// We are going to call these metrics (like free-space),
// non-partitionable metrics. This is going to be the default
// (for backwards compat reasons).
//
// The informers must set the Partitionable field accordingly
// when two metrics with the same value must be grouped in the
// same partition.
//
// Note: aggregatedWeight is the same as weight here (sum of
// weight of all metrics in partitions), and gets updated
// later in partitionMetrics with the aggregated weight of
// sub-partitions.

// Any other case, we partition by value.

// Returns a list of peers sorted by never choosing twice from the same
// partition if there is some other partition to choose from.
func (pnedm *partitionedMetric) sortedPeers() []peer.ID { _ = "STUB: not implemented"; return nil }

// This means we are done.

func (pnedm *partitionedMetric) chooseNext() peer.ID {
	_ = "STUB: not implemented"
	return *new(peer.ID)
}

// Choose something from the sub-partitionedMetric

// We are a bottom-partition. Choose one of our peers

// mark as used

// look in next partition next time

// no peer and we have looked in as many partitions as we have

// Allocate produces a sorted list of cluster peer IDs based on different
// metrics provided for those peer IDs.
// It works as follows:
//
//   - First, it buckets each peer metrics based on the AllocateBy list. The
//     metric name must match the bucket name, otherwise they are put at the end.
//   - Second, based on the AllocateBy order, it orders the first bucket and
//     groups peers by ordered value.
//   - Third, it selects metrics on the second bucket for the most prioritary
//     peers of the first bucket and orders their metrics. Then for the peers in
//     second position etc.
//   - It repeats the process until there is no more buckets to sort.
//   - Finally, it returns the first peer of the first
//   - Third, based on the AllocateBy order, it select the first metric
func (a *Allocator) Allocate(
	ctx context.Context,
	c api.Cid,
	current, candidates, priority api.MetricsSet,
) ([]peer.ID, error) {
	_ = "STUB: not implemented"

	// For the allocation to work well, there have to be metrics of all
	// the types for all the peers. There cannot be a metric of one type
	// for a peer that does not appear in the other types.
	//
	// Removing such occurrences is done in allocate.go, before the
	// allocator is called.
	//
	// Otherwise, the sorting might be funny.
	return nil, nil
}

//fmt.Println(printPartition(candidatePartition, 0))

// Metrics returns the names of the metrics that have been registered
// with this allocator.
func (a *Allocator) Metrics() []string { _ = "STUB: not implemented"; return nil }

func printPartition(m *partitionedMetric, ind int) string { _ = "STUB: not implemented"; return "" }
