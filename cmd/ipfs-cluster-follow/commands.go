package main

import (
	"github.com/ipfs-cluster/ipfs-cluster/api"
	"github.com/ipfs-cluster/ipfs-cluster/cmdutils"
	cli "github.com/urfave/cli/v2"
)

func printFirstStart() { _ = "STUB: not implemented"; return }

func printNotInitialized(clusterName string) { _ = "STUB: not implemented"; return }

func setLogLevels(lvl string) { _ = "STUB: not implemented"; return }

// returns whether the config folder exists
func isInitialized(absPath string) bool { _ = "STUB: not implemented"; return false }

func listClustersCmd(c *cli.Context) error { _ = "STUB: not implemented"; return nil }

func infoCmd(c *cli.Context) error { _ = "STUB: not implemented"; return nil }

// Avoid pollution of the screen

// Either we loaded a valid config, or we are using a default. Worth
// applying env vars in the second case.

func initCmd(c *cli.Context) error { _ = "STUB: not implemented"; return nil }

func initCluster(c *cli.Context, ignoreReinit bool, cfgURL string) error {
	_ = "STUB: not implemented"
	return nil
}

// Setting the datastore here is useless, as we initialize with remote
// config and we will have an empty service.json with the source only.
// That source will decide which datastore is actually used.

func runCmd(c *cli.Context) error { _ = "STUB: not implemented"; return nil }

// set to "info" by default.
// Avoid API logs polluting the screen everytime we
// run some "list" command.

// Always run followers in follower mode.

// Do not let trusted peers GC this peer
// Defaults to Trusted otherwise.

// Discard API configurations and create our own (unix socket)

// Allow customization via env vars

// Hardcode disabled tracing and metrics to avoid mistakenly
// exposing any user data.

// This does nothing since we are not calling SetupMetrics anyways
// But stays just to be explicit.

// We are going to run a cluster peer and should do an
// oderly shutdown if we are interrupted: cancel default
// signal handling and leave things to HandleSignals.

// List
func listCmd(c *cli.Context) error { _ = "STUB: not implemented"; return nil }

// There was an error. Try offline status

// We are on offline mode so we cannot rely on IPFS being
// running and most probably our configuration is remote and
// to be loaded from IPFS. Thus we need to find a different
// way to decide whether to load badger/leveldb, and once we
// know, do it with the default settings.

// Since things were initialized, assume there is one at least.

// not needed

// we have a default crdt config with either leveldb or badger registered.

func printStatusOnline(absPath, clusterName string) error { _ = "STUB: not implemented"; return nil }

// do this once
// PeerMap will only have one key

func printStatusOffline(cfgHelper *cmdutils.ConfigHelper) error {
	_ = "STUB: not implemented"
	return nil
}

func printPin(c api.Cid, status, name, err string) { _ = "STUB: not implemented"; return }
