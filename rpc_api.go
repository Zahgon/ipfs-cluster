package ipfscluster

import (
	"context"

	"github.com/ipfs-cluster/ipfs-cluster/api"

	rpc "github.com/libp2p/go-libp2p-gorpc"
	peer "github.com/libp2p/go-libp2p/core/peer"
)

// RPC endpoint types w.r.t. trust level
const (
	// RPCClosed endpoints can only be called by the local cluster peer
	// on itself.
	RPCClosed RPCEndpointType = iota
	// RPCTrusted endpoints can be called by "trusted" peers.
	// It depends which peers are considered trusted. For example,
	// in "raft" mode, Cluster will all peers as "trusted". In "crdt" mode,
	// trusted peers are those specified in the configuration.
	RPCTrusted
	// RPCOpen endpoints can be called by any peer in the Cluster swarm.
	RPCOpen
)

// RPCEndpointType controls how access is granted to an RPC endpoint
type RPCEndpointType int

const rpcStreamBufferSize = 1024

// A trick to find where something is used (i.e. Cluster.Pin):
// grep -R -B 3 '"Pin"' | grep -C 1 '"Cluster"'.
// This does not cover globalPinInfo*(...) broadcasts nor redirects to leader
// in Raft.

// newRPCServer returns a new RPC Server for Cluster.
func newRPCServer(c *Cluster) (*rpc.Server, error) { _ = "STUB: not implemented"; return nil, nil }

// RPCServiceID returns the Service ID for the given RPCAPI object.
func RPCServiceID(rpcSvc interface{}) string { _ = "STUB: not implemented"; return "" }

// ClusterRPCAPI is a go-libp2p-gorpc service which provides the internal peer
// API for the main cluster component.
type ClusterRPCAPI struct {
	c *Cluster
}

// PinTrackerRPCAPI is a go-libp2p-gorpc service which provides the internal
// peer API for the PinTracker component.
type PinTrackerRPCAPI struct {
	tracker PinTracker
}

// IPFSConnectorRPCAPI is a go-libp2p-gorpc service which provides the
// internal peer API for the IPFSConnector component.
type IPFSConnectorRPCAPI struct {
	ipfs IPFSConnector
}

// ConsensusRPCAPI is a go-libp2p-gorpc service which provides the
// internal peer API for the Consensus component.
type ConsensusRPCAPI struct {
	cons Consensus
}

// PeerMonitorRPCAPI is a go-libp2p-gorpc service which provides the
// internal peer API for the PeerMonitor component.
type PeerMonitorRPCAPI struct {
	mon PeerMonitor
	pid peer.ID
}

/*
   Cluster component methods
*/

// ID runs Cluster.ID()
func (rpcapi *ClusterRPCAPI) ID(ctx context.Context, in struct{}, out *api.ID) error {
	_ = "STUB: not implemented"
	return nil
}

// IDStream runs Cluster.ID() but in streaming form.
func (rpcapi *ClusterRPCAPI) IDStream(ctx context.Context, in <-chan struct{}, out chan<- api.ID) error {
	_ = "STUB: not implemented"
	return nil
}

// Pin runs Cluster.pin().
func (rpcapi *ClusterRPCAPI) Pin(ctx context.Context, in api.Pin, out *api.Pin) error {
	_ = "STUB: not implemented"
	// we do not call the Pin method directly since that method does not
	// allow to pin other than regular DataType pins. The adder will
	// however send Meta, Shard and ClusterDAG pins.
	return nil
}

// Unpin runs Cluster.Unpin().
func (rpcapi *ClusterRPCAPI) Unpin(ctx context.Context, in api.Pin, out *api.Pin) error {
	_ = "STUB: not implemented"
	return nil
}

// PinPath resolves path into a cid and runs Cluster.Pin().
func (rpcapi *ClusterRPCAPI) PinPath(ctx context.Context, in api.PinPath, out *api.Pin) error {
	_ = "STUB: not implemented"
	return nil
}

// UnpinPath resolves path into a cid and runs Cluster.Unpin().
func (rpcapi *ClusterRPCAPI) UnpinPath(ctx context.Context, in api.PinPath, out *api.Pin) error {
	_ = "STUB: not implemented"
	return nil
}

// Pins runs Cluster.Pins().
func (rpcapi *ClusterRPCAPI) Pins(ctx context.Context, in <-chan struct{}, out chan<- api.Pin) error {
	_ = "STUB: not implemented"
	return nil
}

// PinGet runs Cluster.PinGet().
func (rpcapi *ClusterRPCAPI) PinGet(ctx context.Context, in api.Cid, out *api.Pin) error {
	_ = "STUB: not implemented"
	return nil
}

// Version runs Cluster.Version().
func (rpcapi *ClusterRPCAPI) Version(ctx context.Context, in struct{}, out *api.Version) error {
	_ = "STUB: not implemented"
	return nil
}

// Peers runs Cluster.Peers().
func (rpcapi *ClusterRPCAPI) Peers(ctx context.Context, in <-chan struct{}, out chan<- api.ID) error {
	_ = "STUB: not implemented"
	return nil
}

// PeersWithFilter runs Cluster.peersWithFilter().
func (rpcapi *ClusterRPCAPI) PeersWithFilter(ctx context.Context, in <-chan []peer.ID, out chan<- api.ID) error {
	_ = "STUB: not implemented"
	return nil
}

// PeerAdd runs Cluster.PeerAdd().
func (rpcapi *ClusterRPCAPI) PeerAdd(ctx context.Context, in peer.ID, out *api.ID) error {
	_ = "STUB: not implemented"
	return nil
}

// ConnectGraph runs Cluster.GetConnectGraph().
func (rpcapi *ClusterRPCAPI) ConnectGraph(ctx context.Context, in struct{}, out *api.ConnectGraph) error {
	_ = "STUB: not implemented"
	return nil
}

// PeerRemove runs Cluster.PeerRm().
func (rpcapi *ClusterRPCAPI) PeerRemove(ctx context.Context, in peer.ID, out *struct{}) error {
	_ = "STUB: not implemented"
	return nil
}

// Join runs Cluster.Join().
func (rpcapi *ClusterRPCAPI) Join(ctx context.Context, in api.Multiaddr, out *struct{}) error {
	_ = "STUB: not implemented"
	return nil
}

// StatusAll runs Cluster.StatusAll().
func (rpcapi *ClusterRPCAPI) StatusAll(ctx context.Context, in <-chan api.TrackerStatus, out chan<- api.GlobalPinInfo) error {
	_ = "STUB: not implemented"
	return nil
}

// StatusAllLocal runs Cluster.StatusAllLocal().
func (rpcapi *ClusterRPCAPI) StatusAllLocal(ctx context.Context, in <-chan api.TrackerStatus, out chan<- api.PinInfo) error {
	_ = "STUB: not implemented"
	return nil
}

// Status runs Cluster.Status().
func (rpcapi *ClusterRPCAPI) Status(ctx context.Context, in api.Cid, out *api.GlobalPinInfo) error {
	_ = "STUB: not implemented"
	return nil
}

// StatusLocal runs Cluster.StatusLocal().
func (rpcapi *ClusterRPCAPI) StatusLocal(ctx context.Context, in api.Cid, out *api.PinInfo) error {
	_ = "STUB: not implemented"
	return nil
}

// RecoverAll runs Cluster.RecoverAll().
func (rpcapi *ClusterRPCAPI) RecoverAll(ctx context.Context, in <-chan struct{}, out chan<- api.GlobalPinInfo) error {
	_ = "STUB: not implemented"
	return nil
}

// RecoverAllLocal runs Cluster.RecoverAllLocal().
func (rpcapi *ClusterRPCAPI) RecoverAllLocal(ctx context.Context, in <-chan struct{}, out chan<- api.PinInfo) error {
	_ = "STUB: not implemented"
	return nil
}

// Recover runs Cluster.Recover().
func (rpcapi *ClusterRPCAPI) Recover(ctx context.Context, in api.Cid, out *api.GlobalPinInfo) error {
	_ = "STUB: not implemented"
	return nil
}

// RecoverLocal runs Cluster.RecoverLocal().
func (rpcapi *ClusterRPCAPI) RecoverLocal(ctx context.Context, in api.Cid, out *api.PinInfo) error {
	_ = "STUB: not implemented"
	return nil
}

// BlockAllocate returns allocations for blocks. This is used in the adders.
// It's different from pin allocations when ReplicationFactor < 0.
func (rpcapi *ClusterRPCAPI) BlockAllocate(ctx context.Context, in api.Pin, out *[]peer.ID) error {
	_ = "STUB: not implemented"
	return nil
}

// Allocating for a existing pin. Usually the adder calls this with
// cid.Undef.

// Return the current peer list.

// Returned metrics are Valid and belong to current
// Cluster peers.

// blacklist
// prio list

// RepoGC performs garbage collection sweep on all peers' repos.
func (rpcapi *ClusterRPCAPI) RepoGC(ctx context.Context, in struct{}, out *api.GlobalRepoGC) error {
	_ = "STUB: not implemented"
	return nil
}

// RepoGCLocal performs garbage collection sweep only on the local peer's IPFS daemon.
func (rpcapi *ClusterRPCAPI) RepoGCLocal(ctx context.Context, in struct{}, out *api.RepoGC) error {
	_ = "STUB: not implemented"
	return nil
}

// SendInformerMetrics runs Cluster.sendInformerMetric().
func (rpcapi *ClusterRPCAPI) SendInformerMetrics(ctx context.Context, in struct{}, out *struct{}) error {
	_ = "STUB: not implemented"
	return nil
}

// SendInformersMetrics runs Cluster.sendInformerMetric() on all informers.
func (rpcapi *ClusterRPCAPI) SendInformersMetrics(ctx context.Context, in struct{}, out *struct{}) error {
	_ = "STUB: not implemented"
	return nil
}

// Alerts runs Cluster.Alerts().
func (rpcapi *ClusterRPCAPI) Alerts(ctx context.Context, in struct{}, out *[]api.Alert) error {
	_ = "STUB: not implemented"
	return nil
}

// BandwidthByProtocol runs Cluster.BandwidthByProtocol().
func (rpcapi *ClusterRPCAPI) BandwidthByProtocol(ctx context.Context, in struct{}, out *api.BandwidthByProtocol) error {
	_ = "STUB: not implemented"
	return nil
}

// IPFSID returns the current cached IPFS ID for a peer.
func (rpcapi *ClusterRPCAPI) IPFSID(ctx context.Context, in peer.ID, out *api.IPFSID) error {
	_ = "STUB: not implemented"
	return nil
}

/*
   Tracker component methods
*/

// Track runs PinTracker.Track().
func (rpcapi *PinTrackerRPCAPI) Track(ctx context.Context, in api.Pin, out *struct{}) error {
	_ = "STUB: not implemented"
	return nil
}

// Untrack runs PinTracker.Untrack().
func (rpcapi *PinTrackerRPCAPI) Untrack(ctx context.Context, in api.Pin, out *struct{}) error {
	_ = "STUB: not implemented"
	return nil
}

// StatusAll runs PinTracker.StatusAll().
func (rpcapi *PinTrackerRPCAPI) StatusAll(ctx context.Context, in <-chan api.TrackerStatus, out chan<- api.PinInfo) error {
	_ = "STUB: not implemented"
	return nil
}

// Status runs PinTracker.Status().
func (rpcapi *PinTrackerRPCAPI) Status(ctx context.Context, in api.Cid, out *api.PinInfo) error {
	_ = "STUB: not implemented"
	return nil
}

// RecoverAll runs PinTracker.RecoverAll().f
func (rpcapi *PinTrackerRPCAPI) RecoverAll(ctx context.Context, in <-chan struct{}, out chan<- api.PinInfo) error {
	_ = "STUB: not implemented"
	return nil
}

// Recover runs PinTracker.Recover().
func (rpcapi *PinTrackerRPCAPI) Recover(ctx context.Context, in api.Cid, out *api.PinInfo) error {
	_ = "STUB: not implemented"
	return nil
}

// PinQueueSize runs PinTracker.PinQueueSize().
func (rpcapi *PinTrackerRPCAPI) PinQueueSize(ctx context.Context, in struct{}, out *int64) error {
	_ = "STUB: not implemented"
	return nil
}

/*
   IPFS Connector component methods
*/

// Pin runs IPFSConnector.Pin().
func (rpcapi *IPFSConnectorRPCAPI) Pin(ctx context.Context, in api.Pin, out *struct{}) error {
	_ = "STUB: not implemented"
	return nil
}

// Unpin runs IPFSConnector.Unpin().
func (rpcapi *IPFSConnectorRPCAPI) Unpin(ctx context.Context, in api.Pin, out *struct{}) error {
	_ = "STUB: not implemented"
	return nil
}

// PinLsCid runs IPFSConnector.PinLsCid().
func (rpcapi *IPFSConnectorRPCAPI) PinLsCid(ctx context.Context, in api.Pin, out *api.IPFSPinStatus) error {
	_ = "STUB: not implemented"
	return nil
}

// PinLs runs IPFSConnector.PinLs().
func (rpcapi *IPFSConnectorRPCAPI) PinLs(ctx context.Context, in <-chan []string, out chan<- api.IPFSPinInfo) error {
	_ = "STUB: not implemented"
	return nil
}

// ConfigKey runs IPFSConnector.ConfigKey().
func (rpcapi *IPFSConnectorRPCAPI) ConfigKey(ctx context.Context, in string, out *interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// RepoStat runs IPFSConnector.RepoStat().
func (rpcapi *IPFSConnectorRPCAPI) RepoStat(ctx context.Context, in struct{}, out *api.IPFSRepoStat) error {
	_ = "STUB: not implemented"
	return nil
}

// SwarmPeers runs IPFSConnector.SwarmPeers().
func (rpcapi *IPFSConnectorRPCAPI) SwarmPeers(ctx context.Context, in struct{}, out *[]peer.ID) error {
	_ = "STUB: not implemented"
	return nil
}

// BlockStream runs IPFSConnector.BlockStream().
func (rpcapi *IPFSConnectorRPCAPI) BlockStream(ctx context.Context, in <-chan api.NodeWithMeta, out chan<- struct{}) error {
	_ = "STUB: not implemented"
	// very important to do at the end
	return nil
}

// BlockGet runs IPFSConnector.BlockGet().
func (rpcapi *IPFSConnectorRPCAPI) BlockGet(ctx context.Context, in api.Cid, out *[]byte) error {
	_ = "STUB: not implemented"
	return nil
}

// Resolve runs IPFSConnector.Resolve().
func (rpcapi *IPFSConnectorRPCAPI) Resolve(ctx context.Context, in string, out *api.Cid) error {
	_ = "STUB: not implemented"
	return nil
}

/*
   Consensus component methods
*/

// LogPin runs Consensus.LogPin().
func (rpcapi *ConsensusRPCAPI) LogPin(ctx context.Context, in api.Pin, out *struct{}) error {
	_ = "STUB: not implemented"
	return nil
}

// LogUnpin runs Consensus.LogUnpin().
func (rpcapi *ConsensusRPCAPI) LogUnpin(ctx context.Context, in api.Pin, out *struct{}) error {
	_ = "STUB: not implemented"
	return nil
}

// AddPeer runs Consensus.AddPeer().
func (rpcapi *ConsensusRPCAPI) AddPeer(ctx context.Context, in peer.ID, out *struct{}) error {
	_ = "STUB: not implemented"
	return nil
}

// RmPeer runs Consensus.RmPeer().
func (rpcapi *ConsensusRPCAPI) RmPeer(ctx context.Context, in peer.ID, out *struct{}) error {
	_ = "STUB: not implemented"
	return nil
}

// Peers runs Consensus.Peers().
func (rpcapi *ConsensusRPCAPI) Peers(ctx context.Context, in struct{}, out *[]peer.ID) error {
	_ = "STUB: not implemented"
	return nil
}

/*
   PeerMonitor
*/

// LatestMetrics runs PeerMonitor.LatestMetrics().
func (rpcapi *PeerMonitorRPCAPI) LatestMetrics(ctx context.Context, in string, out *[]api.Metric) error {
	_ = "STUB: not implemented"
	return nil
}

// MetricNames runs PeerMonitor.MetricNames().
func (rpcapi *PeerMonitorRPCAPI) MetricNames(ctx context.Context, in struct{}, out *[]string) error {
	_ = "STUB: not implemented"
	return nil
}
