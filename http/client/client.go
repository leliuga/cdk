package client

import (
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"net"
	nethttp "net/http"
	"strings"
	"time"

	"github.com/leliuga/cdk/http"
	"github.com/leliuga/cdk/http/schema"
	"github.com/leliuga/cdk/types"
	"github.com/pkg/errors"
	"golang.org/x/time/rate"
)

// NewClient creates a new client.
func NewClient(options *Options) *Client {
	proxyURL := nethttp.ProxyFromEnvironment

	if options.ProxyURL.Validate() {
		proxyURL = nethttp.ProxyURL(options.ProxyURL.URL)
	}

	c := &Client{
		Options: options,
		client: &nethttp.Client{
			Transport: &CharsetRoundTripper{
				tripper: &EncoderRoundTripper{
					tripper: &LimiterRoundTripper{
						limiter: rate.NewLimiter(rate.Limit(options.QPS), options.Burst),
						tripper: &nethttp.Transport{
							Proxy: proxyURL,
							DialContext: defaultTransportDialContext(&net.Dialer{
								Timeout:   options.Timeout,
								KeepAlive: options.KeepAlive,
							}),
							TLSHandshakeTimeout:   options.TLSHandshake,
							ExpectContinueTimeout: options.ExpectContinue,
							IdleConnTimeout:       options.IdleConnection,
							ResponseHeaderTimeout: options.ResponseHeader,
							MaxIdleConns:          options.MaxIdleConnections,
							MaxConnsPerHost:       options.MaxConnectionsPerHost,
							WriteBufferSize:       options.WriteBufferSize,
							ReadBufferSize:        options.ReadBufferSize,
							TLSClientConfig:       &tls.Config{InsecureSkipVerify: true},
							ForceAttemptHTTP2:     true,
						},
					},
				},
			},
		},
	}

	return c
}

// Do sends an HTTP request and returns an HTTP response.
func (c *Client) Do(ctx context.Context, endpoint *schema.Endpoint) (*Response, error) {
	for k, v := range endpoint.Headers {
		c.Headers.Set(*k, v)
	}

	req, err := c.newRequest(ctx, *endpoint.Method, endpoint.Path, endpoint.Payload)
	if err != nil {
		return nil, err
	}

	response, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	res := FromResponse(response)
	c.cookies = res.Cookies()

	if err = endpoint.Expect.Validate(res.Status(), res.Headers()); err != nil {
		b, _ := io.ReadAll(res.Body())

		return nil, errors.Errorf("%s The response contains - %s", err, string(b))
	}

	return res, nil
}

// newRequest creates a new HTTP request.
func (c *Client) newRequest(ctx context.Context, method http.Method, path *types.Path, payload any) (req *nethttp.Request, err error) {
	var reader io.Reader
	url := c.BaseUri.String() + path.String()

	if payload != nil {
		contentType := c.Headers.Get(http.HeaderContentType)
		reader, err = types.ContentTypeMarshal(contentType, payload)
		if err != nil {
			return nil, errors.Errorf("failed to marshal request payload as %s: %#v: %s", contentType, payload, err)
		}
	}

	req, err = nethttp.NewRequestWithContext(ctx, method.String(), url, reader)
	if err != nil {
		return nil, err
	}

	req.Header = c.Headers.ToHeaders()
	for _, cookie := range c.cookies {
		req.AddCookie(cookie)
	}
	defaultHeaders := types.Map[string]{
		"Accept":                    "text/html,application/xhtml+xml,application/xml,image/avif,image/webp,*/*;q=0.8",
		"Accept-Language":           "en-US,en;q=0.5",
		"Cache-Control":             "no-cache",
		"Connection":                "keep-alive",
		"Dnt":                       "1",
		"Pragma":                    "no-cache",
		"Sec-Fetch-Dest":            "document",
		"Sec-Fetch-Mode":            "navigate",
		"Sec-Fetch-Site":            "none",
		"Sec-Fetch-User":            "?1",
		"Upgrade-Insecure-Requests": "1",
	}

	for k, v := range defaultHeaders {
		if req.Header.Get(string(k)) == "" {
			req.Header.Set(string(k), v)
		}
	}

	accepts := strings.Split(req.Header.Get("Accept"), ",")
	for _, v := range types.ContentTypeNames {
		accepts = append(accepts, v)
	}
	req.Header.Set("Accept", strings.Join(uniqueStrings(accepts), ","))

	if req.Header.Get("User-Agent") == "" {
		req.Header.Set("User-Agent", fmt.Sprintf("%s/%d.%d", DefaultUserAgent, req.ProtoMajor, req.ProtoMinor))
	}

	return req, nil
}

func defaultTransportDialContext(dialer *net.Dialer) func(context.Context, string, string) (net.Conn, error) {
	return dialer.DialContext
}

func uniqueStrings(input []string) []string {
	uniqueMap := make(map[string]bool)
	var uniqueSlice []string

	for _, str := range input {
		if !uniqueMap[str] {
			uniqueMap[str] = true
			uniqueSlice = append(uniqueSlice, str)
		}
	}

	return uniqueSlice
}

// Download downloads a file.
func Download(ctx context.Context, uri *types.URI, filename string) error {
	res, err := NewClient(NewOptions(
		WithBaseUri(uri.BaseUri().String()),
	)).Do(ctx, schema.NewEndpoint("Download file", http.MethodGet, uri.Path))

	if err != nil {
		return err
	}
	defer res.Close()

	progress := NewProgress(filename, uint64(res.ContentLength()))
	progress.Start(1 + time.Second)
	defer progress.Stop()

	return res.Save(filename, progress)
}
