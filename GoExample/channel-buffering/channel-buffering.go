package main

import "fmt"

func main() {
	// 默认通道是 无缓冲 的，这意味着只有在对应的接收（<- chan）通道准备好接收时，才允许进行发送（chan <-）
	// 有缓存通道允许在没有对应接收方的情况下，缓存限定数量的值

	// 这里 make 了一个字符串通道，最多允许缓存 2 个值
	messages := make(chan string, 2)

	// 由于这个通道是有缓冲的，因此可以将这些值发送到通道中，而无需并发的接收
	messages <- "buffered"
	messages <- "channel"

	// 可以正常接收这两个值
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
