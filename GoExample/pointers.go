package main

import "fmt"

// Go 支持 指针
// 允许在程序中通过引用传递值或者数据结构

// zeroVal 有一个 int 型参数，所以使用值传递
// zeroVal 将从调用它的那个函数中得到一个 iVal
func zeroVal(iVal int) {
	iVal = 0
}

// zeroPtr 有一和上面不同的 *int 参数，意味着它用了一个 int指针
// 函数体内的 *iPtr 接着解引用 这个指针，从它内存地址得到这个地址对应的当前值
// 对一个解引用的指针赋值将会改变这个指针引用的真实地址的值
func zeroPtr(iPtr *int) {
	*iPtr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroVal(i)
	fmt.Println("zeroVal:", i)

	// 通过 &i 语法来取得 i 的内存地址，例如一个变量i 的指针。
	zeroPtr(&i)
	fmt.Println("zeroPtr:", i)

	// 指针也是可以被打印的
	fmt.Println("pointer:", &i)

	// zeroVal 在 main 函数中不能改变 i 的值
	// 但是 zeroPtr 可以，因为它有一个这个变量的内存地址的引用。
}
