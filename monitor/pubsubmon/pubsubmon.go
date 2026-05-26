// Package pubsubmon implements a PeerMonitor component for IPFS Cluster that
// uses PubSub to send and receive metrics.
package pubsubmon

import (
	"context"

	"sync"

	"github.com/ipfs-cluster/ipfs-cluster/api"
	"github.com/ipfs-cluster/ipfs-cluster/monitor/metrics"

	logging "github.com/ipfs/go-log/v2"
	rpc "github.com/libp2p/go-libp2p-gorpc"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	peer "github.com/libp2p/go-libp2p/core/peer"
	gocodec "github.com/ugorji/go/codec"
)

var logger = logging.Logger("monitor")

// PubsubTopic specifies the topic used to publish Cluster metrics.
var PubsubTopic = "monitor.metrics"

var msgpackHandle = &gocodec.MsgpackHandle{}

// Monitor is a component in charge of monitoring peers, logging
// metrics and detecting failures
type Monitor struct {
	ctx       context.Context
	cancel    func()
	rpcClient *rpc.Client
	rpcReady  chan struct{}

	pubsub       *pubsub.PubSub
	topic        *pubsub.Topic
	subscription *pubsub.Subscription
	peers        PeersFunc

	metrics *metrics.Store
	checker *metrics.Checker

	config *Config

	shutdownLock sync.Mutex
	shutdown     bool
	wg           sync.WaitGroup
}

// PeersFunc allows the Monitor to filter and discard metrics
// that do not belong to a given peerset.
type PeersFunc func(context.Context) ([]peer.ID, error)

// New creates a new PubSub monitor, using the given host, config and
// PeersFunc. The PeersFunc can be nil. In this case, no metric filtering is
// done based on peers (any peer is considered part of the peerset).
func New(
	ctx context.Context,
	cfg *Config,
	psub *pubsub.PubSub,
	peers PeersFunc,
) (*Monitor, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (mon *Monitor) run() { _ = "STUB: not implemented"; return }

// logFromPubsub logs metrics received in the subscribed topic.
func (mon *Monitor) logFromPubsub() { _ = "STUB: not implemented"; return }

// Previous versions use multicodec with the following header, which
// we need to remove.

// context canceled enters here

// managed to decode an older version metric. Warn about it once.

// SetClient saves the given rpc.Client  for later use
func (mon *Monitor) SetClient(c *rpc.Client) { _ = "STUB: not implemented"; return }

// Shutdown stops the peer monitor. It particular, it will
// not deliver any alerts.
func (mon *Monitor) Shutdown(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// LogMetric stores a metric so it can later be retrieved.
func (mon *Monitor) LogMetric(ctx context.Context, m api.Metric) error {
	_ = "STUB: not implemented"
	return nil
}

// We received a valid metric so avoid alerting.

// PublishMetric broadcasts a metric to all current cluster peers.
func (mon *Monitor) PublishMetric(ctx context.Context, m api.Metric) error {
	_ = "STUB: not implemented"
	return nil
}

// LatestMetrics returns last known VALID metrics of a given type. A metric
// is only valid if it has not expired and belongs to a current cluster peer.
func (mon *Monitor) LatestMetrics(ctx context.Context, name string) []api.Metric {
	_ = "STUB: not implemented"
	return nil
}

// Make sure we only return metrics in the current peerset if we have
// a peerset provider.

// LatestForPeer returns the latest metric received for a peer (it may have
// expired). It returns nil if no metric exists.
func (mon *Monitor) LatestForPeer(ctx context.Context, name string, pid peer.ID) api.Metric {
	_ = "STUB: not implemented"
	return *new(api.Metric)
}

// Alerts returns a channel on which alerts are sent when the
// monitor detects a failure.
func (mon *Monitor) Alerts() <-chan api.Alert { _ = "STUB: not implemented"; return nil }

// MetricNames lists all metric names.
func (mon *Monitor) MetricNames(ctx context.Context) []string {
	_ = "STUB: not implemented"
	return nil
}

func debug(event string, m api.Metric) { _ = "STUB: not implemented"; return }
