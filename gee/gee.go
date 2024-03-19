package gee

import (
	"fmt"
	. "net/http"
)

type Engine struct {
	router map[string]HandlerFunc
}

func New() *Engine {

	return &Engine{router: make(map[string]HandlerFunc, 0)}
}

func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	engine.router[key] = handler
}

func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

func (slf *Engine) ServeHTTP(w ResponseWriter, r *Request) {
	key := r.Method + "-" + r.URL.Path
	if handlerFunc, ok := slf.router[key]; ok {
		handlerFunc(w, r)
	} else {
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", r.URL)
	}
}

func (g *Engine) Run(prot string) error {
	return ListenAndServe(prot, g)
}
