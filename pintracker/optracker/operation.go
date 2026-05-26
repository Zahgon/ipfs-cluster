package optracker

import (
	"context"
	"sync"
	"time"

	"github.com/ipfs-cluster/ipfs-cluster/api"
)

//go:generate stringer -type=OperationType

// OperationType represents the kinds of operations that the PinTracker
// performs and the operationTracker tracks the status of.
type OperationType int

const (
	// OperationUnknown represents an unknown operation.
	OperationUnknown OperationType = iota
	// OperationPin represents a pin operation.
	OperationPin
	// OperationUnpin represents an unpin operation.
	OperationUnpin
	// OperationRemote represents an noop operation
	OperationRemote
	// OperationShard represents a meta pin. We don't
	// pin these.
	OperationShard
)

//go:generate stringer -type=Phase

// Phase represents the multiple phase that an operation can be in.
type Phase int

const (
	// PhaseError represents an error state.
	PhaseError Phase = iota
	// PhaseQueued represents the queued phase of an operation.
	PhaseQueued
	// PhaseInProgress represents the operation as in progress.
	PhaseInProgress
	// PhaseDone represents the operation once finished.
	PhaseDone
)

// Operation represents an ongoing operation involving a
// particular Cid. It provides the type and phase of operation
// and a way to mark the operation finished (also used to cancel).
type Operation struct {
	ctx    context.Context
	cancel func()

	tracker *OperationTracker

	opType OperationType
	pin    api.Pin

	// RW fields
	mu           sync.RWMutex
	phase        Phase
	attemptCount int
	priority     bool
	error        string
	ts           time.Time
}

// newOperation creates a new Operation.
func newOperation(ctx context.Context, pin api.Pin, typ OperationType, ph Phase, tracker *OperationTracker) *Operation {
	_ = "STUB: not implemented"
	return nil
}

// String returns a string representation of an Operation.
func (op *Operation) String() string { _ = "STUB: not implemented"; return "" }

// Cid returns the Cid associated to this operation.
func (op *Operation) Cid() api.Cid {
	_ = "STUB: not implemented"

	// Context returns the context associated to this operation.
	return *new(api.Cid)
}

func (op *Operation) Context() context.Context {
	_ = "STUB: not implemented"

	// Cancel will cancel the context associated to this operation.
	return *new(context.Context)
}

func (op *Operation) Cancel() {
	_ = "STUB: not implemented"

	// Phase returns the Phase.
	return
}

func (op *Operation) Phase() Phase { _ = "STUB: not implemented"; return *new(Phase) }

// SetPhase changes the Phase and updates the timestamp.
func (op *Operation) SetPhase(ph Phase) { _ = "STUB: not implemented"; return }

// AttemptCount returns the number of times that this operation has been in
// progress.
func (op *Operation) AttemptCount() int { _ = "STUB: not implemented"; return 0 }

// IncAttempt does a plus-one on the AttemptCount.
func (op *Operation) IncAttempt() { _ = "STUB: not implemented"; return }

// PriorityPin returns true if the pin has been marked as priority pin.
func (op *Operation) PriorityPin() bool { _ = "STUB: not implemented"; return false }

// SetPriorityPin returns true if the pin has been marked as priority pin.
func (op *Operation) SetPriorityPin(p bool) { _ = "STUB: not implemented"; return }

// Error returns any error message attached to the operation.
func (op *Operation) Error() string { _ = "STUB: not implemented"; return "" }

// SetError sets the phase to PhaseError along with
// an error message. It updates the timestamp.
func (op *Operation) SetError(err error) { _ = "STUB: not implemented"; return }

// Type returns the operation Type.
func (op *Operation) Type() OperationType {
	_ = "STUB: not implemented"

	// Pin returns the Pin object associated to the operation.
	return *new(OperationType)
}

func (op *Operation) Pin() api.Pin {
	_ = "STUB: not implemented"

	// Timestamp returns the time when this operation was
	// last modified (phase changed, error was set...).
	return *new(api.Pin)
}

func (op *Operation) Timestamp() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

// Canceled returns whether the context for this
// operation has been canceled.
func (op *Operation) Canceled() bool { _ = "STUB: not implemented"; return false }

// ToTrackerStatus returns an api.TrackerStatus reflecting
// the current status of this operation. It's a translation
// from the Type and the Phase.
func (op *Operation) ToTrackerStatus() api.TrackerStatus {
	_ = "STUB: not implemented"
	return *new(api.TrackerStatus)
}

// TrackerStatusToOperationPhase takes an api.TrackerStatus and
// converts it to an OpType and Phase.
func TrackerStatusToOperationPhase(status api.TrackerStatus) (OperationType, Phase) {
	_ = "STUB: not implemented"
	return *new(OperationType), *new(Phase)
}
