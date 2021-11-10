package main

import (
	"fmt"
	"netimpale/utils/log"
)

var LOG = log.LOG

func main() {
	fmt.Println("Hello, NetImpale!")
	LOG.Debugf("i am debug, using %s", "sugar")
	LOG.Infof("i am info, using %s", "sugar")
	LOG.Warnf("i am warn, using %s", "sugar")
	LOG.Errorf("i am error, using %s", "sugar")
}
