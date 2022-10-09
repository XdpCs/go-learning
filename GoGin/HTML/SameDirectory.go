package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 同一目录的模版
func main() {
	engine := gin.Default()
	engine.LoadHTMLGlob("templates/*")
	// engine.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
	engine.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "fyy xdp",
		})
	})
	engine.Run(":8080")
}
