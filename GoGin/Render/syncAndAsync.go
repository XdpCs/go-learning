package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func main() {
	// goroutine机制可以方便地实现异步处理
	// 另外，在启动新的goroutine时，不应该使用原始上下文，必须使用它的只读副本
	r := gin.Default()
	// 异步
	r.GET("/longAsync", func(c *gin.Context) {
		copyContext := c.Copy()
		go func() {
			time.Sleep(3 * time.Second)
			log.Println("异步执行：" + copyContext.Request.URL.Path)
		}()
		c.JSON(http.StatusOK, gin.H{"msg": "异步请求成功"})
	})
	// 同步
	r.GET("/longSync", func(c *gin.Context) {
		time.Sleep(3 * time.Second)
		log.Println("同步执行：" + c.Request.URL.Path)
		c.JSON(http.StatusOK, gin.H{"msg": "同步请求成功"})
	})

}
