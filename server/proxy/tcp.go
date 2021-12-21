package proxy

import (
	"io"
	"net"
	"netimpale/utils/log"
)

var LOG = log.LOG

type TCPProxy struct {
	serverAddr     *net.TCPAddr
	serverListener *net.TCPListener

	connCh chan *net.TCPConn
}

// NewTCPProxy 初始化TCPProxy
func NewTCPProxy(network, serverAddr string) *TCPProxy {
	LOG.Infof("Creata TCP Proxy")
	tcpAddr, err := net.ResolveTCPAddr(network, serverAddr)
	if err != nil {
		LOG.Errorf("Create TCPProxy TCPServerAddr Failed: %v", err)
	}
	tcpListener, err := net.ListenTCP(network, tcpAddr)
	if err != nil {
		LOG.Errorf("Create TCPProxy TCPServerListener Failed: %v", err)
	}

	return &TCPProxy{
		serverAddr:     tcpAddr,
		serverListener: tcpListener,
		connCh:         make(chan *net.TCPConn),
	}
}

// StartListener 开始监听远程的连接
func (proxy *TCPProxy) StartListener() {
	for {
		remoteConn, err := proxy.serverListener.AcceptTCP()
		if err != nil {
			LOG.Errorf("TCPProxy Can't Handle Connect: %v", err)
		}
		go func(ch chan *net.TCPConn, Conn *net.TCPConn) {
			ch <- Conn
		}(proxy.connCh, remoteConn)
	}
}

// HandleTCPConnection 远程连接并转发
func (proxy *TCPProxy) HandleTCPConnection(clientCh chan *net.TCPConn) {
	for {
		serverConn := <-proxy.connCh
		clientConn := <-clientCh
		LOG.Infof("Now TCP Get Both Conn")
		go func(serverConn, clientConn *net.TCPConn) {
			proxyReq, err := io.Copy(serverConn, clientConn)
			if err != nil {
				LOG.Error("Can't Proxy TCP Connect from remote")
			}
			LOG.Infof("Proxy TCP Data Size: %d", proxyReq)
		}(serverConn, clientConn)
		go func(serverConn, clientConn *net.TCPConn) {
			proxyRsp, err := io.Copy(clientConn, serverConn)
			if err != nil {
				LOG.Error("Can't Proxy TCP Connect from remote")
			}
			LOG.Infof("Proxy TCP Data Size: %d", proxyRsp)
		}(serverConn, clientConn)
	}
}
