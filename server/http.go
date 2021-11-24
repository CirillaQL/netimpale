package server

import (
	"net"
	"time"
)

// HTTPManager Http连接的管理结构体，目前用来处理Http连接的相关请求
type HTTPManager struct {
	RunID    string
	Conns    []*HTTPConn
	Listener net.Listener
}

// HTTPConn 针对HTTP的Conn,添加相关的支持
type HTTPConn struct {
	ConnID string
	Token  string
	Conn   *net.Conn
	Data   []byte
	Status uint32
	Time   time.Time
}

// NewHTTPManager 创建HTTPManager
func NewHTTPManager() *HTTPManager {
	return &HTTPManager{}
}

// ConnHandler 针对连接进行处理的Handler
func (http *HTTPManager) ConnHandler(conn net.Conn) {

}

// Run 运行
func (http *HTTPManager) Run() {
	for {
		conn, err := http.Listener.Accept()
		if err != nil {
			LOG.Errorf("HTTPManager Accept Failed. HTTPManger RunID: %s, Error: %v", http.RunID, err)
		}
		go http.ConnHandler(conn)

	}
}
