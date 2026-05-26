package badger3

import (
	"time"

	"github.com/dgraph-io/badger/v3"
	"github.com/dgraph-io/badger/v3/options"

	"github.com/ipfs-cluster/ipfs-cluster/config"
)

const configKey = "badger3"
const envConfigKey = "cluster_badger3"

// Default values for badger Config
const (
	DefaultSubFolder = "badger3"
)

var (
	// DefaultBadgerOptions has to be a var because badger.DefaultOptions
	// is. Values are customized during Init().
	DefaultBadgerOptions badger.Options

	// DefaultGCDiscardRatio for GC operations. See Badger docs.
	DefaultGCDiscardRatio float64 = 0.2
	// DefaultGCInterval specifies interval between GC cycles.
	DefaultGCInterval time.Duration = 15 * time.Minute
	// DefaultGCSleep specifies sleep time between GC rounds.
	DefaultGCSleep time.Duration = 10 * time.Second
)

func init() {
	DefaultBadgerOptions = badger.DefaultOptions("")
	// Better to slow down starts than shutdowns.
	DefaultBadgerOptions.CompactL0OnClose = false
	// Defaults to 1MB! For us that means everything goes into the LSM
	// tree and the LSM tree is supposed to be loaded into memory in full.
	// We only put very small things on the LSM tree by default (i.e. a
	// single CID).
	DefaultBadgerOptions.ValueThreshold = 100
	// Disable Block Cache: the cluster read-pattern at scale requires
	// looping regularly all keys. The CRDT read-patterm avoids reading
	// something twice. In general, it probably does not add much, and it
	// is recommended to be disabled when not using compression.
	DefaultBadgerOptions.BlockCacheSize = 0
	// Let's disable compression for values, better perf when reading and
	// usually the ratio between data stored by badger and the cluster
	// should be small. Users can always enable.
	DefaultBadgerOptions.Compression = options.None
	// There is a write lock in go-ds-crdt that writes batches one by one.
	// Also NewWriteBatch says that there can never be transaction
	// conflicts when doing batches. And IPFS will only write a block
	// once, or do it with the same values. In general, we probably don't
	// care about conflicts much (rows updated while a commit transaction
	// was open). Increases perf too.
	DefaultBadgerOptions.DetectConflicts = false
	// TODO: Increase memtable size. This will use some more memory, but any
	// normal system should be able to deal with using 256MiB for the
	// memtable. Badger puts a lot of things in memory anyways,
	// i.e. IndexCacheSize is set to 0. Note NumMemTables is 5.
	// DefaultBadgerOptions.MemTableSize = 268435456 // 256MiB

}

// Config is used to initialize a BadgerDB datastore. It implements the
// ComponentConfig interface.
type Config struct {
	config.Saver

	// The folder for this datastore. Non-absolute paths are relative to
	// the base configuration folder.
	Folder string

	// For GC operation. See Badger documentation.
	GCDiscardRatio float64

	// Interval between GC cycles. Each GC cycle runs one or more
	// rounds separated by GCSleep.
	GCInterval time.Duration

	// Time between rounds in a GC cycle
	GCSleep time.Duration

	BadgerOptions badger.Options
}

// badgerOptions is a copy of badger.Options so it can be marshaled by us.
type badgerOptions struct {
	Dir               string `json:"dir"`
	ValueDir          string `json:"value_dir"`
	SyncWrites        bool   `json:"sync_writes"`
	NumVersionsToKeep int    `json:"num_versions_to_keep"`
	ReadOnly          bool   `json:"read_only"`
	// Logger
	Compression    options.CompressionType `json:"compression"`
	InMemory       bool                    `json:"in_memory"`
	MetricsEnabled bool                    `json:"metrics_enabled"`
	NumGoroutines  int                     `json:"num_goroutines"`

	MemTableSize        int64 `json:"mem_table_size"`
	BaseTableSize       int64 `json:"base_table_size"`
	BaseLevelSize       int64 `json:"base_level_size"`
	LevelSizeMultiplier int   `json:"level_size_multiplier"`
	TableSizeMultiplier int   `json:"table_size_multiplier"`
	MaxLevels           int   `json:"max_levels"`

	VLogPercentile     float64 `json:"v_log_percentile"`
	ValueThreshold     int64   `json:"value_threshold"`
	NumMemtables       int     `json:"num_memtables"`
	BlockSize          int     `json:"block_size"`
	BloomFalsePositive float64 `json:"bloom_false_positive"`
	BlockCacheSize     int64   `json:"block_cache_size"`
	IndexCacheSize     int64   `json:"index_cache_size"`

	NumLevelZeroTables      int `json:"num_level_zero_tables"`
	NumLevelZeroTablesStall int `json:"num_level_zero_tables_stall"`

	ValueLogFileSize   int64  `json:"value_log_file_size"`
	ValueLogMaxEntries uint32 `json:"value_log_max_entries"`

	NumCompactors        int  `json:"num_compactors"`
	CompactL0OnClose     bool `json:"compact_l_0_on_close"`
	LmaxCompaction       bool `json:"lmax_compaction"`
	ZSTDCompressionLevel int  `json:"zstd_compression_level"`

	VerifyValueChecksum bool `json:"verify_value_checksum"`

	ChecksumVerificationMode options.ChecksumVerificationMode `json:"checksum_verification_mode"`
	DetectConflicts          bool                             `json:"detect_conflicts"`

	NamespaceOffset int `json:"namespace_offset"`
}

func (bo *badgerOptions) Unmarshal() *badger.Options { _ = "STUB: not implemented"; return nil }

func (bo *badgerOptions) Marshal(badgerOpts *badger.Options) { _ = "STUB: not implemented"; return }

type jsonConfig struct {
	Folder         string        `json:"folder,omitempty"`
	GCDiscardRatio float64       `json:"gc_discard_ratio"`
	GCInterval     string        `json:"gc_interval"`
	GCSleep        string        `json:"gc_sleep"`
	BadgerOptions  badgerOptions `json:"badger_options,omitempty"`
}

// ConfigKey returns a human-friendly identifier for this type of Datastore.
func (cfg *Config) ConfigKey() string {
	_ = "STUB: not implemented"

	// Default initializes this Config with sensible values.
	return ""
}

func (cfg *Config) Default() error { _ = "STUB: not implemented"; return nil }

// ApplyEnvVars fills in any Config fields found as environment variables.
func (cfg *Config) ApplyEnvVars() error { _ = "STUB: not implemented"; return nil }

// Validate checks that the fields of this Config have working values,
// at least in appearance.
func (cfg *Config) Validate() error { _ = "STUB: not implemented"; return nil }

// LoadJSON reads the fields of this Config from a JSON byteslice as
// generated by ToJSON.
func (cfg *Config) LoadJSON(raw []byte) error { _ = "STUB: not implemented"; return nil }

func (cfg *Config) applyJSONConfig(jcfg *jsonConfig) error { _ = "STUB: not implemented"; return nil }

// 0 is an invalid option anyways. In that case, set default (0.2)

// If these durations are set, GC is enabled by default with default
// values.

// ToJSON generates a JSON-formatted human-friendly representation of this
// Config.
func (cfg *Config) ToJSON() (raw []byte, err error) { _ = "STUB: not implemented"; return nil, nil }

func (cfg *Config) toJSONConfig() *jsonConfig { _ = "STUB: not implemented"; return nil }

// GetFolder returns the BadgerDB folder.
func (cfg *Config) GetFolder() string { _ = "STUB: not implemented"; return "" }

// ToDisplayJSON returns JSON config as a string.
func (cfg *Config) ToDisplayJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
