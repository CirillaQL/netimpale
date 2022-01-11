package config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"netimpale/utils/log"
	"os"
)

var LOG = log.LOG

type TCP struct {
	Host string `yaml:"host"`
	Port uint32 `yaml:"port"`
}

type UDP struct {
	Host string `yaml:"host"`
	Port uint32 `yaml:"port"`
}

// Client Config 客户端
type Client struct {
	TCP
	UDP
}

// Server Config 服务端
type Server struct {
	TCP
	UDP
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
	// 首先检查配置文件是否存在
	exist := checkConfigFileIsExist(filepath)
	if exist != true {
		LOG.Errorf("Config file doesn't exit: %v", filepath)
	}
	// 读文件
	yamlFile, err := ioutil.ReadFile(filepath)
	if err != nil || yamlFile == nil {
		LOG.Errorf("Can't Open yaml file. Error: %v", err)
	}
	// 加载对应的配置
	clientConfig := new(Client)
	err = yaml.Unmarshal(yamlFile, clientConfig)
	if err != nil {
		LOG.Errorf("Can't load ")
	}
	return clientConfig
}

// LoadServerConfig 加载服务端配置文件
func LoadServerConfig(filepath string) *Server {
	// 首先检查配置文件是否存在
	exist := checkConfigFileIsExist(filepath)
	if exist != true {
		LOG.Errorf("Config file doesn't exit: %v", filepath)
	}
	// 读文件
	yamlFile, err := ioutil.ReadFile(filepath)
	if err != nil || yamlFile == nil {
		LOG.Errorf("Can't Open yaml file. Error: %v", err)
	}
	// 加载对应的配置
	serverConfig := new(Server)
	err = yaml.Unmarshal(yamlFile, serverConfig)
	if err != nil {
		LOG.Errorf("Can't load ")
	}
	return serverConfig
}

// OutputClientConfig 格式化输出客户端证书的内容
func (client *Client) OutputClientConfig() {
	if client.TCP != (TCP{}) {
		LOG.Infof("TCP Host: %v", client.TCP.Host)
		LOG.Infof("TCP Port: %v", client.TCP.Port)
	}
	if client.UDP != (UDP{}) {
		LOG.Infof("UDP Host: %v", client.UDP.Host)
		LOG.Infof("UDP Port: %v", client.UDP.Port)
	}
}

// OutputServerConfig 格式化输出服务端证书的内容
func (server *Server) OutputServerConfig() {
	if server.TCP != (TCP{}) {
		LOG.Infof("TCP Host: %v", server.TCP.Host)
		LOG.Infof("TCP Port: %v", server.TCP.Port)
	}
	if server.UDP != (UDP{}) {
		LOG.Infof("UDP Host: %v", server.UDP.Host)
		LOG.Infof("UDP Port: %v", server.UDP.Port)
	}
}
