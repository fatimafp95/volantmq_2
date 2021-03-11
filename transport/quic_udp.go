package transport

import (
"context"
"crypto/tls"
"fmt"
"net"
"time"

"github.com/VolantMQ/volantmq/types"
"github.com/lucas-clemente/quic-go"
)

type ConfigQUIC struct {
	Scheme		string
	TLS			*tls.Config
	transport	*Config
}

type quic_udp struct {
	baseConfig 
	tls		*tls.Config
	listener	quic.Listener 
	//listener	quic.EarlyListener //Listener for UDP: 0RTT
}
// NewConfigQUIC crea nueva configuración para QUIC
func NewConfigQuic(transport *Config) *ConfigQUIC {
	return &ConfigQUIC{
		Scheme:    "udp",
		transport: transport,
	}
}

type Token struct {
	// IsRetryToken encodes how the client received the token. There are two ways:
	// * In a Retry packet sent when trying to establish a new connection.
	// * In a NEW_TOKEN frame on a previous connection.
	IsRetryToken bool
	RemoteAddr   string
	SentTime     time.Time
}

func AcceptToken (clientAddr net.Addr,  Token *quic.Token ) bool{
	if clientAddr == nil{
		return true
	}


	if Token == nil {
		return true
	}
	return true
}

//InternalConfig: configuración interna del servidor
//Provider: interfaces "obligatorias"
func NewQUIC(config *ConfigQUIC, internal *InternalConfig) (Provider, error){
	l := &quic_udp{}
	l.quit = make(chan struct{})
	l.InternalConfig = *internal
	l.config = *config.transport
	l.tls = config.TLS

	var err error

	//llamada al socket de quic
	
	//if l.listener, err = quic.ListenAddr(config.transport.Host+":"+config.transport.Port, config.TLS, nil); err != nil{
	quicConfig := &quic.Config{
		AcceptToken: AcceptToken,
	}
	if l.listener, err = quic.ListenAddrEarly(config.transport.Host+":"+config.transport.Port, config.TLS, quicConfig); err != nil{
		return nil ,err
	}

	sess_chann := make(chan quic.Session)

	go func() {
		for {
			sess, err := l.listener.Accept(context.Background())
			if err != nil {
				fmt.Println("Error accepting: ", err.Error())
				return
			}
			sess_chann <- sess
		}
	}()
	defer l.listener.Close()
	for{
	select{
	case sess := <- sess_chann: //QUIC SESSION
		fmt.Println("Established QUIC connection")
		stream, err := sess.AcceptStream(context.Background())
		if err != nil{
			panic(err)
		}
		l.handleConnection(stream)
	}
	}
	return l, nil
}

//Funciones provider
//Ready
func (l *quic_udp) Ready() error{
	if err := l.baseReady(); err != nil{
		return err
	}

	return nil
}
//Alive
func (l *quic_udp) Alive() error{
	err := l.baseReady()
	if err != nil {
		return err
	}

	return nil
}
//Close quic listener
func (l *quic_udp) Close() error{
	var err error

	l.onceStop.Do(func(){
		close(l.quit)

		err = l.listener.Close()
		l.listener = nil
		l.log = nil

	})
	return err
}

func (l *quic_udp) Serve() error {

	accept := make(chan error, 1)
	defer close(accept)

	for{
		err := l.AcceptPool.ScheduleTimeout(time.Millisecond,func(){
			var er error
			defer func(){
				accept <-er
			}()
			select{
			case <-l.quit:
				er = types.ErrClosed
				return
			default:
			}
		})

		if err != nil && err != types.ErrScheduleTimeout{
			break
		}else if err == types.ErrScheduleTimeout{
			continue
		}
		err = <- accept

		if err != nil {
			if ne, ok := err.(net.Error); ok && ne.Temporary() {
				delay := 5 * time.Millisecond
				time.Sleep(delay)
			}else{
				break
			}

		}
	}
	return nil
}
