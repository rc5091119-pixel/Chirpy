package main

import (
	"context"
	"crypto/tls"
	"net"
	"time"
)

type Server struct {
	Handler                      Handler
	DisableGeneralOptionsHandler bool
	TLSConfig                    *tls.Config
	ReadTimeout                  time.Duration
	ReadHeaderTimeout            time.Duration
	IdleTimeout                  time.Duration
	MaxHeaderBytes               int
	TLSNextProto                 map[string]func(*Server, *tls.Conn, Handler)
	ConnState                    func(net.Conn, ConnState)
	BaseContext                  func(net.Listener) context.Context
	ConnContext                  func(ctx context.Context, c net.Conn) context.Context
	HTTP2                        *HTTP2Config
	Protocols                    *Protocols
}
