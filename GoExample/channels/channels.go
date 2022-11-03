package main

import "fmt"

// 通道 是连接多个 Go 协程的管道
// 可以从一个 Go 协程将值发送到通道，然后在别的 Go 协程中接收
func main() {
	// 使用 make(chan val-type) 创建一个新的通道
	// 通道类型就是需要传递值的类型
	messages := make(chan string)

	// 使用 channel <- 语法 发送 一个新的值到通道中
	// 这里在一个新的 Go 协程中发送 "ping" 到上面创建的messages 通道中
	go func() {
		messages <- "ping"
	}()

	// 使用 <-channel 语法从通道中 接收 一个值
	// 这里将接收在上面发送的 "ping" 消息并打印出来
	msg := <-messages
	fmt.Println(msg)
	// 运行程序时，通过通道，消息 "ping" 成功的从一个 Go 协程传到另一个 Go 协程
	// 默认发送和接收操作是阻塞的，直到发送方和接收方都准备就绪
	// 这个特性允许不使用任何其它的同步操作，来在程序结尾等待消息 "ping"
}
