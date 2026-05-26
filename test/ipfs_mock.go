package test

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"sync"
	"testing"

	"github.com/ipfs-cluster/ipfs-cluster/api"
	"github.com/ipfs-cluster/ipfs-cluster/state"

	cid "github.com/ipfs/go-cid"
)

// Some values used by the ipfs mock
const (
	IpfsCustomHeaderName  = "X-Custom-Header"
	IpfsTimeHeaderName    = "X-Time-Now"
	IpfsCustomHeaderValue = "42"
	IpfsACAOrigin         = "myorigin"
	IpfsErrFromNotPinned  = "'from' cid was not recursively pinned already"
)

// IpfsMock is an ipfs daemon mock which should sustain the functionality used by ipfscluster.
type IpfsMock struct {
	server     *httptest.Server
	Addr       string
	Port       int
	pinMap     state.State
	BlockStore map[string][]byte
	reqCounter chan string

	reqCountsMux sync.Mutex // guards access to reqCounts
	reqCounts    map[string]int

	closeMux sync.Mutex
	closed   bool
}

type mockPinResp struct {
	Pins     []string
	Progress int `json:",omitempty"`
}

type mockPinType struct {
	Type string
}

type mockPinLsAllResp struct {
	Keys map[string]mockPinType
}

type ipfsErr struct {
	Code    int
	Message string
}

type mockIDResp struct {
	ID        string
	Addresses []string
}

type mockRepoStatResp struct {
	RepoSize   uint64
	NumObjects uint64
	StorageMax uint64
}

type mockConfigResp struct {
	Datastore struct {
		StorageMax string
	}
}

type mockRefsResp struct {
	Ref string
	Err string
}

type mockSwarmPeersResp struct {
	Peers []mockIpfsPeer
}

type mockIpfsPeer struct {
	Peer string
}

type mockBlockPutResp struct {
	Key string
}

type mockDagPutResp struct {
	Cid cid.Cid
}

type mockRepoGCResp struct {
	Key   cid.Cid `json:",omitempty"`
	Error string  `json:",omitempty"`
}

// NewIpfsMock returns a new mock.
func NewIpfsMock(t *testing.T) *IpfsMock { _ = "STUB: not implemented"; return nil }

// because IPFS does it, even if for no reason.

func (m *IpfsMock) countRequests() { _ = "STUB: not implemented"; return }

// GetCount allows to get the number of times and endpoint was called.
func (m *IpfsMock) GetCount(path string) int { _ = "STUB: not implemented"; return 0 }

// FIXME: what if IPFS API changes?
func (m *IpfsMock) handler(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// this a v1 cid. Do not return default-base32 but base58btc encoding of it

// this a v1 cid. Do not return default-base32 but base58btc encoding of it

// Get the data and retun the hash

// Parse cid from data and format and add to mock block-store

// DAG-put is a fake implementation as we are not going to
// parse the input and we are just going to hash it and return
// a response.

// Get the data and retun the hash

// Parse cid from data and format and add to mock block-store

// It assumes `/repo/gc` with parameter `stream-errors=true`

// 10 GB

// Close closes the mock server. It's important to call after each test or
// the listeners are left hanging around.
func (m *IpfsMock) Close() { _ = "STUB: not implemented"; return }

// extractCid extracts the cid argument from a url.URL, either via
// the query string parameters or from the url path itself.
func extractCid(u *url.URL) (string, bool) { _ = "STUB: not implemented"; return "", false }

func extractMode(u *url.URL) api.PinMode { _ = "STUB: not implemented"; return *new(api.PinMode) }
