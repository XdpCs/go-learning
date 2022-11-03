package main

import "fmt"

// Go 原生支持多返回值
// 这个特性在 Go 语言中经常被用到
// 例如用来同时返回一个函数的结果和错误信息

// (int,int)在这个函数中标志着这个函数返回2个int
func vals() (int, int) {
	return 3, 7
}

func main() {
	// 这里通过多赋值 操作来使用这两个不同的返回值
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	// 如果仅仅想返回值的一部分的话，可以使用空白定义符 _。
	_, c := vals()
	fmt.Println(c)
}
