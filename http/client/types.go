// Package client provides an HTTP client.
package client

import (
	nethttp "net/http"
	"sync"
	"time"

	"github.com/leliuga/cdk/http"
	"github.com/leliuga/cdk/types"
	"golang.org/x/time/rate"
)

type (
	// Client is an HTTP client.
	Client struct {
		*Options
		client  *nethttp.Client
		cookies []*nethttp.Cookie
	}

	// Response is an HTTP response.
	Response struct {
		response *nethttp.Response
		headers  http.Headers
		body     []byte
	}

	// Progress is a progress.
	Progress struct {
		name      string
		total     uint64
		current   uint64
		startTime time.Time
		stopCh    chan struct{}
		wg        sync.WaitGroup
	}

	// Options is an HTTP client options.
	Options struct {
		Dsn                   types.URI      `json:"dsn"`
		Headers               nethttp.Header `json:"headers"`
		ProxyURL              types.URI      `json:"proxy_url"`
		MaxIdleConnections    int            `json:"max_idle_connections"`
		MaxConnectionsPerHost int            `json:"max_connections_per_host"`
		WriteBufferSize       int            `json:"write_buffer_size"`
		ReadBufferSize        int            `json:"read_buffer_size"`
		Timeout               time.Duration  `json:"timeout"`
		KeepAlive             time.Duration  `json:"keep_alive"`
		TLSHandshake          time.Duration  `json:"tls_handshake"`
		ExpectContinue        time.Duration  `json:"expect_continue"`
		IdleConnection        time.Duration  `json:"idle_connection"`
		ResponseHeader        time.Duration  `json:"response_header"`
		QPS                   float32        `json:"qps"`
		Burst                 int            `json:"burst"`
	}

	// LimiterRoundTripper is a tripper that applies rate limiting to requests.
	LimiterRoundTripper struct {
		limiter *rate.Limiter
		tripper nethttp.RoundTripper
	}

	// EncoderRoundTripper is a tripper that applies encoding to response.
	EncoderRoundTripper struct {
		tripper nethttp.RoundTripper
	}

	// CharsetRoundTripper is a tripper that applies charset to response.
	CharsetRoundTripper struct {
		tripper nethttp.RoundTripper
	}

	// Option represents the service option.
	Option func(o *Options)
)
