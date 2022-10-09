package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func main() {
	r := gin.Default()
	// 可以通过Context的Param方法来获取API参数
	r.GET("/user/:name/*action", func(context *gin.Context) {
		name := context.Param("name")
		action := context.Param("action")
		// 截取
		action = strings.Trim(action, "/")
		context.String(http.StatusOK, name+" is "+action)
	})
	r.Run(":8080")
}
