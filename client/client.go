package client

import "netimpale/pkg/pool"

// Client 部署在客户端机器上执行的Client总管控
type Client struct {
	Pool *pool.Pool
}
