package http

import (
	"net/http"
	"strings"
)

// FromHeaders returns a Headers from the given http.Header.
func FromHeaders(header http.Header) Headers {
	headers := Headers{}
	for k, v := range header {
		if h, err := ParseHeader(k); err == nil {
			headers[h] = strings.Join(v, ",")
		}
	}

	return headers
}

// Set sets the header field to value. It replaces any existing values associated with key.
func (h Headers) Set(key Header, value string) {
	h[&key] = value
}

// Get gets the first value associated with the given key.
func (h Headers) Get(key Header) string {
	if value, ok := h[&key]; ok {
		return value
	}

	return ""
}

// Del deletes the values associated with key.
func (h Headers) Del(key Header) {
	delete(h, &key)
}

// Clone returns a copy of the Headers.
func (h Headers) Clone() Headers {
	headers := Headers{}
	for k, v := range h {
		headers[k] = v
	}

	return headers
}

// Merge merges the given Headers into the Headers.
func (h Headers) Merge(headers Headers) {
	for k, v := range headers {
		h[k] = v
	}
}

// ToHeaders returns the Headers as a http.Header.
func (h Headers) ToHeaders() http.Header {
	headers := http.Header{}
	for k, v := range h {
		headers.Set(k.String(), v)
	}

	return headers
}
