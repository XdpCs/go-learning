package main

import (
	"fmt"
	"strings"
)

// 标准库的 strings 包提供了很多有用的字符串相关的函数
// 这儿有一些用来对 strings 包有一个初步了解的例子

// 给 fmt.Println 一个较短的别名， 因为随后会大量的使用它
var p = fmt.Println

func main() {
	// 这是一些 strings 中有用的函数例子
	// 由于它们都是包的函数，而不是字符串对象自身的方法
	// 这意味着需要在调用函数时，将字符串作为第一个参数进行传递
	// 可以在 strings 包文档中找到更多的函数
	p("Contains: ", strings.Contains("fyyfyy", "fy"))
	p("Count: ", strings.Count("fyyfyy", "f"))
	p("HasPrefix: ", strings.HasPrefix("fyyfyy", "fy"))
	p("HasSuffix: ", strings.HasSuffix("fyyfyy", "yy"))
	p("Index: ", strings.Index("fyyfyy", "f"))
	p("Join: ", strings.Join([]string{"fyy", "xdp"}, "loves"))
	p("Repeat: ", strings.Repeat("fyy", 5))
	p("Replace: ", strings.Replace("fyyfyy", "fy", "dp", -1))
	p("Replace: ", strings.Replace("fyyfyy", "fy", "dp", 1))
	p("Spilt: ", strings.Split("f-y-y", "-"))
	p("ToLower:   ", strings.ToLower("FYY"))
	p("ToUpper:   ", strings.ToUpper("fyy"))
	p()

	// 虽然不是 strings 的函数，但获取字符串长度（以字节为单位）以及通过索引获取一个字节的机制
	p("Len: ", len("fyy"))
	p("Char:", "fyy"[1])
	// 注意，上面的 len 以及索引工作在字节级别上
	// Go 使用 UTF-8 编码字符串，因此通常按原样使用
	// 如果可能使用多字节的字符，则需要使用可识别编码的操作
}
