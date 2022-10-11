package main

import (
	"GoRedis/db"
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func main() {
	conn, err := db.InitRedis()
	if err != nil {
		fmt.Println("conn failed,", err)
		return
	}
	defer conn.Close()
	_, err = conn.Do("lpush", "book_list", "fyy", "xdp", 1118, 1114)
	if err != nil {
		fmt.Println(err)
		return
	}
	result, err := redis.String(conn.Do("lpop", "book_list"))
	if err != nil {
		fmt.Println("lpop failed,", err)
		return
	}
	fmt.Println(result)
}
