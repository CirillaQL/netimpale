package main

import (
	"fmt"
	"io/ioutil"
	"netimpale/utils/log"
	myyaml "netimpale/utils/yaml"

	"gopkg.in/yaml.v3"
)

var LOG = log.LOG

func main() {
	// 创建 listener
	//listener, err := net.Listen("tcp", "localhost:9090")
	//if err != nil {
	//	fmt.Println("Error listening", err.Error())
	//	return //终止程序
	//}
	//// 监听并接受来自客户端的连接
	//for {
	//	conn, err := listener.Accept()
	//	if err != nil {
	//		fmt.Println("Error accepting", err.Error())
	//		return // 终止程序
	//	}
	//	go doServerStuff(conn)
	//}
	conf := new(myyaml.Client)
	yamlFile, err := ioutil.ReadFile("./assets/client.yaml")
	if err != nil {
		LOG.Errorf("%v", err)
	}
	err = yaml.Unmarshal(yamlFile, conf)
	if err != nil {
		LOG.Errorf("%v", err)
	}
	fmt.Println(conf.TCP.Host)
	fmt.Printf("%v", conf)
}

//func doServerStuff(conn net.Conn) {
//	for {
//		buf := make([]byte, 512)
//		len, err := conn.Read(buf)
//		if err != nil {
//			fmt.Println("Error reading", err.Error())
//			return //终止程序
//		}
//		//fmt.Printf("%v", string(buf[:len]))
//		network_utils.GenerateHTTPRequest(buf, len)
//	}
//}
