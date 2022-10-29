package main

import (
	"fmt"
	"time"
)

// 速率限制(英) 是一个重要的控制服务资源利用和质量的途径
// Go 通过 Go 协程、通道和打点器优美的支持了速率限制

// 假设想限制接收请求的处理，将这些请求发送给一个相同的通道
func main() {
	requests := make(chan int, 5)
	for i := 0; i < 5; i++ {
		requests <- i
	}
	close(requests)

	// 这个 limiter 通道将每 200ms 接收一个值
	// 这个是速率限制任务中的管理器
	limiter := time.Tick(time.Millisecond * 200)

	// 通过在每次请求前阻塞 limiter 通道的一个接收，限制自己每 200ms 执行一次请求
	for request := range requests {
		<-limiter
		fmt.Println("request", request, time.Now())
	}

	// 有时候想临时进行速率限制，并且不影响整体的速率控制可以通过通道缓冲来实现
	// 这个 burstyLimiter 通道用来进行 3 次临时的脉冲型速率限制
	burstLimiter := make(chan time.Time, 3)

	// 将通道填充需要临时改变3次的值
	for i := 0; i < 3; i++ {
		burstLimiter <- time.Now()
	}

	// 每 200 ms 将添加一个新的值到 burstyLimiter中，直到达到 3 个的限制
	go func() {
		for t := range time.Tick(time.Millisecond * 200) {
			burstLimiter <- t
		}
	}()

	// 现在模拟超过 5 个的接入请求
	// 它们中刚开始的 3 个将由于受 burstyLimiter 的“脉冲”影响
	burstRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstRequests <- i
	}
	close(burstRequests)

	for request := range burstRequests {
		<-burstLimiter
		fmt.Println("request", request, time.Now())
	}
	// 运行程序，第一批请求意料之中的大约每 200ms 处理一次
	// 第二批请求，直接连续处理了 3 次，这是由于这个“脉冲”速率控制，然后大约每 200ms 处理其余的 2 个
}
