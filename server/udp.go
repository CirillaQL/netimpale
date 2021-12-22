package server

import (
	"net"
	"netimpale/pkg/network_utils"
	"netimpale/server/proxy"
)

// UDPManager UDP连接的管理结构体，目前用来处理UDP连接的相关请求
type UDPManager struct {
	proxy      *proxy.UDPProxy
	clientAddr *net.UDPAddr

	clientCh chan *net.UDPConn
}

// NewUDPManager 创建UDPManager
func NewUDPManager(network, ServerAddr, ClientAddr string) *UDPManager {
	LOG.Info("Create UDPManager")
	serverAddress, serverPort := network_utils.ParseIPv4(ServerAddr)
	udpProxy := proxy.NewUDPProxy(serverAddress, serverPort)
	clientAddress, clientPort := network_utils.ParseIPv4(ClientAddr)
	client := net.UDPAddr{
		IP:   net.IP(clientAddress),
		Port: clientPort,
	}

	return &UDPManager{
		proxy:      udpProxy,
		clientAddr: &client,
		clientCh:   make(chan *net.UDPConn),
	}
}

// ConnectClient 开始接受客户端方向的UDP连接
func (udp *UDPManager) ConnectClient() {
	for {
		clientConn, err := net.ListenUDP("udp", udp.clientAddr)
		if err != nil {
			LOG.Errorf("Get Remote UDP Package Failed")
		}
		go func(ch chan *net.UDPConn, Conn *net.UDPConn) {
			ch <- Conn
		}(udp.clientCh, clientConn)
	}
}

// Run 运行转发
func (udp *UDPManager) Run() {
	go udp.proxy.StartListen()
	go udp.ConnectClient()
	go udp.proxy.Proxy(udp.clientCh)
}
