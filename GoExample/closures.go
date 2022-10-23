package main

import "fmt"

// Go 支持通过 闭包来使用 匿名函数
// 匿名函数在你想定义一个不需要命名的内联函数时是很实用的。

// 这个 intSeq 函数返回另一个在 intSeq 函数体内定义的匿名函数
// 这个返回的函数使用闭包的方式 隐藏 变量 i。
func intSeq() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func main() {
	// 我们调用 intSeq 函数，将返回值（也是一个函数）赋给nextInt
	// 这个函数的值包含了自己的值 i
	// 这样在每次调用 nextInt 时都会更新 i 的值
	nextInt := intSeq()

	// 通过多次调用 nextInt 来看看闭包的效果
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	// 为了确认这个状态对于这个特定的函数是唯一的
	// 我们重新创建并测试一下
	newInts := intSeq()
	fmt.Println(newInts())
}
