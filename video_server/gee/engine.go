package gee

import (
	"net/http"
	"video_server/loggo"
)

// HandlerFunc defines the request handler used by gee
type HandlerFunc func(c *Context)

// Engine implement the interface of ServeHTTP
type Engine struct {
	router *Router
	log *loggo.Loggo
	handlers []HandlerFunc
}

// New is the constructor of gee.Engine
func New(logPath string) *Engine {
	log := loggo.New(logPath, "", loggo.LstdFlags, loggo.Debug)
	return &Engine{router: NewRouter(), log: log, handlers: make([]HandlerFunc, 0)}
}

func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	engine.router.AddHandler(method, pattern, handler)
}

// GET defines the method to add GET request
func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

// POST defines the method to add POST request
func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

// Run defines the method to start a http server
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

// Run defines the method to start a http server
func (engine *Engine) Use(handler HandlerFunc)  {
	engine.handlers = append(engine.handlers, handler)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	c.index = -1
	c.Log = engine.log
	c.handlers = make([]HandlerFunc, 0)
	for _, h := range engine.handlers {
		c.handlers = append(c.handlers, h)
	}
	engine.router.Handle(c)
}
