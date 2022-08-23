package main

import (
	"gee"
	"net/http"
)

func main() {
	engine := gee.New()
	engine.GET("/", func(context *gee.Context) {
		context.HTML(http.StatusOK, "<h1>Hello fyy</h1>")
	})
	engine.GET("/hello", func(context *gee.Context) {
		context.String(http.StatusOK, "hello %s,you're at %s\n", context.Query("name"), context.Path)
	})
	engine.GET("/hello/:name", func(context *gee.Context) {
		context.String(http.StatusOK, "hello %s,you're at %s\n", context.Param("name"), context.Path)
	})
	engine.Run(":9999")
}
