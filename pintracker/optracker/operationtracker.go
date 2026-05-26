// Package optracker implements functionality to track the status of pin and
// operations as needed by implementations of the pintracker component.
// It particularly allows to obtain status information for a given Cid,
// to skip re-tracking already ongoing operations, or to cancel ongoing
// operations when opposing ones arrive.
package optracker

import (
	"context"
	"sync"

	"github.com/ipfs-cluster/ipfs-cluster/api"

	logging "github.com/ipfs/go-log/v2"
	peer "github.com/libp2p/go-libp2p/core/peer"
)

var logger = logging.Logger("optracker")

// OperationTracker tracks and manages all inflight Operations.
type OperationTracker struct {
	// struct alignment. This fields must be upfront!
	pinningCount   int64
	pinErrorCount  int64
	pinQueuedCount int64

	ctx      context.Context // parent context for all ops
	pid      peer.ID
	peerName string

	mu         sync.RWMutex
	operations map[api.Cid]*Operation
}

func (opt *OperationTracker) String() string { _ = "STUB: not implemented"; return "" }

// NewOperationTracker creates a new OperationTracker.
func NewOperationTracker(ctx context.Context, pid peer.ID, peerName string) *OperationTracker {
	_ = "STUB: not implemented"
	return nil
}

// TrackNewOperation will create, track and return a new operation unless
// one already exists to do the same thing, in which case nil is returned.
//
// If an operation exists it is of different type, it is
// canceled and the new one replaces it in the tracker.
func (opt *OperationTracker) TrackNewOperation(ctx context.Context, pin api.Pin, typ OperationType, ph Phase) *Operation {
	_ = "STUB: not implemented"
	return nil
}

// operation exists for the CID

// an ongoing operation of the same
// type. i.e. pinning, or queued.  Update the pin
// object though, as it may have different options.

// i.e. operations in error phase
// i.e. pin operations that need to be canceled for unpinning

// cancel ongoing operation and replace it

// IMPORTANT: the operations must have the OperationTracker context,
// as otherwise their context would be canceled after being added.

// Carry over the attempt count when doing an operation of the
// same type.  The old operation exists and was canceled.
// carry the count

// Clean deletes an operation from the tracker if it is the one we are tracking
// (compares pointers).
func (opt *OperationTracker) Clean(ctx context.Context, op *Operation) {
	_ = "STUB: not implemented"
	return
}

// same pointer

// Status returns the TrackerStatus associated to the last operation known
// with the given Cid. It returns false if we are not tracking any operation
// for the given Cid.
func (opt *OperationTracker) Status(ctx context.Context, c api.Cid) (api.TrackerStatus, bool) {
	_ = "STUB: not implemented"
	return *new(api.TrackerStatus), false
}

// SetError transitions an operation for a Cid into PhaseError if its Status
// is PhaseDone. Any other phases are considered in-flight and not touched.
// For things already in error, the error message is updated.
// Remote pins are ignored too.
// Only used in tests right now.
func (opt *OperationTracker) SetError(ctx context.Context, c api.Cid, err error) {
	_ = "STUB: not implemented"
	return
}

func (opt *OperationTracker) unsafePinInfo(ctx context.Context, op *Operation, ipfs api.IPFSID) api.PinInfo {
	_ = "STUB: not implemented"
	return *new(api.PinInfo)
}

//Created:  0,

// Get returns a PinInfo object for Cid.
func (opt *OperationTracker) Get(ctx context.Context, c api.Cid, ipfs api.IPFSID) api.PinInfo {
	_ = "STUB: not implemented"
	return *new(api.PinInfo)
}

// GetExists returns a PinInfo object for a Cid only if there exists
// an associated Operation.
func (opt *OperationTracker) GetExists(ctx context.Context, c api.Cid, ipfs api.IPFSID) (api.PinInfo, bool) {
	_ = "STUB: not implemented"
	return *new(api.PinInfo), false
}

// GetAll returns PinInfo objects for all known operations.
func (opt *OperationTracker) GetAll(ctx context.Context, ipfs api.IPFSID) []api.PinInfo {
	_ = "STUB: not implemented"
	return nil
}

// GetAllChannel returns all known operations that match the filter on the
// provided channel. Blocks until done.
func (opt *OperationTracker) GetAllChannel(ctx context.Context, filter api.TrackerStatus, ipfs api.IPFSID, out chan<- api.PinInfo) error {
	_ = "STUB: not implemented"
	return nil
}

// CleanAllDone deletes any operation from the tracker that is in PhaseDone.
func (opt *OperationTracker) CleanAllDone(ctx context.Context) { _ = "STUB: not implemented"; return }

// OpContext gets the context of an operation, if any.
func (opt *OperationTracker) OpContext(ctx context.Context, c api.Cid) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// Filter returns a slice of api.PinInfos that had associated
// Operations that matched the provided filter. Note, only supports
// filters of type OperationType or Phase, any other type
// will result in a nil slice being returned.
func (opt *OperationTracker) Filter(ctx context.Context, ipfs api.IPFSID, filters ...interface{}) []api.PinInfo {
	_ = "STUB: not implemented"
	return nil
}

// filterOps returns a slice that only contains operations
// with the matching filter. Note, only supports
// filters of type OperationType or Phase, any other type
// will result in a nil slice being returned.
// Only used in tests right now.
func (opt *OperationTracker) filterOps(ctx context.Context, filters ...interface{}) []*Operation {
	_ = "STUB: not implemented"
	return nil
}

func filterOpsMap(ctx context.Context, ops map[api.Cid]*Operation, filters []interface{}) map[api.Cid]*Operation {
	_ = "STUB: not implemented"
	return nil
}

func filter(ctx context.Context, in, out map[api.Cid]*Operation, filter interface{}) {
	_ = "STUB: not implemented"
	return
}

func initializeMetrics(ctx context.Context) { _ = "STUB: not implemented"; return }

func (opt *OperationTracker) recordMetricUnsafe(op *Operation, val int64) {
	_ = "STUB: not implemented"
	return
}

// we have no metric to log anything

func (opt *OperationTracker) recordMetric(op *Operation, val int64) {
	_ = "STUB: not implemented"
	return
}

// PinQueueSize returns the current number of items queued to pin.
func (opt *OperationTracker) PinQueueSize() int64 { _ = "STUB: not implemented"; return 0 }
