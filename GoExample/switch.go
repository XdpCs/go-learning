package main

import (
	"fmt"
	"time"
)

// switch 是多分支情况时快捷的条件语句
func main() {

	// 一个基本的 switch
	i := 2
	fmt.Print("write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	// 在同一个 case 语句中，你可以使用逗号来分隔多个表达式
	// 在这个例子中，还使用了可选的default 分支。
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("it's the weekend")
	default:
		fmt.Println("it's a weekday")
	}

	// 不带表达式的 switch 是实现 if/else 逻辑的另一种方式
	// 这里展示了 case 表达式是如何使用非常量的
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("it's before noon")
	default:
		fmt.Println("it's after noon")
	}

	// 类型开关 (type switch) 比较类型而非值
	// 可以用来发现一个接口值的类型
	// 在这个例子中，变量 t 在每个分支中会有相应的类型
}
