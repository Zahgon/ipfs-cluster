// Package raft implements a Consensus component for IPFS Cluster which uses
// Raft (go-libp2p-raft).
package raft

import (
	"context"
	"sync"

	"github.com/ipfs-cluster/ipfs-cluster/api"
	"github.com/ipfs-cluster/ipfs-cluster/state"

	ds "github.com/ipfs/go-datastore"
	logging "github.com/ipfs/go-log/v2"
	consensus "github.com/libp2p/go-libp2p-consensus"
	rpc "github.com/libp2p/go-libp2p-gorpc"
	host "github.com/libp2p/go-libp2p/core/host"
	peer "github.com/libp2p/go-libp2p/core/peer"
)

var logger = logging.Logger("raft")

// Consensus handles the work of keeping a shared-state between
// the peers of an IPFS Cluster, as well as modifying that state and
// applying any updates in a thread-safe manner.
type Consensus struct {
	ctx    context.Context
	cancel func()
	config *Config

	host host.Host

	consensus consensus.OpLogConsensus
	actor     consensus.Actor
	baseOp    *LogOp
	raft      *raftWrapper

	rpcClient *rpc.Client
	rpcReady  chan struct{}
	readyCh   chan struct{}

	shutdownLock sync.RWMutex
	shutdown     bool
}

// NewConsensus builds a new ClusterConsensus component using Raft.
//
// Raft saves state snapshots regularly and persists log data in a bolt
// datastore. Therefore, unless memory usage is a concern, it is recommended
// to use an in-memory go-datastore as store parameter.
//
// The staging parameter controls if the Raft peer should start in
// staging mode (used when joining a new Raft peerset with other peers).
//
// The store parameter should be a thread-safe datastore.
func NewConsensus(
	host host.Host,
	cfg *Config,
	store ds.Datastore,
	staging bool, // this peer must not be bootstrapped if no state exists
) (*Consensus, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// WaitForSync waits for a leader and for the state to be up to date, then returns.
func (cc *Consensus) WaitForSync(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// 1 - wait for leader
// 2 - wait until we are a Voter
// 3 - wait until last index is applied

// From raft docs:

// once a staging server receives enough log entries to be sufficiently
// caught up to the leader's log, the leader will invoke a  membership
// change to change the Staging server to a Voter

// Thus, waiting to be a Voter is a guarantee that we have a reasonable
// up to date state. Otherwise, we might return too early (see
// https://github.com/ipfs-cluster/ipfs-cluster/issues/378)

// waits until there is a consensus leader and syncs the state
// to the tracker. If errors happen, this will return and never
// signal the component as Ready.
func (cc *Consensus) finishBootstrap() {
	_ = "STUB: not implemented"
	// wait until we have RPC to perform any actions.
	return
}

// Sometimes bootstrap is a no-op. It only applies when
// no state exists and staging=false.

// Shutdown stops the component so it will not process any
// more updates. The underlying consensus is permanently
// shutdown, along with the libp2p transport.
func (cc *Consensus) Shutdown(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Raft Shutdown

// SetClient makes the component ready to perform RPC requets
func (cc *Consensus) SetClient(c *rpc.Client) { _ = "STUB: not implemented"; return }

// Ready returns a channel which is signaled when the Consensus
// algorithm has finished bootstrapping and is ready to use
func (cc *Consensus) Ready(ctx context.Context) <-chan struct{} {
	_ = "STUB: not implemented"
	return nil
}

// IsTrustedPeer returns true. In Raft we trust all peers.
func (cc *Consensus) IsTrustedPeer(ctx context.Context, p peer.ID) bool {
	_ = "STUB: not implemented"

	// Trust is a no-op.
	return false
}

func (cc *Consensus) Trust(ctx context.Context, pid peer.ID) error {
	_ = "STUB: not implemented"

	// Distrust is a no-op.
	return nil
}

func (cc *Consensus) Distrust(ctx context.Context, pid peer.ID) error {
	_ = "STUB: not implemented"
	return nil
}

func (cc *Consensus) op(ctx context.Context, pin api.Pin, t LogOpType) *LogOp {
	_ = "STUB: not implemented"
	return nil
}

// returns true if the operation was redirected to the leader
// note that if the leader just disappeared, the rpc call will
// fail because we haven't heard that it's gone.
func (cc *Consensus) redirectToLeader(ctx context.Context, method string, arg interface{}) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Retry redirects

// No leader, wait for one

// means we timed out waiting for a leader
// we don't retry in this case

// We are the leader. Do not redirect

// We tried to redirect, but something happened

// commit submits a cc.consensus commit. It retries upon failures.
func (cc *Consensus) commit(ctx context.Context, op *LogOp, rpcOp string, redirectArg interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// required to cross the serialized boundary

// this means we are retrying

// try to send it to the leader
// redirectToLeader has it's own retry loop. If this fails
// we're done here.

// Being here means we are the LEADER. We can commit.

// now commit the changes to our state
// do not shut down while committing

// LogPin submits a Cid to the shared state of the cluster. It will forward
// the operation to the leader if this is not it.
func (cc *Consensus) LogPin(ctx context.Context, pin api.Pin) error {
	_ = "STUB: not implemented"
	return nil
}

// LogUnpin removes a Cid from the shared state of the cluster.
func (cc *Consensus) LogUnpin(ctx context.Context, pin api.Pin) error {
	_ = "STUB: not implemented"
	return nil
}

// AddPeer adds a new peer to participate in this consensus. It will
// forward the operation to the leader if this is not it.
func (cc *Consensus) AddPeer(ctx context.Context, pid peer.ID) error {
	_ = "STUB: not implemented"
	return nil
}

// Being here means we are the leader and can commit
// do not shutdown while committing

// RmPeer removes a peer from this consensus. It will
// forward the operation to the leader if this is not it.
func (cc *Consensus) RmPeer(ctx context.Context, pid peer.ID) error {
	_ = "STUB: not implemented"
	return nil
}

// Being here means we are the leader and can commit
// do not shutdown while committing

// State retrieves the current consensus State. It may error if no State has
// been agreed upon or the state is not consistent. The returned State is the
// last agreed-upon State known by this node. No writes are allowed, as all
// writes to the shared state should happen through the Consensus component
// methods.
func (cc *Consensus) State(ctx context.Context) (state.ReadOnly, error) {
	_ = "STUB: not implemented"
	return *new(state.ReadOnly), nil
}

// Leader returns the peerID of the Leader of the
// cluster. It returns an error when there is no leader.
func (cc *Consensus) Leader(ctx context.Context) (peer.ID, error) {
	_ = "STUB: not implemented"
	return *new(peer.ID), nil
}

// Note the hard-dependency on raft here...

// Clean removes the Raft persisted state.
func (cc *Consensus) Clean(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Rollback replaces the current agreed-upon
// state with the state provided. Only the consensus leader
// can perform this operation.
func (cc *Consensus) Rollback(state state.State) error {
	_ = "STUB: not implemented"
	// This is unused. It *might* be used for upgrades.
	// There is rather untested magic in libp2p-raft's FSM()
	// to make this possible.
	return nil
}

// Peers return the current list of peers in the consensus.
// The list will be sorted alphabetically.
func (cc *Consensus) Peers(ctx context.Context) ([]peer.ID, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// prevent shutdown while here

// things hang a lot in this case

// OfflineState state returns a cluster state by reading the Raft data and
// writing it to the given datastore which is then wrapped as a state.State.
// Usually an in-memory datastore suffices. The given datastore should be
// thread-safe.
func OfflineState(cfg *Config, store ds.Datastore) (state.State, error) {
	_ = "STUB: not implemented"
	return *new(state.State), nil
}
