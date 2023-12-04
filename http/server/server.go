package server

import (
	"context"
	"fmt"
	"net"
	"net/http"
)

// NewServer creates a new Server instance.
func NewServer(options *Options) *Server {
	s := &Server{
		Options: options,
		server: &http.Server{
			Addr:              fmt.Sprintf(":%d", options.Port),
			ReadTimeout:       options.ReadTimeout,
			ReadHeaderTimeout: options.ReadHeaderTimeout,
			WriteTimeout:      options.WriteTimeout,
			IdleTimeout:       options.IdleTimeout,
		},
	}

	s.server.Handler = s
	s.pool.New = func() any {
		return &Context{
			response: &Response{},
		}
	}

	s.server.ConnState = func(conn net.Conn, state http.ConnState) {
		switch state {
		case http.StateNew:
			fmt.Println("new connection")
		case http.StateActive:
			fmt.Println("active connection")
		case http.StateIdle:
			fmt.Println("idle connection")
		case http.StateHijacked, http.StateClosed:
			fmt.Println("closed connection")
		}
	}

	return s
}

// ServeHTTP implements the http.Handler interface.
func (s *Server) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Server", s.Name)

	// acquire context
	c := s.pool.Get().(*Context)
	c.acquire(writer, request)

	// Release context
	s.pool.Put(c)
}

// Serve serves the server.
func (s *Server) Serve() error {
	listener, err := NewListener(s.Port, s.KeepAliveTimeout)
	if err != nil {
		return err
	}

	if (s.CertificateFile != "") && (s.CertificateKeyFile != "") {
		return s.server.ServeTLS(listener, s.CertificateFile, s.CertificateKeyFile)
	}

	return s.server.Serve(listener)
}

// Shutdown gracefully shuts down the server without interrupting any active connections.
func (s *Server) Shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
