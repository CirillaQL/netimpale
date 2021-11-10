package network

import (
	"net"
)

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
