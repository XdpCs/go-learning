package main

import (
	"github.com/gin-gonic/gin"
	"gogin/ParseAndBind/Login"
	"net/http"
)

func main() {
	r := gin.Default()
	r.POST("/:user/:password", func(c *gin.Context) {
		var login Login.Login
		if err := c.ShouldBindUri(&login); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if login.User != "root" || login.Password != "fyy" {
			c.JSON(http.StatusBadRequest, gin.H{"status": "400"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "200"})
	})
	r.Run()
}
