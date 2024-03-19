package gee

import (
	"log"
	"net/http"
)

type ContextHandler func(c *Context)

type router struct {
	handler map[string]ContextHandler
}

func newRouter() *router {
	return &router{handler: make(map[string]ContextHandler, 0)}
}

func (engine *router) addRoute(method string, pattern string, handler ContextHandler) {
	log.Printf("Route %4s - %s", method, pattern)
	key := method + "-" + pattern
	engine.handler[key] = handler
}

func (r *router) handle(c *Context) {
	key := c.Method + "-" + c.Path

	if handlerFunc, ok := r.handler[key]; ok {
		handlerFunc(c)
	} else {
		c.String(http.StatusNotFound, "404 Not Found: %s\n", c.Path)
	}

}
