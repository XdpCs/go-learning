package main

import "fmt"

// Go 中，数组 是一个具有编号且长度固定的元素序列
func main() {
	// 建了一个数组 a 来存放刚好 5 个 int
	// 元素的类型和长度都是数组类型的一部分
	// 数组默认是零值的，对于 int 数组来说,元素的零值是 0
	var a [5]int
	fmt.Println("emp:", a)

	// 可以使用 array[index] = value 语法来设置数组指定位置的值
	// 或者用 array[index] 得到值
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	// 使用内置函数 len 返回数组的长度
	fmt.Println("len:", len(a))

	// 使用这个语法在一行内初始化一个数组
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// 通过编译器 计数数组长度
	c := [...]int{1, 2, 3, 4, 5}
	fmt.Println(len(c))

	// 数组类型是一维的
	// 但可以组合构造多维的数据结构
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}

	// 在使用 fmt.Println 来打印数组的时候，会使用[v1 v2 v3 ...] 的格式显示
	fmt.Println("2d: ", twoD)
	// 在 Go 程序中，相较于数组，用得更多的是切片(slice)
}
