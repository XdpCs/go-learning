package main

import "fmt"

// 在 Go 中，你可以不适用圆括号，但是花括号是需要的
// Go 里没有三目运算符，你只需要基本的条件判断的时候
// 你仍需要使用完整的 if 语句
func main() {
	// 基本例子
	if 1118%2 == 0 {
		fmt.Println("1118 is even")
	} else {
		fmt.Println("1118 is odd")
	}
	// 可以只用 if 语句
	if 1114%2 == 0 {
		fmt.Println("1114 is divisible by 2")
	}

	// 在条件语句之前可以有一个语句
	// 任何在这里声明的变量都可以在所有的条件分支中使用
	if num := 1118; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}
