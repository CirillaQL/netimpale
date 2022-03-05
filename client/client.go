package client

import (
	"io"
	"net"
	"netimpale/pkg/pool"
	"netimpale/utils/log"
)

var LOG = log.LOG

// Client 部署在客户端机器上执行的Client总管控
type Client struct {
	Pool *pool.Pool
}

// NewClient 初始化Client类
func NewClient(size uint8) (c *Client, err error) {
	_pool, err := pool.NewClientPool(size)
	if err != nil {
		LOG.Errorf("Client Init Pool Failed. Error: %v", err)
		return nil, err
	}
	c = &Client{Pool: _pool}
	return c, nil
}

// TransTCP 转发TCP连接，实现穿透
func (c *Client) TransTCP(remoteConn *net.TCPConn) {
	poolConn, err := c.Pool.Get()
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
