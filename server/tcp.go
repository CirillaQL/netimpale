package server

import (
	"net"
)

// TCPManager Http连接的管理结构体，目前用来处理Http连接的相关请求
type TCPManager struct {
	TCPAddr     *net.TCPAddr
	TCPListener *net.TCPListener
}

// NewTCPManager 创建HTTPManager
func NewTCPManager(network, addr string) *TCPManager {
	LOG.Infof("Create TCPManager")
	tcpAddr, err := net.ResolveTCPAddr(network, addr)
	if err != nil {
		LOG.Errorf("Create TCPManger TCPAddr Failed: %v", err)
	}
	tcpListener, err := net.ListenTCP(network, tcpAddr)
	if err != nil {
		LOG.Errorf("Create TCPManger TCPListener Failed: %v", err)
	}
	return &TCPManager{
		TCPAddr:     tcpAddr,
		TCPListener: tcpListener,
	}
}

// Run 运行
func (tcp *TCPManager) Run() {
	if tcp.TCPAddr == nil || tcp.TCPListener == nil {
		LOG.Errorf("TCPManager Init Failed.")
	}
	LOG.Infof("TCPManager Run...")
	for {
		conn, err := tcp.TCPListener.AcceptTCP()
		if err != nil {
			LOG.Infof("Accept TCP Conn Failed, Error: %v", err)
		}
		go tcp.ConnHandler(conn)
	}
}

// ConnHandler 处理来的TCPConn连接
func (tcp *TCPManager) ConnHandler(conn *net.TCPConn) {

}
