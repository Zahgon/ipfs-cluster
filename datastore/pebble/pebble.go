// Package pebble provides a configurable Pebble database backend for use with
// IPFS Cluster.
package pebble

import (
	ds "github.com/ipfs/go-datastore"
	logging "github.com/ipfs/go-log/v2"
)

var logger = logging.Logger("pebble")

// New returns a Pebble datastore configured with the given
// configuration.
func New(cfg *Config) (ds.Datastore, error) {
	_ = "STUB: not implemented"
	return *new(ds.Datastore), nil
}

// Deal with Pebble updates... user should try to be up to date with
// latest Pebble table formats.

// Calling regularly DB's DiskUsage is a way to printout debug
// database statistics.

// Cleanup deletes the pebble datastore.
func Cleanup(cfg *Config) error { _ = "STUB: not implemented"; return nil }
