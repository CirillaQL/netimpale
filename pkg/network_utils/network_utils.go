package network_utils

import (
	"net"
	"netimpale/utils/log"
	"strconv"
	"strings"
)

var LOG = log.LOG

// GetIPAddress 获取IP地址
func GetIPAddress() (map[string]string, error) {
	ips := make(map[string]string)
	interfaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	for _, i := range interfaces {
		address, err := i.Addrs()
		if err != nil {
			return nil, err
		}
		ipv4_address := address[1].String()
		ips[i.Name] = ipv4_address
	}
	return ips, nil
}

func ParseIPv4(address string) (string, int) {
	IPandPort := strings.Split(address, ":")
	port, err := strconv.Atoi(IPandPort[1])
	if err != nil {
		LOG.Errorf("Parse IPv4 Address Failed: %v", err)
	}
	return IPandPort[0], port
}
