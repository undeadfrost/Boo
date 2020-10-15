package Boo

import (
	"log"
	"net/http"
)

type router struct {
	handlers map[string]HandlerFunc
}

func createRouter() *router {
	return &router{make(map[string]HandlerFunc)}
}

func (r *router) addRouter(method string, path string, handler HandlerFunc) {
	log.Printf("Router %s-%s\n", method, path)
	key := method + "-" + path
	r.handlers[key] = handler
}

func (r *router) handler(c *Context) {
	key := c.Method + "-" + c.Path
	if handler, ok := r.handlers[key]; ok {
		handler(c)
	} else {
		c.String(http.StatusNotFound, "404 Not Fount: %s\n", c.Path)
	}
}
