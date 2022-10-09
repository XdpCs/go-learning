package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// 创建路由
	r := gin.Default()
	// 绑定路由规则和执行的函数
	// gin.Context,封装了request和response
	r.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "hello,fyy!")
	})
	// 鉴定端口，默认在8080
	// Run("")里面不指定端口号默认是8080
	r.Run(":8080")
}
