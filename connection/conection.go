package connection

import "time"

// Connection 连接结构体的定义
type Connection struct {
	ConnectID   string    `json:"connect_id"`
	ConnectType string    `json:"connect_type"`
	Time        time.Time `json:"time"`
	IsEncrypted bool      `json:"is_encrypted"`
	Size        uint32    `json:"size"`
	Data        []byte    `json:"data"`
	ClientAddr  string    `json:"client_addr"`
	ServerAddr  string    `json:"server_addr"`
	ClientPort  uint32    `json:"client_port"`
	ServerPort  uint32    `json:"server_port"`
}
