package main

import (
	"fmt"
	"time"
)

// 可以使用通道来同步 Go 协程间的执行状态
// 这儿有一个例子，使用阻塞接收的方式，实现了等待另一个协程完成
// 如果需要等待多个协程，WaitGroup 是一个更好的选择

// 将要在 Go 协程中运行这个函数
// done 通道将被用于通知其他 Go 协程这个函数已经工作完成
func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	// 发送一个值来通知已经完工
	done <- true
}

func main() {
	// 运行一个 worker Go协程，并给予用于通知的通道
	done := make(chan bool)
	go worker(done)

	// 程序将一直阻塞，直至收到 worker 使用通道发送的通知
	<-done

	// 如果把 <- done 这行代码从程序中移除
	// 程序甚至会在 worker还没开始运行时就结束了
}
