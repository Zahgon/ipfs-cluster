// Package pstoremgr provides a Manager that simplifies handling
// addition, listing and removal of cluster peer multiaddresses from
// the libp2p Host. This includes resolving DNS addresses, decapsulating
// and encapsulating the /p2p/ (/ipfs/) protocol as needed, listing, saving
// and loading addresses.
package pstoremgr

import (
	"context"
	"sync"
	"time"

	logging "github.com/ipfs/go-log/v2"
	host "github.com/libp2p/go-libp2p/core/host"
	peer "github.com/libp2p/go-libp2p/core/peer"
	peerstore "github.com/libp2p/go-libp2p/core/peerstore"
	ma "github.com/multiformats/go-multiaddr"
)

var logger = logging.Logger("pstoremgr")

// PriorityTag is used to attach metadata to peers in the peerstore
// so they can be sorted.
var PriorityTag = "cluster"

// Timeouts for network operations triggered by the Manager.
var (
	DNSTimeout     = 5 * time.Second
	ConnectTimeout = 5 * time.Second
)

// Manager provides utilities for handling cluster peer addresses
// and storing them in a libp2p Host peerstore.
type Manager struct {
	ctx           context.Context
	host          host.Host
	peerstoreLock sync.Mutex
	peerstorePath string
}

// New creates a Manager with the given libp2p Host and peerstorePath.
// The path indicates the place to persist and read peer addresses from.
// If empty, these operations (LoadPeerstore, SavePeerstore) will no-op.
func New(ctx context.Context, h host.Host, peerstorePath string) *Manager {
	_ = "STUB: not implemented"
	return nil
}

// ImportPeer adds a new peer address to the host's peerstore, optionally
// dialing to it. The address is expected to include the /p2p/<peerID>
// protocol part or to be a /dnsaddr/multiaddress
// Peers are added with the given ttl.
func (pm *Manager) ImportPeer(addr ma.Multiaddr, connect bool, ttl time.Duration) (peer.ID, error) {
	_ = "STUB: not implemented"
	return *new(peer.ID), nil
}

// We need to pre-resolve this

// returns the last peer ID

// Do not add ourselves

// RmPeer clear all addresses for a given peer ID from the host's peerstore.
func (pm *Manager) RmPeer(pid peer.ID) error { _ = "STUB: not implemented"; return nil }

// if the peer has dns addresses, return only those, otherwise
// return all.
func (pm *Manager) filteredPeerAddrs(p peer.ID) []ma.Multiaddr {
	_ = "STUB: not implemented"
	return nil
}

// PeerInfos returns a slice of peerinfos for the given set of peers in order
// of priority. For peers for which we know DNS
// multiaddresses, we only include those. Otherwise, the AddrInfo includes all
// the multiaddresses known for that peer. Peers without addresses are not
// included.
func (pm *Manager) PeerInfos(peers []peer.ID) []peer.AddrInfo {
	_ = "STUB: not implemented"
	return nil
}

// Sort from highest to lowest priority

// ImportPeers calls ImportPeer for every address in the given slice, using the
// given connect parameter. Peers are tagged with priority as given
// by their position in the list.
func (pm *Manager) ImportPeers(addrs []ma.Multiaddr, connect bool, ttl time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

// ImportPeersWithPriority calls ImportPeer for every address in the given
// slice, using the given connect parameter. Peers are tagged with the given priority.
func (pm *Manager) ImportPeersWithPriority(addrs []ma.Multiaddr, connect bool, ttl time.Duration, priority int) error {
	_ = "STUB: not implemented"
	return nil
}

// ImportPeersFromPeerstore reads the peerstore file and calls ImportPeers with
// the addresses obtained from it.
func (pm *Manager) ImportPeersFromPeerstore(connect bool, ttl time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

// LoadPeerstore parses the peerstore file and returns the list
// of addresses read from it.
func (pm *Manager) LoadPeerstore() (addrs []ma.Multiaddr) { _ = "STUB: not implemented"; return nil }

// nothing to load

// skip anything that is not going to be a multiaddress

// SavePeerstore stores a slice of multiaddresses in the peerstore file, one
// per line.
func (pm *Manager) SavePeerstore(pinfos []peer.AddrInfo) error {
	_ = "STUB: not implemented"
	return nil
}

// SavePeerstoreForPeers calls PeerInfos and then saves the peerstore
// file using the result.
func (pm *Manager) SavePeerstoreForPeers(peers []peer.ID) error {
	_ = "STUB: not implemented"
	return nil
}

// Bootstrap attempts to get up to "count" connected peers.  When byPriority
// is true it will order the peers based on their priority tags and attempt to
// connect in that order.  When force is true it will attempt to connect to
// "count" number of peers regardless of the number of existing
// connections. It returns the list of peers it managed to connect to.
func (pm *Manager) Bootstrap(count int, byPriority bool, force bool) []peer.ID {
	_ = "STUB: not implemented"
	return nil
}

// short-cut if there will be nothing to do.

// Sort from highest to lowest priority

// keep conecting while we have peers in the store
// and we have not reached count.

// We are connected, assume success and do not try
// to re-connect

// SetPriority attaches a priority to a peer. 0 means more priority than
// 1. 1 means more priority than 2 etc.
func (pm *Manager) SetPriority(pid peer.ID, prio int) error { _ = "STUB: not implemented"; return nil }

// HandlePeerFound implements the Notifee interface for discovery (mdns).
func (pm *Manager) HandlePeerFound(p peer.AddrInfo) { _ = "STUB: not implemented"; return }

// actually mdns returns a single address but let's do things
// as if there were several

// peerSort is used to sort a slice of PinInfos given the PriorityTag in the
// peerstore, from the lowest tag value (0 is the highest priority) to the
// highest, Peers without a valid priority tag are considered as having a tag
// with value 0, so they will be among the first elements in the resulting
// slice.
type peerSort struct {
	pinfos []peer.AddrInfo
	pstore peerstore.Peerstore
}

func (ps *peerSort) Len() int { _ = "STUB: not implemented"; return 0 }

func (ps *peerSort) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

// Randomize order

func (ps *peerSort) Swap(i, j int) { _ = "STUB: not implemented"; return }

// byString can sort multiaddresses by its string
type byString []ma.Multiaddr

func (m byString) Len() int           { _ = "STUB: not implemented"; return 0 }
func (m byString) Swap(i, j int)      { _ = "STUB: not implemented"; return }
func (m byString) Less(i, j int) bool { _ = "STUB: not implemented"; return false }
