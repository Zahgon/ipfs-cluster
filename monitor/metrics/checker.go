package metrics

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/ipfs-cluster/ipfs-cluster/api"

	peer "github.com/libp2p/go-libp2p/core/peer"
)

// AlertChannelCap specifies how much buffer the alerts channel has.
var AlertChannelCap = 256

// MaxAlertThreshold specifies how many alerts will occur per a peer is
// removed from the list of monitored peers.
var MaxAlertThreshold = 1

// ErrAlertChannelFull is returned if the alert channel is full.
var ErrAlertChannelFull = errors.New("alert channel is full")

// Checker provides utilities to find expired metrics
// for a given peerset and send alerts if it proceeds to do so.
type Checker struct {
	ctx     context.Context
	alertCh chan api.Alert
	metrics *Store

	failedPeersMu sync.Mutex
	failedPeers   map[peer.ID]map[string]int
}

// NewChecker creates a Checker using the given
// MetricsStore. The threshold value indicates when a
// monitored component should be considered to have failed.
// The greater the threshold value the more leniency is granted.
//
// A value between 2.0 and 4.0 is suggested for the threshold.
func NewChecker(ctx context.Context, metrics *Store) *Checker {
	_ = "STUB: not implemented"
	return nil
}

// CheckPeers will trigger alerts based on the latest metrics from the given peerset
// when they have expired and no alert has been sent before.
func (mc *Checker) CheckPeers(peers []peer.ID) error { _ = "STUB: not implemented"; return nil }

// CheckAll will trigger alerts for all latest metrics when they have expired
// and no alert has been sent before.
func (mc *Checker) CheckAll() error { _ = "STUB: not implemented"; return nil }

// ResetAlerts clears up how many time a peer alerted for a given metric.
// Thus, if it was over the threshold, it will start alerting again.
func (mc *Checker) ResetAlerts(pid peer.ID, metricName string) { _ = "STUB: not implemented"; return }

func (mc *Checker) alert(pid peer.ID, metricName string) error {
	_ = "STUB: not implemented"
	return nil
}

// If above threshold, do not send alert

// Cleanup old metrics eventually

// Alerts returns a channel which gets notified by CheckPeers.
func (mc *Checker) Alerts() <-chan api.Alert {
	_ = "STUB: not implemented"

	// Watch will trigger regular CheckPeers on the given interval. It will call
	// peersF to obtain a peerset. It can be stopped by canceling the context.
	// Usually you want to launch this in a goroutine.
	return nil
}

func (mc *Checker) Watch(ctx context.Context, peersF func(context.Context) ([]peer.ID, error), interval time.Duration) {
	_ = "STUB: not implemented"
	return
}

// FailedMetric returns if a peer is marked as failed for a particular metric.
func (mc *Checker) FailedMetric(metric string, pid peer.ID) bool {
	_ = "STUB: not implemented"
	return false
}
