// Package cmdutils contains utilities to facilitate building of command line
// applications launching cluster peers.
package cmdutils

import (
	"context"
	"io"

	ipfscluster "github.com/ipfs-cluster/ipfs-cluster"
	"github.com/ipfs/go-datastore"
	dual "github.com/libp2p/go-libp2p-kad-dht/dual"
	host "github.com/libp2p/go-libp2p/core/host"
	ma "github.com/multiformats/go-multiaddr"
)

// RandomizePorts replaces TCP and UDP ports with random, but valid port
// values, on the given multiaddresses
func RandomizePorts(addrs []ma.Multiaddr) ([]ma.Multiaddr, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ip always comes in the prev component to /tcp or /udp

// ipv6 needs bracketing

// returns the listener so it can be closed later and port
func listenTCP(name, ip string) (io.Closer, int, error) {
	_ = "STUB: not implemented"
	return *new(io.Closer), 0, nil
}

// returns the listener so it can be cloesd later and port
func listenUDP(name, ip string) (io.Closer, int, error) {
	_ = "STUB: not implemented"
	return *new(io.Closer), 0, nil
}

// HandleSignals orderly shuts down an IPFS Cluster peer
// on SIGINT, SIGTERM, SIGHUP. It forces command termination
// on the 3rd-signal count.
func HandleSignals(
	ctx context.Context,
	cancel context.CancelFunc,
	cluster *ipfscluster.Cluster,
	host host.Host,
	dht *dual.DHT,
	store datastore.Datastore,
) error {
	_ = "STUB: not implemented"
	return nil
}

func handleCtrlC(ctx context.Context, cluster *ipfscluster.Cluster, ctrlcCount int) {
	_ = "STUB: not implemented"
	return
}

// ErrorOut formats something and prints it to sdterr.
func ErrorOut(m string, a ...interface{}) { _ = "STUB: not implemented"; return }

// WaitForIPFS hangs until IPFS API becomes available or the given context is
// canceled.  The IPFS API location is determined by the default ipfshttp
// component configuration and can be overridden using environment variables
// that affect that configuration.  Note that we have to do this in the blind,
// since we want to wait for IPFS before we even fetch the IPFS component
// configuration (because the configuration might be hosted on IPFS itself)
func WaitForIPFS(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// sleep an extra second and quit
