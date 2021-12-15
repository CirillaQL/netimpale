package server

import (
	"net"
)

// TCPManager Http连接的管理结构体，目前用来处理Http连接的相关请求
type TCPManager struct {
	TCPServerAddr     *net.TCPAddr
	TCPServerListener *net.TCPListener
	TCPClientAddr     *net.TCPAddr
	TCPClientListener *net.TCPListener

	serverChan chan *net.TCPConn
	clientChan chan *net.TCPConn
}

// NewTCPManager 创建HTTPManager
func NewTCPManager(network, ServerAddr, ClientAddr string) *TCPManager {
	LOG.Infof("Create TCPManager")
	tcpServerAddr, err := net.ResolveTCPAddr(network, ServerAddr)
	if err != nil {
		LOG.Errorf("Create TCPManager TCPServerAddr Failed: %v", err)
	}
	tcpServerListener, err := net.ListenTCP(network, tcpServerAddr)
	if err != nil {
		LOG.Errorf("Create TCPManager TCPServerListener Failed: %v", err)
	}
	tcpClientAddr, err := net.ResolveTCPAddr(network, ClientAddr)
	if err != nil {
		LOG.Errorf("Create TCPManager TCPClientAddr Failed: %v", err)
	}
	tcpClientListener, err := net.ListenTCP(network, tcpClientAddr)
	if err != nil {
		LOG.Errorf("Create TCPManager TCPClientListener Failed: %v", err)
	}
	serverChan := make(chan *net.TCPConn)
	clientChan := make(chan *net.TCPConn)
	return &TCPManager{
		TCPServerAddr:     tcpServerAddr,
		TCPServerListener: tcpServerListener,
		TCPClientAddr:     tcpClientAddr,
		TCPClientListener: tcpClientListener,

		serverChan: serverChan,
		clientChan: clientChan,
	}
}

// handleConn 获取TCPConnect连接，并将其置于对应的
func (tcp *TCPManager) handlerTCPConn(ch chan *net.TCPConn, listener *net.TCPListener) {
	for {
		conn, err := (*listener).AcceptTCP()
		if err != nil {
			LOG.Errorf("TCPManager Can't Handle Connect: %v", err)
		}
		ch <- conn
	}
}

// Run 运行
func (tcp *TCPManager) Run() {
	if tcp.TCPServerAddr == nil || tcp.TCPServerListener == nil {
		LOG.Errorf("TCPManager Server Init Failed.")
	}
	if tcp.TCPClientAddr == nil || tcp.TCPClientListener == nil {
		LOG.Errorf("TCPManager Client Init Failed.")
	}
	LOG.Infof("TCPManager Run...")
	// 将conn写入对应的channel中
	go tcp.handlerTCPConn(tcp.serverChan, tcp.TCPServerListener)
	go tcp.handlerTCPConn(tcp.clientChan, tcp.TCPClientListener)
}

func (tcp *TCPManager) Proxy() {

}
