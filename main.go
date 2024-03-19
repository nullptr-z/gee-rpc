package main

import (
	"flag"
	"fmt"
	"gee-rpc/gee"
)

func main() {
	port := flag.Int("port", 9999, "service port")

	r := gee.New()
	r.GET("/", func(c *gee.Context) {
		fmt.Fprintf(c.Resp, "URL.Path = %q\n", c.Req.URL.Path)
	})

	r.GET("/hello", func(c *gee.Context) {
		for k, v := range c.Req.Header {
			fmt.Fprintf(c.Resp, "%s: %s\n", k, v)
		}
	})

	fmt.Printf("Listening on: http://localhost:%d", *port)
	r.Run(fmt.Sprintf(":%d", *port))
}
