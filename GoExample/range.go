package main

import "fmt"

// range 迭代各种各样的数据结构
func main() {
	// 使用 range 来统计一个 slice 的元素的和
	// 数组也可以采用这种方法
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	// range 在数组和 slice 中都提供每个项的索引和值的访问
	// 上面不需要索引，所以使用 空值定义符_ 来忽略它
	// 有时候实际上是需要这个索引的
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	// range 在 map 中迭代键值对
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// range 在字符串中迭代 unicode 编码
	// 第一个返回值是rune 的起始字节位置，然后第二个是 rune 自己
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
