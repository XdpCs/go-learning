package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/testdata/protoexample"
	"net/http"
)

func main() {
	r := gin.Default()
	// json
	r.GET("/Json", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Json", "status": 200})
	})
	// struct
	r.GET("/Struct", func(c *gin.Context) {
		var msg struct {
			Name    string
			Message string
			Number  int
		}
		msg.Name = "xdp"
		msg.Message = "fyy"
		msg.Number = 1118
		c.JSON(http.StatusOK, msg)
	})
	// XML
	r.GET("/XML", func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{"message": "fyy"})
	})
	// YAML
	r.GET("/YAML", func(c *gin.Context) {
		c.YAML(http.StatusOK, gin.H{"name": "fyy"})
	})
	// protobuf
	r.GET("/protoBuf", func(c *gin.Context) {
		reps := []int64{int64(1), int64(2)}
		label := "label"
		// 传protobuf格式数据
		data := &protoexample.Test{
			Label: &label,
			Reps:  reps,
		}
		c.ProtoBuf(http.StatusOK, data)
	})
	r.Run()
}
