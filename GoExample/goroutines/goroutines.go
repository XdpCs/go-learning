package main

import (
	"fmt"
	"time"
)

// Go 协程 是轻量级的执行线程

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	// 假设有一个函数叫做 f(s)
	// 一般会这样 同步地 调用它
	f("direct")

	// 使用 go f(s) 在一个 Go 协程中调用这个函数
	// 这个新的 Go 协程将会并发地执行这个函数
	go f("goroutine")

	// 可以为匿名函数启动一个 Go 协程
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// 现在这两个 Go 协程在独立的 Go 协程中异步地运行，所以需要等它们执行结束
	// 更好的方法是使用WaitGroup
	time.Sleep(time.Second)
	fmt.Println("done")
	// 当运行这个程序时，将首先看到阻塞式调用的输出，然后是两个 Go 协程的交替输出
	// 这种交替的情况表示 Go 运行时是以并发的方式运行协程的
}
