package api

import "video_server/gee"

func RegisterHandler(engine *gee.Engine)  {
	engine.GET("/index/hello", hello)
}
