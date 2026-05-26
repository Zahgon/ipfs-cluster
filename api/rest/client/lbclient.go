package client

import (
	"context"

	"github.com/ipfs-cluster/ipfs-cluster/api"

	files "github.com/ipfs/boxo/files"
	shell "github.com/ipfs/go-ipfs-api"

	peer "github.com/libp2p/go-libp2p/core/peer"
)

// loadBalancingClient is a client to interact with IPFS Cluster APIs
// that balances the load by distributing requests among peers.
type loadBalancingClient struct {
	strategy LBStrategy
	retries  int
}

// LBStrategy is a strategy to load balance requests among clients.
type LBStrategy interface {
	Next(count int) Client
	SetClients(clients []Client)
}

// RoundRobin is a load balancing strategy that would use clients in a sequence
// for all methods, throughout the lifetime of the lb client.
type RoundRobin struct {
	clients []Client
	counter uint32
	length  uint32
}

// Next return the next client to be used.
func (r *RoundRobin) Next(count int) Client { _ = "STUB: not implemented"; return *new(Client) }

// SetClients sets a list of clients for this strategy.
func (r *RoundRobin) SetClients(cl []Client) { _ = "STUB: not implemented"; return }

// Failover is a load balancing strategy that would try the first cluster peer
// first. If the first call fails it would try other clients for that call in a
// round robin fashion.
type Failover struct {
	clients []Client
}

// Next returns the next client to be used.
func (f *Failover) Next(count int) Client { _ = "STUB: not implemented"; return *new(Client) }

// SetClients sets a list of clients for this strategy.
func (f *Failover) SetClients(cl []Client) {
	_ = "STUB: not implemented"

	// NewLBClient returns a new client that would load balance requests among
	// clients.
	return
}

func NewLBClient(strategy LBStrategy, cfgs []*Config, retries int) (Client, error) {
	_ = "STUB: not implemented"
	return *new(Client), nil
}

// retry tries the request until it is successful or tries `lc.retries` times.
func (lc *loadBalancingClient) retry(count int, call func(Client) error) error {
	_ = "STUB: not implemented"
	return nil
}

// successful request

// It is a safety check. This error should never occur.
// All errors returned by client methods are of type `api.Error`.

// ID returns information about the cluster Peer.
func (lc *loadBalancingClient) ID(ctx context.Context) (api.ID, error) {
	_ = "STUB: not implemented"
	return *new(api.ID), nil
}

// Peers requests ID information for all cluster peers.
func (lc *loadBalancingClient) Peers(ctx context.Context, out chan<- api.ID) error {
	_ = "STUB: not implemented"
	return nil
}

// this blocks until done

// wait for cout to be closed

// retries call as needed.

// PeerAdd adds a new peer to the cluster.
func (lc *loadBalancingClient) PeerAdd(ctx context.Context, pid peer.ID) (api.ID, error) {
	_ = "STUB: not implemented"
	return *new(api.ID), nil
}

// PeerRm removes a current peer from the cluster.
func (lc *loadBalancingClient) PeerRm(ctx context.Context, id peer.ID) error {
	_ = "STUB: not implemented"
	return nil
}

// Pin tracks a Cid with the given replication factor and a name for
// human-friendliness.
func (lc *loadBalancingClient) Pin(ctx context.Context, ci api.Cid, opts api.PinOptions) (api.Pin, error) {
	_ = "STUB: not implemented"
	return *new(api.Pin), nil
}

// Unpin untracks a Cid from cluster.
func (lc *loadBalancingClient) Unpin(ctx context.Context, ci api.Cid) (api.Pin, error) {
	_ = "STUB: not implemented"
	return *new(api.Pin), nil
}

// PinPath allows to pin an element by the given IPFS path.
func (lc *loadBalancingClient) PinPath(ctx context.Context, path string, opts api.PinOptions) (api.Pin, error) {
	_ = "STUB: not implemented"
	return *new(api.Pin), nil
}

// UnpinPath allows to unpin an item by providing its IPFS path.
// It returns the unpinned api.Pin information of the resolved Cid.
func (lc *loadBalancingClient) UnpinPath(ctx context.Context, p string) (api.Pin, error) {
	_ = "STUB: not implemented"
	return *new(api.Pin), nil
}

// Allocations returns the consensus state listing all tracked items and
// the peers that should be pinning them.
func (lc *loadBalancingClient) Allocations(ctx context.Context, filter api.PinType, out chan<- api.Pin) error {
	_ = "STUB: not implemented"
	return nil
}

// this blocks until done

// wait for cout to be closed

// Allocation returns the current allocations for a given Cid.
func (lc *loadBalancingClient) Allocation(ctx context.Context, ci api.Cid) (api.Pin, error) {
	_ = "STUB: not implemented"
	return *new(api.Pin), nil
}

// Status returns the current ipfs state for a given Cid. If local is true,
// the information affects only the current peer, otherwise the information
// is fetched from all cluster peers.
func (lc *loadBalancingClient) Status(ctx context.Context, ci api.Cid, local bool) (api.GlobalPinInfo, error) {
	_ = "STUB: not implemented"
	return *new(api.GlobalPinInfo), nil
}

// StatusCids returns Status() information for the given Cids. If local is
// true, the information affects only the current peer, otherwise the
// information is fetched from all cluster peers.
func (lc *loadBalancingClient) StatusCids(ctx context.Context, cids []api.Cid, local bool, out chan<- api.GlobalPinInfo) error {
	_ = "STUB: not implemented"
	return nil
}

// this blocks until done

// wait for cout to be closed

// StatusAll gathers Status() for all tracked items. If a filter is
// provided, only entries matching the given filter statuses
// will be returned. A filter can be built by merging TrackerStatuses with
// a bitwise OR operation (st1 | st2 | ...). A "0" filter value (or
// api.TrackerStatusUndefined), means all.
func (lc *loadBalancingClient) StatusAll(ctx context.Context, filter api.TrackerStatus, local bool, out chan<- api.GlobalPinInfo) error {
	_ = "STUB: not implemented"
	return nil
}

// this blocks until done

// wait for cout to be closed

// Recover retriggers pin or unpin ipfs operations for a Cid in error state.
// If local is true, the operation is limited to the current peer, otherwise
// it happens on every cluster peer.
func (lc *loadBalancingClient) Recover(ctx context.Context, ci api.Cid, local bool) (api.GlobalPinInfo, error) {
	_ = "STUB: not implemented"
	return *new(api.GlobalPinInfo), nil
}

// RecoverAll triggers Recover() operations on all tracked items. If local is
// true, the operation is limited to the current peer. Otherwise, it happens
// everywhere.
func (lc *loadBalancingClient) RecoverAll(ctx context.Context, local bool, out chan<- api.GlobalPinInfo) error {
	_ = "STUB: not implemented"
	return nil
}

// this blocks until done

// wait for cout to be closed

// Alerts returns things that are wrong with cluster.
func (lc *loadBalancingClient) Alerts(ctx context.Context) ([]api.Alert, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (lc *loadBalancingClient) BandwidthByProtocol(ctx context.Context) (api.BandwidthByProtocol, error) {
	_ = "STUB: not implemented"
	return *new(api.BandwidthByProtocol), nil
}

// Version returns the ipfs-cluster peer's version.
func (lc *loadBalancingClient) Version(ctx context.Context) (api.Version, error) {
	_ = "STUB: not implemented"
	return *new(api.Version), nil
}

// GetConnectGraph returns an ipfs-cluster connection graph.
// The serialized version, strings instead of pids, is returned.
func (lc *loadBalancingClient) GetConnectGraph(ctx context.Context) (api.ConnectGraph, error) {
	_ = "STUB: not implemented"
	return *new(api.ConnectGraph), nil
}

// Metrics returns a map with the latest valid metrics of the given name
// for the current cluster peers.
func (lc *loadBalancingClient) Metrics(ctx context.Context, name string) ([]api.Metric, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// MetricNames returns the list of metric types.
func (lc *loadBalancingClient) MetricNames(ctx context.Context) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// RepoGC runs garbage collection on IPFS daemons of cluster peers and
// returns collected CIDs. If local is true, it would garbage collect
// only on contacted peer, otherwise on all peers' IPFS daemons.
func (lc *loadBalancingClient) RepoGC(ctx context.Context, local bool) (api.GlobalRepoGC, error) {
	_ = "STUB: not implemented"
	return *new(api.GlobalRepoGC), nil
}

// Add imports files to the cluster from the given paths. A path can
// either be a local filesystem location or an web url (http:// or https://).
// In the latter case, the destination will be downloaded with a GET request.
// The AddParams allow to control different options, like enabling the
// sharding the resulting DAG across the IPFS daemons of multiple cluster
// peers. The output channel will receive regular updates as the adding
// process progresses.
func (lc *loadBalancingClient) Add(
	ctx context.Context,
	paths []string,
	params api.AddParams,
	out chan<- api.AddedOutput,
) error {
	_ = "STUB: not implemented"
	return nil
}

// this blocks until done

// wait for cout to be closed

// AddMultiFile imports new files from a MultiFileReader. See Add().
func (lc *loadBalancingClient) AddMultiFile(
	ctx context.Context,
	multiFileR *files.MultiFileReader,
	params api.AddParams,
	out chan<- api.AddedOutput,
) error {
	_ = "STUB: not implemented"
	return nil
}

// this blocks until done

// wait for cout to be closed

// IPFS returns an instance of go-ipfs-api's Shell, pointing to the
// configured ProxyAddr (or to the default Cluster's IPFS proxy port).
// It re-uses this Client's HTTP client, thus will be constrained by
// the same configurations affecting it (timeouts...).
func (lc *loadBalancingClient) IPFS(ctx context.Context) *shell.Shell {
	_ = "STUB: not implemented"
	return nil
}

func (lc *loadBalancingClient) Health(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}
