package server

import "net"

// TCPManager TCP连接的管理结构体，目前用来处理TCP连接的相关请求
type TCPManager struct {
	RemoteListener *net.TCPListener //外部TCP的监听器
}

// NewTCPManager 创建TCPManager
func NewTCPManager(serverAddr string) *TCPManager {
	tcpAddr, err := net.ResolveTCPAddr("tcp", serverAddr)
	if err != nil {
		LOG.Errorf("Create TCPManager TCPAddr Failed. Error: %v", err)
	}
	tcpListener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		LOG.Errorf("Create TCPManager TCPListener Failed. Error: %v", err)
	}
	return &TCPManager{RemoteListener: tcpListener}
}

// Start 启动服务端TCPManager，主要开始维护与客户端的连接
func (tcpManager *TCPManager) Start() {
	// TODO：从与客户端的TCP连接池中取出对应的TCPConn

}

// Run 执行转发
func (tcpManager *TCPManager) Run() {

}
