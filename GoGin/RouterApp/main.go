package main

import (
	"fmt"
	"gogin/RouterApp/app/blog"
	"gogin/RouterApp/app/shop"
	"gogin/RouterApp/routers"
)

func main() {
	// 加载多个app的路由配置
	routers.Include(shop.Routers, blog.Routers)
	r := routers.Init()
	if err := r.Run(); err != nil {
		fmt.Printf("startup service failed,err:%v\n\n", err)
	}
}
