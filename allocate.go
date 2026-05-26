package ipfscluster

import (
	"context"

	peer "github.com/libp2p/go-libp2p/core/peer"

	"github.com/ipfs-cluster/ipfs-cluster/api"
)

// This file gathers allocation logic used when pinning or re-pinning
// to find which peers should be allocated to a Cid. Allocation is constrained
// by ReplicationFactorMin and ReplicationFactorMax parameters obtained
// from the Pin object.

// The allocation process has several steps:
//
// * Find which peers are pinning a CID
// * Obtain the last values for the configured informer metrics from the
//   monitor component
// * Divide the metrics between "current" (peers already pinning the CID)
//   and "candidates" (peers that could pin the CID), as long as their metrics
//   are valid.
// * Given the candidates:
//   * Check if we are overpinning an item
//   * Check if there are not enough candidates for the "needed" replication
//     factor.
//   * If there are enough candidates:
//     * Call the configured allocator, which sorts the candidates (and
//       may veto some depending on the allocation strategy.
//     * The allocator returns a list of final candidate peers sorted by
//       order of preference.
//     * Take as many final candidates from the list as we can, until
//       ReplicationFactorMax is reached. Error if there are less than
//       ReplicationFactorMin.

// A wrapper to carry peer metrics that have been classified.
type classifiedMetrics struct {
	current        api.MetricsSet
	currentPeers   []peer.ID
	candidate      api.MetricsSet
	candidatePeers []peer.ID
	priority       api.MetricsSet
	priorityPeers  []peer.ID
}

// allocate finds peers to allocate a hash using the informer and the monitor
// it should only be used with valid replicationFactors (if rplMin and rplMax
// are > 0, then rplMin <= rplMax).
// It always returns allocations, but if no new allocations are needed,
// it will return the current ones. Note that allocate() does not take
// into account if the given CID was previously in a "pin everywhere" mode,
// and will consider such Pins as currently unallocated ones, providing
// new allocations as available.
func (c *Cluster) allocate(ctx context.Context, hash api.Cid, currentPin api.Pin, rplMin, rplMax int, blacklist []peer.ID, priorityList []peer.ID) ([]peer.ID, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// allocate everywhere

// Figure out who is holding the CID

// Get Metrics that the allocator is interested on

// Filter and divide metrics.  The resulting sets only have peers that
// have all the metrics needed and are not blacklisted.

// if current allocations are above the minimal threshold,
// obtainAllocations returns nil and we just leave things as they are.
// This is what makes repinning do nothing if items are still above
// rmin.

// Given metrics from all informers, split them into 3 MetricsSet:
// - Those corresponding to currently allocated peers
// - Those corresponding to priority allocations
// - Those corresponding to "candidate" allocations
// And return also an slice of the peers in those groups.
//
// Peers from untrusted peers are left out if configured.
//
// For a metric/peer to be included in a group, it is necessary that it has
// metrics for all informers.
func (c *Cluster) filterMetrics(ctx context.Context, mSet api.MetricsSet, numMetrics int, currentAllocs, priorityList, blacklist []peer.ID) classifiedMetrics {
	_ = "STUB: not implemented"
	return *new(classifiedMetrics)
}

// Divide the metric by current/candidate/prio and by peer

// discard blacklisted peers

// discard peers that are not trusted

// discard peers that are trusted

// Put the metrics in their sets if peers have metrics for all
// informers Record peers. This relies on LatestMetrics
// returning exactly one metric per peer. Thus, a peer with
// all the needed metrics should have exactly numMetrics.
// Otherwise, they are ignored.

// otherwise this peer will be ignored.

// allocationError logs an allocation error
func allocationError(hash api.Cid, needed, wanted int, candidatesValid []peer.ID) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Cluster) obtainAllocations(
	ctx context.Context,
	hash api.Cid,
	rplMin, rplMax int,
	metrics classifiedMetrics,
) ([]peer.ID, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// The minimum we need
// The maximum we want

// Reminder: rplMin <= rplMax AND >0

// allocations above maximum threshold: drop some
// This could be done more intelligently by dropping them
// according to the allocator order (i.e. free-ing peers
// with most used space first).

// allocations are above minimal threshold
// We don't provide any new allocations

// not enough candidates

// We can allocate from this point. Use the allocator to decide
// on the priority of candidates grab as many as "wanted"

// the allocator returns a list of peers ordered by priority

// check that we have enough as the allocator may have returned
// less candidates than provided.

// the final result is the currently valid allocations
// along with the ones provided by the allocator
