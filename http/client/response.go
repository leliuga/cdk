package client

import (
	"fmt"
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

func (r *Response) Status() http.Status        { return http.Status(r.response.StatusCode) }
func (r *Response) Headers() http.Headers      { return r.headers }
func (r *Response) Cookies() []*nethttp.Cookie { return r.response.Cookies() }
func (r *Response) Body() io.ReadCloser        { return r.response.Body }
func (r *Response) ContentLength() int64       { return r.response.ContentLength }

// Unmarshal unmarshals the response body into the given value.
func (r *Response) Unmarshal(v any) error {
	cType := r.headers[http.HeaderContentType]
	ct := types.ParseContentType(cType)
	if !ct.Validate() {
		return fmt.Errorf("unsupported content type: %s", cType)
	}

	return ct.Unmarshal(r.Body(), v)
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
func (r *Response) Save(filename string, writer io.Writer) error {
	if err := os.MkdirAll(filepath.Dir(filename), 0755); err != nil {
		return err
	}

	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = io.Copy(f, io.TeeReader(r.Body(), writer))

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
