package cmdutils

import (
	"io"

	"github.com/ipfs-cluster/ipfs-cluster/api"
	"github.com/ipfs-cluster/ipfs-cluster/config"
	"github.com/ipfs-cluster/ipfs-cluster/state"

	ds "github.com/ipfs/go-datastore"
)

// StateManager is the interface that allows to import, export and clean
// different cluster states depending on the consensus component used.
type StateManager interface {
	ImportState(io.Reader, api.PinOptions) error
	ExportState(io.Writer) error
	GetStore() (ds.Datastore, error)
	GetOfflineState(ds.Datastore) (state.State, error)
	Clean() error
}

// NewStateManager returns an state manager implementation for the given
// consensus ("raft" or "crdt"). It will need initialized configs.
func NewStateManager(consensus string, datastore string, ident *config.Identity, cfgs *Configs) (StateManager, error) {
	_ = "STUB: not implemented"
	return *new(StateManager), nil
}

// NewStateManagerWithHelper returns a state manager initialized using the
// configuration and identity provided by the given config helper.
func NewStateManagerWithHelper(cfgHelper *ConfigHelper) (StateManager, error) {
	_ = "STUB: not implemented"
	return *new(StateManager), nil
}

type raftStateManager struct {
	ident *config.Identity
	cfgs  *Configs
}

func (raftsm *raftStateManager) GetStore() (ds.Datastore, error) {
	_ = "STUB: not implemented"
	return *new(ds.Datastore), nil
}

func (raftsm *raftStateManager) GetOfflineState(store ds.Datastore) (state.State, error) {
	_ = "STUB: not implemented"
	return *new(state.State), nil
}

func (raftsm *raftStateManager) ImportState(r io.Reader, opts api.PinOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (raftsm *raftStateManager) ExportState(w io.Writer) error {
	_ = "STUB: not implemented"
	return nil
}

func (raftsm *raftStateManager) Clean() error { _ = "STUB: not implemented"; return nil }

type crdtStateManager struct {
	cfgs      *Configs
	datastore string
}

func (crdtsm *crdtStateManager) GetStore() (ds.Datastore, error) {
	_ = "STUB: not implemented"
	return *new(ds.Datastore), nil
}

func (crdtsm *crdtStateManager) GetOfflineState(store ds.Datastore) (state.State, error) {
	_ = "STUB: not implemented"
	return *new(state.State), nil
}

func (crdtsm *crdtStateManager) ImportState(r io.Reader, opts api.PinOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (crdtsm *crdtStateManager) ExportState(w io.Writer) error {
	_ = "STUB: not implemented"
	return nil
}

func (crdtsm *crdtStateManager) Clean() error { _ = "STUB: not implemented"; return nil }

func importState(r io.Reader, st state.State, opts api.PinOptions) error {
	_ = "STUB: not implemented"
	return nil
}

// We are injecting directly to the state.
// UserAllocation option is not stored in the state.
// We need to set Allocations directly.

// ExportState saves a json representation of a state
func exportState(w io.Writer, st state.State) error { _ = "STUB: not implemented"; return nil }
