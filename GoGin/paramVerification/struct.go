package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type PersonStruct struct {
	// 不能为空并且大于10
	Age      int       `form:"age" binding:"required,gt=10"`
	Name     string    `form:"name" binding:"required"`
	Birthday time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
}

func main() {
	// 使用gin框架进行数据验证，可以不用解析数据
	r := gin.Default()
	r.GET("/person", func(c *gin.Context) {
		var person PersonStruct
		if err := c.ShouldBind(&person); err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		c.String(http.StatusOK, fmt.Sprintf("%#v", person))
	})
	r.Run()
}
