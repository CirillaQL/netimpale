package server

import (
	"io"
	"net"
	"sync"
)

// TCPManager Http连接的管理结构体，目前用来处理Http连接的相关请求
type TCPManager struct {
	TCPServerAddr     *net.TCPAddr
	TCPServerListener *net.TCPListener
	TCPClientAddr     *net.TCPAddr
	TCPClientListener *net.TCPListener
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
	return &TCPManager{
		TCPServerAddr:     tcpServerAddr,
		TCPServerListener: tcpServerListener,
		TCPClientAddr:     tcpClientAddr,
		TCPClientListener: tcpClientListener,
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
	// 首先应该监控与客户端的TCP连接是否已经建立，否则无法接受外界的TCP的连接
	for {
		wg := &sync.WaitGroup{}
		conn, err := tcp.TCPServerListener.AcceptTCP()
		if err != nil {
			LOG.Infof("Accept TCP Conn Failed, Error: %v", err)
		}
		go tcp.ConnHandler(conn)
	}
}

// ConnHandler 处理来的TCPConn连接
func (tcp *TCPManager) ConnHandler(conn *net.TCPConn) {
	io.Copy()
}
