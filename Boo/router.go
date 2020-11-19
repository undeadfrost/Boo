package Boo

import (
	"log"
	"net/http"
	"strings"
)

type router struct {
	trees    map[string]*node
	handlers map[string]HandlerFunc
}

func createRouter() *router {
	return &router{
		trees:    make(map[string]*node),
		handlers: make(map[string]HandlerFunc),
	}
}

func parsePattern(pattern string) []string {
	vs := strings.Split(pattern, "/")

	parts := make([]string, 0)
	for _, item := range vs {
		if item != "" {
			parts = append(parts, item)
			if item[0] == '*' {
				break
			}
		}
	}
	return parts
}

func (r *router) addRouter(method string, path string, handler HandlerFunc) {
	log.Printf("Router %s-%s\n", method, path)

	// 创建树
	if _, ok := r.trees[method]; !ok {
		r.trees[method] = new(node)
	}

	paths := parsePattern(path)

	r.trees[method].insert(path, paths, handler)
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
