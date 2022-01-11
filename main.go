package main

import (
	"netimpale/utils/config"
)

func main() {
	clientConfig := config.LoadClientConfig("./assets/client.yaml")
	clientConfig.OutputClientConfig()
	serverConfig := config.LoadServerConfig("./assets/server.yaml")
	serverConfig.OutputServerConfig()
}
