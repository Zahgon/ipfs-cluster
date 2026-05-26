// Package test provides utility methods to test APIs based on the common
// API.
package test

import (
	"io"
	"net/http"
	"testing"

	"github.com/libp2p/go-libp2p/core/host"
)

var (
	// SSLCertFile is the location of the certificate file.
	// Used in HTTPClient to set the right certificate when
	// creating an HTTPs client. Might need adjusting depending
	// on where the tests are running.
	SSLCertFile = "test/server.crt"

	// ClientOrigin sets the Origin header for requests to this.
	ClientOrigin = "myorigin"
)

// ProcessResp puts a response into a given type or fails the test.
func ProcessResp(t *testing.T, httpResp *http.Response, err error, resp interface{}) {
	_ = "STUB: not implemented"
	return
}

// ProcessStreamingResp decodes a streaming response into the given type
// and fails the test on error.
func ProcessStreamingResp(t *testing.T, httpResp *http.Response, err error, resp interface{}, trailerError bool) {
	_ = "STUB: not implemented"
	return
}

// normal response with error

// If we passed a slice we fill it in, otherwise we just decode
// on top of the passed value.

// CheckHeaders checks that all the headers are set to what is expected.
func CheckHeaders(t *testing.T, expected map[string][]string, url string, headers http.Header) {
	_ = "STUB: not implemented"
	return
}

// API represents what an API is to us.
type API interface {
	HTTPAddresses() ([]string, error)
	Host() host.Host
	Headers() map[string][]string
}

// URLFunc is a function that given an API returns a url string.
type URLFunc func(a API) string

// HTTPURL returns the http endpoint of the API.
func HTTPURL(a API) string { _ = "STUB: not implemented"; return "" }

// P2pURL returns the libp2p endpoint of the API.
func P2pURL(a API) string { _ = "STUB: not implemented"; return "" }

// HttpsURL returns the HTTPS endpoint of the API
func httpsURL(a API) string { _ = "STUB: not implemented"; return "" }

// IsHTTPS returns true if a url string uses HTTPS.
func IsHTTPS(url string) bool { _ = "STUB: not implemented"; return false }

// HTTPClient returns a client that supporst both http/https and
// libp2p-tunneled-http.
func HTTPClient(t *testing.T, h host.Host, isHTTPS bool) *http.Client {
	_ = "STUB: not implemented"
	return nil
}

// MakeHost makes a libp2p host that knows how to talk to the given API.
func MakeHost(t *testing.T, api API) host.Host { _ = "STUB: not implemented"; return *new(host.Host) }

// MakeGet performs a GET request against the API.
func MakeGet(t *testing.T, api API, url string, resp interface{}) {
	_ = "STUB: not implemented"
	return
}

// MakePost performs a POST request against the API with the given body.
func MakePost(t *testing.T, api API, url string, body []byte, resp interface{}) {
	_ = "STUB: not implemented"
	return
}

// MakePostWithContentType performs a POST with the given body and content-type.
func MakePostWithContentType(t *testing.T, api API, url string, body []byte, contentType string, resp interface{}) {
	_ = "STUB: not implemented"
	return
}

// MakeDelete performs a DELETE request against the given API.
func MakeDelete(t *testing.T, api API, url string, resp interface{}) {
	_ = "STUB: not implemented"
	return
}

// MakeOptions performs an OPTIONS request against the given api.
func MakeOptions(t *testing.T, api API, url string, reqHeaders http.Header) http.Header {
	_ = "STUB: not implemented"
	return *new(http.Header)
}

// MakeStreamingPost performs a POST request and uses ProcessStreamingResp
func MakeStreamingPost(t *testing.T, api API, url string, body io.Reader, contentType string, resp interface{}) {
	_ = "STUB: not implemented"
	return
}

// MakeStreamingGet performs a GET request and uses ProcessStreamingResp
func MakeStreamingGet(t *testing.T, api API, url string, resp interface{}, trailerError bool) {
	_ = "STUB: not implemented"
	return
}

// Func is a function that runs a test with a given URL.
type Func func(t *testing.T, url URLFunc)

// BothEndpoints runs a test.Func against the http and p2p endpoints.
func BothEndpoints(t *testing.T, test Func) { _ = "STUB: not implemented"; return }

// HTTPSEndPoint runs the given test.Func against an HTTPs endpoint.
func HTTPSEndPoint(t *testing.T, test Func) { _ = "STUB: not implemented"; return }
