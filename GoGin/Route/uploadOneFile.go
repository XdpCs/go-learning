package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// multipart/form-data格式用于文件上传
	// gin文件上传与原生的net/http方法类似，不同在于gin把原生的request封装到c.Request中
	r := gin.Default()
	// 限制处理上传文件的最大内存
	r.MaxMultipartMemory = 8 << 20
	r.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.String(http.StatusInternalServerError, "上传图片出错")
		}
		// 上传到当前项目的相对路径下
		c.SaveUploadedFile(file, file.Filename)
		c.String(http.StatusOK, file.Filename)
	})
	r.Run()
}
