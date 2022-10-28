package main

import (
	"fmt"
	"math"
)

// 接口 是方法特征的命名集合

// 几何体的基本接口
type geometry interface {
	area() float64
	perim() float64
}

type rectangle struct {
	width, height float64
}

type circle struct {
	radius float64
}

// 要在 Go 中实现一个接口
// 只需要实现接口中的所有方法
// rectangle 实现了 geometry 接口
func (r rectangle) area() float64 {
	return r.width * r.height
}

func (r rectangle) perim() float64 {
	return 2 * (r.width + r.height)
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// 如果一个变量的是接口类型，那么可以调用这个被命名的接口中的方法
// 这里有一个一通用的 measure 函数，利用这个特性，它可以用在任何 geometry 上
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rectangle{width: 3, height: 4}
	c := circle{radius: 5}

	// 结构体类型 circle 和 rectangle 都实现了 geometry接口
	// 我们可以使用它们的实例作为 measure 的参数。
	measure(r)
	measure(c)
}
