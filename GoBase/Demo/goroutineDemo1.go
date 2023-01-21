package main

import (
	"fmt"
	"time"
)

func testDemoOne() {
	for i := 0; i < 10; i++ {
		fmt.Println("test()", i)
		time.Sleep(time.Microsecond * 100)
	}
}

// 主线程结束 协程没执行完，协程还是会被结束
func main() {
	go testDemoOne()
	for i := 0; i < 10; i++ {
		fmt.Println("main()", i)
		time.Sleep(time.Microsecond * 50)
	}
}
