package gee

import "net/http"

//只有刚启动服务器的时候会写入，后面就算是高并发也都是读，所以并发安全
type router struct {
	handlers map[string]*UrlHandler
}

type UrlHandler struct {
	f HandlerFunc
	midF []HandlerFunc
}

func (uh *UrlHandler) Use(handler HandlerFunc)  {
	uh.midF = append(uh.midF, handler)
}

func newRouter() *router {
	return &router{handlers: make(map[string]*UrlHandler)}
}

func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	uh := UrlHandler{f: handler, midF: make([]HandlerFunc, 0)}
	r.handlers[key] = &uh
}

func (r *router) GetRoute(method string, pattern string) *UrlHandler{
	key := method + "-" + pattern
	return r.handlers[key]
}

func (r *router) handle(c *Context) {
	key := c.Method + "-" + c.Path
	if handler, ok := r.handlers[key]; !ok {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	} else {
		err := c.Req.ParseForm()
		if err != nil {
			c.Log.Error(err)
		}
		//加入额外处理
		for _, h := range handler.midF {
			c.handlers = append(c.handlers, h)
		}
		//加入业务处理
		c.handlers = append(c.handlers, handler.f)
		c.Next()
	}
}
