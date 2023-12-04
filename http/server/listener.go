package server

import (
	"fmt"
	"net"
	"time"
)

// NewListener returns a new Listener instance.
func NewListener(port uint32, keepAlive time.Duration) (*Listener, error) {
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return nil, err
	}

	return &Listener{
		TCPListener: l.(*net.TCPListener),
		keepAlive:   keepAlive,
	}, nil
}

// Accept implements the Accept method in the Listener interface; it waits for
func (l *Listener) Accept() (c net.Conn, err error) {
	if c, err = l.AcceptTCP(); err != nil {
		return
	} else if err = c.(*net.TCPConn).SetKeepAlive(true); err != nil {
		return
	}
	// Ignore error from setting the KeepAlivePeriod as some systems, such as
	// OpenBSD, do not support setting TCP_USER_TIMEOUT on IPPROTO_TCP
	_ = c.(*net.TCPConn).SetKeepAlivePeriod(l.keepAlive)
	return
}
