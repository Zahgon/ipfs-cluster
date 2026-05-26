// Package numpin implements an ipfs-cluster informer which determines how many
// items this peer is pinning and returns it as api.Metric
package numpin

import (
	"context"
	"sync"

	"github.com/ipfs-cluster/ipfs-cluster/api"

	rpc "github.com/libp2p/go-libp2p-gorpc"
)

// MetricName specifies the name of our metric
var MetricName = "numpin"

// Informer is a simple object to implement the ipfscluster.Informer
// and Component interfaces
type Informer struct {
	config *Config

	mu        sync.Mutex
	rpcClient *rpc.Client
}

// NewInformer returns an initialized Informer.
func NewInformer(cfg *Config) (*Informer, error) { _ = "STUB: not implemented"; return nil, nil }

// SetClient provides us with an rpc.Client which allows
// contacting other components in the cluster.
func (npi *Informer) SetClient(c *rpc.Client) { _ = "STUB: not implemented"; return }

// Shutdown is called on cluster shutdown. We just invalidate
// any metrics from this point.
func (npi *Informer) Shutdown(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Name returns the name of this informer
func (npi *Informer) Name() string {
	_ = "STUB: not implemented"

	// GetMetrics contacts the IPFSConnector component and requests the `pin ls`
	// command. We return the number of pins in IPFS. It must always return at
	// least one metric.
	return ""
}

func (npi *Informer) GetMetrics(ctx context.Context) []api.Metric {
	_ = "STUB: not implemented"
	return nil
}

// make use of the RPC API to obtain information
// about the number of pins in IPFS. See RPCAPI docs.

// Local call
// Service name
// Method name
