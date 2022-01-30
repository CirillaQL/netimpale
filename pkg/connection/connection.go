/*
	Connect是用来封装客户端与服务端之间通信的Message桥梁的
	新Struct，其中包含ID与Context等相关信息，方便后续连接
	池的使用。
*/
package connection

import (
	"context"
	"github.com/satori/go.uuid"
	"net"
	"netimpale/utils/log"
	"time"
)

var LOG = log.LOG

type Conn struct {
	ID        string             //标识连接的ID
	TCPConn   *net.TCPConn       //具体连接的TCPConn实例
	Ctx       context.Context    //添加操作Context
	CtxCancel context.CancelFunc //ctx结束通知
	err       error
}

// NewConn 创建连接时客户端向服务端发起创建连接
func NewConn(serverAddr string) (c *Conn, err error) {
	// 初始化连接
	c = &Conn{
		ID:  uuid.Must(uuid.NewV4(), nil).String(),
		err: nil,
	}
	// 拨号
	var conn net.Conn
	if conn, err = net.DialTimeout("tcp", serverAddr, 5*time.Second); err != nil {
		// TODO：Support config dialTimeout time
		LOG.Errorf("Create Conn Failed. Failed to Connect: %v", err)
	} else {
		c.TCPConn = conn.(*net.TCPConn)
	}
	if err = c.TCPConn.SetKeepAlive(true); err != nil {
		LOG.Errorf("Conn Set KeepAlive Failed. Error: %v", err)
	}
	c.Ctx, c.CtxCancel = context.WithCancel(context.Background())
	return c, nil
}
