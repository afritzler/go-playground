package main

import (
	"fmt"
	"os"
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
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}
