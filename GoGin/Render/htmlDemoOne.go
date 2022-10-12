package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	// 注意必须按照module的路径
	r.LoadHTMLGlob("./Render/tem/**/*")
	r.Static("/static", "./static")
	r.GET("/start", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{"title": "xdp", "birth": 1118})
	})
	r.Run()
}
