package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 适用于条目比较少的简单项目
func main() {
	r := gin.Default()
	r.GET("/toMyGithub", myGithubHandler)
	if err := r.Run(); err != nil {
		fmt.Printf("startup service failed,err:%v\n", err)
	}
}

func myGithubHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "https://github.com/XdpCS",
	})
}
