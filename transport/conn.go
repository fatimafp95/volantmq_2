package transport

import (
	"github.com/lucas-clemente/quic-go"
	"github.com/VolantMQ/volantmq/auth"
	"github.com/VolantMQ/volantmq/metrics"
)

// Conn is wrapper to net.Conn
// implemented to encapsulate bytes statistic

type Sess interface {
	quic.Stream
}
type sess struct {
	Sess quic.Stream
	stat metrics.Bytes
}

// Handler ...
type Handler interface {
	OnConnection(Sess, *auth.Manager) error
}


func newSess(cn quic.Stream, stat metrics.Bytes) *sess {
	c := &sess{
		Sess: cn,
		stat: stat,
	}
	return c
}

// Read ...


func (c *sess) Read(b []byte) (int, error){
	n, err := c.Sess.Read(b)
	c.stat.OnRecv(n)

	return n, err
}


func (c *sess) Write(b []byte) (int, error){
	n, err := c.Sess.Write(b)
	c.stat.OnSent(n)

	return n, err
}
