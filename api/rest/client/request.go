package client

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
)

type responseDecoder func(d *json.Decoder) error

func (c *defaultClient) do(
	ctx context.Context,
	method, path string,
	headers map[string]string,
	body io.Reader,
	obj interface{},
) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *defaultClient) doStream(
	ctx context.Context,
	method, path string,
	headers map[string]string,
	body io.Reader,
	outHandler responseDecoder,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *defaultClient) doRequest(
	ctx context.Context,
	method, path string,
	headers map[string]string,
	body io.Reader,
) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// this lets go use "chunked".

func (c *defaultClient) handleResponse(resp *http.Response, obj interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// not json. 404s etc.

func (c *defaultClient) handleStreamResponse(resp *http.Response, handler responseDecoder) error {
	_ = "STUB: not implemented"
	return nil
}

// we need to check trailers
