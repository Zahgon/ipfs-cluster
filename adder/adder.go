// Package adder implements functionality to add content to IPFS daemons
// managed by the Cluster.
package adder

import (
	"context"
	"mime/multipart"

	"github.com/ipfs-cluster/ipfs-cluster/adder/ipfsadd"
	"github.com/ipfs-cluster/ipfs-cluster/api"
	peer "github.com/libp2p/go-libp2p/core/peer"

	files "github.com/ipfs/boxo/files"
	merkledag "github.com/ipfs/boxo/ipld/merkledag"
	cid "github.com/ipfs/go-cid"
	ipld "github.com/ipfs/go-ipld-format"
	ipldlegacy "github.com/ipfs/go-ipld-legacy"
	logging "github.com/ipfs/go-log/v2"
	dagpb "github.com/ipld/go-codec-dagpb"
	"github.com/ipld/go-ipld-prime/codec/dagcbor"
	"github.com/ipld/go-ipld-prime/codec/raw"
	cidlink "github.com/ipld/go-ipld-prime/linking/cid"
	"github.com/ipld/go-ipld-prime/multicodec"
	"github.com/ipld/go-ipld-prime/node/basicnode"
)

var logger = logging.Logger("adder")

var ipldDecoder *ipldlegacy.Decoder

// create an ipld registry specific to this package
func init() {
	mcReg := multicodec.Registry{}
	mcReg.RegisterDecoder(cid.DagProtobuf, dagpb.Decode)
	mcReg.RegisterDecoder(cid.Raw, raw.Decode)
	mcReg.RegisterDecoder(cid.DagCBOR, dagcbor.Decode)
	ls := cidlink.LinkSystemUsingMulticodecRegistry(mcReg)

	ipldDecoder = ipldlegacy.NewDecoderWithLS(ls)
	ipldDecoder.RegisterCodec(cid.DagProtobuf, dagpb.Type.PBNode, merkledag.ProtoNodeConverter)
	ipldDecoder.RegisterCodec(cid.Raw, basicnode.Prototype.Bytes, merkledag.RawNodeConverter)
}

// ClusterDAGService is an implementation of ipld.DAGService plus a Finalize
// method. ClusterDAGServices can be used to provide Adders with a different
// add implementation.
type ClusterDAGService interface {
	ipld.DAGService
	// Finalize receives the IPFS content root CID as
	// returned by the ipfs adder.
	Finalize(ctx context.Context, ipfsRoot api.Cid) (api.Cid, error)
	// Close performs any necessary cleanups and should be called
	// whenever the DAGService is not going to be used anymore.
	Close() error
	// Allocations returns the allocations made by the cluster DAG service
	// for the added content.
	Allocations() []peer.ID
}

// A dagFormatter can create dags from files.Node. It can keep state
// to add several files to the same dag.
type dagFormatter interface {
	Add(ctx context.Context, name string, f files.Node) (api.Cid, error)
}

// Adder is used to add content to IPFS Cluster using an implementation of
// ClusterDAGService.
type Adder struct {
	ctx    context.Context
	cancel context.CancelFunc

	dgs ClusterDAGService

	params api.AddParams

	// AddedOutput updates are placed on this channel
	// whenever a block is processed. They contain information
	// about the block, the CID, the Name etc. and are mostly
	// meant to be streamed back to the user.
	output chan api.AddedOutput
}

// New returns a new Adder with the given ClusterDAGService, add options and a
// channel to send updates during the adding process.
//
// An Adder may only be used once.
func New(ds ClusterDAGService, p api.AddParams, out chan api.AddedOutput) *Adder {
	_ = "STUB: not implemented"
	// Discard all progress update output as the caller has not provided
	// a channel for them to listen on.
	return nil
}

func (a *Adder) setContext(ctx context.Context) {
	_ = "STUB: not implemented"
	// only allows first context
	return
}

// FromMultipart adds content from a multipart.Reader. The adder will
// no longer be usable after calling this method.
func (a *Adder) FromMultipart(ctx context.Context, r *multipart.Reader) (api.Cid, error) {
	_ = "STUB: not implemented"
	return *new(api.Cid), nil
}

// FromFiles adds content from a files.Directory. The adder will no longer
// be usable after calling this method.
func (a *Adder) FromFiles(ctx context.Context, f files.Directory) (api.Cid, error) {
	_ = "STUB: not implemented"
	return *new(api.Cid), nil
}

// don't allow running twice

// setup wrapping

// TODO (hector): We can only add a single CAR file for the
// moment.

// A wrapper around the ipfsadd.Adder to satisfy the dagFormatter interface.
type ipfsAdder struct {
	*ipfsadd.Adder
}

func newIpfsAdder(ctx context.Context, dgs ClusterDAGService, params api.AddParams, out chan api.AddedOutput) (*ipfsAdder, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Set up prefi

func (ia *ipfsAdder) Add(ctx context.Context, name string, f files.Node) (api.Cid, error) {
	_ = "STUB: not implemented"
	// In order to set the AddedOutput names right, we use
	// OutputPrefix:
	//
	// When adding a folder, this is the root folder name which is
	// prepended to the addedpaths.  When adding a single file,
	// this is the name of the file which overrides the empty
	// AddedOutput name.
	//
	// After coreunix/add.go was refactored in go-ipfs and we
	// followed suit, it no longer receives the name of the
	// file/folder being added and does not emit AddedOutput
	// events with the right names. We addressed this by adding
	// OutputPrefix to our version. go-ipfs modifies emitted
	// events before sending to user).
	return *new(api.Cid), nil
}

// An adder to add CAR files. It is at the moment very basic, and can
// add a single CAR file with a single root. Ideally, it should be able to
// add more complex, or several CARs by wrapping them with a single root.
// But for that we would need to keep state and track an MFS root similarly to
// what the ipfsadder does.
type carAdder struct {
	ctx    context.Context
	dgs    ClusterDAGService
	params api.AddParams
	output chan api.AddedOutput
}

func newCarAdder(ctx context.Context, dgs ClusterDAGService, params api.AddParams, out chan api.AddedOutput) (*carAdder, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Add takes a node which should be a CAR file and nothing else and
// adds its blocks using the ClusterDAGService.
func (ca *carAdder) Add(ctx context.Context, name string, fn files.Node) (api.Cid, error) {
	_ = "STUB: not implemented"
	return *new(api.Cid), nil
}

// If the root is in the CAR and the root is a UnixFS
// node, then set the size in the output object.
