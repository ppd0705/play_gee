package main

import (
	"fmt"
	"gee"
	"net/http"
)

func indexHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
}

func helloHandler(w http.ResponseWriter, req *http.Request) {
	for k, v := range req.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
}

func main() {
	r := gee.New()

	r.GET("/", indexHandler)
	r.GET("/hello", helloHandler)

	r.Run(":9999")
}
