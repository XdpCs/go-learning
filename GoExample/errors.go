package main

import (
	"errors"
	"fmt"
)

// Go 语言使用一个独立的·明确的返回值来传递错误信息的
// Go 语言的处理方式能清楚的知道哪个函数返回了错误，并能像调用那些没有出错的函数一样调用

// 错误通常是最后一个返回值并且是 error 类型，一个内建的接口
func f1(arg int) (int, error) {
	if arg == 42 {
		// errors.New 构造一个使用给定的错误信息的基本error 值
		return -1, errors.New("can't work with 42")
	}

	// 返回错误值为 nil 代表没有错误
	return arg + 3, nil
}

// 通过实现 Error 方法来自定义 error 类型
// 这里使用自定义错误类型来表示上面的参数错误
type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		// 使用 &argError 语法来建立一个新的结构体，并提供了 arg 和 prob 这个两个字段的值
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func main() {
	for _, i := range []int{7, 42} {
		if r, err := f1(i); err != nil {
			fmt.Println("f1 failed:", err)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}
	for _, i := range []int{7, 42} {
		if r, err := f2(i); err != nil {
			fmt.Println("f2 failed:", err)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}

	// 如果想在程序中使用一个自定义错误类型中的数据
	// 需要通过类型断言来得到这个错误类型的实例
	_, err := f2(42)
	if a, ok := err.(*argError); ok {
		fmt.Println(a.arg)
		fmt.Println(a.prob)
	}
}
