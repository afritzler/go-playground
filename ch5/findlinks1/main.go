package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visits(nil, doc) {
		fmt.Println(link)
	}
}

func visits(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			links = append(links, a.Val)
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visits(links, c)
	}
	return links
}
