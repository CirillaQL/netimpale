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
