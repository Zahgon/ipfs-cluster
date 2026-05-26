package raft

import (
	"context"
	"errors"
	"io"
	"time"

	"github.com/ipfs-cluster/ipfs-cluster/state"

	p2praft "github.com/libp2p/go-libp2p-raft"
	host "github.com/libp2p/go-libp2p/core/host"
	peer "github.com/libp2p/go-libp2p/core/peer"

	hraft "github.com/hashicorp/raft"
	raftboltdb "github.com/hashicorp/raft-boltdb"
)

var raftLogger = &p2praft.HcLogToLogger{}

// ErrWaitingForSelf is returned when we are waiting for ourselves to depart
// the peer set, which won't happen
var errWaitingForSelf = errors.New("waiting for ourselves to depart")

// RaftMaxSnapshots indicates how many snapshots to keep in the consensus data
// folder.
// TODO: Maybe include this in Config. Not sure how useful it is to touch
// this anyways.
var RaftMaxSnapshots = 5

// RaftLogCacheSize is the maximum number of logs to cache in-memory.
// This is used to reduce disk I/O for the recently committed entries.
var RaftLogCacheSize = 512

// How long we wait for updates during shutdown before snapshotting
var waitForUpdatesShutdownTimeout = 5 * time.Second
var waitForUpdatesInterval = 400 * time.Millisecond

// How many times to retry snapshotting when shutting down
var maxShutdownSnapshotRetries = 5

// raftWrapper wraps the hraft.Raft object and related things like the
// different stores used or the hraft.Configuration.
// Its methods provide functionality for working with Raft.
type raftWrapper struct {
	ctx           context.Context
	cancel        context.CancelFunc
	raft          *hraft.Raft
	config        *Config
	host          host.Host
	serverConfig  hraft.Configuration
	transport     *hraft.NetworkTransport
	snapshotStore hraft.SnapshotStore
	logStore      hraft.LogStore
	stableStore   hraft.StableStore
	boltdb        *raftboltdb.BoltStore
	staging       bool
}

// newRaftWrapper creates a Raft instance and initializes
// everything leaving it ready to use. Note, that Bootstrap() should be called
// to make sure the raft instance is usable.
func newRaftWrapper(
	host host.Host,
	cfg *Config,
	fsm hraft.FSM,
	staging bool,
) (*raftWrapper, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Set correct LocalID

// makeDataFolder creates the folder that is meant to store Raft data. Ensures
// we always set 0700 mode.
func makeDataFolder(folder string) error { _ = "STUB: not implemented"; return nil }

func (rw *raftWrapper) makeTransport() (err error) { _ = "STUB: not implemented"; return nil }

func (rw *raftWrapper) makeStores() error { _ = "STUB: not implemented"; return nil }

// wraps the store in a LogCache to improve performance.
// See consul/agent/consul/server.go

// Bootstrap calls BootstrapCluster on the Raft instance with a valid
// Configuration (generated from InitPeerset) when Raft has no state
// and we are not setting up a staging peer. It returns if Raft
// was bootstrapped (true) and an error.
func (rw *raftWrapper) Bootstrap() (bool, error) { _ = "STUB: not implemented"; return false, nil }

// Inform the user that we are working with a pre-existing peerset

// create Raft servers configuration. The result is used
// by Bootstrap() when it proceeds to Bootstrap.
func (rw *raftWrapper) makeServerConfig() { _ = "STUB: not implemented"; return }

// creates a server configuration with all peers as Voters.
func makeServerConf(peers []peer.ID) hraft.Configuration {
	_ = "STUB: not implemented"
	return *new(hraft.Configuration)
}

// Servers are peers + self. We avoid duplicate entries below

// avoid dups

// WaitForLeader holds until Raft says we have a leader.
// Returns if ctx is canceled.
func (rw *raftWrapper) WaitForLeader(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (rw *raftWrapper) WaitForVoter(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func isVoter(srvID hraft.ServerID, cfg hraft.Configuration) bool {
	_ = "STUB: not implemented"
	return false
}

// WaitForUpdates holds until Raft has synced to the last index in the log
func (rw *raftWrapper) WaitForUpdates(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (rw *raftWrapper) WaitForPeer(ctx context.Context, pid string, depart bool) error {
	_ = "STUB: not implemented"
	return nil
}

// departing

// joining

// Snapshot tells Raft to take a snapshot.
func (rw *raftWrapper) Snapshot() error { _ = "STUB: not implemented"; return nil }

// snapshotOnShutdown attempts to take a snapshot before a shutdown.
// Snapshotting might fail if the raft applied index is not the last index.
// This waits for the updates and tries to take a snapshot when the
// applied index is up to date.
// It will retry if the snapshot still fails, in case more updates have arrived.
// If waiting for updates times-out, it will not try anymore, since something
// is wrong. This is a best-effort solution as there is no way to tell Raft
// to stop processing entries because we want to take a snapshot before
// shutting down.
func (rw *raftWrapper) snapshotOnShutdown() error { _ = "STUB: not implemented"; return nil }

// things worked

// There was an error

// Shutdown shutdown Raft and closes the BoltDB.
func (rw *raftWrapper) Shutdown(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// important!

// AddPeer adds a peer to Raft
func (rw *raftWrapper) AddPeer(ctx context.Context, peer string) error {
	_ = "STUB: not implemented"
	return nil
}

// Check that we don't have it to not waste
// log entries if so.

// TODO: Extra cfg value?

// RemovePeer removes a peer from Raft
func (rw *raftWrapper) RemovePeer(ctx context.Context, peer string) error {
	_ = "STUB: not implemented"
	return nil
}

// Check that we have it to not waste
// log entries if we don't.

// TODO: Extra cfg value?

// Leader returns Raft's leader. It may be an empty string if
// there is no leader or it is unknown.
func (rw *raftWrapper) Leader(ctx context.Context) string { _ = "STUB: not implemented"; return "" }

func (rw *raftWrapper) Peers(ctx context.Context) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// latestSnapshot looks for the most recent raft snapshot stored at the
// provided basedir.  It returns the snapshot's metadata, and a reader
// to the snapshot's bytes
func latestSnapshot(raftDataFolder string) (*hraft.SnapshotMeta, io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return nil, *new(io.ReadCloser), nil
}

// no error if snapshot isn't found

// LastStateRaw returns the bytes of the last snapshot stored, its metadata,
// and a flag indicating whether any snapshot was found.
func LastStateRaw(cfg *Config) (io.Reader, bool, error) {
	_ = "STUB: not implemented"
	// Read most recent snapshot
	return *new(io.Reader), false, nil
}

// nothing to read

// no snapshots could be read

// SnapshotSave saves the provided state to a snapshot in the
// raft data path.  Old raft data is backed up and replaced
// by the new snapshot.  pids contains the config-specified
// peer ids to include in the snapshot metadata if no snapshot exists
// from which to copy the raft metadata
func SnapshotSave(cfg *Config, newState state.State, pids []peer.ID) error {
	_ = "STUB: not implemented"
	return nil
}

// make a new raft snapshot
// As of hraft v1.0.0 this is always 1

// Begin the log after the index of a fresh start so that
// the snapshot's state propagate's during bootstrap

// CleanupRaft moves the current data folder to a backup location
func CleanupRaft(cfg *Config) error { _ = "STUB: not implemented"; return nil }

// no snapshots at all. Avoid creating backups
// from empty state folders.

// only call when Raft is shutdown
func (rw *raftWrapper) Clean() error { _ = "STUB: not implemented"; return nil }

func find(s []string, elem string) bool { _ = "STUB: not implemented"; return false }

func (rw *raftWrapper) observePeers() { _ = "STUB: not implemented"; return }
