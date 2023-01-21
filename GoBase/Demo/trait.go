package main

// Go语言的并发是基于 goroutine 的，goroutine 类似于线程，但并非线程。可以将 goroutine 理解为一种虚拟线程。
// Go 语言运行时会参与调度 goroutine，并将 goroutine 合理地分配到每个 CPU 中，最大限度地使用CPU性能。
// 开启一个goroutine的消耗非常小（大约2KB的内存），你可以轻松创建数百万个goroutine。
// goroutine的特点：
//    1.具有可增长的分段堆栈。这意味着它们只在需要时才会使用更多内存
//    2.启动时间比线程快
//    3.原生支持利用channel安全地进行通信。
//    4.共享数据结构时无需使用互斥锁。
// 目前流行的项目结构：使用个人的github用户名来区分不同的包
// 自带gc
// 交叉编译，仅需更改环境变量
// 同一个目录中所有的go文件的package声明必须相同，所以main方法要单独放一个文件
