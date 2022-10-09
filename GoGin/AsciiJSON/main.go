package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 使用AsciiJSON生成具有转义的非Ascii字符的ASCII-only JSON
func main() {
	engine := gin.Default()
	engine.GET("/someJson", func(c *gin.Context) {
		data := map[string]interface{}{
			"lang": "go语言",
			"tag":  "<br>",
		}
		c.AsciiJSON(http.StatusOK, data)
	})
	engine.Run(":8080")
}
