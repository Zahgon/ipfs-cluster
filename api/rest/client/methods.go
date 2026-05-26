package client

import (
	"context"
	"time"

	"github.com/ipfs-cluster/ipfs-cluster/api"

	files "github.com/ipfs/boxo/files"
	peer "github.com/libp2p/go-libp2p/core/peer"
)

// ID returns information about the cluster Peer.
func (c *defaultClient) ID(ctx context.Context) (api.ID, error) {
	_ = "STUB: not implemented"
	return *new(api.ID), nil
}

// Peers requests ID information for all cluster peers.
func (c *defaultClient) Peers(ctx context.Context, out chan<- api.ID) error {
	_ = "STUB: not implemented"
	return nil
}

type peerAddBody struct {
	PeerID string `json:"peer_id"`
}

// PeerAdd adds a new peer to the cluster.
func (c *defaultClient) PeerAdd(ctx context.Context, pid peer.ID) (api.ID, error) {
	_ = "STUB: not implemented"
	return *new(api.ID), nil
}

// PeerRm removes a current peer from the cluster
func (c *defaultClient) PeerRm(ctx context.Context, id peer.ID) error {
	_ = "STUB: not implemented"
	return nil
}

// Pin tracks a Cid with the given replication factor and a name for
// human-friendliness.
func (c *defaultClient) Pin(ctx context.Context, ci api.Cid, opts api.PinOptions) (api.Pin, error) {
	_ = "STUB: not implemented"
	return *new(api.Pin), nil
}

// Unpin untracks a Cid from cluster.
func (c *defaultClient) Unpin(ctx context.Context, ci api.Cid) (api.Pin, error) {
	_ = "STUB: not implemented"
	return *new(api.Pin), nil
}

// PinPath allows to pin an element by the given IPFS path.
func (c *defaultClient) PinPath(ctx context.Context, path string, opts api.PinOptions) (api.Pin, error) {
	_ = "STUB: not implemented"
	return *new(api.Pin), nil
}

// UnpinPath allows to unpin an item by providing its IPFS path.
// It returns the unpinned api.Pin information of the resolved Cid.
func (c *defaultClient) UnpinPath(ctx context.Context, p string) (api.Pin, error) {
	_ = "STUB: not implemented"
	return *new(api.Pin), nil
}

// Allocations returns the consensus state listing all tracked items and
// the peers that should be pinning them.
func (c *defaultClient) Allocations(ctx context.Context, filter api.PinType, out chan<- api.Pin) error {
	_ = "STUB: not implemented"
	return nil
}

// the filter includes this type

// Allocation returns the current allocations for a given Cid.
func (c *defaultClient) Allocation(ctx context.Context, ci api.Cid) (api.Pin, error) {
	_ = "STUB: not implemented"
	return *new(api.Pin), nil
}

// Status returns the current ipfs state for a given Cid. If local is true,
// the information affects only the current peer, otherwise the information
// is fetched from all cluster peers.
func (c *defaultClient) Status(ctx context.Context, ci api.Cid, local bool) (api.GlobalPinInfo, error) {
	_ = "STUB: not implemented"
	return *new(api.GlobalPinInfo), nil
}

// StatusCids returns Status() information for the given Cids. If local is
// true, the information affects only the current peer, otherwise the
// information is fetched from all cluster peers.
func (c *defaultClient) StatusCids(ctx context.Context, cids []api.Cid, local bool, out chan<- api.GlobalPinInfo) error {
	_ = "STUB: not implemented"
	return nil
}

// StatusAll gathers Status() for all tracked items. If a filter is
// provided, only entries matching the given filter statuses
// will be returned. A filter can be built by merging TrackerStatuses with
// a bitwise OR operation (st1 | st2 | ...). A "0" filter value (or
// api.TrackerStatusUndefined), means all.
func (c *defaultClient) StatusAll(ctx context.Context, filter api.TrackerStatus, local bool, out chan<- api.GlobalPinInfo) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *defaultClient) statusAllWithCids(ctx context.Context, filter api.TrackerStatus, cids []api.Cid, local bool, out chan<- api.GlobalPinInfo) error {
	_ = "STUB: not implemented"
	return nil
}

// undefined filter means "all"

// Recover retriggers pin or unpin ipfs operations for a Cid in error state.
// If local is true, the operation is limited to the current peer, otherwise
// it happens on every cluster peer.
func (c *defaultClient) Recover(ctx context.Context, ci api.Cid, local bool) (api.GlobalPinInfo, error) {
	_ = "STUB: not implemented"
	return *new(api.GlobalPinInfo), nil
}

// RecoverAll triggers Recover() operations on all tracked items. If local is
// true, the operation is limited to the current peer. Otherwise, it happens
// everywhere.
func (c *defaultClient) RecoverAll(ctx context.Context, local bool, out chan<- api.GlobalPinInfo) error {
	_ = "STUB: not implemented"
	return nil
}

// Alerts returns information health events in the cluster (expired metrics
// etc.).
func (c *defaultClient) Alerts(ctx context.Context) ([]api.Alert, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// BandwidthByProtocol returns bandwidth stats for each libp2p protocol used.
func (c *defaultClient) BandwidthByProtocol(ctx context.Context) (api.BandwidthByProtocol, error) {
	_ = "STUB: not implemented"
	return *new(api.BandwidthByProtocol), nil
}

// Version returns the ipfs-cluster peer's version.
func (c *defaultClient) Version(ctx context.Context) (api.Version, error) {
	_ = "STUB: not implemented"
	return *new(api.Version), nil
}

// GetConnectGraph returns an ipfs-cluster connection graph.
// The serialized version, strings instead of pids, is returned
func (c *defaultClient) GetConnectGraph(ctx context.Context) (api.ConnectGraph, error) {
	_ = "STUB: not implemented"
	return *new(api.ConnectGraph), nil
}

// Metrics returns a map with the latest valid metrics of the given name
// for the current cluster peers.
func (c *defaultClient) Metrics(ctx context.Context, name string) ([]api.Metric, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// MetricNames lists names of all metrics.
func (c *defaultClient) MetricNames(ctx context.Context) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// RepoGC runs garbage collection on IPFS daemons of cluster peers and
// returns collected CIDs. If local is true, it would garbage collect
// only on contacted peer, otherwise on all peers' IPFS daemons.
func (c *defaultClient) RepoGC(ctx context.Context, local bool) (api.GlobalRepoGC, error) {
	_ = "STUB: not implemented"
	return *new(api.GlobalRepoGC), nil
}

// WaitFor is a utility function that allows for a caller to wait until a CID
// status target is reached (as given in StatusFilterParams).
// It returns the final status for that CID and an error, if there was one.
//
// WaitFor works by calling Status() repeatedly and checking that returned
// peers have transitioned to the target TrackerStatus. It immediately returns
// an error when the an error is among the statuses (and an empty
// GlobalPinInfo).
//
// A special case exists for TrackerStatusPinned targets: in this case,
// TrackerStatusRemote statuses are ignored, so WaitFor will return when
// all Statuses are Pinned or Remote by default.
//
// The Limit parameter allows to specify finer-grained control to, for
// example, only wait until a number of peers reaches a status.
func WaitFor(ctx context.Context, c Client, fp StatusFilterParams) (api.GlobalPinInfo, error) {
	_ = "STUB: not implemented"
	return *new(api.GlobalPinInfo), nil
}

// channel closed

// StatusFilterParams contains the parameters required
// to filter a stream of status results.
type StatusFilterParams struct {
	Cid       api.Cid
	Local     bool // query status from the local peer only
	Target    api.TrackerStatus
	Limit     int // wait for N peers reaching status. 0 == all
	CheckFreq time.Duration
}

type statusFilter struct {
	In, Out chan api.GlobalPinInfo
	Done    chan struct{}
	Err     chan error
}

func newStatusFilter() *statusFilter { _ = "STUB: not implemented"; return nil }

func (sf *statusFilter) filter(ctx context.Context, fp StatusFilterParams) {
	_ = "STUB: not implemented"
	return
}

func (sf *statusFilter) pollStatus(ctx context.Context, c Client, fp StatusFilterParams) {
	_ = "STUB: not implemented"
	return
}

func statusReached(target api.TrackerStatus, gblPinInfo api.GlobalPinInfo, limit int) (bool, error) {
	_ = "STUB: not implemented"
	// Specific case: return error if there are errors
	return false, nil
}

// Specific case: when limit it set, just count how many targets we
// reached.

// General case: all statuses should be the target.
// Specific case: when looking for Pinned, ignore status remote.

// All statuses are the target, as otherwise we would have returned
// false.

// logic drawn from go-ipfs-cmds/cli/parse.go: appendFile
func makeSerialFile(fpath string, params api.AddParams) (string, files.Node, error) {
	_ = "STUB: not implemented"
	return "", *new(files.Node), nil
}

// Add imports files to the cluster from the given paths. A path can
// either be a local filesystem location or an web url (http:// or https://).
// In the latter case, the destination will be downloaded with a GET request.
// The AddParams allow to control different options, like enabling the
// sharding the resulting DAG across the IPFS daemons of multiple cluster
// peers. The output channel will receive regular updates as the adding
// process progresses.
func (c *defaultClient) Add(
	ctx context.Context,
	paths []string,
	params api.AddParams,
	out chan<- api.AddedOutput,
) error {
	_ = "STUB: not implemented"
	return nil
}

// If `form` is set to true, the multipart data will have
// a Content-Type of 'multipart/form-data', if `form` is false,
// the Content-Type will be 'multipart/mixed'.

// AddMultiFile imports new files from a MultiFileReader. See Add().
func (c *defaultClient) AddMultiFile(
	ctx context.Context,
	multiFileR *files.MultiFileReader,
	params api.AddParams,
	out chan<- api.AddedOutput,
) error {
	_ = "STUB: not implemented"
	return nil
}

// This method must run with StreamChannels set.

// our handler decodes an AddedOutput and puts it
// in the out channel.

func (c *defaultClient) Health(ctx context.Context) error { _ = "STUB: not implemented"; return nil }
