package main

import "fmt"

// 在Go中，变量被显式声明，并被编译器所用来检查函数调用时的类型正确性
func main() {
	// var 声明 1 个或者多个变量
	// 你可以申明一次性声明多个变量
	var a string = "fyy"
	fmt.Println(a)
	var b, c int = 1, 2
	fmt.Println(b, c)

	// Go 将自动推断已经初始化的变量类型
	var d = true
	fmt.Println(d)

	// 声明变量且没有给出对应的初始值时，变量将会初始化为零值
	// 例如，一个 int 的零值是 0
	var e int
	fmt.Println(e)

	// := 语句是申明并初始化变量的简写
	// var f string = "xdp"
	f := "xdp"
	fmt.Println(f)
}
