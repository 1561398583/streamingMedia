package main

import (
	"fmt"
	_ "gorm.io/gorm"
	"runtime"
	"video_server/gee"
)

func main()  {
	engine := gee.New()
	//添加panic处理
	engine.Use(catchPanic)
	engine.GET("/index/hello", hello)
	engine.GET("/user/*", userController)
	routerNode, err := engine.router.GetNode("GET", "/user")
	if err != nil {
		panic(err)
	}
	routerNode.AddMidHandler(userP)

	engine.Use(recordTime)
	engine.GET("/static/pictrue/*", getPictrue)
	engine.Run("localhost:7000")
}

func catchPanic(c *gee.Context)  {
	defer func() {
		r := recover()
		if r != nil {
			errInfo := fmt.Sprintln(r) + "\n"
			for i := 0; ; i++ {
				_, file, line, ok := runtime.Caller(i)
				if !ok {
					break
				}
				errInfo += fmt.Sprintln(file, line)
			}
			//记录日志
			c.Log.Error(errInfo)
		}
	}()
	c.Next()
}

