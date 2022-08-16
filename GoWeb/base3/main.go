package main

import (
	"fmt"
	"gee"
	"net/http"
)

func main() {
	engine := gee.New() // 遇到导包有问题，找到GO sdk src目录下，存在空包问题
	engine.GET("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)

	})
	engine.GET("/hello", func(writer http.ResponseWriter, request *http.Request) {
		for k, v := range request.Header {
			fmt.Fprintf(writer, "Header[%q] = %q\n", k, v)
		}
	})
	engine.Run(":9999")
}
