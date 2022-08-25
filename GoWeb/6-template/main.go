package main

import (
	"fmt"
	"gee"
	"html/template"
	"net/http"
	"time"
)

type student struct {
	Name string
	Age  int8
}

func FormatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}
func main() {
	//engine := gee.New() demo1
	//engine.Static("/assets", "/Users/xudapeng/static")
	//engine.Run(":9999")
	engine := gee.New()
	engine.Use(gee.Logger())
	engine.SetFuncMap(template.FuncMap{
		"FormatAsDate": FormatAsDate,
	})
	engine.LoadHTMLGlob("templates/*")
	engine.Static("/assets", "./static")
	stu1 := &student{Name: "xudapeng", Age: 22}
	stu2 := &student{Name: "fuyiyi", Age: 22}
	engine.GET("/", func(context *gee.Context) {
		context.HTML(http.StatusOK, "css.tmpl", nil)
	})
	engine.GET("/students", func(context *gee.Context) {
		context.HTML(http.StatusOK, "arr.tmpl", gee.H{
			"title":  "fyy",
			"stuArr": [2]*student{stu1, stu2},
		})
	})
	engine.GET("/date", func(context *gee.Context) {
		context.HTML(http.StatusOK, "custom_func.tmpl", gee.H{
			"title": "xdp",
			"now":   time.Date(2022, 11, 18, 0, 0, 0, 0, time.UTC),
		})
	})
	engine.Run(":9999")
}
