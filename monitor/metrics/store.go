package metrics

import (
	"sync"

	"github.com/ipfs-cluster/ipfs-cluster/api"

	peer "github.com/libp2p/go-libp2p/core/peer"
)

// PeerMetrics maps a peer IDs to a metrics window.
type PeerMetrics map[peer.ID]*Window

// Store can be used to store and access metrics.
type Store struct {
	mux    sync.RWMutex
	byName map[string]PeerMetrics
}

// NewStore can be used to create a Store.
func NewStore() *Store { _ = "STUB: not implemented"; return nil }

// Add inserts a new metric in Metrics.
func (mtrs *Store) Add(m api.Metric) { _ = "STUB: not implemented"; return }

// We always lock the outer map, so we can use unsafe
// Window.

// RemovePeer removes all metrics related to a peer from the Store.
func (mtrs *Store) RemovePeer(pid peer.ID) { _ = "STUB: not implemented"; return }

// RemovePeerMetrics removes all metrics of a given name for a given peer ID.
func (mtrs *Store) RemovePeerMetrics(pid peer.ID, name string) { _ = "STUB: not implemented"; return }

// LatestValid returns all the last known valid metrics of a given type. A metric
// is valid if it has not expired.
func (mtrs *Store) LatestValid(name string) []api.Metric { _ = "STUB: not implemented"; return nil }

// TODO(ajl): for accrual, does it matter if a ping has expired?

// AllMetrics returns the latest metrics for all peers and metrics types.  It
// may return expired metrics.
func (mtrs *Store) AllMetrics() []api.Metric { _ = "STUB: not implemented"; return nil }

// PeerMetrics returns the latest metrics for a given peer ID for
// all known metrics types. It may return expired metrics.
func (mtrs *Store) PeerMetrics(pid peer.ID) []api.Metric { _ = "STUB: not implemented"; return nil }

// PeerMetricAll returns all of a particular metrics for a
// particular peer.
func (mtrs *Store) PeerMetricAll(name string, pid peer.ID) []api.Metric {
	_ = "STUB: not implemented"
	return nil
}

// PeerLatest returns the latest of a particular metric for a
// particular peer. It may return an expired metric.
func (mtrs *Store) PeerLatest(name string, pid peer.ID) api.Metric {
	_ = "STUB: not implemented"
	return *new(api.Metric)
}

// ignoring error, as nil metric is indicative enough

// MetricNames returns all the known metric names
func (mtrs *Store) MetricNames() []string { _ = "STUB: not implemented"; return nil }
