package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 模拟实现权限验证中间件
//有2个路由，login和home
//login用于设置cookie
//home是访问查看信息的请求
//在请求home之前，先跑中间件代码，检验是否存在cookie
func main() {
	r := gin.Default()
	r.GET("/login", func(c *gin.Context) {
		c.SetCookie("login", "fyy", 60, "/", "localhost", false, true)
		c.String(http.StatusOK, "Login success!")
	})
	r.GET("/home", AuthMiddleWare, func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "home"})
	})
	r.Run()
}

func AuthMiddleWare(c *gin.Context) {
	// 获取客户端cookie并校验
	if cookie, err := c.Cookie("login"); err == nil {
		if cookie == "fyy" {
			c.Next()
			return
		}
	}
	// 返回错误
	c.JSON(http.StatusUnauthorized, gin.H{"error": "StatusUnauthorized"})
	// 验证不通过，不再调用后续函数处理
	c.Abort()
}
