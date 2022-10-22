package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func testDemoTwo1() {
	for i := 0; i < 10; i++ {
		fmt.Println("test1()", i)
		time.Sleep(time.Microsecond * 100)
	}
	wg.Done() // 协程计数器-1
}
func testDemoTwo2() {
	for i := 0; i < 10; i++ {
		fmt.Println("test2()", i)
		time.Sleep(time.Microsecond * 100)
	}
	wg.Done() // 协程计数器-1
}

// 主线程结束 协程没执行完，协程还是会被结束 ，通过sync.waitgroup进行解决
func main() {
	wg.Add(1)         // 协程计数器加1
	go testDemoTwo1() // 开启一个协程
	wg.Add(1)
	go testDemoTwo2()
	for i := 0; i < 10; i++ {
		fmt.Println("main()", i)
		time.Sleep(time.Microsecond * 50)
	}
	wg.Wait() // 等待协程执行完毕
	fmt.Println("主线程退出")
}
