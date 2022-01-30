package main

import (
	"fmt"
	"netimpale/pkg/connection"
)

func main() {
	k, _ := connection.NewConn("127.0.0.1:9090")
	fmt.Printf("%v", k)
}
