// Package stateless implements a PinTracker component for IPFS Cluster, which
// aims to reduce the memory footprint when handling really large cluster
// states.
package stateless

import (
	"context"
	"errors"
	"sync"

	"github.com/ipfs-cluster/ipfs-cluster/api"
	"github.com/ipfs-cluster/ipfs-cluster/pintracker/optracker"
	"github.com/ipfs-cluster/ipfs-cluster/state"

	logging "github.com/ipfs/go-log/v2"
	rpc "github.com/libp2p/go-libp2p-gorpc"
	peer "github.com/libp2p/go-libp2p/core/peer"
)

var logger = logging.Logger("pintracker")

const pinsChannelSize = 1024

var (
	// ErrFullQueue is the error used when pin or unpin operation channel is full.
	ErrFullQueue = errors.New("pin/unpin operation queue is full. Try increasing max_pin_queue_size")

	// items with this error should be recovered
	errUnexpectedlyUnpinned = errors.New("the item should be pinned but it is not")
)

// Tracker uses the optracker.OperationTracker to manage
// transitioning shared ipfs-cluster state (Pins) to the local IPFS node.
type Tracker struct {
	config *Config

	optracker *optracker.OperationTracker

	peerID   peer.ID
	peerName string

	ctx    context.Context
	cancel func()

	getState func(ctx context.Context) (state.ReadOnly, error)

	rpcClient *rpc.Client
	rpcReady  chan struct{}

	priorityPinCh chan *optracker.Operation
	pinCh         chan *optracker.Operation
	unpinCh       chan *optracker.Operation

	shutdownMu sync.Mutex
	shutdown   bool
	wg         sync.WaitGroup
}

// New creates a new StatelessPinTracker.
func New(cfg *Config, pid peer.ID, peerName string, getState func(ctx context.Context) (state.ReadOnly, error)) *Tracker {
	_ = "STUB: not implemented"
	return nil
}

// we can get our IPFS id from our own monitor ping metrics which
// are refreshed regularly.
func (spt *Tracker) getIPFSID(ctx context.Context) api.IPFSID {
	_ = "STUB: not implemented"
	// Wait until RPC is ready
	return *new(api.IPFSID)
}

// local peer

// receives a pin Function (pin or unpin) and channels.  Used for both pinning
// and unpinning.
func (spt *Tracker) opWorker(pinF func(*optracker.Operation) error, prioCh, normalCh chan *optracker.Operation) {
	_ = "STUB: not implemented"
	return
}

// Process the priority channel first.

// Then process things on the other channels.
// Block if there are no things to process.

// apply operations that came from some channel

// applyPinF returns true if the operation can be considered "DONE".
func applyPinF(pinF func(*optracker.Operation) error, op *optracker.Operation) bool {
	_ = "STUB: not implemented"

	// operation was canceled. Move on.
	// This saves some time, but not 100% needed.
	return false
}

// call pin/unpin

// there was an error because
// we were canceled. Move on.

// this tells the opWorker to clean the operation from the tracker.

func (spt *Tracker) pin(op *optracker.Operation) error { _ = "STUB: not implemented"; return nil }

func (spt *Tracker) unpin(op *optracker.Operation) error { _ = "STUB: not implemented"; return nil }

// Enqueue puts a new operation on the queue, unless ongoing exists.
func (spt *Tracker) enqueue(ctx context.Context, c api.Pin, typ optracker.OperationType) error {
	_ = "STUB: not implemented"
	return nil
}

// the operation exists and must be queued already.

// SetClient makes the StatelessPinTracker ready to perform RPC requests to
// other components.
func (spt *Tracker) SetClient(c *rpc.Client) { _ = "STUB: not implemented"; return }

// Shutdown finishes the services provided by the StatelessPinTracker
// and cancels any active context.
func (spt *Tracker) Shutdown(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Track tells the StatelessPinTracker to start managing a Cid,
// possibly triggering Pin operations on the IPFS daemon.
func (spt *Tracker) Track(ctx context.Context, c api.Pin) error {
	_ = "STUB: not implemented"
	return nil
}

// Sharded pins are never pinned. A sharded pin cannot turn into
// something else or viceversa like it happens with Remote pins so
// we just ignore them.

// Trigger unpin whenever something remote is tracked
// Note, IPFSConn checks with pin/ls before triggering
// pin/rm.

// ongoing unpin

// Untrack tells the StatelessPinTracker to stop managing a Cid.
// If the Cid is pinned locally, it will be unpinned.
func (spt *Tracker) Untrack(ctx context.Context, c api.Cid) error {
	_ = "STUB: not implemented"
	return nil
}

// StatusAll returns information for all Cids pinned to the local IPFS node.
func (spt *Tracker) StatusAll(ctx context.Context, filter api.TrackerStatus, out chan<- api.PinInfo) error {
	_ = "STUB: not implemented"
	return nil
}

// Any other states are just operation-tracker states, so we just give
// those and return.

// get global state - cluster pinset

// Only query IPFS if we want to status for pinned items

// At some point we need a full map of what we have and what
// we don't. The IPFS pinset is the smallest thing we can keep
// on memory.

// If there was an error listing recursive pins then abort.

// Prepare pinset streaming

// a shorthand for this select.

// For every item in the state.

// if there is an operation, issue that and move on

// next pin

// Preliminary PinInfo for this Pin.

// TBD

// No need to filter. pinnedInIpfs is false
// unless the filter is Pinned |
// UnexpectedlyUnpinned. We filter at the end.

// Not on an operation
// Not a meta pin
// Not a remote pin
// Not a pin on ipfs

// We understand that this is something that
// should be pinned on IPFS and it is not.

// Status returns information for a Cid pinned to the local IPFS node.
func (spt *Tracker) Status(ctx context.Context, c api.Cid) api.PinInfo {
	_ = "STUB: not implemented"
	return *new(api.PinInfo)
}

// check if c has an inflight operation or errorred operation in optracker

// etc to be filled later

// check global state to see if cluster should even be caring about
// the provided cid

// The pin IS in the state.

// check if pin is a meta pin

// check if pin is a remote pin

// else attempt to get status from ipfs node

// The item is in the state but not in IPFS:
// PinError. Should be pinned.

// RecoverAll attempts to recover all items tracked by this peer. It returns
// any errors or when it is done re-tracking.
func (spt *Tracker) RecoverAll(ctx context.Context, out chan<- api.PinInfo) error {
	_ = "STUB: not implemented"
	return nil
}

// Break out if we shutdown. We might be going through
// a very long list of statuses.

// Recover will trigger pinning or unpinning for items in
// PinError or UnpinError states.
func (spt *Tracker) Recover(ctx context.Context, c api.Cid) (api.PinInfo, error) {
	_ = "STUB: not implemented"
	return *new(api.PinInfo), nil
}

// if it was not enqueued, no updated pin-info is returned.
// Use the one we had.

func (spt *Tracker) recoverWithPinInfo(ctx context.Context, pi api.PinInfo) (api.PinInfo, error) {
	_ = "STUB: not implemented"
	return *new(api.PinInfo), nil
}

// ignore error - in case pin was removed while recovering

// We do not return any information when recover was a no-op

// This status call should be cheap as it would normally come from the
// optracker and does not need to hit ipfs.

func (spt *Tracker) ipfsPins(ctx context.Context) (<-chan api.IPFSPinInfo, <-chan error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// type filter.

// PinQueueSize returns the current size of the pinning queue.
func (spt *Tracker) PinQueueSize(ctx context.Context) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// func (spt *Tracker) getErrorsAll(ctx context.Context) []api.PinInfo {
// 	return spt.optracker.Filter(ctx, optracker.PhaseError)
// }

// OpContext exports the internal optracker's OpContext method.
// For testing purposes only.
func (spt *Tracker) OpContext(ctx context.Context, c api.Cid) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func addError(pinInfo *api.PinInfo, err error) { _ = "STUB: not implemented"; return }
