package main

import "fmt"

// for 和 range 使用这个语法来遍历从通道中取得的值
func main() {
	// 将遍历在 queue 通道中的两个值
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	// 这个 range 迭代从 queue 中得到的每个值。因为我们在前面 close 了这个通道，这个迭代会在接收完 2 个值之后结束
	// 如果没有 close 它，将在这个循环中继续阻塞执行，等待接收第三个值
	for elem := range queue {
		fmt.Println(elem)
	}
	// 这个例子可以看到一个非空的通道也是可以关闭的，但是通道中剩下的值仍然可以被接收到
}
