package main

import (
	"fmt"
	"time"
)

// 超时 对于一个连接外部资源，或者其它一些需要花费执行时间的操作的程序而言是很重要的
// 得益于通道和 select，在 Go中实现超时操作是简洁而优雅的
func main() {
	// 假如执行一个外部调用，并在 2 秒后通过通道 c1 返回它的执行结果
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "result 1"
	}()

	// 使用 select 实现一个超时操作
	// res := <- c1 等待结果，<-Time.After 等待超时时间 1 秒后发送的值
	// 由于 select 默认处理第一个已准备好的接收操作，如果这个操作超过了允许的 1 秒的话，将会执行超时 case
	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(time.Second * 1):
		fmt.Println("timeout 1")
	}

	// 如果允许一个长一点的超时时间 3 秒，将会成功的从 c2接收到值，并且打印出结果
	c2 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "result 2"
	}()
	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(time.Second * 3):
		fmt.Println("timeout 2")
	}
	// 使用这个 select 超时方式，需要使用通道传递结果
	// 这对于一般情况是个好的方式，因为其他重要的 Go 特性是基于通道和select 的
}
