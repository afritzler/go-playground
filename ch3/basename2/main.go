package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	path := os.Args[1:]
	if len(path) == 0 {
		fmt.Println("nothing to do here")
	} else {
		for i := range path {
			fmt.Printf("%s\n", basename(path[i]))
		}
	}
}

func basename(s string) string {
	slash := strings.LastIndex(s, "/")
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}
