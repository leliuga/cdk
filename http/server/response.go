package server

import (
	"bufio"
	"net"
	nethttp "net/http"

	"github.com/leliuga/cdk/http"
)

// NewResponse creates a new instance of Response.
func NewResponse(w nethttp.ResponseWriter) (r *Response) {
	return &Response{Writer: w}
}

// Header returns the header map for the writer that will be sent by
// WriteHeader. Changing the header after a call to WriteHeader (or Write) has
// no effect unless the modified headers were declared as trailers by setting
// the "Trailer" header before the call to WriteHeader (see example)
// To suppress implicit response headers, set their value to nil.
// Example: https://golang.org/pkg/net/http/#example_ResponseWriter_trailers
func (r *Response) Header() nethttp.Header {
	return r.Writer.Header()
}

// WriteHeader sends an HTTP response header with status code. If WriteHeader is
// not called explicitly, the first call to Write will trigger an implicit
// WriteHeader(http.StatusOK). Thus explicit calls to WriteHeader are mainly
// used to send error codes.
func (r *Response) WriteHeader(status http.Status) {
	if r.Committed {
		// r.echo.Logger.Warn("response already committed")
		return
	}
	r.Status = status
	r.Writer.WriteHeader(int(r.Status))
	r.Committed = true
}

// Write writes the data to the connection as part of an HTTP reply.
func (r *Response) Write(b []byte) (n int, err error) {
	if !r.Committed {
		if r.Status == 0 {
			r.Status = http.StatusOK
		}
		r.WriteHeader(r.Status)
	}
	n, err = r.Writer.Write(b)
	r.Size += int64(n)

	return
}

// Flush implements the http.Flusher interface to allow an HTTP handler to flush
// buffered data to the client.
// See [http.Flusher](https://golang.org/pkg/net/http/#Flusher)
func (r *Response) Flush() {
	r.Writer.(nethttp.Flusher).Flush()
}

// Hijack implements the http.Hijacker interface to allow an HTTP handler to
// take over the connection.
// See [http.Hijacker](https://golang.org/pkg/net/http/#Hijacker)
func (r *Response) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	return r.Writer.(nethttp.Hijacker).Hijack()
}

// Unwrap returns the original http.ResponseWriter.
// ResponseController can be used to access the original http.ResponseWriter.
// See [https://go.dev/blog/go1.20]
func (r *Response) Unwrap() nethttp.ResponseWriter {
	return r.Writer
}

// acquire sets ResponseWriter.
func (r *Response) acquire(w nethttp.ResponseWriter) {
	r.Writer = w
	r.Size = 0
	r.Status = http.StatusOK
	r.Committed = false
}
