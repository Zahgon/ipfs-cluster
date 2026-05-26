// Package api holds declarations for types used in ipfs-cluster APIs to make
// them re-usable across differen tools. This include RPC API "Serial[izable]"
// versions for types. The Go API uses natives types, while RPC API,
// REST APIs etc use serializable types (i.e. json format). Conversion methods
// exists between types.
//
// Note that all conversion methods ignore any parsing errors. All values must
// be validated first before initializing any of the types defined here.
package api

import (
	"net/url"
	"time"

	pb "github.com/ipfs-cluster/ipfs-cluster/api/pb"

	cid "github.com/ipfs/go-cid"
	logging "github.com/ipfs/go-log/v2"
	peer "github.com/libp2p/go-libp2p/core/peer"
	protocol "github.com/libp2p/go-libp2p/core/protocol"
	multiaddr "github.com/multiformats/go-multiaddr"

	// needed to parse /ws multiaddresses
	_ "github.com/libp2p/go-libp2p/p2p/transport/websocket"
	// needed to parse /dns* multiaddresses
	_ "github.com/multiformats/go-multiaddr-dns"
)

var logger = logging.Logger("apitypes")

var unixZero = time.Unix(0, 0)

func init() {
	// initialize trackerStatusString
	stringTrackerStatus = make(map[string]TrackerStatus)
	for k, v := range trackerStatusString {
		stringTrackerStatus[v] = k
	}
}

// TrackerStatus values
const (
	// IPFSStatus should never take this value.
	// When used as a filter. It means "all".
	TrackerStatusUndefined TrackerStatus = 0
	// The cluster node is offline or not responding
	TrackerStatusClusterError TrackerStatus = 1 << iota
	// An error occurred pinning
	TrackerStatusPinError
	// An error occurred unpinning
	TrackerStatusUnpinError
	// The IPFS daemon has pinned the item
	TrackerStatusPinned
	// The IPFS daemon is currently pinning the item
	TrackerStatusPinning
	// The IPFS daemon is currently unpinning the item
	TrackerStatusUnpinning
	// The IPFS daemon is not pinning the item
	TrackerStatusUnpinned
	// The IPFS daemon is not pinning the item but it is being tracked
	TrackerStatusRemote
	// The item has been queued for pinning on the IPFS daemon
	TrackerStatusPinQueued
	// The item has been queued for unpinning on the IPFS daemon
	TrackerStatusUnpinQueued
	// The IPFS daemon is not pinning the item through this cid but it is
	// tracked in a cluster dag
	TrackerStatusSharded
	// The item is in the state and should be pinned, but
	// it is however not pinned and not queued/pinning.
	TrackerStatusUnexpectedlyUnpinned
)

// Composite TrackerStatus.
const (
	TrackerStatusError  = TrackerStatusClusterError | TrackerStatusPinError | TrackerStatusUnpinError
	TrackerStatusQueued = TrackerStatusPinQueued | TrackerStatusUnpinQueued
)

// TrackerStatus represents the status of a tracked Cid in the PinTracker
type TrackerStatus int

var trackerStatusString = map[TrackerStatus]string{
	TrackerStatusUndefined:            "undefined",
	TrackerStatusClusterError:         "cluster_error",
	TrackerStatusPinError:             "pin_error",
	TrackerStatusUnpinError:           "unpin_error",
	TrackerStatusError:                "error",
	TrackerStatusPinned:               "pinned",
	TrackerStatusPinning:              "pinning",
	TrackerStatusUnpinning:            "unpinning",
	TrackerStatusUnpinned:             "unpinned",
	TrackerStatusRemote:               "remote",
	TrackerStatusPinQueued:            "pin_queued",
	TrackerStatusUnpinQueued:          "unpin_queued",
	TrackerStatusQueued:               "queued",
	TrackerStatusSharded:              "sharded",
	TrackerStatusUnexpectedlyUnpinned: "unexpectedly_unpinned",
}

// values autofilled in init()
var stringTrackerStatus map[string]TrackerStatus

// String converts a TrackerStatus into a readable string.
// If the given TrackerStatus is a filter (with several
// bits set), it will return a comma-separated list.
func (st TrackerStatus) String() string {
	_ = "STUB: not implemented"

	// simple and known composite values
	return ""
}

// other filters

// Match returns true if the tracker status matches the given filter.
// For example TrackerStatusPinError will match TrackerStatusPinError
// and TrackerStatusError.
func (st TrackerStatus) Match(filter TrackerStatus) bool { _ = "STUB: not implemented"; return false }

// MarshalJSON uses the string representation of TrackerStatus for JSON
// encoding.
func (st TrackerStatus) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// UnmarshalJSON sets a tracker status from its JSON representation.
func (st *TrackerStatus) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

// TrackerStatusFromString parses a string and returns the matching
// TrackerStatus value. The string can be a comma-separated list
// representing a TrackerStatus filter. Unknown status names are
// ignored.
func TrackerStatusFromString(str string) TrackerStatus {
	_ = "STUB: not implemented"
	return *new(TrackerStatus)
}

// TrackerStatusAll all known TrackerStatus values.
func TrackerStatusAll() []TrackerStatus { _ = "STUB: not implemented"; return nil }

// IPFSPinStatus values
// FIXME include maxdepth
const (
	IPFSPinStatusBug IPFSPinStatus = iota
	IPFSPinStatusError
	IPFSPinStatusDirect
	IPFSPinStatusRecursive
	IPFSPinStatusIndirect
	IPFSPinStatusUnpinned
)

// IPFSPinStatus represents the status of a pin in IPFS (direct, recursive etc.)
type IPFSPinStatus int

// IPFSPinStatusFromString parses a string and returns the matching
// IPFSPinStatus.
func IPFSPinStatusFromString(t string) IPFSPinStatus {
	_ = "STUB: not implemented"
	// Since indirect statuses are of the form "indirect through <cid>"
	// use a prefix match
	return *new(IPFSPinStatus)
}

// FIXME: Maxdepth?

// String returns the string form of the status as written by IPFS.
func (ips IPFSPinStatus) String() string { _ = "STUB: not implemented"; return "" }

// UnmarshalJSON parses a status from JSON
func (ips *IPFSPinStatus) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

// MarshalJSON converts a status to JSON.
func (ips IPFSPinStatus) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// IsPinned returns true if the item is pinned as expected by the
// maxDepth parameter.
func (ips IPFSPinStatus) IsPinned(maxDepth PinDepth) bool { _ = "STUB: not implemented"; return false }

// FIXME: when we know how ipfs returns partial pins.

// ToTrackerStatus converts the IPFSPinStatus value to the
// appropriate TrackerStatus value.
func (ips IPFSPinStatus) ToTrackerStatus() TrackerStatus {
	_ = "STUB: not implemented"
	return *new(TrackerStatus)
}

var ipfsPinStatus2TrackerStatusMap = map[IPFSPinStatus]TrackerStatus{
	IPFSPinStatusDirect:    TrackerStatusPinned,
	IPFSPinStatusRecursive: TrackerStatusPinned,
	IPFSPinStatusIndirect:  TrackerStatusUnpinned,
	IPFSPinStatusUnpinned:  TrackerStatusUnpinned,
	IPFSPinStatusBug:       TrackerStatusUndefined,
	IPFSPinStatusError:     TrackerStatusClusterError, //TODO(ajl): check suitability
}

// Cid embeds a cid.Cid with the MarshalJSON/UnmarshalJSON methods overwritten.
type Cid struct {
	cid.Cid
}

// CidUndef is an Undefined CID.
var CidUndef = Cid{cid.Undef}

// NewCid wraps a cid.Cid in a Cid.
func NewCid(c cid.Cid) Cid {
	_ = "STUB: not implemented"
	return *

	// DecodeCid parses a CID from its string form.
	new(Cid)
}

func DecodeCid(str string) (Cid, error) { _ = "STUB: not implemented"; return *new(Cid), nil }

// CastCid returns a CID from its bytes.
func CastCid(bs []byte) (Cid, error) { _ = "STUB: not implemented"; return *new(Cid), nil }

// MarshalJSON marshals a CID as JSON as a normal CID string.
func (c Cid) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// UnmarshalJSON reads a CID from its representation as JSON string.
func (c *Cid) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

// Equals returns true if two Cids are equal.
func (c Cid) Equals(c2 Cid) bool { _ = "STUB: not implemented"; return false }

// IPFSPinInfo represents an IPFS Pin, which only has a CID and type.
// Its JSON form is what IPFS returns when querying a pinset.
type IPFSPinInfo struct {
	Cid  Cid           `json:"Cid" codec:"c"`
	Type IPFSPinStatus `json:"Type" codec:"t"`
}

// GlobalPinInfo contains cluster-wide status information about a tracked Cid,
// indexed by cluster peer.
type GlobalPinInfo struct {
	Cid         Cid               `json:"cid" codec:"c"`
	Name        string            `json:"name" codec:"n"`
	Allocations []peer.ID         `json:"allocations" codec:"a,omitempty"`
	Origins     []Multiaddr       `json:"origins" codec:"g,omitempty"`
	Created     time.Time         `json:"created" codec:"t,omitempty"`
	Metadata    map[string]string `json:"metadata" codec:"m,omitempty"`

	// https://github.com/golang/go/issues/28827
	// Peer IDs are of string Kind(). We can't use peer IDs here
	// as Go ignores TextMarshaler.
	PeerMap map[string]PinInfoShort `json:"peer_map" codec:"pm,omitempty"`
}

// String returns the string representation of a GlobalPinInfo.
func (gpi GlobalPinInfo) String() string { _ = "STUB: not implemented"; return "" }

// Add adds a PinInfo object to a GlobalPinInfo
func (gpi *GlobalPinInfo) Add(pi PinInfo) { _ = "STUB: not implemented"; return }

// Defined returns if the object is not empty.
func (gpi *GlobalPinInfo) Defined() bool { _ = "STUB: not implemented"; return false }

// Match returns true if one of the statuses in GlobalPinInfo matches
// the given filter.
func (gpi GlobalPinInfo) Match(filter TrackerStatus) bool { _ = "STUB: not implemented"; return false }

// PinInfoShort is a subset of PinInfo which is embedded in GlobalPinInfo
// objects and does not carry redundant information as PinInfo would.
type PinInfoShort struct {
	PeerName      string        `json:"peername" codec:"pn,omitempty"`
	IPFS          peer.ID       `json:"ipfs_peer_id,omitempty" codec:"i,omitempty"`
	IPFSAddresses []Multiaddr   `json:"ipfs_peer_addresses,omitempty" codec:"ia,omitempty"`
	Status        TrackerStatus `json:"status" codec:"st,omitempty"`
	TS            time.Time     `json:"timestamp" codec:"ts,omitempty"`
	Error         string        `json:"error" codec:"e,omitempty"`
	AttemptCount  int           `json:"attempt_count" codec:"a,omitempty"`
	PriorityPin   bool          `json:"priority_pin" codec:"y,omitempty"`
}

// String provides a string representation of PinInfoShort.
func (pis PinInfoShort) String() string { _ = "STUB: not implemented"; return "" }

// PinInfo holds information about local pins. This is used by the Pin
// Trackers.
type PinInfo struct {
	Cid         Cid               `json:"cid" codec:"c"`
	Name        string            `json:"name" codec:"m,omitempty"`
	Peer        peer.ID           `json:"peer" codec:"p,omitempty"`
	Allocations []peer.ID         `json:"allocations" codec:"o,omitempty"`
	Origins     []Multiaddr       `json:"origins" codec:"g,omitempty"`
	Created     time.Time         `json:"created" codec:"t,omitempty"`
	Metadata    map[string]string `json:"metadata" codec:"md,omitempty"`

	PinInfoShort
}

// ToGlobal converts a PinInfo object to a GlobalPinInfo with
// a single peer corresponding to the given PinInfo.
func (pi PinInfo) ToGlobal() GlobalPinInfo { _ = "STUB: not implemented"; return *new(GlobalPinInfo) }

// Defined returns if the PinInfo is not zero.
func (pi PinInfo) Defined() bool { _ = "STUB: not implemented"; return false }

// String provides a string representation of PinInfo.
func (pi PinInfo) String() string { _ = "STUB: not implemented"; return "" }

// Version holds version information
type Version struct {
	Version string `json:"version" codec:"v"`
}

// ConnectGraph holds information about the connectivity of the cluster To
// read, traverse the keys of ClusterLinks.  Each such id is one of the peers
// of the "ClusterID" peer running the query.  ClusterLinks[id] in turn lists
// the ids that peer "id" sees itself connected to.  It is possible that id is
// a peer of ClusterID, but ClusterID can not reach id over rpc, in which case
// ClusterLinks[id] == [], as id's view of its connectivity can not be
// retrieved.
//
// Iff there was an error reading the IPFSID of the peer then id will not be a
// key of ClustertoIPFS or IPFSLinks. Finally iff id is a key of ClustertoIPFS
// then id will be a key of IPFSLinks.  In the event of a SwarmPeers error
// IPFSLinks[id] == [].
type ConnectGraph struct {
	ClusterID    peer.ID           `json:"cluster_id" codec:"id"`
	IDtoPeername map[string]string `json:"id_to_peername" codec:"ip,omitempty"`
	// ipfs to ipfs links
	IPFSLinks map[string][]peer.ID `json:"ipfs_links" codec:"il,omitempty"`
	// cluster to cluster links
	ClusterLinks map[string][]peer.ID `json:"cluster_links" codec:"cl,omitempty"`
	// cluster trust links
	ClusterTrustLinks map[string]bool `json:"cluster_trust_links" codec:"ctl,omitempty"`
	// cluster to ipfs links
	ClustertoIPFS map[string]peer.ID `json:"cluster_to_ipfs" codec:"ci,omitempty"`
}

// Multiaddr is a concrete type to wrap a Multiaddress so that it knows how to
// serialize and deserialize itself.
type Multiaddr struct {
	multiaddr.Multiaddr
}

// NewMultiaddr returns a cluster Multiaddr wrapper creating the
// multiaddr.Multiaddr with the given string.
func NewMultiaddr(mstr string) (Multiaddr, error) {
	_ = "STUB: not implemented"
	return *new(Multiaddr), nil
}

// NewMultiaddrWithValue returns a new cluster Multiaddr wrapper using the
// given multiaddr.Multiaddr.
func NewMultiaddrWithValue(ma multiaddr.Multiaddr) Multiaddr {
	_ = "STUB: not implemented"
	return *new(Multiaddr)
}

// MarshalJSON returns a JSON-formatted multiaddress.
func (maddr Multiaddr) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// UnmarshalJSON parses a cluster Multiaddr from the JSON representation.
func (maddr *Multiaddr) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

// null multiaddresses not allowed

// MarshalBinary returs the bytes of the wrapped multiaddress.
func (maddr Multiaddr) MarshalBinary() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// UnmarshalBinary casts some bytes as a multiaddress wraps it with
// the given cluster Multiaddr.
func (maddr *Multiaddr) UnmarshalBinary(data []byte) error { _ = "STUB: not implemented"; return nil }

// This is super important

// null multiaddresses not allowed

// Value returns the wrapped multiaddr.Multiaddr.
func (maddr Multiaddr) Value() multiaddr.Multiaddr {
	_ = "STUB: not implemented"
	return *

	// ID holds information about the Cluster peer
	new(multiaddr.Multiaddr)
}

type ID struct {
	ID                    peer.ID     `json:"id" codec:"i,omitempty"`
	Addresses             []Multiaddr `json:"addresses" codec:"a,omitempty"`
	ClusterPeers          []peer.ID   `json:"cluster_peers" codec:"cp,omitempty"`
	ClusterPeersAddresses []Multiaddr `json:"cluster_peers_addresses" codec:"cpa,omitempty"`
	Version               string      `json:"version" codec:"v,omitempty"`
	Commit                string      `json:"commit" codec:"c,omitempty"`
	RPCProtocolVersion    protocol.ID `json:"rpc_protocol_version" codec:"rv,omitempty"`
	Error                 string      `json:"error" codec:"e,omitempty"`
	IPFS                  IPFSID      `json:"ipfs,omitempty" codec:"ip,omitempty"`
	Peername              string      `json:"peername" codec:"pn,omitempty"`
	//PublicKey          crypto.PubKey
}

// IPFSID is used to store information about the underlying IPFS daemon
type IPFSID struct {
	ID        peer.ID     `json:"id,omitempty" codec:"i,omitempty"`
	Addresses []Multiaddr `json:"addresses" codec:"a,omitempty"`
	Error     string      `json:"error" codec:"e,omitempty"`
}

// PinType specifies which sort of Pin object we are dealing with.
// In practice, the PinType decides how a Pin object is treated by the
// PinTracker.
// See descriptions above.
// A sharded Pin would look like:
//
// [ Meta ] (not pinned on IPFS, only present in cluster state)
//
//	|
//	v
//
// [ Cluster DAG ] (pinned everywhere in "direct")
//
//	|      ..  |
//	v          v
//
// [Shard1] .. [ShardN] (allocated to peers and pinned with max-depth=1
// | | .. |    | | .. |
// v v .. v    v v .. v
// [][]..[]    [][]..[] Blocks (indirectly pinned on ipfs, not tracked in cluster)
type PinType uint64

// PinType values. See PinType documentation for further explanation.
const (
	// BadType type showing up anywhere indicates a bug
	BadType PinType = 1 << iota
	// DataType is a regular, non-sharded pin. It is pinned recursively.
	// It has no associated reference.
	DataType
	// MetaType tracks the original CID of a sharded DAG. Its Reference
	// points to the Cluster DAG CID.
	MetaType
	// ClusterDAGType pins carry the CID of the root node that points to
	// all the shard-root-nodes of the shards in which a DAG has been
	// divided. Its Reference carries the MetaType CID.
	// ClusterDAGType pins are pinned directly everywhere.
	ClusterDAGType
	// ShardType pins carry the root CID of a shard, which points
	// to individual blocks on the original DAG that the user is adding,
	// which has been sharded.
	// They carry a Reference to the previous shard.
	// ShardTypes are pinned with MaxDepth=1 (root and
	// direct children only).
	ShardType
)

// AllType is a PinType used for filtering all pin types
const AllType PinType = DataType | MetaType | ClusterDAGType | ShardType

// PinTypeFromString is the inverse of String.  It returns the PinType value
// corresponding to the input string
func PinTypeFromString(str string) PinType { _ = "STUB: not implemented"; return *new(PinType) }

// String returns a printable value to identify the PinType
func (pT PinType) String() string { _ = "STUB: not implemented"; return "" }

// MarshalJSON provides json-representation of the pin type.
func (pT PinType) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// UnmarshalJSON provides json-representation of the pin type.
func (pT *PinType) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

var pinOptionsMetaPrefix = "meta-"

// PinMode is a PinOption that indicates how to pin something on IPFS,
// recursively or direct.
type PinMode int

// PinMode values
const (
	PinModeRecursive PinMode = 0
	PinModeDirect    PinMode = 1
)

// PinModeFromString converts a string to PinMode.
func PinModeFromString(s string) PinMode { _ = "STUB: not implemented"; return *new(PinMode) }

// String returns a human-readable value for PinMode.
func (pm PinMode) String() string { _ = "STUB: not implemented"; return "" }

// ToIPFSPinStatus converts a PinMode to IPFSPinStatus.
func (pm PinMode) ToIPFSPinStatus() IPFSPinStatus {
	_ = "STUB: not implemented"
	return *new(IPFSPinStatus)
}

// MarshalJSON converts the PinMode into a readable string in JSON.
func (pm PinMode) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// UnmarshalJSON takes a JSON value and parses it into PinMode.
func (pm *PinMode) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

// ToPinDepth converts the Mode to Depth.
func (pm PinMode) ToPinDepth() PinDepth { _ = "STUB: not implemented"; return *new(PinDepth) }

// PinOptions wraps user-defined options for Pins
type PinOptions struct {
	ReplicationFactorMin int               `json:"replication_factor_min" codec:"rn,omitempty"`
	ReplicationFactorMax int               `json:"replication_factor_max" codec:"rx,omitempty"`
	Name                 string            `json:"name" codec:"n,omitempty"`
	Mode                 PinMode           `json:"mode" codec:"o,omitempty"`
	ShardSize            uint64            `json:"shard_size" codec:"s,omitempty"`
	UserAllocations      []peer.ID         `json:"user_allocations" codec:"ua,omitempty"`
	ExpireAt             time.Time         `json:"expire_at" codec:"e,omitempty"`
	Metadata             map[string]string `json:"metadata" codec:"m,omitempty"`
	PinUpdate            Cid               `json:"pin_update,omitempty" codec:"pu,omitempty"`
	Origins              []Multiaddr       `json:"origins" codec:"g,omitempty"`
}

// Equals returns true if two PinOption objects are equivalent. po and po2 may
// be nil.
func (po PinOptions) Equals(po2 PinOptions) bool { _ = "STUB: not implemented"; return false }

// avoid side effects in the original objects

// deliberately ignore Update

// ToQuery returns the PinOption as query arguments.
func (po PinOptions) ToQuery() (string, error) { _ = "STUB: not implemented"; return "", nil }

// FromQuery is the inverse of ToQuery().
func (po *PinOptions) FromQuery(q url.Values) error { _ = "STUB: not implemented"; return nil }

// override

// PinDepth indicates how deep a pin should be pinned, with
// -1 meaning "to the bottom", or "recursive".
type PinDepth int

// ToPinMode converts PinDepth to PinMode
func (pd PinDepth) ToPinMode() PinMode { _ = "STUB: not implemented"; return *new(PinMode) }

// Pin carries all the information associated to a CID that is pinned
// in IPFS Cluster. It also carries transient information (that may not
// get protobuffed, like UserAllocations).
type Pin struct {
	PinOptions

	Cid Cid `json:"cid" codec:"c"`

	// See PinType comments
	Type PinType `json:"type" codec:"t,omitempty"`

	// The peers to which this pin is allocated
	Allocations []peer.ID `json:"allocations" codec:"a,omitempty"`

	// MaxDepth associated to this pin. -1 means
	// recursive.
	MaxDepth PinDepth `json:"max_depth" codec:"d,omitempty"`

	// We carry a reference CID to this pin. For
	// ClusterDAGs, it is the MetaPin CID. For the
	// MetaPin it is the ClusterDAG CID. For Shards,
	// it is the previous shard CID.
	// When not needed the pointer is nil
	Reference *Cid `json:"reference" codec:"r,omitempty"`

	// The time that the pin was submitted to the consensus layer.
	Timestamp time.Time `json:"timestamp" codec:"i,omitempty"`
}

// String is a string representation of a Pin.
func (pin Pin) String() string { _ = "STUB: not implemented"; return "" }

// IsPinEverywhere returns when the both replication factors are set to -1.
func (pin Pin) IsPinEverywhere() bool { _ = "STUB: not implemented"; return false }

// PinPath is a wrapper for holding pin options and path of the content.
type PinPath struct {
	PinOptions
	Path string `json:"path"`
}

// Defined returns if the path has a value.
func (pp PinPath) Defined() bool { _ = "STUB: not implemented"; return false }

// PinCid is a shortcut to create a Pin only with a Cid.  Default is for pin to
// be recursive and the pin to be of DataType.
func PinCid(c Cid) Pin { _ = "STUB: not implemented"; return *new(Pin) }

// Recursive

// PinWithOpts creates a new Pin calling PinCid(c) and then sets its
// PinOptions fields with the given options. Pin fields that are linked to
// options are set accordingly (MaxDepth from Mode).
func PinWithOpts(c Cid, opts PinOptions) Pin { _ = "STUB: not implemented"; return *new(Pin) }

func convertPinType(t PinType) pb.Pin_PinType {
	_ = "STUB: not implemented"
	return *new(pb.Pin_PinType)
}

// ProtoMarshal marshals this Pin using probobuf.
func (pin Pin) ProtoMarshal() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// Cursory google search says len=0 slices will be
// decoded as null, which is fine.

// Only set the protobuf field with non-zero times.

// Only set the protobuf field with non-zero times.

// Our metadata needs to always be serialized in exactly the same way,
// and that is why we use an array sorted by key and deprecated using
// a protobuf map.

// Metadata:             pin.Metadata,

// Mode:                 pin.Mode,
// UserAllocations:      pin.UserAllocations,

// ProtoUnmarshal unmarshals this fields from protobuf-encoded bytes.
func (pin *Pin) ProtoUnmarshal(data []byte) error { _ = "STUB: not implemented"; return nil }

// pin.UserAllocations = opts.GetUserAllocations()

// Use whatever metadata is available.
//lint:ignore SA1019 we keed to keep backwards compat

// We do not store the PinMode option but we can
// derive it from the MaxDepth setting.

// Equals checks if two pins are the same (with the same allocations).
// If allocations are the same but in different order, they are still
// considered equivalent.
func (pin Pin) Equals(pin2 Pin) bool { _ = "STUB: not implemented"; return false }

// IsRemotePin determines whether a Pin's ReplicationFactor has
// been met, so as to either pin or unpin it from the peer.
func (pin Pin) IsRemotePin(pid peer.ID) bool { _ = "STUB: not implemented"; return false }

// ExpiredAt returns whether the pin has expired at the given time.
func (pin Pin) ExpiredAt(t time.Time) bool { _ = "STUB: not implemented"; return false }

// Defined returns true if this is not a zero-object pin (the CID must be set).
func (pin Pin) Defined() bool { _ = "STUB: not implemented"; return false }

// NodeWithMeta specifies a block of data and a set of optional metadata fields
// carrying information about the encoded ipld node
type NodeWithMeta struct {
	Data    []byte `codec:"d,omitempty"`
	Cid     Cid    `codec:"c,omitempty"`
	CumSize uint64 `codec:"s,omitempty"` // Cumulative size
}

// Size returns how big is the block. It is different from CumSize, which
// records the size of the underlying tree.
func (n *NodeWithMeta) Size() uint64 { _ = "STUB: not implemented"; return 0 }

// MetricsSet is a map to carry slices of metrics indexed by type.
type MetricsSet map[string][]Metric

// Metric transports information about a peer.ID. It is used to decide
// pin allocations by a PinAllocator. IPFS cluster is agnostic to
// the Value, which should be interpreted by the PinAllocator.
// The ReceivedAt value is a timestamp representing when a peer has received
// the metric value.
type Metric struct {
	Name          string  `json:"name" codec:"n,omitempty"`
	Peer          peer.ID `json:"peer" codec:"p,omitempty"`
	Value         string  `json:"value" codec:"v,omitempty"`
	Expire        int64   `json:"expire" codec:"e,omitempty"`
	Valid         bool    `json:"valid" codec:"d,omitempty"`
	Weight        int64   `json:"weight" codec:"w,omitempty"`
	Partitionable bool    `json:"partitionable" codec:"o,omitempty"`
	ReceivedAt    int64   `json:"received_at" codec:"t,omitempty"` // ReceivedAt contains a UnixNano timestamp
}

func (m Metric) String() string { _ = "STUB: not implemented"; return "" }

// Defined returns true if the metric name is set.
func (m Metric) Defined() bool { _ = "STUB: not implemented"; return false }

// SetTTL sets Metric to expire after the given time.Duration
func (m *Metric) SetTTL(d time.Duration) { _ = "STUB: not implemented"; return }

// GetTTL returns the time left before the Metric expires
func (m Metric) GetTTL() time.Duration { _ = "STUB: not implemented"; return *new(time.Duration) }

// Expired returns if the Metric has expired
func (m Metric) Expired() bool { _ = "STUB: not implemented"; return false }

// Discard returns if the metric not valid or has expired
func (m Metric) Discard() bool { _ = "STUB: not implemented"; return false }

// GetWeight returns the weight of the metric.
// This is for compatibility.
func (m Metric) GetWeight() int64 {
	_ = "STUB: not implemented"

	// MetricSlice is a sortable Metric array.
	return 0
}

type MetricSlice []Metric

func (es MetricSlice) Len() int           { _ = "STUB: not implemented"; return 0 }
func (es MetricSlice) Swap(i, j int)      { _ = "STUB: not implemented"; return }
func (es MetricSlice) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

// Alert carries alerting information about a peer.
type Alert struct {
	Metric
	TriggeredAt time.Time `json:"triggered_at" codec:"r,omitempty"`
}

// Bandwidth carries bandwidth information per libp2p/core/metrics.Stats.
type Bandwidth struct {
	TotalIn  int64   `json:"total_in" codec:"i,omitempty"`
	TotalOut int64   `json:"total_out" codec:"o,omitempty"`
	RateIn   float64 `json:"rate_in" codec:"ri,omitempty"`
	RateOut  float64 `json:"rate_out" codec:"ro,omitempty"`
}

// BandwidthByProtocol carries Bandwidth information indexed by libp2p
// protocol tag.
type BandwidthByProtocol map[protocol.ID]Bandwidth

// Error can be used by APIs to return errors.
type Error struct {
	Code    int    `json:"code" codec:"o,omitempty"`
	Message string `json:"message" codec:"m,omitempty"`
}

// Error implements the error interface and returns the error's message.
func (e Error) Error() string { _ = "STUB: not implemented"; return "" }

// IPFSRepoStat wraps information about the IPFS repository.
type IPFSRepoStat struct {
	RepoSize   uint64 `codec:"r,omitempty"`
	StorageMax uint64 `codec:"s, omitempty"`
}

// IPFSRepoGC represents the streaming response sent from repo gc API of IPFS.
type IPFSRepoGC struct {
	Key   Cid    `json:"key,omitempty" codec:"k,omitempty"`
	Error string `json:"error,omitempty" codec:"e,omitempty"`
}

// RepoGC contains garbage collected CIDs from a cluster peer's IPFS daemon.
type RepoGC struct {
	Peer     peer.ID      `json:"peer" codec:"p,omitempty"` // the Cluster peer ID
	Peername string       `json:"peername" codec:"pn,omitempty"`
	Keys     []IPFSRepoGC `json:"keys" codec:"k"`
	Error    string       `json:"error,omitempty" codec:"e,omitempty"`
}

// GlobalRepoGC contains cluster-wide information about garbage collected CIDs
// from IPFS.
type GlobalRepoGC struct {
	PeerMap map[string]RepoGC `json:"peer_map" codec:"pm,omitempty"`
}
