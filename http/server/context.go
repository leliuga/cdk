package server

import (
	"fmt"
	nethttp "net/http"

	"github.com/leliuga/cdk/http"
	"github.com/leliuga/cdk/types"
)

func NewContext() *Context {
	return &Context{
		response: &Response{},
	}
}

// acquire acquires a context from the pool.
func (c *Context) acquire(writer nethttp.ResponseWriter, request *nethttp.Request) {
	c.request = request
	c.response.acquire(writer)
	c.handler = NotFoundHandler
}

// Unmarshal unmarshals the request body into the given value.
func (c *Context) Unmarshal(out any) error {
	contentType := c.Header("Content-Type")

	if ct := types.ParseContentType(contentType); ct.Validate() {
		return ct.Unmarshal(c.request.Body, out)
	}

	return fmt.Errorf("unsupported content type: %s", contentType)
}

// Method returns the request method.
func (c *Context) Method() string {
	return c.request.Method
}

// Host returns the request host.
func (c *Context) Host() string {
	return c.request.Host
}

// Path returns the request path.
func (c *Context) Path() string {
	return c.request.URL.Path
}

// SetStatus sets the response status code.
func (c *Context) SetStatus(status http.Status) *Context {
	c.response.WriteHeader(status)

	return c
}

// Header returns the value of a header field.
func (c *Context) Header(key string) string {
	return c.request.Header.Get(key)
}

// Headers returns the values of a header field.
func (c *Context) Headers(key string) []string {
	return c.request.Header.Values(key)
}

// SetHeader sets the value of a header field.
func (c *Context) SetHeader(key, value string) *Context {
	c.response.Header().Set(key, value)

	return c
}

// Cookie returns the named cookie provided in the request or nil if not found.
func (c *Context) Cookie(name string) (*nethttp.Cookie, error) {
	return c.request.Cookie(name)
}

// Cookies returns the HTTP cookies sent with the request.
func (c *Context) Cookies() []*nethttp.Cookie {
	return c.request.Cookies()
}

// SetCookie sets a cookie.
func (c *Context) SetCookie(cookie *nethttp.Cookie) *Context {
	c.response.Header().Add("Set-Cookie", cookie.String())

	return c
}
