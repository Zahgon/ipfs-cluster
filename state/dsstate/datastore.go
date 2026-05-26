// Package dsstate implements the IPFS Cluster state interface using
// an underlying go-datastore.
package dsstate

import (
	"context"
	"io"

	"github.com/ipfs-cluster/ipfs-cluster/api"
	"github.com/ipfs-cluster/ipfs-cluster/state"

	ds "github.com/ipfs/go-datastore"
	logging "github.com/ipfs/go-log/v2"
	codec "github.com/ugorji/go/codec"
)

var _ state.State = (*State)(nil)
var _ state.BatchingState = (*BatchingState)(nil)

var logger = logging.Logger("dsstate")

// State implements the IPFS Cluster "state" interface by wrapping
// a go-datastore and choosing how api.Pin objects are stored
// in it. It also provides serialization methods for the whole
// state which are datastore-independent.
type State struct {
	dsRead      ds.Read
	dsWrite     ds.Write
	codecHandle codec.Handle
	namespace   ds.Key
	// version     int

	totalPins int64
}

// DefaultHandle returns the codec handler of choice (Msgpack).
func DefaultHandle() codec.Handle { _ = "STUB: not implemented"; return *new(codec.Handle) }

// New returns a new state using the given datastore.
//
// All keys are namespaced with the given string when written. Thus the same
// go-datastore can be sharded for different uses.
//
// The Handle controls options for the serialization of the full state
// (marshaling/unmarshaling).
func New(ctx context.Context, dstore ds.Datastore, namespace string, handle codec.Handle) (*State, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Add adds a new Pin or replaces an existing one.
func (st *State) Add(ctx context.Context, c api.Pin) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// Rm removes an existing Pin. It is a no-op when the
// item does not exist.
func (st *State) Rm(ctx context.Context, c api.Cid) error { _ = "STUB: not implemented"; return nil }

// Get returns a Pin from the store and whether it
// was present. When not present, a default pin
// is returned.
func (st *State) Get(ctx context.Context, c api.Cid) (api.Pin, error) {
	_ = "STUB: not implemented"
	return *new(api.Pin), nil
}

// Has returns whether a Cid is stored.
func (st *State) Has(ctx context.Context, c api.Cid) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// List sends all the pins on the pinset on the given channel.
// Returns and closes channel when done.
func (st *State) List(ctx context.Context, out chan<- api.Pin) error {
	_ = "STUB: not implemented"
	return nil
}

// Abort if we shutdown.

// Migrate migrates an older state version to the current one.
// This is a no-op for now.
func (st *State) Migrate(ctx context.Context, r io.Reader) error {
	_ = "STUB: not implemented"
	return nil
}

type serialEntry struct {
	Key   string `codec:"k"`
	Value []byte `codec:"v"`
}

// Marshal dumps the state to a writer. It does this by encoding every
// key/value in the store. The keys are stored without the namespace part to
// reduce the size of the snapshot.
func (st *State) Marshal(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// reduce snapshot size by not storing the prefix

// Unmarshal reads and parses a previous dump of the state.
// All the parsed key/values are added to the store. As of now,
// Unmarshal does not empty the existing store from any values
// before unmarshaling from the given reader.
func (st *State) Unmarshal(r io.Reader) error { _ = "STUB: not implemented"; return nil }

// used to be on go-ipfs-ds-help
func cidToDsKey(c api.Cid) ds.Key { _ = "STUB: not implemented"; return *new(ds.Key) }

// used to be on go-ipfs-ds-help
func dsKeyToCid(k ds.Key) (api.Cid, error) { _ = "STUB: not implemented"; return *new(api.Cid), nil }

// convert Cid to /namespace/cid1Key
func (st *State) key(c api.Cid) ds.Key { _ = "STUB: not implemented"; return *new(ds.Key) }

// convert /namespace/cidKey to Cid
func (st *State) unkey(k ds.Key) (api.Cid, error) {
	_ = "STUB: not implemented"
	return *new(api.Cid), nil
}

// this decides how a Pin object is serialized to be stored in the
// datastore. Changing this may require a migration!
func (st *State) serializePin(c api.Pin) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil,

		// this deserializes a Pin object from the datastore. It should be
		// the exact opposite from serializePin.
		nil
}

func (st *State) deserializePin(c api.Cid, buf []byte) (api.Pin, error) {
	_ = "STUB: not implemented"
	return *new(api.Pin), nil
}

// BatchingState implements the IPFS Cluster "state" interface by wrapping a
// batching go-datastore. All writes are batched and only written disk
// when Commit() is called.
type BatchingState struct {
	*State
	batch ds.Batch
}

// NewBatching returns a new batching statate using the given datastore.
//
// All keys are namespaced with the given string when written. Thus the same
// go-datastore can be sharded for different uses.
//
// The Handle controls options for the serialization of the full state
// (marshaling/unmarshaling).
func NewBatching(ctx context.Context, dstore ds.Batching, namespace string, handle codec.Handle) (*BatchingState, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Commit persists the batched write operations.
func (bst *BatchingState) Commit(ctx context.Context) error { _ = "STUB: not implemented"; return nil }
