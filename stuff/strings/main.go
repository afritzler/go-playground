package main

import (
	"fmt"
	"strings"
)

func main() {
	hostname := "dummy-machine.openstack.local"
	fmt.Println(strings.Split(hostname, ".")[0])
}
