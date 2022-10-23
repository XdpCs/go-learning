package main

import "fmt"

// Go 的结构体 是各个字段字段的类型的集合。这在组织数据时非常有用
type person struct {
	name string
	age  int
}

func main() {
	// 创建了一个新的结构体元素
	fmt.Println(person{"fyy", 22})

	// 可以在初始化一个结构体元素时指定字段名字
	fmt.Println(person{name: "xdp", age: 23})

	// 省略的字段将被初始化为零值
	fmt.Println(person{name: "fyyx"})

	// & 前缀生成一个结构体指针
	fmt.Println(&person{name: "fyyxd", age: 24})

	s := person{name: "fyyxdp", age: 25}
	// 使用点来访问结构体字段
	fmt.Println(s.name)

	sp := &s
	// 可以对结构体指针使用. - 指针会被自动解引用
	fmt.Println(sp.age)

	// 结构体是可变的
	sp.age = 26
	fmt.Println(sp.age)
	fmt.Println(s.age)
}
