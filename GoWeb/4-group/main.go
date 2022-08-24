package main

import (
	"gee"
	"net/http"
)

func main() {
	engine := gee.New()
	engine.GET("/index", func(context *gee.Context) {
		context.HTML(http.StatusOK, "<h1>Index Page</h1>")
	})
	v1 := engine.Group("/v1")
	{
		v1.GET("/", func(context *gee.Context) {
			context.HTML(http.StatusOK, "<h1>Hello Fyy</h1>")
		})
		v1.GET("/hello", func(context *gee.Context) {
			context.String(http.StatusOK, "hello %s,you're at %s\n", context.Query("name"), context.Path)
		})
	}
	v2 := engine.Group("/v2")
	{
		v2.GET("/hello/:name", func(context *gee.Context) {
			context.String(http.StatusOK, "hello %s,you're at %s\n", context.Param("name"), context.Path)
		})
		v2.POST("/login", func(context *gee.Context) {
			context.JSON(http.StatusOK, gee.H{
				"username": context.PostForm("username"),
				"password": context.PostForm("password"),
			})
		})
	}
	engine.Run(":9999")
}
