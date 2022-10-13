package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// MiddleWare
// 所有请求都经过此中间件
// 定义中间件
func MiddleWareGlobal() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		fmt.Println("中间件开始执行了")
		// 设置变量到Context的key中，可以通过Get()获取
		c.Set("request", "中间件")
		status := c.Writer.Status()
		fmt.Println("中间件执行完成", status)
		t2 := time.Since(t)
		fmt.Println("time:", t2)
	}
}

func main() {
	r := gin.Default()
	// 注册中间件
	r.Use(MiddleWareGlobal())
	// {}为了代码规范
	{
		r.GET("/middle", func(c *gin.Context) {
			req, _ := c.Get("request")
			fmt.Println("request:", req)
			c.JSON(http.StatusOK, gin.H{"request": req})
		})
	}
	r.Run()
}
