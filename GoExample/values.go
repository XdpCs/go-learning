package main

import "fmt"

// Go 拥有各值类型，包括字符串，整形，浮点型，布尔型等
func main() {
	// 字符串可以通过 + 连接
	fmt.Println("xdp" + "fyy")
	fmt.Println("1+1=", 1+1)
	fmt.Println("1118.0/1114.0=", 1118.0/1114.0)
	// 逻辑运算符
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}
