// Package rest implements an IPFS Cluster API component. It provides
// a REST-ish API to interact with Cluster.
//
// The implented API is based on the common.API component (refer to module
// description there). The only thing this module does is to provide route
// handling for the otherwise common API component.
package rest

import (
	"context"
	"net/http"

	types "github.com/ipfs-cluster/ipfs-cluster/api"
	"github.com/ipfs-cluster/ipfs-cluster/api/common"

	logging "github.com/ipfs/go-log/v2"
	rpc "github.com/libp2p/go-libp2p-gorpc"
	"github.com/libp2p/go-libp2p/core/host"
)

var (
	logger    = logging.Logger("restapi")
	apiLogger = logging.Logger("restapilog")
)

type peerAddBody struct {
	PeerID string `json:"peer_id"`
}

// API implements the REST API Component.
// It embeds a common.API.
type API struct {
	*common.API

	rpcClient *rpc.Client
	config    *Config
}

// NewAPI creates a new REST API component.
func NewAPI(ctx context.Context, cfg *Config) (*API, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewAPIWithHost creates a new REST API component using the given libp2p Host.
func NewAPIWithHost(ctx context.Context, cfg *Config, h host.Host) (*API, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Routes returns endpoints supported by this API.
func (api *API) routes(c *rpc.Client) []common.Route { _ = "STUB: not implemented"; return nil }

func (api *API) idHandler(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (api *API) versionHandler(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (api *API) graphHandler(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (api *API) metricsHandler(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (api *API) metricNamesHandler(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (api *API) alertsHandler(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (api *API) bandwidthByProtocolHandler(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (api *API) addHandler(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// any errors sent as trailer

func (api *API) peerListHandler(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (api *API) peerAddHandler(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (api *API) peerRemoveHandler(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (api *API) pinHandler(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// span.AddAttributes(trace.StringAttribute("cid", pin.Cid))

func (api *API) unpinHandler(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// span.AddAttributes(trace.StringAttribute("cid", pin.Cid))

func (api *API) pinPathHandler(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (api *API) unpinPathHandler(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (api *API) allocationsHandler(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// this means we keep iterating if no filter
// matched

func (api *API) allocationHandler(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (api *API) statusAllHandler(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// FIXME: This is a bit lazy, as "invalidxx,pinned" would result in a
// valid "pinned" filter.

// request statuses for multiple CIDs in parallel.
func (api *API) statusCidsHandler(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// Close channel when done

func (api *API) statusHandler(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (api *API) recoverAllHandler(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (api *API) recoverHandler(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (api *API) repoGCHandler(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func repoGCToGlobal(r types.RepoGC) types.GlobalRepoGC {
	_ = "STUB: not implemented"
	return *new(types.GlobalRepoGC)
}
