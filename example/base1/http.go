package main

import (
	"fmt"
	"log"
	"net/http"
)

type MyHandler struct {
	routes []string
}

func (slf *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		slf.routes = append(slf.routes, r.URL.Path)
		fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	case "/hello":
		slf.routes = append(slf.routes, r.URL.Path)
		for k, v := range r.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	case "/route_show":
		for _, v := range slf.routes {
			fmt.Fprintf(w, "route= %q\n", v)
		}
	default:
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", r.URL)
	}
}

func main() {
	h := new(MyHandler)
	http.Handle("/", h)
	// http.HandleFunc("/hello", helloHandler)
	log.Fatal(http.ListenAndServe(":9999", nil))
}
func indexHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
}

// handler echoes r.URL.Header
func helloHandler(w http.ResponseWriter, req *http.Request) {
	for k, v := range req.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
}
