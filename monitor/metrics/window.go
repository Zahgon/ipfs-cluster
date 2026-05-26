// Package metrics provides common functionality for working with metrics,
// particularly useful for monitoring components. It includes types to store,
// check and filter metrics.
package metrics

import (
	"container/ring"
	"errors"
	"sync"

	"github.com/ipfs-cluster/ipfs-cluster/api"
)

// DefaultWindowCap sets the amount of metrics to store per peer.
var DefaultWindowCap = 25

// ErrNoMetrics is returned when there are no metrics in a Window.
var ErrNoMetrics = errors.New("no metrics have been added to this window")

// Window implements a circular queue to store metrics.
type Window struct {
	wMu    sync.RWMutex
	window *ring.Ring
}

// NewWindow creates an instance with the given
// window capacity.
func NewWindow(windowCap int) *Window { _ = "STUB: not implemented"; return nil }

// Add adds a new metric to the window. If the window capacity
// has been reached, the oldest metric (by the time it was added),
// will be discarded. Add leaves the cursor on the next spot,
// which is either empty or the oldest record.
func (mw *Window) Add(m api.Metric) { _ = "STUB: not implemented"; return }

// Latest returns the last metric added. It returns an error
// if no metrics were added.
func (mw *Window) Latest() (api.Metric, error) {
	_ = "STUB: not implemented"
	return *new(api.Metric), nil
}

// This just returns the previous ring and
// doesn't set the window "cursor" to the previous
// ring. Therefore this is just a read operation
// as well.

// All returns all the metrics in the window, in the inverse order
// they were Added. That is, result[0] will be the last added
// metric.
func (mw *Window) All() []api.Metric { _ = "STUB: not implemented"; return nil }

// append younger values to older value
