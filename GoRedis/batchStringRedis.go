package main

import (
	"GoRedis/db"
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func main() {
	conn, err := db.InitRedis()
	if err != nil {
		fmt.Println("conn")
	}
	defer conn.Close()
	_, err = conn.Do("Mset", "fyy", 1118, "xdp", 1114)
	if err != nil {
		fmt.Println(err)
		return
	}

	result, err := redis.Ints(conn.Do("Mget", "fyy", "xdp"))
	if err != nil {
		fmt.Println("get batch string failed", err)
		return
	}

	for _, v := range result {
		fmt.Println(v)
	}
}
