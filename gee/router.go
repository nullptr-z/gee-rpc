package gee

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"testing"
)

type ContextHandler func(c *Context)

type router struct {
	roots   map[string]*node
	handler map[string]ContextHandler
}

func newRouter() *router {
	return &router{
		roots:   make(map[string]*node),
		handler: make(map[string]ContextHandler, 0),
	}
}

func parsePattern(pattern string) []string {
	// 将pattern使用`/`分隔，如 /p/:lang/doc
	vs := strings.Split(pattern, "/")

	parts := make([]string, 0)
	for _, item := range vs {
		if item != "" {
			parts = append(parts, item)
			// 如果是通配符*，直接退出
			if item[0] == '*' {
				break
			}
		}
	}

	fmt.Println("parts:", parts)
	return parts
}

func (r *router) addRoute(method string, pattern string, handler ContextHandler) {
	parts := parsePattern(pattern)

	key := method + "-" + pattern
	if _, ok := r.roots[method]; !ok {
		// 初始化 node{}
		r.roots[method] = &node{}
	}
	r.roots[method].insert(pattern, parts, 0)
	r.handler[key] = handler
}

func (r *router) getRoute(method string, path string) (*node, map[string]string) {
	searchParts := parsePattern(path)
	params := make(map[string]string)
	root, ok := r.roots[method]

	if !ok {
		return nil, nil
	}

	n := root.search(searchParts, 0)

	if n != nil {
		parts := parsePattern(n.pattern)
		for index, part := range parts {
			if part[0] != ':' {
				params[part[1:]] = searchParts[index]
			}
			if part[0] == '*' && len(part) > 1 {
				params[part[1:]] = strings.Join(searchParts[index:], "/")
				break
			}
		}
		return n, params
	}
	return nil, nil
}

func (r *router) addRoute_old(method string, pattern string, handler ContextHandler) {
	key := method + "-" + pattern
	r.handler[key] = handler
}

func (r *router) handle(c *Context) {
	n, params := r.getRoute(c.Method, c.Path)
	if n != nil {
		c.Params = params
		key := c.Method + "-" + n.pattern
		r.handler[key](c)
	} else {
		c.String(http.StatusNotFound, "404 Not Found: %s\n", c.Path)
	}
}

func (r *router) handle_old(c *Context) {
	key := c.Method + "-" + c.Path

	if handlerFunc, ok := r.handler[key]; ok {
		handlerFunc(c)
	} else {
		c.String(http.StatusNotFound, "404 Not Found: %s\n", c.Path)
	}

}

func NewTestRouter() *router {
	r := newRouter()
	r.addRoute("GET", "/", nil)
	r.addRoute("GET", "/hello/:name", nil)
	r.addRoute("GET", "/hello/b/c", nil)
	r.addRoute("GET", "/hi/:name", nil)
	r.addRoute("GET", "/assets/*filepath", nil)
	return r
}

func TestParsePattern(t *testing.T) {
	ok := reflect.DeepEqual(parsePattern("/p/:name"), []string{"p", "name"})
	ok = ok && reflect.DeepEqual(parsePattern("/p/:name"), []string{"p", "*"})
	ok = ok && reflect.DeepEqual(parsePattern("/p/*name/*"), []string{"p", "name"})
	if !ok {
		t.Fatal("test parsePattern failed")
	}
}

func TestGetRoute(t *testing.T) {
	r := NewTestRouter()
	n, ps := r.getRoute("GET", "/hello/geektutu")

	if n == nil {
		t.Fatal("nil shouldn't be returned")
	}

	if n.pattern != "/hello/:name" {
		t.Fatal("should match /hello/:name")
	}

	if ps["name"] != "geektutu" {
		t.Fatal("name should be equal to 'geektutu'")
	}

	fmt.Printf("matched path: %s, params['name']: %s\n", n.pattern, ps["name"])
}
