package proxy

import (
	"io"
	"net"
)

type UDPProxy struct {
	serverAddr *net.UDPAddr

	connCh chan *net.UDPConn
}

// NewUDPProxy 初始化UDPProxy
func NewUDPProxy(serverIP string, port int) *UDPProxy {
	addr := net.UDPAddr{
		IP:   net.IP(serverIP),
		Port: port,
	}
	return &UDPProxy{
		serverAddr: &addr,
		connCh:     make(chan *net.UDPConn),
	}
}

// StartListen 开始接受远方的UDP连接
func (proxy *UDPProxy) StartListen() {
	serverConn, err := net.ListenUDP("udp", proxy.serverAddr)
	if err != nil {
		LOG.Errorf("Get Remote UDP Package Failed")
	}
	go func(ch chan *net.UDPConn, Conn *net.UDPConn) {
		ch <- Conn
	}(proxy.connCh, serverConn)
}

// Proxy 转发方法，转发UDP的包
func (proxy *UDPProxy) Proxy(clientCh chan *net.UDPConn) {
	for {
		serverConn := <-proxy.connCh
		clientConn := <-clientCh
		LOG.Infof("Now UDP Get Both Package")
		go func(serverConn, clientConn *net.UDPConn) {
			proxyReq, err := io.Copy(serverConn, clientConn)
			if err != nil {
				LOG.Error("Can't Proxy UDP Connect from remote")
			}
			LOG.Infof("Proxy UDP Data Size: %d", proxyReq)
		}(serverConn, clientConn)
		go func(serverConn, clientConn *net.UDPConn) {
			proxyRsp, err := io.Copy(clientConn, serverConn)
			if err != nil {
				LOG.Error("Can't Proxy UDP Connect from remote")
			}
			LOG.Infof("Proxy UDP Data Size: %d", proxyRsp)
		}(serverConn, clientConn)
	}
}
