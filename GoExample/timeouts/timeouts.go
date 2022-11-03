package main

import (
	"fmt"
	"time"
)

// 超时 对于一个需要连接外部资源，或者有耗时较长的操作的程序而言是很重要的
// 得益于通道和 select，在 Go中实现超时操作是简洁而优雅的

func main() {
	// 假如执行一个外部调用，并在 2 秒后，使用通道 c1 返回它的执行结果
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "result 1"
	}()

	// 使用 select 实现一个超时操作
	// res := <- c1 等待结果，<-time.After 等待超时 (1 秒钟）后发送的值
	// 由于 select 默认处理第一个已准备好的接收操作，因此如果操作耗时超过了允许的 1 秒的话，将会执行超时 case
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
	// 运行这个程序，首先显示运行超时的操作，然后是成功接收的
}
