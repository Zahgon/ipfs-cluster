package raft

import (
	"go.opencensus.io/trace"

	"github.com/ipfs-cluster/ipfs-cluster/api"

	consensus "github.com/libp2p/go-libp2p-consensus"
)

// Type of consensus operation
const (
	LogOpPin = iota + 1
	LogOpUnpin
)

// LogOpType expresses the type of a consensus Operation
type LogOpType int

// LogOp represents an operation for the OpLogConsensus system.
// It implements the consensus.Op interface and it is used by the
// Consensus component.
type LogOp struct {
	SpanCtx   trace.SpanContext `codec:"s,omitempty"`
	TagCtx    []byte            `codec:"t,omitempty"`
	Cid       api.Pin           `codec:"c,omitempty"`
	Type      LogOpType         `codec:"p,omitempty"`
	consensus *Consensus        `codec:"-"`
	tracing   bool              `codec:"-"`
}

// ApplyTo applies the operation to the State
func (op *LogOp) ApplyTo(cstate consensus.State) (consensus.State, error) {
	_ = "STUB: not implemented"
	return *new(consensus.State), nil
}

// Should never be here

// Async, we let the PinTracker take care of any problems

// Async, we let the PinTracker take care of any problems

// We failed to apply the operation to the state
// and therefore we need to request a rollback to the
// cluster to the previous state. This operation can only be performed
// by the cluster leader.
