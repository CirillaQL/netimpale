package network_utils

import (
	"fmt"
	"testing"
)

func TestGetIPAddress(t *testing.T) {
	ips, err := GetIPAddress()
	if err != nil {
		t.Errorf("%v", err)
	}
	for k, v := range ips {
		fmt.Printf("%s    %s", k, v)
		fmt.Println()
	}
}
