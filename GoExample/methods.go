package main

import "fmt"

//  Go 支持在结构体类型中定义方法
type rect struct {
	width, height int
}

// area 是一个拥有 *rect 类型接收器(receiver)的方法
func (r *rect) area() int {
	return r.width * r.height
}

// 可以为值类型或者指针类型的接收者定义方法
// 这是一个值类型接收者的例子
func (r rect) perim() int {
	r.width = 100
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}
	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())

	// Go 自动处理方法调用时的值和指针之间的转化
	// 想要避免在调用方法时产生一个拷贝
	// 或者想让方法可以修改接受结构体的值
	// 可以使用指针来调用方法
	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())
}
