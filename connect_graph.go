package ipfscluster

import (
	"github.com/ipfs-cluster/ipfs-cluster/api"
)

// ConnectGraph returns a description of which cluster peers and ipfs
// daemons are connected to each other.
func (c *Cluster) ConnectGraph() (api.ConnectGraph, error) {
	_ = "STUB: not implemented"
	return *new(api.ConnectGraph), nil
}

// one of the entries is for itself, but that shouldn't hurt

// Only setting cluster connections when no error occurs

// IPFS connections

func (c *Cluster) recordClusterLinks(cg *api.ConnectGraph, p string, peers []api.ID) (bool, api.ID) {
	_ = "STUB: not implemented"
	return false, *new(api.ID)
}

func (c *Cluster) recordIPFSLinks(cg *api.ConnectGraph, pID api.ID) {
	_ = "STUB: not implemented"
	return
}

// Only setting ipfs connections when no error occurs
