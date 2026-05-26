// Package inmem provides a in-memory thread-safe datastore for use with
// Cluster.
package inmem

import (
	ds "github.com/ipfs/go-datastore"
)

// New returns a new thread-safe in-memory go-datastore.
func New() ds.Datastore { _ = "STUB: not implemented"; return *new(ds.Datastore) }
