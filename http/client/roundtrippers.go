package client

import (
	"bytes"
	"compress/gzip"
	"io"
	"mime"
	"net/http"
	"strings"

	"github.com/andybalholm/brotli"
	"github.com/mattn/go-encoding"
	"github.com/pkg/errors"
)

// RoundTrip executes the HTTP request and applies rate limiting.
func (t *LimiterRoundTripper) RoundTrip(req *http.Request) (response *http.Response, err error) {
	if err = t.limiter.Wait(req.Context()); err != nil {
		return nil, err
	}

	return t.tripper.RoundTrip(req)
}

// RoundTrip executes the HTTP request and applies encoding to response.
func (t *EncoderRoundTripper) RoundTrip(req *http.Request) (response *http.Response, err error) {
	req.Header.Set("Accept-Encoding", "gzip, br")

	if req.Body != nil {
		if rawBody, err := io.ReadAll(req.Body); err == nil {
			var encodedBody bytes.Buffer
			writer := brotli.NewWriter(&encodedBody)
			writer.Write(rawBody)
			writer.Close()

			req.Body = io.NopCloser(&encodedBody)
			req.ContentLength = int64(encodedBody.Len())
			req.Header.Set("Content-Encoding", "br")
		}
	}

	response, err = t.tripper.RoundTrip(req)
	if err != nil {
		return response, err
	}

	if response.Header.Get("Content-Encoding") == "gzip" {
		r, err := gzip.NewReader(response.Body)
		if err != nil {
			return nil, errors.Errorf("failed to read response body: %s", err)
		}
		response.Body = r
	}

	if response.Header.Get("Content-Encoding") == "br" {
		response.Body = io.NopCloser(
			brotli.NewReader(response.Body),
		)
	}

	return response, err
}

// RoundTrip executes the HTTP request and applies charset to response.
func (t *CharsetRoundTripper) RoundTrip(req *http.Request) (response *http.Response, err error) {
	response, err = t.tripper.RoundTrip(req)
	if err != nil {
		return response, err
	}

	if ct := response.Header.Get("Content-Type"); ct != "" {
		_, params, err := mime.ParseMediaType(strings.Trim(ct, " "))
		if err != nil {
			return nil, errors.Errorf("failed to parse Content-Definitions response header %q: %s", ct, err)
		}
		if name, ok := params["charset"]; ok {
			enc := encoding.GetEncoding(name)
			if enc == nil {
				return nil, errors.Errorf("failed to decode response body: unknown charset %q", name)
			}
			response.Body = io.NopCloser(
				enc.NewDecoder().Reader(response.Body),
			)
		}
	}

	return response, err
}
