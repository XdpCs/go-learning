package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

func main() {
	engine := gin.Default()
	html := template.Must(template.ParseFiles("templates/file1.tmpl", "templates/file2.tmpl"))
	engine.SetHTMLTemplate(html)
	engine.GET("/file1", func(c *gin.Context) {
		c.HTML(http.StatusOK, "file1.tmpl", gin.H{
			"title": "fyy",
		})
	})
	engine.GET("/file2", func(c *gin.Context) {
		c.HTML(http.StatusOK, "file2.tmpl", gin.H{
			"title": "xdp",
		})
	})
	engine.Run(":8080")
}
