package main

import (
	"fmt"
	"time"
)

// 在这个例子中，将看到如何使用 Go 协程和通道实现一个工作池

// worker程序
// 会并发的运行多个worker
// 这些执行者将从 jobs 通道接收工作，并在 results 发送对应的结果
// 将让每个任务间隔 1s 来模仿一个耗时的任务
func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Println("worker", id, "started  job", job)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", job)
		results <- job * 2
	}
}
func main() {
	// 为了使用 worker 工作池并且收集他们的结果，需要2 个通道
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)
	now := time.Now()
	// 这里启动了 3 个 worker，初始是阻塞的，因为还没有传递任务
	for i := 1; i <= 3; i++ {
		go worker(i, jobs, results)
	}

	// 这里发送 5 个 jobs，然后 close 这些通道, 表示这些就是所有的任务了
	for i := 1; i <= numJobs; i++ {
		jobs <- i
	}
	close(jobs)

	// 收集所有这些任务的返回值
	// 这也确保了所有的 worker 协程都已完成
	// 另一个等待多个协程的方法是使用WaitGroup
	for i := 0; i < numJobs; i++ {
		fmt.Println("result:", <-results)
	}

	fmt.Println(time.Since(now))
	// 运行这个程序，显示 5 个任务被多个 worker 执行
	// 整个程序处理所有的任务仅执行了 2s 而不是 5s，是因为 3 个 worker是并行的
}
