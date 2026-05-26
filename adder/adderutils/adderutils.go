// Package adderutils provides some utilities for adding content to cluster.
package adderutils

import (
	"context"
	"mime/multipart"
	"net/http"

	"github.com/ipfs-cluster/ipfs-cluster/api"

	logging "github.com/ipfs/go-log/v2"
	rpc "github.com/libp2p/go-libp2p-gorpc"
)

var logger = logging.Logger("adder")

// AddMultipartHTTPHandler is a helper function to add content
// uploaded using a multipart request. The outputTransform parameter
// allows to customize the http response output format to something
// else than api.AddedOutput objects.
func AddMultipartHTTPHandler(
	ctx context.Context,
	rpc *rpc.Client,
	params api.AddParams,
	reader *multipart.Reader,
	w http.ResponseWriter,
	outputTransform func(api.AddedOutput) interface{},
) (api.Cid, error) {
	_ = "STUB: not implemented"
	return *new(api.Cid), nil
}

// This must be application/json otherwise go-ipfs client
// will break.

// Browsers should not cache these responses.

// We need to ask the clients to close the connection
// (no keep-alive) of things break badly when adding.
// https://github.com/ipfs/go-ipfs-cmds/pull/116

// in this case we buffer responses in memory and
// return them as a valid JSON array.

// a slice of transformed AddedOutput

// Send an error

// handle stream-adding. This should be the default.

// https://github.com/ipfs-shipyard/ipfs-companion/issues/600

// Used by go-ipfs to signal errors half-way through the stream.

// Set trailer with error

func streamOutput(w http.ResponseWriter, output chan api.AddedOutput, transform func(api.AddedOutput) interface{}) {
	_ = "STUB: not implemented"
	return
}

func buildOutput(output chan api.AddedOutput, transform func(api.AddedOutput) interface{}) []interface{} {
	_ = "STUB: not implemented"
	return nil
}
