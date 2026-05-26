package raft

import (
	"time"

	"github.com/ipfs-cluster/ipfs-cluster/config"

	peer "github.com/libp2p/go-libp2p/core/peer"

	hraft "github.com/hashicorp/raft"
)

// ConfigKey is the default configuration key for holding this component's
// configuration section.
var configKey = "raft"
var envConfigKey = "cluster_raft"

// Configuration defaults
var (
	DefaultDataSubFolder        = "raft"
	DefaultWaitForLeaderTimeout = 15 * time.Second
	DefaultCommitRetries        = 1
	DefaultNetworkTimeout       = 10 * time.Second
	DefaultCommitRetryDelay     = 200 * time.Millisecond
	DefaultBackupsRotate        = 6
	DefaultDatastoreNamespace   = "/r" // from "/raft"
)

// Config allows to configure the Raft Consensus component for ipfs-cluster.
// The component's configuration section is represented by ConfigJSON.
// Config implements the ComponentConfig interface.
type Config struct {
	config.Saver

	// will shutdown libp2p host on shutdown. Useful for testing
	hostShutdown bool

	// A folder to store Raft's data.
	DataFolder string

	// InitPeerset provides the list of initial cluster peers for new Raft
	// peers (with no prior state). It is ignored when Raft was already
	// initialized or when starting in staging mode.
	InitPeerset []peer.ID
	// LeaderTimeout specifies how long to wait for a leader before
	// failing an operation.
	WaitForLeaderTimeout time.Duration
	// NetworkTimeout specifies how long before a Raft network
	// operation is timed out
	NetworkTimeout time.Duration
	// CommitRetries specifies how many times we retry a failed commit until
	// we give up.
	CommitRetries int
	// How long to wait between retries
	CommitRetryDelay time.Duration
	// BackupsRotate specifies the maximum number of Raft's DataFolder
	// copies that we keep as backups (renaming) after cleanup.
	BackupsRotate int
	// Namespace to use when writing keys to the datastore
	DatastoreNamespace string

	// A Hashicorp Raft's configuration object.
	RaftConfig *hraft.Config

	// Tracing enables propagation of contexts across binary boundaries.
	Tracing bool
}

// jsonConfig represents a human-friendly Config
// object which can be saved to JSON.  Most configuration keys are converted
// into simple types like strings, and key names aim to be self-explanatory
// for the user.
// Check https://godoc.org/github.com/hashicorp/raft#Config for extended
// description on all Raft-specific keys.
type jsonConfig struct {
	// Storage folder for snapshots, log store etc. Used by
	// the Raft.
	DataFolder string `json:"data_folder,omitempty"`

	// InitPeerset provides the list of initial cluster peers for new Raft
	// peers (with no prior state). It is ignored when Raft was already
	// initialized or when starting in staging mode.
	InitPeerset []string `json:"init_peerset"`

	// How long to wait for a leader before failing
	WaitForLeaderTimeout string `json:"wait_for_leader_timeout"`

	// How long to wait before timing out network operations
	NetworkTimeout string `json:"network_timeout"`

	// How many retries to make upon a failed commit
	CommitRetries int `json:"commit_retries"`

	// How long to wait between commit retries
	CommitRetryDelay string `json:"commit_retry_delay"`

	// BackupsRotate specifies the maximum number of Raft's DataFolder
	// copies that we keep as backups (renaming) after cleanup.
	BackupsRotate int `json:"backups_rotate"`

	DatastoreNamespace string `json:"datastore_namespace,omitempty"`

	// HeartbeatTimeout specifies the time in follower state without
	// a leader before we attempt an election.
	HeartbeatTimeout string `json:"heartbeat_timeout,omitempty"`

	// ElectionTimeout specifies the time in candidate state without
	// a leader before we attempt an election.
	ElectionTimeout string `json:"election_timeout,omitempty"`

	// CommitTimeout controls the time without an Apply() operation
	// before we heartbeat to ensure a timely commit.
	CommitTimeout string `json:"commit_timeout,omitempty"`

	// MaxAppendEntries controls the maximum number of append entries
	// to send at once.
	MaxAppendEntries int `json:"max_append_entries,omitempty"`

	// TrailingLogs controls how many logs we leave after a snapshot.
	TrailingLogs uint64 `json:"trailing_logs,omitempty"`

	// SnapshotInterval controls how often we check if we should perform
	// a snapshot.
	SnapshotInterval string `json:"snapshot_interval,omitempty"`

	// SnapshotThreshold controls how many outstanding logs there must be
	// before we perform a snapshot.
	SnapshotThreshold uint64 `json:"snapshot_threshold,omitempty"`

	// LeaderLeaseTimeout is used to control how long the "lease" lasts
	// for being the leader without being able to contact a quorum
	// of nodes. If we reach this interval without contact, we will
	// step down as leader.
	LeaderLeaseTimeout string `json:"leader_lease_timeout,omitempty"`

	// The unique ID for this server across all time. When running with
	// ProtocolVersion < 3, you must set this to be the same as the network
	// address of your transport.
	// LocalID string `json:local_id`
}

// ConfigKey returns a human-friendly identifier for this Config.
func (cfg *Config) ConfigKey() string {
	_ = "STUB: not implemented"

	// Validate checks that this configuration has working values,
	// at least in appearance.
	return ""
}

func (cfg *Config) Validate() error { _ = "STUB: not implemented"; return nil }

// LoadJSON parses a json-encoded configuration (see jsonConfig).
// The Config will have default values for all fields not explicited
// in the given json object.
func (cfg *Config) LoadJSON(raw []byte) error { _ = "STUB: not implemented"; return nil }

func (cfg *Config) applyJSONConfig(jcfg *jsonConfig) error { _ = "STUB: not implemented"; return nil }

// Parse durations. We ignore errors as 0 will take Default values.

// Set all values in config. For some, take defaults if they are 0.
// Set values from jcfg if they are not 0 values

// Own values

// Raft values

// ToJSON returns the pretty JSON representation of a Config.
func (cfg *Config) ToJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (cfg *Config) toJSONConfig() *jsonConfig { _ = "STUB: not implemented"; return nil }

// otherwise leave empty so it gets omitted.

// Default initializes this configuration with working defaults.
func (cfg *Config) Default() error {
	_ = "STUB: not implemented"
	// empty so it gets omitted
	return nil
}

// These options are imposed over any Default Raft Config.

// Set up logging

// ApplyEnvVars fills in any Config fields found
// as environment variables.
func (cfg *Config) ApplyEnvVars() error { _ = "STUB: not implemented"; return nil }

// GetDataFolder returns the Raft data folder that we are using.
func (cfg *Config) GetDataFolder() string { _ = "STUB: not implemented"; return "" }

// ToDisplayJSON returns JSON config as a string.
func (cfg *Config) ToDisplayJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
