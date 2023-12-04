package client

import (
	"io"
	nethttp "net/http"
	"net/http/httputil"
	"os"
	"path/filepath"

	"github.com/antchfx/htmlquery"
	"github.com/leliuga/cdk/http"
	"github.com/leliuga/cdk/types"
	"golang.org/x/net/html"
)

// FromResponse creates a new Response instance from a net/http response.
func FromResponse(response *nethttp.Response) *Response {
	return &Response{
		response: response,
		headers:  http.FromHeaders(response.Header),
	}
}

func (r *Response) Status() http.Status        { return http.Status(r.response.StatusCode) }
func (r *Response) Headers() http.Headers      { return r.headers }
func (r *Response) Cookies() []*nethttp.Cookie { return r.response.Cookies() }
func (r *Response) Body() io.ReadCloser        { return r.response.Body }
func (r *Response) ContentLength() int64       { return r.response.ContentLength }

// Unmarshal unmarshals the response body into the given value.
func (r *Response) Unmarshal(out any) error {
	contentType := r.headers.Get(http.HeaderContentType)

	return types.ContentTypeUnmarshal(contentType, r.Body(), out)
}

// HtmlQuery returns the html.Node of the response body.
func (r *Response) HtmlQuery(expression string) []*html.Node {
	doc, err := htmlquery.Parse(r.Body())
	if err != nil {
		return nil
	}

	return htmlquery.Find(doc, expression)
}

// Save saves the response body to the given filename.
func (r *Response) Save(filename string, progressWriter io.Writer) error {
	if err := os.MkdirAll(filepath.Dir(filename), 0755); err != nil {
		return err
	}

	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	if progressWriter == nil {
		_, err = io.Copy(f, r.Body())
		return err
	}

	_, err = io.Copy(f, io.TeeReader(r.Body(), progressWriter))
	return err
}

// Close closes the response body.
func (r *Response) Close() error {
	return r.Body().Close()
}

// Dump dumps the response.
func (r *Response) Dump(body bool) ([]byte, error) {
	dump, err := httputil.DumpResponse(r.response, body)
	if err != nil {
		return nil, err
	}

	return dump, err
}
