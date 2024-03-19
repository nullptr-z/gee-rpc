package gee

import (
	"net/http"
)

type Engine struct {
	router *router
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := mewContext(w, r)
	e.router.handle(c)
}

func New() *Engine {
	return &Engine{router: newRouter()}
}

func (e *Engine) GET(pattern string, handler ContextHandler) {
	e.addRoute("GET", pattern, handler)
}

func (e *Engine) POST(pattern string, handler ContextHandler) {
	e.addRoute("POST", pattern, handler)
}

func (e *Engine) addRoute(method string, pattern string, handler ContextHandler) {
	e.router.addRoute(method, pattern, handler)
}

func (g *Engine) Run(prot string) error {
	return http.ListenAndServe(prot, g)
}
