package main

import (
	"errors"
	"io"

	dot "github.com/ipfs-cluster/go-dot"
	peer "github.com/libp2p/go-libp2p/core/peer"

	"github.com/ipfs-cluster/ipfs-cluster/api"
)

/*
   These functions are used to write an IPFS Cluster connectivity graph to a
   graphviz-style dot file.  Input an api.ConnectGraphSerial object, makeDot
   does some preprocessing and then passes all 3 link maps to a
   cluster-dotWriter which handles iterating over the link maps and writing
   dot file node and edge statements to make a dot-file graph.  Nodes are
   labeled with the go-libp2p-peer shortened peer id.  IPFS nodes are rendered
   with turquoise boundaries, Cluster nodes with orange.  Currently preprocessing
   consists of moving IPFS swarm peers not connected to any cluster peer to
   the IPFSLinks map in the event that the function was invoked with the
   allIpfs flag.  This allows all IPFS peers connected to the cluster to be
   rendered as nodes in the final graph.
*/

// nodeType specifies the type of node being represented in the dot file:
// either IPFS or Cluster
type nodeType int

const (
	tSelfCluster    nodeType = iota // cluster self node
	tCluster                        // cluster node
	tTrustedCluster                 // trusted cluster node
	tIPFS                           // IPFS node
	tIPFSMissing                    // Missing IPFS node
)

var errUnknownNodeType = errors.New("unsupported node type. Expected cluster or ipfs")

func makeDot(cg api.ConnectGraph, w io.Writer, allIpfs bool) error {
	_ = "STUB: not implemented"
	return nil
}

// include all swarm peers in the graph

// if id in IPFSLinks this will be overwritten
// if id not in IPFSLinks this will stay blank

type dotWriter struct {
	clusterNodes map[string]*dot.VertexDescription
	ipfsNodes    map[string]*dot.VertexDescription

	w        io.Writer
	dotGraph dot.Graph

	self             string
	idToPeername     map[string]string
	trustMap         map[string]bool
	ipfsEdges        map[string][]peer.ID
	clusterEdges     map[string][]peer.ID
	clusterIpfsEdges map[string]peer.ID
}

func (dW *dotWriter) addSubGraph(sGraph dot.Graph, rank string) { _ = "STUB: not implemented"; return }

// writes nodes to dot file output and creates and stores an ordering over nodes
func (dW *dotWriter) addNode(graph *dot.Graph, id string, nT nodeType) error {
	_ = "STUB: not implemented"
	return nil
}

func shorten(id string) string { _ = "STUB: not implemented"; return "" }

func label(peername, id string) string { _ = "STUB: not implemented"; return "" }

func (dW *dotWriter) print() error { _ = "STUB: not implemented"; return nil }

// Write cluster nodes, use sorted order for consistent labels

// Write ipfs nodes, use sorted order for consistent labels

// Write cluster edges

// Write cluster to ipfs edges

// Write ipfs edges

func sortedKeys(dict map[string][]peer.ID) []string { _ = "STUB: not implemented"; return nil }
