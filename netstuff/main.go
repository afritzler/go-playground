package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"time"
)

func main() {
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(10 * time.Second)
	}))
	defer svr.Close()
	fmt.Println("making request")
	client := http.Client{Timeout: 5 * time.Second}
	req, err := http.NewRequest(http.MethodGet, svr.URL, nil)
	if err != nil {
		panic(err)
	}
	client.Do(req)
	fmt.Println("finished request")
}
