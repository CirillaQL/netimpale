package server

import (
	"net"
	"netimpale/server/proxy"
	"time"
)

// TCPManager TCP连接的管理结构体，目前用来处理TCP连接的相关请求
type TCPManager struct {
	proxy          *proxy.TCPProxy
	clientAddr     *net.TCPAddr
	clientListener *net.TCPListener

	clientCh chan *net.TCPConn
}

// NewTCPManager 创建TCPManager
func NewTCPManager(network, ServerAddr, ClientAddr string) *TCPManager {
	LOG.Infof("Create TCPManager")
	tcpProxy := proxy.NewTCPProxy(network, ServerAddr)
	tcpAddr, err := net.ResolveTCPAddr(network, ClientAddr)
	if err != nil {
		LOG.Errorf("Create TCPManager TCPClientAddr Failed: %v", err)
	}
	tcpListener, err := net.ListenTCP(network, tcpAddr)
	if err != nil {
		LOG.Errorf("Create TCPManager TCPClientListener Failed: %v", err)
	}
	return &TCPManager{
		proxy:          tcpProxy,
		clientAddr:     tcpAddr,
		clientListener: tcpListener,
		clientCh:       make(chan *net.TCPConn),
	}
}

// ConnectClient 连接客户端
func (tcp *TCPManager) ConnectClient() {
	for {
		remoteConn, err := tcp.clientListener.AcceptTCP()
		if err != nil {
			LOG.Errorf("TCPProxy Can't Handle Connect: %v", err)
		}
		err = remoteConn.SetKeepAlive(true)
		if err != nil {
			LOG.Errorf("TCPProxy Can't Handle Connect: %v", err)
		}
		err = remoteConn.SetKeepAlivePeriod(2 * time.Second)
		if err != nil {
			LOG.Errorf("TCPProxy Can't Handle Connect: %v", err)
		}
		go func(ch chan *net.TCPConn, Conn *net.TCPConn) {
			ch <- Conn
		}(tcp.clientCh, remoteConn)
	}
}

// Run 运行转发
func (tcp *TCPManager) Run() {
	//首先运行Proxy的StartListener，用来Handle 远程TCP连接
	go tcp.proxy.StartListener()
	//此时需获取来自客户端的net.TCPConn
	go tcp.ConnectClient()
	go tcp.proxy.HandleTCPConnection(tcp.clientCh)
}
