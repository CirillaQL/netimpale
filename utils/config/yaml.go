package config

import (
	"netimpale/utils/log"
	"os"
)

var LOG = log.LOG

// Client Config 客户端
type Client struct {
	TCP struct {
		Host string `yaml:"host"`
		Port uint32 `yaml:"port"`
	}
	UDP struct {
		Host string `yaml:"host"`
		Port uint32 `yaml:"port"`
	}
}

// Server Config 服务端
type Server struct {
}

// 判断对应文件是否存在
func checkConfigFileIsExist(filepath string) bool {
	var exist = true
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

// LoadClientConfig 加载客户端配置文件
func LoadClientConfig(filepath string) *Client {
	return nil
}

// LoadServerConfig 加载服务端配置文件
func LoadServerConfig(filepath string) *Server {
	return nil
}
