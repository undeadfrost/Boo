package Boo

import (
	"log"
	"net/http"
	"strings"
)

type router struct {
	trees map[string]*node
}

func createRouter() *router {
	return &router{
		trees: make(map[string]*node),
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

	r.trees[method].insert(path, handler)
}

func (r *router) getRouter(method string, path string) (*node, []Param) {
	if _, ok := r.trees[method]; !ok {
		return nil, nil
	}

	n, params := r.trees[method].search(path)
	return n, params
}

func (r *router) handler(c *Context) {
	n, params := r.getRouter(c.Method, c.Path)
	if n != nil {
		c.Params = params
		n.handler(c)
	} else {
		c.String(http.StatusNotFound, "404 Not Fount: %s\n", c.Path)
	}
}
