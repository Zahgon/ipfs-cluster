// Package pinsvcapi implements an IPFS Cluster API component which provides
// an IPFS Pinning Services API to the cluster.
//
// The implented API is based on the common.API component (refer to module
// description there). The only thing this module does is to provide route
// handling for the otherwise common API component.
package pinsvcapi

import (
	"context"
	"net/http"

	types "github.com/ipfs-cluster/ipfs-cluster/api"
	"github.com/ipfs-cluster/ipfs-cluster/api/common"
	"github.com/ipfs-cluster/ipfs-cluster/api/pinsvcapi/pinsvc"

	logging "github.com/ipfs/go-log/v2"
	rpc "github.com/libp2p/go-libp2p-gorpc"
	"github.com/libp2p/go-libp2p/core/host"
)

var (
	logger    = logging.Logger("pinsvcapi")
	apiLogger = logging.Logger("pinsvcapilog")
)

var apiInfo map[string]string = map[string]string{
	"source":   "IPFS cluster API",
	"warning1": "CID used for requestID. Conflicts possible",
	"warning2": "experimental",
}

func trackerStatusToSvcStatus(st types.TrackerStatus) pinsvc.Status {
	_ = "STUB: not implemented"
	return *new(pinsvc.Status)
}

func svcStatusToTrackerStatus(st pinsvc.Status) types.TrackerStatus {
	_ = "STUB: not implemented"
	return *new(types.TrackerStatus)
}

func svcPinToClusterPin(p pinsvc.Pin) (types.Pin, error) {
	_ = "STUB: not implemented"
	return *new(types.Pin), nil
}

func globalPinInfoToSvcPinStatus(
	rID string,
	gpi types.GlobalPinInfo,
) pinsvc.PinStatus {
	_ = "STUB: not implemented"
	return *new(pinsvc.PinStatus)
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

func (api *API) parseBodyOrFail(w http.ResponseWriter, r *http.Request) pinsvc.Pin {
	_ = "STUB: not implemented"
	return *new(pinsvc.Pin)
}

func (api *API) parseRequestIDOrFail(w http.ResponseWriter, r *http.Request) (types.Cid, bool) {
	_ = "STUB: not implemented"
	return *new(types.Cid), false
}

func (api *API) addPin(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

// Pin item

// Unpin old item

func (api *API) getPinSvcStatus(ctx context.Context, c types.Cid) (pinsvc.PinStatus, error) {
	_ = "STUB: not implemented"
	return *new(pinsvc.PinStatus), nil
}

func (api *API) getPin(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

func (api *API) removePin(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (api *API) listPins(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

// copy approach from restapi

// ignore things unpinning

// i.e things unpinning

func (api *API) pinToSvcPinStatus(ctx context.Context, rID string, pin types.Pin) pinsvc.PinStatus {
	_ = "STUB: not implemented"
	return *new(pinsvc.PinStatus)
}

// all cluster peers

// Delegates should come from allocations

// call the local peer

// retrieve ipfs info for this peer
