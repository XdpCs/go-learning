package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"net/http"
)

/*
   对绑定解析到结构体上的参数，自定义验证功能
   比如我们要对 name 字段做校验，要不能为空，并且不等于 admin ，类似这种需求，就无法 binding 现成的方法
   需要我们自己验证方法才能实现 官网示例（https://godoc.org/gopkg.in/go-playground/validator.v8#hdr-Custom_Functions）
   这里需要下载引入下 gopkg.in/go-playground/validator.v8
*/
type PersonDiy struct {
	Age int `form:"age" binding:"required,gt=10"`
	// validate 上使用自定义的校验方法函数注册时候的名称
	Name    string `form:"name" binding:"NotNullAndAdmin"`
	Address string `form:"address" binding:"required"`
}

// 自定义校验方法
func NotNullAndAdmin(fl validator.FieldLevel) bool {
	if value, ok := fl.Field().Interface().(string); ok {
		return value != "" && !("admin" == value)
	}
	return true
}

func main() {
	r := gin.Default()
	// 将自定义的校验方法注册到validator中
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 这里的 key 和 func 可以不一样最终在 struct 使用的是 key
		v.RegisterValidation("NotNullAndAdmin", NotNullAndAdmin)
	}
	r.GET("/person", func(c *gin.Context) {
		var person PersonDiy
		if e := c.ShouldBind(&person); e == nil {
			c.String(http.StatusOK, "%v", person)
		} else {
			c.String(http.StatusOK, "person bind err:%v", e.Error())
		}
	})
	r.Run()
}
