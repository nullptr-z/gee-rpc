package main

import (
	"flag"
	"fmt"
	"gee-rpc/gee"
	"net/http"
)

func main() {
	port := flag.Int("port", 9999, "service port")

	r := gee.New()
	r.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	})

	r.GET("/hello", func(w http.ResponseWriter, r *http.Request) {
		for k, v := range r.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	})

	fmt.Printf(fmt.Sprintf("Listening on: http://localhost:%d", *port))
	r.Run(fmt.Sprintf(":%d", *port))
}
