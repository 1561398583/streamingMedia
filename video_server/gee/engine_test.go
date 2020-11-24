package gee

import (
	"fmt"
	"runtime"
	"strings"
	"testing"
	"time"
)

func TestEngine_Run(t *testing.T) {
	engine := New("E:\\log\\video_server\\")
	engine.GET("/index/hello", hello)
	engine.GET("/user/*", userController)
	routerNode, err := engine.router.GetNode("GET", "/user")
	if err != nil {
		panic(err)
	}
	routerNode.AddMidHandler(userP)
	engine.Use(catchPanic)
	engine.Use(recordTime)
	engine.Run("localhost:7000")
}

func hello(c *Context)  {
	c.String(200, "hello world")
}

func userController(c *Context)  {
	path := c.Path
	parts := strings.Split(path, "/")
	m := make(map[string]string)
	m["name"] = parts[len(parts)-1]
	m["work"] = "golang"
	c.JSON(200, m)
}

func userP(c *Context)  {
	fmt.Println("begin user")
	c.Next()
	fmt.Println("finish user")
}

func recordTime(c *Context)  {
	start := time.Now()
	c.Next()
	spend := time.Since(start)
	fmt.Println("spend ", spend)
}


func catchPanic(c *Context)  {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Println("get a panic")
			for i := 0; ; i++ {
				pc, file, line, ok := runtime.Caller(i)
				if !ok {
					break
				}
				fmt.Println(pc, file, line)
			}
		}
	}()
	c.Next()
}
