// Package ipfsadd is a simplified copy of go-ipfs/core/coreunix/add.go
package ipfsadd

import (
	"context"
	"io"
	"os"
	"time"

	"github.com/ipfs-cluster/ipfs-cluster/api"

	files "github.com/ipfs/boxo/files"
	mfs "github.com/ipfs/boxo/mfs"
	cid "github.com/ipfs/go-cid"
	ipld "github.com/ipfs/go-ipld-format"
	logging "github.com/ipfs/go-log/v2"
	peer "github.com/libp2p/go-libp2p/core/peer"
)

var log = logging.Logger("coreunix")

// how many bytes of progress to wait before sending a progress update message
const progressReaderIncrement = 1024 * 256

// cluster: we need to cache all to be able to output intermediate folders
//var liveCacheSize = uint64(256 << 10)

// NewAdder Returns a new Adder used for a file add operation.
func NewAdder(ctx context.Context, ds ipld.DAGService, allocs func() []peer.ID) (*Adder, error) {
	_ = "STUB: not implemented"
	// Cluster: we don't use pinner nor GCLocker.
	return nil, nil
}

// Adder holds the switches passed to the `add` command.
type Adder struct {
	ctx               context.Context
	dagService        ipld.DAGService
	allocsFun         func() []peer.ID
	Out               chan api.AddedOutput
	Progress          bool
	Trickle           bool
	RawLeaves         bool
	MaxLinks          int
	MaxDirectoryLinks int
	MaxHAMTFanout     int
	Silent            bool
	NoCopy            bool
	Chunker           string
	mroot             *mfs.Root
	tempRoot          cid.Cid
	CidBuilder        cid.Builder
	// liveNodes  uint64 // cluster: we do not clear mfs cache.
	lastFile mfs.FSNode
	// Cluster: ipfs does a hack in commands/add.go to set the filenames
	// in emitted events correctly. We carry a root folder name (or a
	// filename in the case of single files here and emit those events
	// correctly from the beginning).
	OutputPrefix string

	PreserveMode  bool
	PreserveMtime bool
	FileMode      os.FileMode
	FileMtime     time.Time
}

func (adder *Adder) mfsRoot() (*mfs.Root, error) { _ = "STUB: not implemented"; return nil, nil }

// Note, this adds it to DAGService already.

// SetMfsRoot sets `r` as the root for Adder.
func (adder *Adder) SetMfsRoot(r *mfs.Root) {
	_ = "STUB: not implemented"

	// Constructs a node from reader's data, and adds it. Doesn't pin.
	return
}

func (adder *Adder) add(reader io.Reader) (ipld.Node, error) {
	_ = "STUB: not implemented"
	return *new(ipld.Node), nil
}

// Cluster: we don't do batching/use BufferedDS.

// Cluster: commented as it is unused
// // RootNode returns the mfs root node
// func (adder *Adder) curRootNode() (ipld.Node, error) {
// 	mr, err := adder.mfsRoot()
// 	if err != nil {
// 		return nil, err
// 	}
// 	root, err := mr.GetDirectory().GetNode()
// 	if err != nil {
// 		return nil, err
// 	}

// 	// if one root file, use that hash as root.
// 	if len(root.Links()) == 1 {
// 		nd, err := root.Links()[0].GetNode(adder.ctx, adder.dagService)
// 		if err != nil {
// 			return nil, err
// 		}

// 		root = nd
// 	}

// 	return root, err
// }

// PinRoot recursively pins the root node of Adder and
// writes the pin state to the backing datastore.
// Cluster: we don't pin. Former Finalize().
func (adder *Adder) PinRoot(ctx context.Context, root ipld.Node, name string) error {
	_ = "STUB: not implemented"
	return nil
}

func (adder *Adder) outputDirs(path string, fsn mfs.FSNode) error {
	_ = "STUB: not implemented"
	return nil
}

// This fails when Child is of type *mfs.File
// because it tries to get them from the DAG
// service (does not implement this and returns
// a "not found" error)
// *mfs.Files are ignored in the recursive call
// anyway.
// For Cluster, we just ignore errors here.

func (adder *Adder) addNode(node ipld.Node, path string) error {
	_ = "STUB: not implemented"
	// patch it into the root
	return nil
}

// Cluster: cache the last file added.
// This avoids using the DAGService to get the first children
// if the MFS root when not wrapping.

// AddAllAndPin adds the given request's files and pin them.
// Cluster: we don'pin. Former AddFiles.
func (adder *Adder) AddAllAndPin(ctx context.Context, file files.Node) (ipld.Node, error) {
	_ = "STUB: not implemented"
	return *new(ipld.Node), nil
}

// get root

// if adding a file without wrapping, swap the root to it (when adding a
// directory, mfs root is the directory)

// Replace root with the first child

// Cluster: use the last file we added
// if we have one.

// output directory events

// Flush the MFS directories. This must happen after outputDirs as
// otherwise we will have no cached directories in MFS, and we cannot
// fetch from the DAGService.

// Cluster: call PinRoot which adds the root cid to the DAGService.
// Unsure if this a bug in IPFS when not pinning. Or it would get added
// twice.

// Cluster: we don't Pause for GC
func (adder *Adder) addFileNode(ctx context.Context, path string, file files.Node, toplevel bool) error {
	_ = "STUB: not implemented"
	return nil
}

// cluster: flushing MFS will cause issues when outputting intermediary
// mfs folders.
// if adder.liveNodes >= liveCacheSize {
// 	// TODO: A smarter cache that uses some sort of lru cache with an eviction handler
// 	mr, err := adder.mfsRoot()
// 	if err != nil {
// 		return err
// 	}
// 	if err := mr.FlushMemFree(adder.ctx); err != nil {
// 		return err
// 	}

// 	adder.liveNodes = 0
// }
// adder.liveNodes++

func (adder *Adder) addSymlink(ctx context.Context, path string, l *files.Symlink) error {
	_ = "STUB: not implemented"
	return nil
}

func (adder *Adder) addFile(path string, file files.File) error {
	_ = "STUB: not implemented"
	// if the progress flag was specified, wrap the file so that we can send
	// progress updates to the client (over the output channel)
	return nil
}

// patch it into the root

func (adder *Adder) addDir(ctx context.Context, path string, dir files.Directory, toplevel bool) error {
	_ = "STUB: not implemented"
	return nil
}

// if we need to store mode or modification time then create a new root which includes that data

// outputDagnode sends dagnode info over the output channel.
// Cluster: we use api.AddedOutput instead of coreiface events
// and make this an adder method to be able to prefix.
func (adder *Adder) outputDagnode(out chan api.AddedOutput, name string, dn ipld.Node) error {
	_ = "STUB: not implemented"
	return nil
}

// When adding things in a folder: "OutputPrefix/name"
// When adding a single file: "OutputPrefix" (name is unset)
// When adding a single thing with no name: ""
// Note: ipfs sets the name of files received on stdin to the CID,
// but cluster does not support stdin-adding so we do not
// account for this here.

type progressReader struct {
	file         io.Reader
	path         string
	out          chan api.AddedOutput
	bytes        int64
	lastProgress int64
}

func (i *progressReader) Read(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

type progressReader2 struct {
	*progressReader
	files.FileInfo
}

func (i *progressReader2) Read(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }
