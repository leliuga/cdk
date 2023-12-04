package server

import (
	"net"
	nethttp "net/http"
	"sync"
	"time"

	"github.com/leliuga/cdk/http"
	"github.com/leliuga/cdk/render"
	"github.com/leliuga/cdk/types"
)

type (
	// Server is a server
	Server struct {
		*Options

		server *nethttp.Server
		pool   sync.Pool
	}

	// Options is a set of options for the server.
	Options struct {
		Name               string           `json:"name"`
		Port               uint32           `json:"port"                 env:"PORT"`
		CertificateFile    string           `json:"certificate_file"     env:"CERTIFICATE_FILE"`
		CertificateKeyFile string           `json:"certificate_key_file" env:"CERTIFICATE_KEY_FILE"`
		Renderer           render.IRenderer `json:"-"`
		ReadTimeout        time.Duration    `json:"read_timeout"         env:"READ_TIMEOUT"`
		ReadHeaderTimeout  time.Duration    `json:"read_header_timeout"  env:"READ_HEADER_TIMEOUT"`
		WriteTimeout       time.Duration    `json:"write_timeout"        env:"WRITE_TIMEOUT"`
		IdleTimeout        time.Duration    `json:"idle_timeout"         env:"IDLE_TIMEOUT"`
		KeepAliveTimeout   time.Duration    `json:"keep_alive_timeout"   env:"KEEP_ALIVE_TIMEOUT"`
		ShutdownTimeout    time.Duration    `json:"shutdown_timeout"     env:"SHUTDOWN_TIMEOUT"`
	}

	// Context is a request context.
	Context struct {
		request  *nethttp.Request
		response *Response
		path     string
		handler  HandlerFunc
		lock     sync.RWMutex
	}

	// Error is a server error.
	Error struct {
		Status   http.Status `json:"status"`
		Message  interface{} `json:"message"`
		Internal error       `json:"-"` // Stores the error returned by an external dependency
	}

	// Listener is a TCP listener.
	Listener struct {
		*net.TCPListener
		keepAlive time.Duration
	}

	Route struct {
		Name        string
		Description string
		Method      *http.Method
		Path        types.Path
		Handler     HandlerFunc
	}

	Router struct {
		routes []Route
	}

	// Response is a server response.
	Response struct {
		Writer    nethttp.ResponseWriter
		Status    http.Status
		Size      int64
		Committed bool
	}

	// Option defines a function to customize Options.
	Option func(*Options)

	// MiddlewareFunc defines a function to process middleware.
	MiddlewareFunc func(next HandlerFunc) HandlerFunc

	// HandlerFunc defines a function to serve HTTP requests.
	HandlerFunc func(ctx *Context) error

	// HTTPErrorHandler is a centralized HTTP error handler.
	HTTPErrorHandler func(err error, c Context)
)
