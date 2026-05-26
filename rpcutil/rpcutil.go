// Package rpcutil provides utility methods to perform go-libp2p-gorpc calls,
// particularly gorpc.MultiCall().
package rpcutil

import (
	"context"
	"time"

	"github.com/ipfs-cluster/ipfs-cluster/api"
)

// CtxsWithTimeout returns n contexts, derived from the given parent
// using the given timeout.
func CtxsWithTimeout(
	parent context.Context,
	n int,
	timeout time.Duration,
) ([]context.Context, []context.CancelFunc) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CtxsWithCancel returns n cancellable contexts, derived from the given parent.
func CtxsWithCancel(
	parent context.Context,
	n int,
) ([]context.Context, []context.CancelFunc) {
	_ = "STUB: not implemented"
	return nil, nil
}

// MultiCancel calls all the provided CancelFuncs. It
// is useful with "defer Multicancel()"
func MultiCancel(cancels []context.CancelFunc) { _ = "STUB: not implemented"; return }

// The copy functions below are used in calls to Cluster.multiRPC()

// // CopyPIDsToIfaces converts a peer.ID slice to an empty interface
// // slice using pointers to each elements of the original slice.
// // Useful to handle gorpc.MultiCall() replies.
// func CopyPIDsToIfaces(in []peer.ID) []interface{} {
// 	ifaces := make([]interface{}, len(in))
// 	for i := range in {
// 		ifaces[i] = &in[i]
// 	}
// 	return ifaces
// }

// CopyIDsToIfaces converts an api.ID slice to an empty interface
// slice using pointers to each elements of the original slice.
// Useful to handle gorpc.MultiCall() replies.
func CopyIDsToIfaces(in []api.ID) []interface{} { _ = "STUB: not implemented"; return nil }

// CopyIDSliceToIfaces converts an api.ID slice of slices
// to an empty interface slice using pointers to each elements of the
// original slice. Useful to handle gorpc.MultiCall() replies.
func CopyIDSliceToIfaces(in [][]api.ID) []interface{} { _ = "STUB: not implemented"; return nil }

// CopyPinInfoToIfaces converts an api.PinInfo slice to
// an empty interface slice using pointers to each elements of
// the original slice. Useful to handle gorpc.MultiCall() replies.
func CopyPinInfoToIfaces(in []api.PinInfo) []interface{} { _ = "STUB: not implemented"; return nil }

// CopyPinInfoSliceToIfaces converts an api.PinInfo slice of slices
// to an empty interface slice using pointers to each elements of the original
// slice. Useful to handle gorpc.MultiCall() replies.
func CopyPinInfoSliceToIfaces(in [][]api.PinInfo) []interface{} {
	_ = "STUB: not implemented"
	return nil
}

// CopyRepoGCSliceToIfaces converts an api.RepoGC slice to
// an empty interface slice using pointers to each elements of
// the original slice. Useful to handle gorpc.MultiCall() replies.
func CopyRepoGCSliceToIfaces(in []api.RepoGC) []interface{} { _ = "STUB: not implemented"; return nil }

// CopyEmptyStructToIfaces converts an empty struct slice to an empty interface
// slice using pointers to each elements of the original slice.
// Useful to handle gorpc.MultiCall() replies.
func CopyEmptyStructToIfaces(in []struct{}) []interface{} { _ = "STUB: not implemented"; return nil }

// RPCDiscardReplies returns a []interface{} slice made from a []struct{}
// slice of then given length. Useful for RPC methods which have no response
// types (so they use empty structs).
func RPCDiscardReplies(n int) []interface{} { _ = "STUB: not implemented"; return nil }

// CheckErrs returns nil if all the errors in a slice are nil, otherwise
// it returns a single error formed by joining the error messages existing
// in the slice with a line-break.
func CheckErrs(errs []error) error { _ = "STUB: not implemented"; return nil }
