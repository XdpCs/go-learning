package main

import (
	"gee"
	"log"
	"net/http"
	"time"
)

func onlyForV2() gee.HandlerFunc {
	return func(context *gee.Context) {
		// Start timer
		t := time.Now()
		// if a server error occurred
		context.Fail(500, "Internal Server Error")
		// Calculate resolution time
		log.Printf("[%d] %s in %v for group v2", context.StatusCode, context.Req.RequestURI, time.Since(t))
	}
}
func main() {
	engine := gee.New()
	engine.Use(gee.Logger()) // global middleware
	engine.GET("/", func(context *gee.Context) {
		context.HTML(http.StatusOK, "<h1>hello,fyy</h1>")
	})
	v2 := engine.Group("/v2")
	v2.Use(onlyForV2())
	{
		v2.GET("/hello/:name", func(context *gee.Context) {
			context.String(http.StatusOK, "hello %s,you're at %s\n", context.Param("name"), context.Path)
		})
	}
	engine.Run(":9999")
}
