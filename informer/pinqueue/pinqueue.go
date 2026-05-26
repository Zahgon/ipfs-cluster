// Package pinqueue implements an ipfs-cluster informer which issues the
// current size of the pinning queue.
package pinqueue

import (
	"context"
	"sync"

	"github.com/ipfs-cluster/ipfs-cluster/api"

	rpc "github.com/libp2p/go-libp2p-gorpc"
)

// MetricName specifies the name of our metric
var MetricName = "pinqueue"

// Informer is a simple object to implement the ipfscluster.Informer
// and Component interfaces
type Informer struct {
	config *Config

	mu        sync.Mutex
	rpcClient *rpc.Client
}

// New returns an initialized Informer.
func New(cfg *Config) (*Informer, error) { _ = "STUB: not implemented"; return nil, nil }

// SetClient provides us with an rpc.Client which allows
// contacting other components in the cluster.
func (inf *Informer) SetClient(c *rpc.Client) { _ = "STUB: not implemented"; return }

// Shutdown is called on cluster shutdown. We just invalidate
// any metrics from this point.
func (inf *Informer) Shutdown(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Name returns the name of this informer
func (inf *Informer) Name() string {
	_ = "STUB: not implemented"

	// GetMetrics contacts the Pintracker component and requests the number of
	// queued items for pinning.
	return ""
}

func (inf *Informer) GetMetrics(ctx context.Context) []api.Metric {
	_ = "STUB: not implemented"
	return nil
}

// smaller pin queues have more priority
