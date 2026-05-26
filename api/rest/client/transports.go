package client

// This is essentially a http.DefaultTransport. We should not mess
// with it since it's a global variable, and we don't know who else uses
// it, so we create our own.
// TODO: Allow more configuration options.
func (c *defaultClient) defaultTransport() { _ = "STUB: not implemented"; return }

func (c *defaultClient) enableLibp2p() error { _ = "STUB: not implemented"; return nil }

func (c *defaultClient) enableTLS() error {
	_ = "STUB: not implemented"

	// based on https://github.com/denji/golang-tls
	return nil
}

func (c *defaultClient) enableUnix() error { _ = "STUB: not implemented"; return nil }
