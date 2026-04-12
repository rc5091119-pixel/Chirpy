
A "migration" in Goose is just a .sql file with some SQL queries and some special comments. Our first migration should just create a users table. The simplest format for these files is:

psql "postgres://postgres:postgres@localhost:5432/chirpy?sslmode=disable"

goose postgres "postgres://postgres:postgres@localhost:5432/chirpy?sslmode=disable" up

goose -dir sql/schema postgres "$DB_URL" up



















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
