package main

import (
	"github.com/gin-gonic/gin"
	"gogin/ParseAndBind/Login"
	"net/http"
)

func main() {
	r := gin.Default()
	r.POST("/loginForm", func(c *gin.Context) {
		var form Login.Login
		// Bind()默认解析并绑定form格式
		// 根据请求头中content-type自动推断
		if err := c.Bind(&form); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if form.User != "root" || form.Password != "fyy" {
			c.JSON(http.StatusBadRequest, gin.H{"status": "400"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "200"})
	})
	r.Run()
}
