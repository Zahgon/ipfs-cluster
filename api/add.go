package api

import (
	"net/url"

	peer "github.com/libp2p/go-libp2p/core/peer"
)

// DefaultShardSize is the shard size for params objects created with DefaultParams().
var DefaultShardSize = uint64(100 * 1024 * 1024) // 100 MB

// AddedOutput carries information for displaying the standard ipfs output
// indicating a node of a file has been added.
type AddedOutput struct {
	Name        string    `json:"name" codec:"n,omitempty"`
	Cid         Cid       `json:"cid" codec:"c"`
	Bytes       uint64    `json:"bytes,omitempty" codec:"b,omitempty"`
	Size        uint64    `json:"size,omitempty" codec:"s,omitempty"`
	Allocations []peer.ID `json:"allocations,omitempty" codec:"a,omitempty"`
}

// IPFSAddParams groups options specific to the ipfs-adder, which builds
// UnixFS dags with the input files. This struct is embedded in AddParams.
type IPFSAddParams struct {
	Layout     string
	Chunker    string
	RawLeaves  bool
	Progress   bool
	CidVersion int
	HashFun    string
	NoCopy     bool
}

// AddParams contains all of the configurable parameters needed to specify the
// importing process of a file being added to an ipfs-cluster
type AddParams struct {
	PinOptions

	Local          bool
	Recursive      bool
	Hidden         bool
	Wrap           bool
	Shard          bool
	StreamChannels bool
	Format         string // selects with adder
	NoPin          bool

	IPFSAddParams
}

// DefaultAddParams returns a AddParams object with standard defaults
func DefaultAddParams() AddParams { _ = "STUB: not implemented"; return *new(AddParams) }

// corresponds to balanced layout

func parseBoolParam(q url.Values, name string, dest *bool) error {
	_ = "STUB: not implemented"
	return nil
}

func parseIntParam(q url.Values, name string, dest *int) error {
	_ = "STUB: not implemented"
	return nil
}

// AddParamsFromQuery parses the AddParams object from
// a URL.Query().
func AddParamsFromQuery(query url.Values) (AddParams, error) {
	_ = "STUB: not implemented"
	return *new(AddParams), nil
}

// hardcode as does not make sense for adding

// nothing

// This mimics go-ipfs behavior.

// If the raw-leaves param is empty, the default RawLeaves value will
// take place (which may be true or false depending on
// CidVersion). Otherwise, it will be explicitly set.

// ToQueryString returns a url query string (key=value&key2=value2&...)
func (p AddParams) ToQueryString() (string, error) { _ = "STUB: not implemented"; return "", nil }

// Equals checks if p equals p2.
func (p AddParams) Equals(p2 AddParams) bool { _ = "STUB: not implemented"; return false }
