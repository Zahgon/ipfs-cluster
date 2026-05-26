package metrics

import (
	"github.com/ipfs-cluster/ipfs-cluster/api"

	peer "github.com/libp2p/go-libp2p/core/peer"
)

// PeersetFilter removes all metrics not belonging to the given
// peerset
func PeersetFilter(metrics []api.Metric, peerset []peer.ID) []api.Metric {
	_ = "STUB: not implemented"
	return nil
}
