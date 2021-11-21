package server

import (
	"net"
	"netimpale/utils/log"
	"time"
)

var LOG = log.LOG

// TCPManager Http连接的管理结构体，目前用来处理Http连接的相关请求
type TCPManager struct {
	RunID    string
	Conns    []*TCPConn
	Listener net.Listener
}

// TCPConn 针对HTTP的Conn,添加相关的支持
type TCPConn struct {
	ConnID string
	Token  string
	Conn   *net.Conn
	Data   []byte
	Status uint32
	Time   time.Time
}

// NewTCPManager 创建HTTPManager
func NewTCPManager() *HTTPManager {

	return &HTTPManager{}
}

// ConnHandler 针对连接进行处理的Handler
func (tcp *TCPManager) ConnHandler(conn net.Conn) {

}

// Run 运行
func (tcp *TCPManager) Run() {

}
