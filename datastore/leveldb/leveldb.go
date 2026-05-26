// Package leveldb provides a configurable LevelDB go-datastore for use with
// IPFS Cluster.
package leveldb

import (
	ds "github.com/ipfs/go-datastore"
)

// New returns a LevelDB datastore configured with the given
// configuration.
func New(cfg *Config) (ds.Datastore, error) {
	_ = "STUB: not implemented"
	return *new(ds.Datastore), nil
}

// Cleanup deletes the leveldb datastore.
func Cleanup(cfg *Config) error { _ = "STUB: not implemented"; return nil }
