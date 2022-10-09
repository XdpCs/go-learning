package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// URL参数可以通过DefaultQuery()或Query()方法获取
	// DefaultQuery()若参数不传，返回默认值，Query()若不存在，返回空串
	r := gin.Default()
	r.GET("/user", func(c *gin.Context) {
		// 指定默认值
		// http://localhost:8080/user 才会打印出来默认的值
		name := c.DefaultQuery("name", "fyy")
		c.String(http.StatusOK, fmt.Sprintf("hello %s", name))
	})
	r.Run()
}
