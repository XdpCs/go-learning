package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	r := gin.Default()
	r.POST("/upload", func(c *gin.Context) {
		_, headers, err := c.Request.FormFile("file")
		if err != nil {
			log.Printf("Error when try to get file: %v", err)
		}
		// headers.Size 获取文件大小
		if headers.Size > 1024*1024*2 {
			c.String(http.StatusOK, "文件太大了")
			return
		}
		// headers.Header.Get("Content-Type")获取上传文件的类型
		if headers.Header.Get("Content-Type") != "image/png" {
			c.String(http.StatusOK, "只允许上传png图片")
			return
		}
		c.SaveUploadedFile(headers, "./fyy/"+headers.Filename)
		c.String(http.StatusOK, headers.Filename)
	})
	r.Run()
}
