package main

import (
	"github.com/gin-gonic/gin"
	"gogin/ParseAndBind/Login"
	"net/http"
)

func main() {
	r := gin.Default()
	// Json绑定
	r.POST("/loginJson", func(c *gin.Context) {
		// 声明接收的变量
		var json Login.Login
		if err := c.ShouldBindJSON(&json); err != nil {
			// 返回错误信息
			// gin.H 封装生成JSON数据的工具
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if json.User != "root" || json.Password != "fyy" {
			c.JSON(http.StatusBadRequest, gin.H{"status": "400"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": 200})
	})
	r.Run()

}
