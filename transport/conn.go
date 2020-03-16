package transport

import (
//	"errors"
//	"net"
//	"os"
	"github.com/lucas-clemente/quic-go"
	"github.com/VolantMQ/volantmq/auth"
	"github.com/VolantMQ/volantmq/metrics"
)

// Conn is wrapper to net.Conn
// implemented to encapsulate bytes statistic
/*type Conn interface {
	net.Conn
}

type conn struct {
	net.Conn
	stat metrics.Bytes
}

var _ Conn = (*conn)(nil)
*/
type Sess interface {
	quic.Stream
}
type sess struct {
	Sess quic.Stream
	stat metrics.Bytes
}
//var _ Sess = (*sess)(nil)

// Handler ...
type Handler interface {
	OnConnection(Sess, *auth.Manager) error
//	OnConnection(Conn, *auth.Manager) error
}

/*func newConn(cn net.Conn, stat metrics.Bytes) *conn {
	c := &conn{
		Conn: cn,
		stat: stat,
	}

	return c
}*/
func newSess(cn quic.Stream, stat metrics.Bytes) *sess {
	c := &sess{
		Sess: cn,
		stat: stat,
	}
	return c
}

// Read ...

//func (c *conn) Read(b []byte) (int, error) {
//	n, err := c.Conn.Read(b)
func (c *sess) Read(b []byte) (int, error){
	n, err := c.Sess.Read(b)
	c.stat.OnRecv(n)

	return n, err
}

// Write ...
//func (c *conn) Write(b []byte) (int, error) {
//	n, err := c.Conn.Write(b)
func (c *sess) Write(b []byte) (int, error){
	n, err := c.Sess.Write(b)
	c.stat.OnSent(n)

	return n, err
}
/*
// File ...
func (c *conn) File() (*os.File, error) {
	switch t := c.Conn.(type) {
	case *net.TCPConn:
		return t.File()
	}

	return nil, errors.New("not implemented")
}
*/
