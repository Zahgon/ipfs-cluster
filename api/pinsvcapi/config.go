package pinsvcapi

import (
	"net/http"
	"time"

	"github.com/ipfs-cluster/ipfs-cluster/api/common"
)

const configKey = "pinsvcapi"
const envConfigKey = "cluster_pinsvcapi"

const minMaxHeaderBytes = 4096

// Default values for Config.
const (
	DefaultReadTimeout       = 0
	DefaultReadHeaderTimeout = 5 * time.Second
	DefaultWriteTimeout      = 0
	DefaultIdleTimeout       = 120 * time.Second
	DefaultMaxHeaderBytes    = minMaxHeaderBytes
)

// Default values for Config.
var (
	// DefaultHTTPListenAddrs contains default listen addresses for the HTTP API.
	DefaultHTTPListenAddrs = []string{"/ip4/127.0.0.1/tcp/9097"}
	DefaultHeaders         = map[string][]string{}
)

// CORS defaults.
var (
	DefaultCORSAllowedOrigins = []string{"*"}
	DefaultCORSAllowedMethods = []string{
		http.MethodGet,
	}
	// rs/cors this will set sensible defaults when empty:
	// {"Origin", "Accept", "Content-Type", "X-Requested-With"}
	DefaultCORSAllowedHeaders = []string{}
	DefaultCORSExposedHeaders = []string{
		"Content-Type",
		"X-Stream-Output",
		"X-Chunked-Output",
		"X-Content-Length",
	}
	DefaultCORSAllowCredentials = true
	DefaultCORSMaxAge           time.Duration // 0. Means always.
)

// Config fully implements the config.ComponentConfig interface. Use
// NewConfig() to instantiate. Config embeds a common.Config object.
type Config struct {
	common.Config
}

// NewConfig creates a Config object setting the necessary meta-fields in the
// common.Config embedded object.
func NewConfig() *Config { _ = "STUB: not implemented"; return nil }

// ConfigKey returns a human-friendly identifier for this type of
// Config.
func (cfg *Config) ConfigKey() string {
	_ = "STUB: not implemented"

	// Default initializes this Config with working values.
	return ""
}

func (cfg *Config) Default() error { _ = "STUB: not implemented"; return nil }

// Sets all defaults for this config.
func defaultFunc(cfg *common.Config) error {
	_ = "STUB: not implemented"
	// http
	return nil
}

// libp2p

// Auth

// Logs

// Headers
