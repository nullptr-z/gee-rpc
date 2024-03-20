package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type H map[string]any

type Context struct {
	Req  *http.Request
	Resp http.ResponseWriter

	Path   string
	Method string
	Params map[string]string // get param
	// response status
	StatusCode int
}

func mewContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		Req:    r,
		Resp:   w,
		Path:   r.URL.Path,
		Method: r.Method,
	}
}

func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Resp.WriteHeader(code)
}

func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}

func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

func (c *Context) Param(key string) string {
	value, _ := c.Params[key]
	return value
}

// func (c *Context) String(code int, format string, values ...any)
func (c *Context) String(code int, format string, values ...any) {
	c.Status(code)
	c.SetHeader("Content-Type", "text/plain")
	c.Resp.Write([]byte(fmt.Sprintf(format, values...)))
}

func (c *Context) SetHeader(key string, value string) {
	c.Resp.Header().Set(key, value)
}

func (c *Context) JSON(code int, obj any) {
	c.SetHeader("Content-Type", "application/json")
	encoder := json.NewEncoder(c.Resp)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Resp, err.Error(), code)
	}
}

func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.Resp.Write(data)
}

func (c *Context) Html(code int, html string) {
	c.Status(code)
	c.SetHeader("Content-Type", "text/plain")
	c.Resp.Write([]byte(html))
}
