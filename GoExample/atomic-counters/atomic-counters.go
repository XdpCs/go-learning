package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Go 中最主要的状态管理机制是依靠通道间的通信来完成的
// 如何使用 sync/atomic 包在多个协程中进行 原子计数
func main() {
	// 将使用一个无符号整型数来表示（永远是正整数）这个计数器
	var ops uint64 = 0

	// WaitGroup 帮助等待所有协程完成它们的工作
	var wg sync.WaitGroup

	// 会启动 50 个协程，并且每个协程会将计数器递增 1000 次
	for i := 0; i < 50; i++ {
		wg.Add(1)

		go func() {
			for j := 0; j < 1000; j++ {
				// 使用 AddUint64 来让计数器自动增加， 使用 & 语法给定 ops 的内存地址
				atomic.AddUint64(&ops, 1)
			}
			wg.Done()
		}()
	}

	// 等待，直到所有协程完成
	wg.Wait()

	// 现在可以安全的访问 ops，因为此时没有协程写入 ops
	// 此外，还可以使用 atomic.LoadUint64 之类的函数，在原子更新的同时安全地读取它们
	fmt.Println("ops", ops)

	// 预计会进行 50,000 次操作
	// 如果使用非原子的 ops++ 来增加计数器， 由于多个协程会互相干扰，运行时值会改变
	// 可能会导致得到一个不同的数字
	// 此外，运行程序时带上 -race 标志，可以获取数据竞争失败的详情
}
