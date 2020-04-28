package main

import (
	"fmt"
	"gee"
)

func indexHandler(c *gee.Context) {
	fmt.Fprintf(c.Writer, "URL.Path = %q\n", c.Req.URL.Path)
}

func helloHandler(c *gee.Context) {
	for k, v := range c.Req.Header {
		fmt.Fprintf(c.Writer, "Header[%q] = %q\n", k, v)
	}
}

func main() {
	r := gee.New()

	r.GET("/", indexHandler)
	r.GET("/hello", helloHandler)

	r.Run(":9999")
}
