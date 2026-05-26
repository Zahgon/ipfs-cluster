package ipfscluster

import (
	blake2b "golang.org/x/crypto/blake2b"

	"github.com/ipfs-cluster/ipfs-cluster/api"
	"github.com/ipfs-cluster/ipfs-cluster/config"
	peer "github.com/libp2p/go-libp2p/core/peer"
	ma "github.com/multiformats/go-multiaddr"
)

// PeersFromMultiaddrs returns all the different peers in the given addresses.
// each peer only will appear once in the result, even if several
// multiaddresses for it are provided.
func PeersFromMultiaddrs(addrs []ma.Multiaddr) []peer.ID { _ = "STUB: not implemented"; return nil }

// // connect to a peer ID.
// func connectToPeer(ctx context.Context, h host.Host, id peer.ID, addr ma.Multiaddr) error {
// 	err := h.Connect(ctx, peerstore.PeerInfo{
// 		ID:    id,
// 		Addrs: []ma.Multiaddr{addr},
// 	})
// 	return err
// }

// // return the local multiaddresses used to communicate to a peer.
// func localMultiaddrsTo(h host.Host, pid peer.ID) []ma.Multiaddr {
// 	var addrs []ma.Multiaddr
// 	conns := h.Network().ConnsToPeer(pid)
// 	logger.Debugf("conns to %s are: %s", pid, conns)
// 	for _, conn := range conns {
// 		addrs = append(addrs, multiaddrJoin(conn.LocalMultiaddr(), h.ID()))
// 	}
// 	return addrs
// }

func logError(fmtstr string, args ...interface{}) error { _ = "STUB: not implemented"; return nil }

func containsPeer(list []peer.ID, peer peer.ID) bool { _ = "STUB: not implemented"; return false }

func minInt(x, y int) int { _ = "STUB: not implemented"; return 0 }

// // updatePinParents modifies the api.Pin input to give it the correct parents
// // so that previous additions to the pins parents are maintained after this
// // pin is committed to consensus.  If this pin carries new parents they are
// // merged with those already existing for this CID.
// func updatePinParents(pin *api.Pin, existing *api.Pin) {
// 	// no existing parents this pin is up to date
// 	if existing.Parents == nil || len(existing.Parents.Keys()) == 0 {
// 		return
// 	}
// 	for _, c := range existing.Parents.Keys() {
// 		pin.Parents.Add(c)
// 	}
// }

type distance [blake2b.Size256]byte

type distanceChecker struct {
	local      peer.ID
	otherPeers []peer.ID
	cache      map[peer.ID]distance
}

func (dc distanceChecker) isClosest(ci api.Cid) bool { _ = "STUB: not implemented"; return false }

// if myDistance is larger than for other peers...

// convertPeerID hashes a Peer ID (Multihash).
func (dc distanceChecker) convertPeerID(id peer.ID) distance {
	_ = "STUB: not implemented"
	return *new(distance)
}

// convertKey hashes a key.
func convertKey(id string) distance { _ = "STUB: not implemented"; return *new(distance) }

func xor(a, b distance) distance { _ = "STUB: not implemented"; return *new(distance) }

// peersSubtract subtracts peers ID slice b from peers ID slice a.
func peersSubtract(a []peer.ID, b []peer.ID) []peer.ID { _ = "STUB: not implemented"; return nil }

// pingValue describes the value carried by ping metrics
type pingValue struct {
	Peername      string          `json:"peer_name,omitempty"`
	IPFSID        peer.ID         `json:"ipfs_id,omitempty"`
	IPFSAddresses []api.Multiaddr `json:"ipfs_addresses,omitempty"`
}

// Valid returns true if the PingValue has IPFSID set.
func (pv pingValue) Valid() bool { _ = "STUB: not implemented"; return false }

// PingValue from metric parses a ping value from the value of a given metric,
// if possible.
func pingValueFromMetric(m api.Metric) (pv pingValue) {
	_ = "STUB: not implemented"
	return *new(pingValue)
}

func publicIPFSAddresses(in []api.Multiaddr) []api.Multiaddr { _ = "STUB: not implemented"; return nil }

// a dns multiaddress: take it

// We have an IP in the multiaddress. Only include
// global unicast.

func toMultiAddrs(addrs config.Strings) ([]ma.Multiaddr, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func multiAddrstoStrings(mAddrs []ma.Multiaddr) []string { _ = "STUB: not implemented"; return nil }
