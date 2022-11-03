package main

import (
	"fmt"
	"time"
)

// 速率限制是控制服务资源利用和质量的重要机制
// Go 通过 Go 协程、通道和打点器优美的支持了速率限制

// 将看一个基本的速率限制
// 假设想限制接收请求的处理，将这些请求发送给一个相同的通道
func main() {
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// limiter 通道每 200ms 接收一个值
	// 这是任务速率限制的调度器
	limiter := time.Tick(time.Millisecond * 200)

	// 通过在每次请求前阻塞 limiter 通道的一个接收
	// 可以将频率限制为，每 200ms 执行一次请求
	for request := range requests {
		<-limiter
		fmt.Println("request", request, time.Now())
	}

	// 有时候可能希望在速率限制方案中允许短暂的并发请求，并同时保留总体速率限制
	// 可以通过缓冲通道来完成此任务
	// burstLimiter 通道允许最多 3 个爆发（bursts）事件
	burstLimiter := make(chan time.Time, 3)

	// 填充通道，表示允许的爆发（bursts）
	for i := 0; i < 3; i++ {
		burstLimiter <- time.Now()
	}

	// 每 200 ms 将尝试添加一个新的值到 burstLimiter中，直到达到 3 个的限制
	go func() {
		for t := range time.Tick(time.Millisecond * 200) {
			burstLimiter <- t
		}
	}()

	// 现在模拟另外 5 个的传入请求
	// 受益于 burstLimiter 的爆发（bursts）能力，前 3 个请求可以快速完成
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
	// 第二批请求，由于爆发（burstable）速率控制
	// 直接连续处理了 3 个请求， 然后以大约每 200ms 一次的速度，处理了剩余的 2 个请求
}
