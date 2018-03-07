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
			fmt.Printf("%s\n", comma(path[i]))
		}
	}
}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}
