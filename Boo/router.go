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

func (r *router) getRouter(method string, path string) (*node, map[string]string) {
	paths := parsePattern(path)

	if _, ok := r.trees[method]; !ok {
		return nil, nil
	}

	n, params := r.trees[method].search(path, paths)
	return n, params
}

func (r *router) handler(c *Context) {
	//key := c.Method + "-" + c.Path
	n, params := r.getRouter(c.Method, c.Path)
	if n != nil {
		c.Params = params
		n.handler(c)
	} else {
		c.String(http.StatusNotFound, "404 Not Fount: %s\n", c.Path)
	}
	//if handler, ok := r.handlers[key]; ok {
	//	handler(c)
	//} else {
	//	c.String(http.StatusNotFound, "404 Not Fount: %s\n", c.Path)
	//}
}
