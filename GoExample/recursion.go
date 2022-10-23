package main

import "fmt"

// Go 支持 递归
// 这里是一个经典的阶乘示例

// face 函数在到达 face(0) 前一直调用自身
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	fmt.Println(fact(7))
}
