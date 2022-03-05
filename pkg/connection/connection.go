/*
	Connect是用来封装客户端与服务端之间通信的Message桥梁的
	新Struct，其中包含ID与Context等相关信息，方便后续连接
	池的使用。
*/
package connection

import (
	"bufio"
	"context"
	"net"
	"netimpale/utils/log"
	"time"

	uuid "github.com/satori/go.uuid"
)

var LOG = log.LOG

type Conn struct {
	ID        string             //标识连接的ID
	TCPConn   *net.TCPConn       //具体连接的TCPConn实例
	Ctx       context.Context    //添加操作Context
	CtxCancel context.CancelFunc //ctx结束通知
}

// NewConn 创建连接时客户端向服务端发起创建连接
func NewConn(serverAddr string) (c *Conn, err error) {
	// 初始化连接
	c = &Conn{
		ID: uuid.Must(uuid.NewV4(), nil).String(),
	}
	// 拨号
	var conn net.Conn
	if conn, err = net.DialTimeout("tcp", serverAddr, 5*time.Second); err != nil {
		// TODO：支持后续可以针对建立连接时的Timeout进行设置
		LOG.Errorf("Create Conn Failed. Failed to Connect: %v", err)
		return nil, err
	} else {
		c.TCPConn = conn.(*net.TCPConn)
		LOG.Infof("Now Client connect to Server success, Conn ID: %s", c.ID)
	}
	if err = c.TCPConn.SetKeepAlive(true); err != nil {
		LOG.Errorf("Conn Set KeepAlive Failed. Error: %v", err)
		return nil, err
	}
	c.Ctx, c.CtxCancel = context.WithCancel(context.Background())
	return c, nil
}

// Listen 当连接建立后，调用Listen方法可以
func (c *Conn) Handle(msgChan chan []byte) {
	for {
		reader := bufio.NewReader(c.TCPConn)
		var buf [8192]byte
		n, err := reader.Read(buf[:])
		if err != nil {
			LOG.Errorf("Conn: %s read data failed. Error: %+v", c.ID, err)
			break
		}
		recvinfo := buf[:n]
		msgChan <- recvinfo
	}
}

// Handshaking 连接池判断存活
func (c *Conn) Handshaking(conn *Conn) bool {
	if c.TCPConn == nil {
		LOG.Errorf("Conn ID:%s is closed", c.ID)
	}
	buf := make([]byte, 256)
	c.TCPConn.Write([]byte("ping"))
	for {
		_size, err := c.TCPConn.Read(buf)
		if err != nil {
			LOG.Errorf("Conn {%s} Handshaking Failed. Error: %v", c.ID, err)
		}
		getData := string(buf[0:_size])
		if getData == "pong" {
			return true
		} else {
			return false
		}
	}
}
