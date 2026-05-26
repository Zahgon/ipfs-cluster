// Package disk implements an ipfs-cluster informer which can provide different
// disk-related metrics from the IPFS daemon as an api.Metric.
package disk

import (
	"context"
	"sync"

	"github.com/ipfs-cluster/ipfs-cluster/api"

	logging "github.com/ipfs/go-log/v2"
	rpc "github.com/libp2p/go-libp2p-gorpc"
)

// MetricType identifies the type of metric to fetch from the IPFS daemon.
type MetricType int

const (
	// MetricFreeSpace provides the available space reported by IPFS
	MetricFreeSpace MetricType = iota
	// MetricRepoSize provides the used space reported by IPFS
	MetricRepoSize
)

// String returns a string representation for MetricType.
func (t MetricType) String() string { _ = "STUB: not implemented"; return "" }

var logger = logging.Logger("diskinfo")

// Informer is a simple object to implement the ipfscluster.Informer
// and Component interfaces.
type Informer struct {
	config *Config // set when created, readonly

	mu        sync.Mutex // guards access to following fields
	rpcClient *rpc.Client
}

// NewInformer returns an initialized informer using the given InformerConfig.
func NewInformer(cfg *Config) (*Informer, error) { _ = "STUB: not implemented"; return nil, nil }

// Name returns the name of the metric issued by this informer.
func (disk *Informer) Name() string { _ = "STUB: not implemented"; return "" }

// SetClient provides us with an rpc.Client which allows
// contacting other components in the cluster.
func (disk *Informer) SetClient(c *rpc.Client) { _ = "STUB: not implemented"; return }

// Shutdown is called on cluster shutdown. We just invalidate
// any metrics from this point.
func (disk *Informer) Shutdown(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// GetMetrics returns the metric obtained by this Informer. It must always
// return at least one metric.
func (disk *Informer) GetMetrics(ctx context.Context) []api.Metric {
	_ = "STUB: not implemented"
	return nil
}

// Make sure we don't underflow and stop
// sending this metric when space is exhausted.

// smaller repositories have more priority
