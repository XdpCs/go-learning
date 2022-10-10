package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/toMyGithub", myGithubHandler)
	return r
}

func myGithubHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "https://github.com/XdpCS",
	})
}
