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
	_, err = conn.Do("HSet", "books", "fyy", 1118)
	if err != nil {
		fmt.Println(err)
		return
	}
	result, err := redis.Int(conn.Do("Hget", "books", "fyy"))
	if err != nil {
		fmt.Println("get fyy failed,", err)
		return
	}
	fmt.Println(result)
}
