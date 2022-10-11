package main

import (
	"GoRedis/db"
	"fmt"
)

func main() {
	conn, err := db.InitRedis()
	if err != nil {
		fmt.Println("conn")
	}
	defer conn.Close()
	_, err = conn.Do("expire", "fyy", 10)
	if err != nil {
		fmt.Println(err)
		return
	}
}
