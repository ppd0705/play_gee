package main

import (
	"fmt"
	"gee"
	"net/http"
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

	r.Use(gee.Logger())
	r.GET("/", indexHandler)
	r.GET("/hello", helloHandler)

	r.GET("/hello/:name", func(c *gee.Context) {
		c.String(http.StatusOK, "hello %s, you are at %s\n", c.Params["name"], c.Path)
	})

	r.GET("/assets/*filepath", func(c *gee.Context) {
		c.Json(http.StatusOK, gee.H{"filepath": c.Params["filepath"]})
	})

	r.Run(":9999")
}
