// Package crdt implements the IPFS Cluster consensus interface using
// CRDT-datastore to replicate the cluster global state to every peer.
package crdt

import (
	"context"
	"errors"
	"sync"

	"github.com/ipfs-cluster/ipfs-cluster/api"
	"github.com/ipfs-cluster/ipfs-cluster/pstoremgr"
	"github.com/ipfs-cluster/ipfs-cluster/state"

	ds "github.com/ipfs/go-datastore"
	crdt "github.com/ipfs/go-ds-crdt"
	logging "github.com/ipfs/go-log/v2"
	rpc "github.com/libp2p/go-libp2p-gorpc"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	host "github.com/libp2p/go-libp2p/core/host"
	peer "github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/routing"

	ipfslite "github.com/hsanjuan/ipfs-lite"
)

var logger = logging.Logger("crdt")

var (
	// BlocksNs is the namespace to use as blockstore with ipfs-lite.
	BlocksNs   = "b"
	connMgrTag = "crdt"
)

// Common variables for the module.
var (
	ErrNoLeader            = errors.New("crdt consensus component does not provide a leader")
	ErrRmPeer              = errors.New("crdt consensus component cannot remove peers")
	ErrMaxQueueSizeReached = errors.New("batching max_queue_size reached. Too many operations are waiting to be batched. Try increasing the max_queue_size or adjusting the batching options")
)

// wraps pins so that they can be batched.
type batchItem struct {
	ctx     context.Context
	isPin   bool // pin or unpin
	pin     api.Pin
	batched chan error // notify if item was sent for batching
}

// Consensus implement ipfscluster.Consensus and provides the facility to add
// and remove pins from the Cluster shared state. It uses a CRDT-backed
// implementation of go-datastore (go-ds-crdt).
type Consensus struct {
	ctx            context.Context
	cancel         context.CancelFunc
	batchingCtx    context.Context
	batchingCancel context.CancelFunc

	config *Config

	trustedPeers sync.Map

	host        host.Host
	peerManager *pstoremgr.Manager

	store     ds.Datastore
	namespace ds.Key

	state         state.State
	batchingState state.BatchingState
	crdt          *crdt.Datastore
	ipfs          *ipfslite.Peer

	dht    routing.Routing
	pubsub *pubsub.PubSub

	rpcClient  *rpc.Client
	rpcReady   chan struct{}
	stateReady chan struct{}
	readyCh    chan struct{}

	sendToBatchCh chan batchItem
	batchItemCh   chan batchItem
	batchingDone  chan struct{}

	shutdownLock sync.RWMutex
	shutdown     bool
}

// New creates a new crdt Consensus component. The given PubSub will be used to
// broadcast new heads. The given thread-safe datastore will be used to persist
// data and all will be prefixed with cfg.DatastoreNamespace.
func New(
	host host.Host,
	dht routing.Routing,
	pubsub *pubsub.PubSub,
	cfg *Config,
	store ds.Datastore,
) (*Consensus, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (css *Consensus) setup() { _ = "STUB: not implemented"; return }

// Set up a fast-lookup trusted peers cache.
// Protect these peers in the ConnMgr

// Hash the cluster name and produce the topic name from there
// as a way to avoid pubsub topic collisions with other
// pubsub applications potentially when both potentially use
// simple names like "test".

// Validate pubsub messages for our topic (only accept
// from trusted sources)

// subscription name

// TODO: tracing for this context

// unsure if we should set something else but crdt is already
// namespaced and this would only namespace the keys, which only
// complicates things.

// launch batching workers

// notifies State() it is safe to return

// Shutdown closes this component, canceling the pubsub subscription and
// closing the datastore.
func (css *Consensus) Shutdown(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Cancel the batching code

// Only close crdt after canceling the context, otherwise
// the pubsub broadcaster stays on and locks it.

// SetClient gives the component the ability to communicate and
// leaves it ready to use.
func (css *Consensus) SetClient(c *rpc.Client) { _ = "STUB: not implemented"; return }

// Ready returns a channel which is signaled when the component
// is ready to use.
func (css *Consensus) Ready(ctx context.Context) <-chan struct{} {
	_ = "STUB: not implemented"
	return nil

	// IsTrustedPeer returns whether the given peer is taken into account
	// when submitting updates to the consensus state.
}

func (css *Consensus) IsTrustedPeer(ctx context.Context, pid peer.ID) bool {
	_ = "STUB: not implemented"
	return false
}

// Trust marks a peer as "trusted". It makes sure it is trusted as issuer
// for pubsub updates, it is protected in the connection manager, it
// has the highest priority when the peerstore is saved, and it's addresses
// are always remembered.
func (css *Consensus) Trust(ctx context.Context, pid peer.ID) error {
	_ = "STUB: not implemented"
	return nil
}

// Distrust removes a peer from the "trusted" set.
func (css *Consensus) Distrust(ctx context.Context, pid peer.ID) error {
	_ = "STUB: not implemented"
	return nil
}

// LogPin adds a new pin to the shared state.
func (css *Consensus) LogPin(ctx context.Context, pin api.Pin) error {
	_ = "STUB: not implemented"
	return nil
}

// LogUnpin removes a pin from the shared state.
func (css *Consensus) LogUnpin(ctx context.Context, pin api.Pin) error {
	_ = "STUB: not implemented"
	return nil
}

func (css *Consensus) sendToBatchWorker() { _ = "STUB: not implemented"; return }

// This will stay here forever to catch any pins sent
// while shutting down.

// no error
// queue is full

// Launched in setup as a goroutine.
func (css *Consensus) batchWorker() { _ = "STUB: not implemented"; return }

// Create the timer but stop it. It will reset when
// items start arriving.

// Add/Rm from state

// Drain batchItemCh for missing things to be batched

// First item in batch. Start the timer

// Stop timer and commit. Leave ready to reset on next
// item.

// Commit

// timer is expired at this point, it will have to be
// reset.

// Peers returns the current known peerset. It uses
// the monitor component and considers every peer with
// valid known metrics a member.
func (css *Consensus) Peers(ctx context.Context) ([]peer.ID, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Always include self

// WaitForSync is a no-op as it is not necessary to be fully synced for the
// component to be usable.
func (css *Consensus) WaitForSync(ctx context.Context) error {
	_ = "STUB: not implemented"

	// AddPeer is a no-op as we do not need to do peerset management with
	// Merkle-CRDTs. Therefore adding a peer to the peerset means doing nothing.
	return nil
}

func (css *Consensus) AddPeer(ctx context.Context, pid peer.ID) error {
	_ = "STUB: not implemented"

	// RmPeer is a no-op which always errors, as, since we do not do peerset
	// management, we also have no ability to remove a peer from it.
	return nil
}

func (css *Consensus) RmPeer(ctx context.Context, pid peer.ID) error {
	_ = "STUB: not implemented"

	// State returns the cluster shared state. It will block until the consensus
	// component is ready, shutdown or the given context has been canceled.
	return nil
}

func (css *Consensus) State(ctx context.Context) (state.ReadOnly, error) {
	_ = "STUB: not implemented"
	return *new(state.ReadOnly), nil
}

// Clean deletes all crdt-consensus datas from the datastore.
func (css *Consensus) Clean(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Clean deletes all crdt-consensus datas from the given datastore.
func Clean(ctx context.Context, cfg *Config, store ds.Datastore) error {
	_ = "STUB: not implemented"
	return nil
}

// do not die, continue cleaning

// Leader returns ErrNoLeader.
func (css *Consensus) Leader(ctx context.Context) (peer.ID, error) {
	_ = "STUB: not implemented"
	return *

	// OfflineState returns an offline, batching state using the given
	// datastore. This allows to inspect and modify the shared state in offline
	// mode.
	new(peer.ID), nil
}

func OfflineState(cfg *Config, store ds.Datastore) (state.BatchingState, error) {
	_ = "STUB: not implemented"
	return *new(state.BatchingState), nil
}
