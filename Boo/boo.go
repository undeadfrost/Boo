package Boo

import (
	"net/http"
)

type HandlerFunc func(c *Context)

type Engine struct {
	router *router
}

func New() *Engine {
	return &Engine{router: createRouter()}
}

func (engine *Engine) addRouter(method string, path string, handler HandlerFunc) {
	engine.router.addRouter(method, path, handler)
}

func (engine *Engine) GET(path string, handler HandlerFunc) {
	engine.addRouter("GET", path, handler)
}

func (engine *Engine) POST(path string, handler HandlerFunc) {
	engine.addRouter("POST", path, handler)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	context := createContext(w, r)
	engine.router.handler(context)
}

func (engine *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, engine)
}
