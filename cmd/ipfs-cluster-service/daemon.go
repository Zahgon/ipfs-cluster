package main

import (
	"context"

	ipfscluster "github.com/ipfs-cluster/ipfs-cluster"
	"github.com/ipfs-cluster/ipfs-cluster/cmdutils"

	ds "github.com/ipfs/go-datastore"
	dual "github.com/libp2p/go-libp2p-kad-dht/dual"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	host "github.com/libp2p/go-libp2p/core/host"
	metrics "github.com/libp2p/go-libp2p/core/metrics"

	ma "github.com/multiformats/go-multiaddr"

	cli "github.com/urfave/cli"
)

func parseBootstraps(flagVal []string) (bootstraps []ma.Multiaddr) {
	_ = "STUB: not implemented"
	return nil
}

// Runs the cluster peer
func daemon(c *cli.Context) error { _ = "STUB: not implemented"; return nil }

// Execution lock

// Load all the configurations and identity

// Setup bootstrapping

// Cleanup state if bootstrapping

// noop if no bootstraps
// if bootstrapping fails, consensus will never be ready
// and timeout. So this can happen in background and we
// avoid worrying about error handling here (since Cluster
// will realize).

// send readiness notification to systemd

// createCluster creates all the necessary things to produce the cluster
// object and returns it along the datastore so the lifecycle can be handled
// (the datastore needs to be Closed after shutting down the Cluster).
func createCluster(
	ctx context.Context,
	c *cli.Context,
	cfgHelper *cmdutils.ConfigHelper,
	host host.Host,
	bwc metrics.Reporter,
	pubsub *pubsub.PubSub,
	dht *dual.DHT,
	store ds.Datastore,
	raftStaging bool,
) (*ipfscluster.Cluster, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Do NOT enable default Libp2p API endpoint on CRDT
// clusters. Collaborative clusters are likely to share the
// secret with untrusted peers, thus the API would be open for
// anyone.

// For legacy compatibility we need to make the allocator
// automatically compatible with informers that have been loaded. For
// simplicity we assume that anyone that does not specify an allocator
// configuration (legacy configs), will be using "freespace"

// bootstrap will bootstrap this peer to one of the bootstrap addresses
// if there are any.
func bootstrap(ctx context.Context, cluster *ipfscluster.Cluster, bootstraps []ma.Multiaddr) {
	_ = "STUB: not implemented"
	return
}

func setupDatastore(cfgHelper *cmdutils.ConfigHelper) ds.Datastore {
	_ = "STUB: not implemented"
	return *new(ds.Datastore)
}

func setupConsensus(
	cfgHelper *cmdutils.ConfigHelper,
	h host.Host,
	dht *dual.DHT,
	pubsub *pubsub.PubSub,
	store ds.Datastore,
	raftStaging bool,
) (ipfscluster.Consensus, error) {
	_ = "STUB: not implemented"
	return *new(ipfscluster.Consensus), nil
}

// go-ds-crdt migrations are the main cause that may need
// additional time for this consensus layer to be ready.
