package main

import (
	"fmt"
	"runtime"
	"sync/atomic"
	"time"
)

// Go 中最主要的状态管理方式是通过通道间的沟通来完成的，还是有一些其他的方法来管理状态的
// 这里我们将看看如何使用 sync/atomic包在多个 Go 协程中进行 原子计数
func main() {
	// 将使用一个无符号整型数来表示（永远是正整数）这个计数器
	var ops uint64 = 0

	// 为了模拟并发更新，启动 50 个 Go 协程，对计数器每隔 1ms （译者注：应为非准确时间）进行一次加一操作
	for i := 0; i < 50; i++ {
		go func() {
			// 使用 AddUint64 来让计数器自动增加，使用& 语法来给出 ops 的内存地址
			atomic.AddUint64(&ops, 1)
			runtime.Gosched()
		}()
	}

	time.Sleep(time.Second)

	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("ops:", opsFinal)
}
