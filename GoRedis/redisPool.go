package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

// Redis 是基于内存的数据库，使用之前需要建立连接，建立断开连接需要消耗大量的时间
// 使用连接池可以实现在客户端建立多个连接，需要的时候从连接池拿过来，用完了再放回去。这样就节省了建立、断开连接所消耗的时间

var pool *redis.Pool

func init() {
	pool = &redis.Pool{ // 实例化一个连接池
		MaxIdle:     16,  //最初的连接数量
		MaxActive:   0,   //连接池最大连接数量，不确定可以用0（0表示自动定义），按需分配
		IdleTimeout: 300, //连接关闭时间 （300秒不使用自动关闭）
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "localhost:6379")
		},
	}
}

func main() {
	conn := pool.Get() //从连接池，取一个连接
	defer conn.Close()
	_, err := conn.Do("Set", "fyyxdp", 1118)
	if err != nil {
		fmt.Println(err)
		return
	}
	result, err := redis.Int(conn.Do("Get", "fyyxdp"))
	if err != nil {
		fmt.Println("get fyyxdp failed :", err)
		return
	}
	fmt.Println(result)
	pool.Close() // 关闭连接池
}
