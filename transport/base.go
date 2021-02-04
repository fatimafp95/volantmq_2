package transport

import (
	"errors"
	"sync"

	"go.uber.org/zap"
	////
	"github.com/lucas-clemente/quic-go"
	///
	"github.com/VolantMQ/volantmq/auth"
	"github.com/VolantMQ/volantmq/metrics"
	"github.com/VolantMQ/volantmq/types"
)

// Config is base configuration object used by all transports
type Config struct {
	// AuthManager
	AuthManager *auth.Manager
	Host string
	// Port tcp port to listen on
	Port string
}

// InternalConfig used by server implementation to configure internal specific needs
type InternalConfig struct {
	Handler
	AcceptPool types.Pool
	Metrics    metrics.Bytes
}

type baseConfig struct {
	InternalConfig
	config       Config
	onConnection sync.WaitGroup // nolint: structcheck
	onceStop     sync.Once      // nolint: structcheck
	quit         chan struct{}  // nolint: structcheck
	log          *zap.SugaredLogger
	protocol     string
}

// Provider is interface that all of transports must implement
type Provider interface {
	Protocol() string
	Serve() error
	Close() error
	Port() string
	Ready() error
	Alive() error
}

var (
	// ErrListenerIsOff ...
	ErrListenerIsOff = errors.New("listener is off")
)

// Port return tcp port used by transport
func (c *baseConfig) Port() string {
	return c.config.Port
}

// Protocol return protocol name used by transport
func (c *baseConfig) Protocol() string {
	return c.protocol
}

func (c *baseConfig) baseReady() error {
	select {
	case <-c.quit:
		return ErrListenerIsOff
	default:
	}

	return nil
}

// handleConnection is for the broker to handle an incoming connection from a client
func (c *baseConfig) handleConnection(sess quic.Stream) {
	if c == nil {
		c.log.Error("Invalid connection type")
		return
	}

	var err error

	defer func(){
		if err != nil{
			_ = sess.Close()
		}
	}()
	err = c.OnConnection(sess, c.config.AuthManager)
}

