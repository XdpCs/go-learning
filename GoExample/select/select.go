package main

import (
	"fmt"
	"time"
)

// Go 的选择器 让你可以同时等待多个通道操作
// Go 协程、通道和选择器的结合是 Go 的一个强大特性
func main() {
	// 在这个例子中，将从两个通道中选择
	c1 := make(chan string)
	c2 := make(chan string)

	// 各个通道将在一定时间后接收一个值，通过这种方式来模拟并行的协程执行（例如，RPC 操作）时造成的阻塞（耗时）
	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "two"
	}()

	// 使用 select 关键字来同时等待这两个值，并打印各自接收到的值
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}

	// 跟预期的一样，首先接收到值 "one"，然后是 "two"
	// 注意，程序总共仅运行了两秒左右
	// 因为 1 秒 和 2 秒的 Sleeps 是并发执行的
}
