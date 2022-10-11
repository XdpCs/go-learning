package db

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func InitRedis() (redis.Conn, error) {
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("conn redis failed,", err)
		return conn, err
	}
	fmt.Println("redis conn success")
	return conn, err
}
