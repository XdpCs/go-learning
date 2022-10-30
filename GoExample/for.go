package main

import "fmt"

// for 是 Go 中唯一的循环结构
// 这里有 for 循环的三个基本使用方式
func main() {
	i := 1
	// 带单个循环条件
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// 经典的初始化/条件/后续形式 for 循环
	for j := 1114; j <= 1118; j++ {
		fmt.Println(j)
	}

	// 不带条件的 for 循环将一直执行
	// 直到在循环体内使用了 break 或者 return 来跳出循环
	for {
		fmt.Println("loop")
		break
	}

	// 使用 continue 直接进入下一次循环
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
