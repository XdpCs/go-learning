package main

import (
	"fmt"
	"gogin/Route/routers"
)

func main() {
	r := routers.SetupRouter()
	if err := r.Run(); err != nil {
		fmt.Printf("startup service failed,err:%v\n", err)
	}
}
