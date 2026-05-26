// Package badger3 provides a configurable BadgerDB v3 go-datastore for use with
// IPFS Cluster.
package badger3

import (
	ds "github.com/ipfs/go-datastore"
	logging "github.com/ipfs/go-log/v2"
)

var logger = logging.Logger("badger3")

// New returns a BadgerDB datastore configured with the given
// configuration.
func New(cfg *Config) (ds.Datastore, error) {
	_ = "STUB: not implemented"
	return *new(ds.Datastore), nil
}

// Cleanup deletes the badger datastore.
func Cleanup(cfg *Config) error { _ = "STUB: not implemented"; return nil }
