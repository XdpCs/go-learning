package main

import (
	"fmt"
	"sync"
)

// 对于更加复杂的情况，可以使用一个互斥量 来在 Go 协程间安全的访问数据

// Container 中定义了 counters 的 map ，由于希望从多个 goroutine 同时更新它
// 因此添加了一个 互斥锁Mutex 来同步访问
// 请注意不能复制互斥锁，如果需要传递这个 struct，应使用指针完成
type Container struct {
	mu       sync.Mutex
	counters map[string]int
}

func (c *Container) inc(name string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[name]++
}

func main() {
	// 请注意，互斥量的零值是可用的
	// 因此这里不需要初始化
	c := Container{
		counters: map[string]int{"a": 0, "b": 0},
	}

	var wg sync.WaitGroup

	// 这个函数在循环中递增对 name 的计数
	doIncrement := func(name string, n int) {
		for i := 0; i < n; i++ {
			c.inc(name)
		}
		wg.Done()
	}

	// 同时运行多个 goroutines
	// 请注意，它们都访问相同的 Container，其中两个访问相同的计数器
	wg.Add(3)
	go doIncrement("a", 10000)
	go doIncrement("a", 10000)
	go doIncrement("b", 10000)

	// 等待上面的 goroutines 都执行结束
	wg.Wait()
	fmt.Println(c.counters)
}
