package main

import (
	"fmt"
	"time"
)

// 在这个例子中，将看到如何使用 Go 协程和通道实现一个工作池

// 将要在多个并发实例中支持的任务
// 这些执行者将从 jobs 通道接收任务，并且通过 results 发送对应的结果
// 我们将让每个任务间隔 1s 来模仿一个耗时的任务
func workerForWorkPools(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Println("worker", id, "processing job", job)
		time.Sleep(time.Second)
		results <- job * 2
	}
}
func main() {
	// 为了使用 worker 工作池并且收集他们的结果，我们需要2 个通道
	jobs := make(chan int, 100)
	results := make(chan int, 100)
	now := time.Now()
	// 这里启动了 3 个 worker，初始是阻塞的，因为还没有传递任务
	for i := 1; i <= 3; i++ {
		go workerForWorkPools(i, jobs, results)
	}

	// 这里发送 9 个 jobs，然后 close 这些通道来表示这些就是所有的任务了
	for i := 1; i <= 9; i++ {
		jobs <- i
	}
	close(jobs)

	// 收集所有这些任务的返回值
	for i := 0; i < 9; i++ {
		fmt.Println("result:", <-results)
	}

	fmt.Println(time.Since(now))
	// 执行这个程序，显示 9 个任务被多个 worker 执行
	// 整个程序处理所有的任务仅执行了 3s 而不是 9s，是因为 3 个 worker是并行的
}
