package main

import (
	"GoRedis/db"
	"fmt"
	"github.com/gomodule/redigo/redis"
)

// redis 启动命令
// redis-cli -h 127.0.0.1 -p 6379
func main() {
	conn, err := db.InitRedis()
	if err != nil {
		fmt.Println("conn")
	}
	defer conn.Close()
	_, err = conn.Do("Set", "fyy", 1118)
	if err != nil {
		fmt.Println(err)
		return
	}
	result, err := redis.Int(conn.Do("Get", "fyy"))
	if err != nil {
		fmt.Println("get fyy failed", err)
		return
	}
	fmt.Println(result)
}
