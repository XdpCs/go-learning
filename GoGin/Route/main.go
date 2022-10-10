package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gogin/Route/routers"
)

func main() {
	r := gin.Default()
	routers.LoadBlog(r)
	routers.LoadShop(r)
	if err := r.Run(); err != nil {
		fmt.Println("startup service failed,err:%v\n", err)
	}
}
