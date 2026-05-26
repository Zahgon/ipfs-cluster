package main

import (
	"io"
)

// lock logic heavily inspired by go-ipfs/repo/fsrepo/lock/lock.go

// The name of the file used for locking
const lockFileName = "cluster.lock"

var locker *lock

// lock helps to coordinate proceeds via a lock file
type lock struct {
	lockCloser io.Closer
	path       string
}

func (l *lock) lock() { _ = "STUB: not implemented"; return }

// we should have a config folder whenever we try to lock

// set the lock file within this function

func (l *lock) tryUnlock() error {
	_ = "STUB: not implemented"
	// Noop in the uninitialized case
	return nil
}
