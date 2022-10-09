package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 使用不同目录下名称相同的模版
func main() {
	engine := gin.Default()
	engine.LoadHTMLGlob("templates/**/*")
	engine.GET("/posts/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "posts/index.tmpl", gin.H{
			"title": "Posts",
			"name":  "fyy",
		})
	})
	engine.GET("/users/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "users/index.tmpl", gin.H{
			"title": "Users",
			"name":  "fyy",
		})
	})
	engine.Run(":8080")

}
