package test

import (
	"context"
	"errors"
	"testing"

	"github.com/ipfs-cluster/ipfs-cluster/api"

	rpc "github.com/libp2p/go-libp2p-gorpc"
	host "github.com/libp2p/go-libp2p/core/host"
	peer "github.com/libp2p/go-libp2p/core/peer"
)

var (
	// ErrBadCid is returned when using ErrorCid. Operations with that CID always
	// fail.
	ErrBadCid = errors.New("this is an expected error when using ErrorCid")
	// ErrLinkNotFound is error returned when no link is found
	ErrLinkNotFound = errors.New("no link by that name")
)

// NewMockRPCClient creates a mock ipfs-cluster RPC server and returns
// a client to it.
func NewMockRPCClient(t testing.TB) *rpc.Client { _ = "STUB: not implemented"; return nil }

// NewMockRPCClientWithHost returns a mock ipfs-cluster RPC server
// initialized with a given host.
func NewMockRPCClientWithHost(t testing.TB, h host.Host) *rpc.Client {
	_ = "STUB: not implemented"
	return nil
}

type mockCluster struct{}
type mockPinTracker struct{}
type mockIPFSConnector struct{}
type mockConsensus struct{}
type mockPeerMonitor struct{}

func (mock *mockCluster) Pin(ctx context.Context, in api.Pin, out *api.Pin) error {
	_ = "STUB: not implemented"
	return nil
}

// a pin is never returned the replications set to 0.

func (mock *mockCluster) Unpin(ctx context.Context, in api.Pin, out *api.Pin) error {
	_ = "STUB: not implemented"
	return nil
}

func (mock *mockCluster) PinPath(ctx context.Context, in api.PinPath, out *api.Pin) error {
	_ = "STUB: not implemented"
	return nil
}

// no need to resolve

// must resolve

func (mock *mockCluster) UnpinPath(ctx context.Context, in api.PinPath, out *api.Pin) error {
	_ = "STUB: not implemented"
	return nil
}

// Mock-Unpin behaves like pin (doing nothing).

func (mock *mockCluster) Pins(ctx context.Context, in <-chan struct{}, out chan<- api.Pin) error {
	_ = "STUB: not implemented"
	return nil
}

func (mock *mockCluster) PinGet(ctx context.Context, in api.Cid, out *api.Pin) error {
	_ = "STUB: not implemented"
	return nil
}

// This is a remote pin

func (mock *mockCluster) ID(ctx context.Context, in struct{}, out *api.ID) error {
	_ = "STUB: not implemented"
	//_, pubkey, _ := crypto.GenerateKeyPair(
	//	DefaultConfigCrypto,
	//	DefaultConfigKeyLength)
	return nil
}

//PublicKey: pubkey,

func (mock *mockCluster) IDStream(ctx context.Context, in <-chan struct{}, out chan<- api.ID) error {
	_ = "STUB: not implemented"
	return nil
}

func (mock *mockCluster) Version(ctx context.Context, in struct{}, out *api.Version) error {
	_ = "STUB: not implemented"
	return nil
}

func (mock *mockCluster) Peers(ctx context.Context, in <-chan struct{}, out chan<- api.ID) error {
	_ = "STUB: not implemented"
	return nil
}

func (mock *mockCluster) PeersWithFilter(ctx context.Context, in <-chan []peer.ID, out chan<- api.ID) error {
	_ = "STUB: not implemented"
	return nil
}

func (mock *mockCluster) PeerAdd(ctx context.Context, in peer.ID, out *api.ID) error {
	_ = "STUB: not implemented"
	return nil
}

func (mock *mockCluster) PeerRemove(ctx context.Context, in peer.ID, out *struct{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (mock *mockCluster) ConnectGraph(ctx context.Context, in struct{}, out *api.ConnectGraph) error {
	_ = "STUB: not implemented"
	return nil
}

func (mock *mockCluster) StatusAll(ctx context.Context, in <-chan api.TrackerStatus, out chan<- api.GlobalPinInfo) error {
	_ = "STUB: not implemented"
	return nil
}

// If there is no filter match, we will not return that status and we
// will not have an entry for that peer in the peerMap.  In turn, when
// a single peer, we will not have an entry for the cid at all.

func (mock *mockCluster) StatusAllLocal(ctx context.Context, in <-chan api.TrackerStatus, out chan<- api.PinInfo) error {
	_ = "STUB: not implemented"
	return nil
}

func (mock *mockCluster) Status(ctx context.Context, in api.Cid, out *api.GlobalPinInfo) error {
	_ = "STUB: not implemented"
	return nil
}

func (mock *mockCluster) StatusLocal(ctx context.Context, in api.Cid, out *api.PinInfo) error {
	_ = "STUB: not implemented"
	return nil
}

func (mock *mockCluster) RecoverAll(ctx context.Context, in <-chan struct{}, out chan<- api.GlobalPinInfo) error {
	_ = "STUB: not implemented"
	return nil
}

func (mock *mockCluster) RecoverAllLocal(ctx context.Context, in <-chan struct{}, out chan<- api.PinInfo) error {
	_ = "STUB: not implemented"
	return nil
}

func (mock *mockCluster) Recover(ctx context.Context, in api.Cid, out *api.GlobalPinInfo) error {
	_ = "STUB: not implemented"
	return nil
}

func (mock *mockCluster) RecoverLocal(ctx context.Context, in api.Cid, out *api.PinInfo) error {
	_ = "STUB: not implemented"
	return nil
}

func (mock *mockCluster) BlockAllocate(ctx context.Context, in api.Pin, out *[]peer.ID) error {
	_ = "STUB: not implemented"
	return nil
}

// allocate to local peer

func (mock *mockCluster) RepoGC(ctx context.Context, in struct{}, out *api.GlobalRepoGC) error {
	_ = "STUB: not implemented"
	return nil
}

func (mock *mockCluster) RepoGCLocal(ctx context.Context, in struct{}, out *api.RepoGC) error {
	_ = "STUB: not implemented"
	return nil
}

func (mock *mockCluster) SendInformerMetrics(ctx context.Context, in struct{}, out *struct{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (mock *mockCluster) Alerts(ctx context.Context, in struct{}, out *[]api.Alert) error {
	_ = "STUB: not implemented"
	return nil
}

func (mock *mockCluster) BandwidthByProtocol(ctx context.Context, in struct{}, out *api.BandwidthByProtocol) error {
	_ = "STUB: not implemented"
	return nil
}

func (mock *mockCluster) IPFSID(ctx context.Context, in peer.ID, out *api.IPFSID) error {
	_ = "STUB: not implemented"
	return nil
}

/* Tracker methods */

func (mock *mockPinTracker) Track(ctx context.Context, in api.Pin, out *struct{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (mock *mockPinTracker) Untrack(ctx context.Context, in api.Pin, out *struct{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (mock *mockPinTracker) StatusAll(ctx context.Context, in <-chan api.TrackerStatus, out chan<- api.PinInfo) error {
	_ = "STUB: not implemented"
	return nil
}

func (mock *mockPinTracker) Status(ctx context.Context, in api.Cid, out *api.PinInfo) error {
	_ = "STUB: not implemented"
	return nil
}

func (mock *mockPinTracker) RecoverAll(ctx context.Context, in <-chan struct{}, out chan<- api.PinInfo) error {
	_ = "STUB: not implemented"
	return nil
}

func (mock *mockPinTracker) Recover(ctx context.Context, in api.Cid, out *api.PinInfo) error {
	_ = "STUB: not implemented"
	return nil
}

func (mock *mockPinTracker) PinQueueSize(ctx context.Context, in struct{}, out *int64) error {
	_ = "STUB: not implemented"
	return

	/* PeerMonitor methods */
	nil
}

// LatestMetrics runs PeerMonitor.LatestMetrics().
func (mock *mockPeerMonitor) LatestMetrics(ctx context.Context, in string, out *[]api.Metric) error {
	_ = "STUB: not implemented"
	return nil
}

// MetricNames runs PeerMonitor.MetricNames().
func (mock *mockPeerMonitor) MetricNames(ctx context.Context, in struct{}, out *[]string) error {
	_ = "STUB: not implemented"
	return nil
}

/* IPFSConnector methods */

func (mock *mockIPFSConnector) Pin(ctx context.Context, in api.Pin, out *struct{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (mock *mockIPFSConnector) Unpin(ctx context.Context, in api.Pin, out *struct{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (mock *mockIPFSConnector) PinLsCid(ctx context.Context, in api.Pin, out *api.IPFSPinStatus) error {
	_ = "STUB: not implemented"
	return nil
}

func (mock *mockIPFSConnector) PinLs(ctx context.Context, in <-chan []string, out chan<- api.IPFSPinInfo) error {
	_ = "STUB: not implemented"
	return nil
}

func (mock *mockIPFSConnector) SwarmPeers(ctx context.Context, in struct{}, out *[]peer.ID) error {
	_ = "STUB: not implemented"
	return nil
}

func (mock *mockIPFSConnector) ConfigKey(ctx context.Context, in string, out *interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (mock *mockIPFSConnector) RepoStat(ctx context.Context, in struct{}, out *api.IPFSRepoStat) error {
	_ = "STUB: not implemented"
	// since we have two pins. Assume each is 1000B.
	return nil
}

func (mock *mockIPFSConnector) BlockStream(ctx context.Context, in <-chan api.NodeWithMeta, out chan<- struct{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (mock *mockIPFSConnector) Resolve(ctx context.Context, in string, out *api.Cid) error {
	_ = "STUB: not implemented"
	return nil
}

func (mock *mockConsensus) AddPeer(ctx context.Context, in peer.ID, out *struct{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (mock *mockConsensus) RmPeer(ctx context.Context, in peer.ID, out *struct{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (mock *mockConsensus) Peers(ctx context.Context, in struct{}, out *[]peer.ID) error {
	_ = "STUB: not implemented"
	return nil
}
