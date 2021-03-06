package server

import (
	"io"
	"net"
	"net/http"
	"netimpale/pkg/pool"
	"netimpale/utils/log"
)

var LOG = log.LOG

// Server 服务器结构体
type Server struct {
	Pool       *pool.Pool
	httpServer *http.Server
}

// NewServer 初始化Server类
func NewServer(size uint8, addr string) (s *Server, err error) {
	_pool, err := pool.NewServerPool(size)
	if err != nil {
		LOG.Errorf("Server Init Pool Failed. Error: %v", err)
		return nil, err
	}
	_server := http.Server{
		Addr: addr,
	}
	s = &Server{
		Pool:       _pool,
		httpServer: &_server,
	}
	return s, nil
}

// TransTCP 转发TCP
func (s *Server) TransTCP(remoteConn *net.TCPConn) {
	poolConn, err := s.Pool.Get()
	if err != nil {
		LOG.Errorf("Get Conn from Pool Failed. Error: %v", err)
	}
	go func() {
		_, err := io.Copy(poolConn.TCPConn, remoteConn)
		if err != nil {
			LOG.Errorf("Handler Transport TCP failed. Error: %v", err)
		}
	}()
	go func() {
		_, err := io.Copy(remoteConn, poolConn.TCPConn)
		if err != nil {
			LOG.Errorf("Handler Transport TCP failed. Error: %v", err)
		}
	}()
}
