package pebble

import (
	"time"

	"github.com/cockroachdb/pebble/v2"
	"github.com/cockroachdb/pebble/v2/bloom"

	"github.com/ipfs-cluster/ipfs-cluster/config"
)

const configKey = "pebble"
const envConfigKey = "cluster_pebble"

// Default values for Pebble Config
const (
	DefaultSubFolder = "pebble"
)

var (
	// DefaultPebbleOptions for convenience.
	DefaultPebbleOptions pebble.Options
	// DefaultCacheSize sets the maximum size of the block cache.
	DefaultCacheSize int64 = 1 << 30 // Pebble's default: 8MiB
	// DefaultMemTableSize sets the size of the memtables and affecst total
	// size of the WAL. It must be under 4GB.
	DefaultMemTableSize uint64 = 64 << 20 // Pebble's default: 4MiB
	// DefaultMemTableStopWritesThreshold defines how many memtables can
	// be queued for writing before stopping writes (memtable memory
	// consumption should approach
	// MemTableStopWritesThreshold*MemTableSize in that case).
	DefaultMemTableStopWritesThreshold = 20 // Pebble's default: 12
	// DefaultBytesPerSync controls how often to call the Filesystem
	// Sync.
	DefaultBytesPerSync = 1 << 20 // Pebble's default: 512KiB
	// DefaultMaxOpenFiles controls how many files can be kept open by
	// Pebble.
	DefaultMaxOpenFiles = 1000 // Pebble's default: 500
	// DefaultL0CompactionThreshold defines the read amplification on L0
	// that triggers compaction
	DefaultL0CompactionThreshold = 4 // Pebble's default: 4
	// DefaultL0CompactionFileThreshold defines the number of files that
	// trigger compactions of L0
	DefaultL0CompactionFileThreshold = 750 // Pebble's default: 500
	// DefaultL0StopWritesThreshold defines the critical threshold for
	// read amplification on L0, which stops writes until compaction
	// reduces it.
	DefaultL0StopWritesThreshold = 12 // Pebble's default : 4
	// DefaultLBaseMaxBytes defines maximum size of LBase, where memtables
	// are temporally written to
	DefaultLBaseMaxBytes int64 = 128 << 20 // Pebble's default: 64MiB
	// DefaultL0TargetFileSize defines the target filesize for L0. It is
	// multiplied by 2 for every subsequent level.
	DefaultL0TargetFileSize int64 = 4 << 20 // Pebble's default: 4M
	// DefaultBlockSize defines the target size for table blocks (used in
	// all levels).
	DefaultBlockSize int = 4 << 10 // Pebble's default: 4KiB
	// DefaultFilterPolicy defines the number of bits used per key for
	// bloom filters. 10 yields a 1% false positive rate.
	DefaultFilterPolicy bloom.FilterPolicy = 10 // Pebble's default: 10
	// DefaultFormatMajorVersion sets the format of Pebble on-disk files.
	DefaultFormatMajorVersion = pebble.FormatNewest
)

func init() {
	DefaultPebbleOptions.EnsureDefaults()
}

// Config is used to initialize a Pebble datastore. It implements the
// ComponentConfig interface.
type Config struct {
	config.Saver

	// The folder for this datastore. Non-absolute paths are relative to
	// the base configuration folder.
	Folder string

	PebbleOptions pebble.Options
}

// pebbleOptions is a subset of pebble.Options so it can be marshaled by us in
// the cluster configuration.
type pebbleOptions struct {
	EventListener               *pebble.EventListener     `json:"-"`
	CacheSizeBytes              int64                     `json:"cache_size_bytes"`
	BytesPerSync                int                       `json:"bytes_per_sync"`
	DisableWAL                  bool                      `json:"disable_wal"`
	FlushDelayDeleteRange       time.Duration             `json:"flush_delay_delete_range"`
	FlushDelayRangeKey          time.Duration             `json:"flush_delay_range_key"`
	FlushSplitBytes             int64                     `json:"flush_split_bytes"`
	FormatMajorVersion          pebble.FormatMajorVersion `json:"format_major_version"`
	L0CompactionFileThreshold   int                       `json:"l0_compaction_file_threshold"`
	L0CompactionThreshold       int                       `json:"l0_compaction_threshold"`
	L0StopWritesThreshold       int                       `json:"l0_stop_writes_threshold"`
	LBaseMaxBytes               int64                     `json:"l_base_max_bytes"`
	MaxOpenFiles                int                       `json:"max_open_files"`
	MemTableSize                uint64                    `json:"mem_table_size"`
	MemTableStopWritesThreshold int                       `json:"mem_table_stop_writes_threshold"`
	ReadOnly                    bool                      `json:"read_only"`
	WALBytesPerSync             int                       `json:"wal_bytes_per_sync"`
	Levels                      []levelOptions            `json:"levels"`
}

func (po *pebbleOptions) Unmarshal() *pebble.Options { _ = "STUB: not implemented"; return nil }

// pebbleOpts.Levels

func (po *pebbleOptions) Marshal(pebbleOpts *pebble.Options) { _ = "STUB: not implemented"; return }

// levelOptions carries options for pebble's per-level parameters.
// Compression used to be:
//
// const (
//
//	DefaultCompression Compression = iota
//	NoCompression
//	SnappyCompression
//	ZstdCompression
//	NCompression -- which means NoCompression it seems
//
// )
type levelOptions struct {
	BlockRestartInterval int                `json:"block_restart_interval"`
	BlockSize            int                `json:"block_size"`
	BlockSizeThreshold   int                `json:"block_size_threshold"`
	Compression          int                `json:"compression"`
	FilterType           pebble.FilterType  `json:"filter_type"`
	FilterPolicy         bloom.FilterPolicy `json:"filter_policy"`
	IndexBlockSize       int                `json:"index_block_size"`
	TargetFileSize       int64              `json:"target_file_size"`
}

func (lo *levelOptions) Unmarshal() (*pebble.LevelOptions, int64) {
	_ = "STUB: not implemented"
	return nil, 0
}

// On prevous relesases we could return the int. Now we need to provide a profile.

func (lo *levelOptions) Marshal(levelOpts *pebble.LevelOptions, targetFileSize int64) {
	_ = "STUB: not implemented"
	return
}

// NoCompression

type jsonConfig struct {
	Folder        string        `json:"folder,omitempty"`
	PebbleOptions pebbleOptions `json:"pebble_options,omitempty"`
}

// ConfigKey returns a human-friendly identifier for this type of Datastore.
func (cfg *Config) ConfigKey() string {
	_ = "STUB: not implemented"

	// Default initializes this Config with sensible values.
	return ""
}

func (cfg *Config) Default() error { _ = "STUB: not implemented"; return nil }

// cfg.PebbleOptions.Levels = make([]pebble.LevelOptions, 7) // fixed to [7]LevelOptions

// Deprecated: cfg.PebbleOptions.Levels[0].TargetFileSize = DefaultL0TargetFileSize
//added

func defaultLevelOpts(l, prev *pebble.LevelOptions, i int) { _ = "STUB: not implemented"; return }

// does not overwite, only sets the rest.

// ApplyEnvVars fills in any Config fields found as environment variables.
func (cfg *Config) ApplyEnvVars() error { _ = "STUB: not implemented"; return nil }

// Validate checks that the fields of this Config have working values,
// at least in appearance.
func (cfg *Config) Validate() error { _ = "STUB: not implemented"; return nil }

// LoadJSON reads the fields of this Config from a JSON byteslice as
// generated by ToJSON.
func (cfg *Config) LoadJSON(raw []byte) error { _ = "STUB: not implemented"; return nil }

func (cfg *Config) applyJSONConfig(jcfg *jsonConfig) error { _ = "STUB: not implemented"; return nil }

// ToJSON generates a JSON-formatted human-friendly representation of this
// Config.
func (cfg *Config) ToJSON() (raw []byte, err error) { _ = "STUB: not implemented"; return nil, nil }

func (cfg *Config) toJSONConfig() *jsonConfig { _ = "STUB: not implemented"; return nil }

// GetFolder returns the Pebble folder.
func (cfg *Config) GetFolder() string { _ = "STUB: not implemented"; return "" }

// ToDisplayJSON returns JSON config as a string.
func (cfg *Config) ToDisplayJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
