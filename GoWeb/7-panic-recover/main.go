package main

import (
	"gee"
	"net/http"
)

func main() {
	engine := gee.Default()
	engine.GET("/", func(context *gee.Context) {
		context.String(http.StatusOK, "hello fyy\n")
	})
	engine.GET("/panic", func(context *gee.Context) {
		names := []string{"fuyiyi"}
		context.String(http.StatusOK, names[100])
	})
	engine.Run(":9999")
}
