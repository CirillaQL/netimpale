package server

import (
	"io"
	"net"
	"netimpale/pkg/pool"
	"netimpale/utils/log"
)

var LOG = log.LOG

// Server 服务器结构体
type Server struct {
	Pool *pool.Pool
}

// NewServer 初始化Server类
func NewServer(size uint8) (s *Server, err error) {
	_pool, err := pool.NewServerPool(size)
	if err != nil {
		LOG.Errorf("Server Init Pool Failed. Error: %v", err)
		return nil, err
	}
	s = &Server{Pool: _pool}
	return s, nil
}

// TransTCP 转发TCP
func (s *Server) TransTCP(remoteConn *net.TCPConn) {
	poolConn, err := s.Pool.Get()
	if err != nil {
		LOG.Errorf("Get Conn from Pool Failed. Error: %v", err)
	}
	go io.Copy(poolConn.TCPConn, remoteConn)
	go io.Copy(remoteConn, poolConn.TCPConn)
}
