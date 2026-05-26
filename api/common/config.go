package common

import (
	"crypto/tls"
	"io"
	"time"

	logging "github.com/ipfs/go-log/v2"
	crypto "github.com/libp2p/go-libp2p/core/crypto"
	peer "github.com/libp2p/go-libp2p/core/peer"
	ma "github.com/multiformats/go-multiaddr"

	"github.com/rs/cors"

	"github.com/ipfs-cluster/ipfs-cluster/config"
)

const minMaxHeaderBytes = 4096

const defaultMaxHeaderBytes = minMaxHeaderBytes

// Config provides common API configuration values and allows to customize its
// behavior. It implements most of the config.ComponentConfig interface
// (except the Default() and ConfigKey() methods). Config should be embedded
// in a Config object that implements the missing methods and sets the
// meta options.
type Config struct {
	config.Saver

	// These are meta-options and should be set by actual Config
	// implementations as early as possible.
	DefaultFunc   func(*Config) error
	ConfigKey     string
	EnvConfigKey  string
	Logger        *logging.ZapEventLogger
	RequestLogger *logging.ZapEventLogger
	APIErrorFunc  func(err error, status int) error

	// Listen address for the HTTP REST API endpoint.
	HTTPListenAddr []ma.Multiaddr

	// TLS configuration for the HTTP listener
	TLS *tls.Config

	// pathSSLCertFile is a path to a certificate file used to secure the
	// HTTP API endpoint. We track it so we can write it in the JSON.
	PathSSLCertFile string

	// pathSSLKeyFile is a path to the private key corresponding to the
	// SSLKeyFile. We track it so we can write it in the JSON.
	PathSSLKeyFile string

	// Maximum duration before timing out reading a full request
	ReadTimeout time.Duration

	// Maximum duration before timing out reading the headers of a request
	ReadHeaderTimeout time.Duration

	// Maximum duration before timing out write of the response
	WriteTimeout time.Duration

	// Server-side amount of time a Keep-Alive connection will be
	// kept idle before being reused
	IdleTimeout time.Duration

	// Maximum cumulative size of HTTP request headers in bytes
	// accepted by the server
	MaxHeaderBytes int

	// Listen address for the Libp2p REST API endpoint.
	Libp2pListenAddr []ma.Multiaddr

	// ID and PrivateKey are used to create a libp2p host if we
	// want the API component to do it (not by default).
	ID         peer.ID
	PrivateKey crypto.PrivKey

	// BasicAuthCredentials is a map of username-password pairs
	// which are authorized to use Basic Authentication
	BasicAuthCredentials map[string]string

	// HTTPLogFile is path of the file that would save HTTP API logs. If this
	// path is empty, HTTP logs would be sent to standard output. This path
	// should either be absolute or relative to cluster base directory. Its
	// default value is empty.
	HTTPLogFile string

	// Headers provides customization for the headers returned
	// by the API on existing routes.
	Headers map[string][]string

	// CORS header management
	CORSAllowedOrigins   []string
	CORSAllowedMethods   []string
	CORSAllowedHeaders   []string
	CORSExposedHeaders   []string
	CORSAllowCredentials bool
	CORSMaxAge           time.Duration

	// Tracing flag used to skip tracing specific paths when not enabled.
	Tracing bool
}

type jsonConfig struct {
	HTTPListenMultiaddress config.Strings `json:"http_listen_multiaddress"`
	SSLCertFile            string         `json:"ssl_cert_file,omitempty"`
	SSLKeyFile             string         `json:"ssl_key_file,omitempty"`
	ReadTimeout            string         `json:"read_timeout"`
	ReadHeaderTimeout      string         `json:"read_header_timeout"`
	WriteTimeout           string         `json:"write_timeout"`
	IdleTimeout            string         `json:"idle_timeout"`
	MaxHeaderBytes         int            `json:"max_header_bytes"`

	Libp2pListenMultiaddress config.Strings `json:"libp2p_listen_multiaddress,omitempty"`
	ID                       string         `json:"id,omitempty"`
	PrivateKey               string         `json:"private_key,omitempty" hidden:"true"`

	BasicAuthCredentials map[string]string   `json:"basic_auth_credentials"  hidden:"true"`
	HTTPLogFile          string              `json:"http_log_file"`
	Headers              map[string][]string `json:"headers"`

	CORSAllowedOrigins   []string `json:"cors_allowed_origins"`
	CORSAllowedMethods   []string `json:"cors_allowed_methods"`
	CORSAllowedHeaders   []string `json:"cors_allowed_headers"`
	CORSExposedHeaders   []string `json:"cors_exposed_headers"`
	CORSAllowCredentials bool     `json:"cors_allow_credentials"`
	CORSMaxAge           string   `json:"cors_max_age"`
}

// GetHTTPLogPath gets full path of the file where http logs should be
// saved.
func (cfg *Config) GetHTTPLogPath() string { _ = "STUB: not implemented"; return "" }

// ApplyEnvVars fills in any Config fields found as environment variables.
func (cfg *Config) ApplyEnvVars() error { _ = "STUB: not implemented"; return nil }

// Validate makes sure that all fields in this Config have
// working values, at least in appearance.
func (cfg *Config) Validate() error { _ = "STUB: not implemented"; return nil }

func (cfg *Config) validateLibp2p() error { _ = "STUB: not implemented"; return nil }

// if one is set, all should be

// LoadJSON parses a raw JSON byte slice created by ToJSON() and sets the
// configuration fields accordingly.
func (cfg *Config) LoadJSON(raw []byte) error { _ = "STUB: not implemented"; return nil }

func (cfg *Config) applyJSONConfig(jcfg *jsonConfig) error { _ = "STUB: not implemented"; return nil }

// Other options

func (cfg *Config) loadHTTPOptions(jcfg *jsonConfig) error { _ = "STUB: not implemented"; return nil }

// CORS

// compatibility

func (cfg *Config) tlsOptions(jcfg *jsonConfig) error { _ = "STUB: not implemented"; return nil }

func (cfg *Config) loadLibp2pOptions(jcfg *jsonConfig) error { _ = "STUB: not implemented"; return nil }

// ToJSON produce a human-friendly JSON representation of the Config
// object.
func (cfg *Config) ToJSON() (raw []byte, err error) { _ = "STUB: not implemented"; return nil, nil }

func (cfg *Config) toJSONConfig() (jcfg *jsonConfig, err error) {
	_ = "STUB: not implemented"
	// Multiaddress String() may panic
	return nil, nil
}

// CorsOptions returns cors.Options setup from the configured values.
func (cfg *Config) CorsOptions() *cors.Options { _ = "STUB: not implemented"; return nil }

// ToDisplayJSON returns JSON config as a string.
func (cfg *Config) ToDisplayJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// LogWriter returns a writer to write logs to. If a log path is configured,
// it creates a file.  Otherwise, uses the given logger.
func (cfg *Config) LogWriter() (io.Writer, error) {
	_ = "STUB: not implemented"
	return *new(io.Writer), nil
}

func newTLSConfig(certFile, keyFile string) (*tls.Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// based on https://github.com/denji/golang-tls
