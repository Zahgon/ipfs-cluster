// Package tags implements an ipfs-cluster informer publishes user-defined
// tags as metrics.
package tags

import (
	"context"
	"sync"

	"github.com/ipfs-cluster/ipfs-cluster/api"

	logging "github.com/ipfs/go-log/v2"
	rpc "github.com/libp2p/go-libp2p-gorpc"
)

var logger = logging.Logger("tags")

// MetricName specifies the name of our metric
var MetricName = "tags"

// Informer is a simple object to implement the ipfscluster.Informer
// and Component interfaces.
type Informer struct {
	config *Config // set when created, readonly

	mu        sync.Mutex // guards access to following fields
	rpcClient *rpc.Client
}

// New returns an initialized informer using the given InformerConfig.
func New(cfg *Config) (*Informer, error) { _ = "STUB: not implemented"; return nil, nil }

// Name returns the name of this informer. Note the informer issues metrics
// with custom names.
func (tags *Informer) Name() string {
	_ = "STUB: not implemented"

	// SetClient provides us with an rpc.Client which allows
	// contacting other components in the cluster.
	return ""
}

func (tags *Informer) SetClient(c *rpc.Client) { _ = "STUB: not implemented"; return }

// Shutdown is called on cluster shutdown. We just invalidate
// any metrics from this point.
func (tags *Informer) Shutdown(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// GetMetrics returns one metric for each tag defined in the configuration.
// The metric name is set as "tags:<tag_name>". When no tags are defined,
// a single invalid metric is returned.
func (tags *Informer) GetMetrics(ctx context.Context) []api.Metric {
	_ = "STUB: not implemented"
	// Note we could potentially extend the tag:value syntax to include manual weights
	// ie: { "region": "us:100", ... }
	// This would potentially allow to always give priority to peers of a certain group
	return nil
}
