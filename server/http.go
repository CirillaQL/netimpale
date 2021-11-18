package server

import "netimpale/utils/log"

var LOG = log.LOG

// HTTPManager Http连接的管理结构体，目前用来处理Http连接的相关请求
type HTTPManager struct {
	RunID string
}

// NewHTTPManager 创建HTTPManager
func NewHTTPManager() *HTTPManager {
	return &HTTPManager{}
}

func (http *HTTPManager) Run() {

}
