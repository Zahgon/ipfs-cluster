package ipfscluster

import (
	"context"
	"log/slog"

	config "github.com/ipfs-cluster/ipfs-cluster/config"

	ds "github.com/ipfs/go-datastore"
	logging "github.com/ipfs/go-log/v2"
	libp2p "github.com/libp2p/go-libp2p"
	dual "github.com/libp2p/go-libp2p-kad-dht/dual"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	pubsub_pb "github.com/libp2p/go-libp2p-pubsub/pb"
	crypto "github.com/libp2p/go-libp2p/core/crypto"
	host "github.com/libp2p/go-libp2p/core/host"
	metrics "github.com/libp2p/go-libp2p/core/metrics"
	network "github.com/libp2p/go-libp2p/core/network"
	corepnet "github.com/libp2p/go-libp2p/core/pnet"
	"github.com/libp2p/go-libp2p/gologshim"
	autorelay "github.com/libp2p/go-libp2p/p2p/host/autorelay"
	p2pbhost "github.com/libp2p/go-libp2p/p2p/host/basic"
	"github.com/libp2p/go-libp2p/p2p/host/observedaddrs"
	libp2pquic "github.com/libp2p/go-libp2p/p2p/transport/quic"
	ma "github.com/multiformats/go-multiaddr"
)

const dhtNamespace = "dht"

var _ = libp2pquic.NewTransport

func init() {
	// Cluster peers should advertise their public IPs as soon as they
	// learn about them. Default for this is 4, which prevents clusters
	// with less than 4 peers to advertise an external address they know
	// of, therefore they cannot be remembered by other peers asap. This
	// affects dockerized setups mostly. This may announce non-dialable
	// NATed addresses too eagerly, but they should progressively be
	// cleaned up.
	observedaddrs.ActivationThresh = 1

	// Route all slog logs through go-log
	slog.SetDefault(slog.New(logging.SlogHandler()))

	// Connect go-libp2p to go-log
	gologshim.SetDefaultHandler(logging.SlogHandler())
}

// NewClusterHost creates a fully-featured libp2p Host with the options from
// the provided cluster configuration. Using that host, it creates pubsub and
// a DHT instances (persisting to the given datastore), for shared use by all
// cluster components. The returned host uses the DHT for routing. Relay and
// NATService are additionally setup for this host.
func NewClusterHost(
	ctx context.Context,
	ident *config.Identity,
	cfg *Config,
	ds ds.Datastore,
) (host.Host, metrics.Reporter, *pubsub.PubSub, *dual.DHT, error) {
	_ = "STUB: not implemented"

	// Set the default dial timeout for all libp2p connections.  It is not
	// very good to touch this global variable here, but the alternative
	// is to used a modify context everywhere, even if the user supplies
	// it.
	return *new(host.Host), *new(metrics.Reporter), nil, nil, nil
}

// a channel to wait until these variables have been set
// (or left unset on errors). Mostly to avoid reading while writing.

// closed when we finish NewClusterHost

// closed when we finish NewClusterHost

// newHost creates a base cluster host without dht, pubsub, relay or nat etc.
// mostly used for testing.
func newHost(ctx context.Context, psk corepnet.PSK, priv crypto.PrivKey, opts ...libp2p.Option) (host.Host, error) {
	_ = "STUB: not implemented"
	return *new(host.Host), nil
}

func baseOpts(psk corepnet.PSK) []libp2p.Option { _ = "STUB: not implemented"; return nil }

// TODO: quic does not support private networks
// libp2p.DefaultTransports,

func newDHT(ctx context.Context, h host.Host, store ds.Datastore, extraopts ...dual.Option) (*dual.DHT, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// this function reduces the size of the pubsub message IDs to about 24 bytes.
func hashMsgID(m *pubsub_pb.Message) string {
	_ = "STUB: not implemented"
	// hash := blake2b.Sum256(m.Data)
	return ""
}

func newPubSub(ctx context.Context, cfg *Config, h host.Host) (*pubsub.PubSub, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Let's configure pubsub in an attempt to support networks with
// thousands of peers which may be bandwidth and processing constrained.

// Hearbeat should be rather slow as some peers just may not manage
// to handle all incoming publications, so we don't have to deal with
// additional IHAVEs too often either.

// We increase the gap between History gossip and length to give peers
// more time to grab messages, but expire announcements rather quick
// x * Hearbeat seconds of availability of messages for IWANTs (def 5)
// x * Heartbeat seconds of announcement of messages via IHAVEs (def 3)

// Longer Hearbeat intervals means more messages queue up.  IHAVEs
// lists may have more message at given moment, so increase this
// default.
// Note: it may be better to reduce this to a minimum and reduce
// heartbeat interval, but then, this is a list of IDs so 10000x16byte
// x D is not too much to transfer out is it?
// 10000 message IDs. (def 5000)

// For my taste, default mesh parameters are a bit low. We already
// have connections to a bunch of peers given DHT etc and having a
// tiny mesh only makes more replay necessary. Thus allow me to
// increase 4x the defaults.
// default 6
// default 5
// default 12
// default 4
// default 2
// default 6

// FloodPublish is disabled as every peer publishes and we
// prefer messages to just follow the mesh rather than risking
// that small peers saturate themselves every time they
// publish a metric.

// Custom hash function reduces size of messages and thus traffic.

// default is 32. Give more leeway to large clusters.
//default also 32

// Keep a long cache of seen messages. Otherwise, if they
// don't propagate under two minutes, they are re-requested
// and re-broadcasted.  1000 peers * 3 metrics * 1 minute
// interval * 30 minutes = 180000 messages * 16 bytes = 1.3MiB
// of memory needed, tops.
// default 120 seconds

// future work
//pubsub.WithDefaultValidator(
//	pubsub.NewBasicSeqnoValidator(h.Peerstore())),

// Inspired in Kubo's
// https://github.com/ipfs/go-ipfs/blob/9327ee64ce96ca6da29bb2a099e0e0930b0d9e09/core/node/libp2p/relay.go#L79-L103
// and https://github.com/ipfs/go-ipfs/blob/9327ee64ce96ca6da29bb2a099e0e0930b0d9e09/core/node/libp2p/routing.go#L242-L317
// but simplified and adapted:
//   - Everytime we need peers for relays we do a DHT lookup.
//   - We return the peers from that lookup.
//   - No need to do it async, since we have to wait for the full lookup to
//     return anyways. We put them on a buffered channel and be done.
func newPeerSource(hostGetter func() host.Host, dhtGetter func() *dual.DHT) autorelay.PeerSource {
	_ = "STUB: not implemented"
	return *new(autorelay.PeerSource)
}

// make a channel to return, and put items from numPeers on
// that channel up to numPeers. Then close it.

// Because the Host, DHT are initialized after relay, we need to
// obtain them indirectly this way.

// context canceled etc.

// context canceled etc.

// length of closest peers is K.

// Bail out. Usually a "no peers found".

// Attempt to put peers on r if we have space,
// otherwise return (we reached numPeers)

// We are here if numPeers > closestPeers

// EncodeProtectorKey converts a byte slice to its hex string representation.
func EncodeProtectorKey(secretBytes []byte) string { _ = "STUB: not implemented"; return "" }

func makeAddrsFactory(announce []ma.Multiaddr, noAnnounce []ma.Multiaddr) (p2pbhost.AddrsFactory, error) {
	_ = "STUB: not implemented"
	return *new(p2pbhost.AddrsFactory), nil
}

// check for exact matches

// check for /ipcidr matches

// mostly copy/pasted from https://github.com/ipfs/rainbow/blob/main/rcmgr.go
// which is itself copy-pasted from Kubo, because libp2p does not have
// a sane way of doing this.
func makeResourceMgr(enabled bool, maxMemory, maxFD uint64, connMgrHighWater int) (network.ResourceManager, error) {
	_ = "STUB: not implemented"
	return *new(network.ResourceManager), nil
}

// Auto-scaled limits based on available memory/fds.

// 1 GiB

// At least as of 2023-01-25, it's possible to open a connection that
// doesn't ask for any memory usage with the libp2p Resource Manager/Accountant
// (see https://github.com/libp2p/go-libp2p/issues/2010#issuecomment-1404280736).
// As a result, we can't currently rely on Memory limits to full protect us.
// Until https://github.com/libp2p/go-libp2p/issues/2010 is addressed,
// we take a proxy now of restricting to 1 inbound connection per MB.
// Note: this is more generous than go-libp2p's default autoscaled limits which do
// 64 connections per 1GB
// (see https://github.com/libp2p/go-libp2p/blob/master/p2p/host/resource-manager/limit_defaults.go#L357 ).

// Transient connections won't cause any memory to be accounted for by the resource manager/accountant.
// Only established connections do.
// As a result, we can't rely on System.Memory to protect us from a bunch of transient connection being opened.
// We limit the same values as the System scope, but only allow the Transient scope to take 25% of what is allowed for the System scope.

// Lets get out of the way of the allow list functionality.
// If someone specified "Swarm.ResourceMgr.Allowlist" we should let it go through.

// Keep it simple by not having Service, ServicePeer, Protocol, ProtocolPeer, Conn, or Stream limits.

// Limit the resources consumed by a peer.
// This doesn't protect us against intentional DoS attacks since an attacker can easily spin up multiple peers.
// We specify this limit against unintentional DoS attacks (e.g., a peer has a bug and is sending too much traffic intentionally).
// In that case we want to keep that peer's resource consumption contained.
// To keep this simple, we only constrain inbound connections and streams.

// Anything set above in partialLimits that had a value of rcmgr.DefaultLimit will be overridden.
// Anything in scalingLimitConfig that wasn't defined in partialLimits above will be added (e.g., libp2p's default service limits).

// Simple checks to override autoscaling ensuring limits make sense versus the connmgr values.
// There are ways to break this, but this should catch most problems already.
// We might improve this in the future.
// See: https://github.com/ipfs/kubo/issues/9545

// Scale System.StreamsInbound as well, but use the existing ratio of StreamsInbound to ConnsInbound

// We already have a complete value thus pass in an empty ConcreteLimitConfig.
