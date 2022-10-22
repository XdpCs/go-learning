package main

import (
	"fmt"
	"runtime"
)

func main() {
	// 获取当前计算机上的cpu个数
	numCPU := runtime.NumCPU()
	fmt.Println("cpuNum=", numCPU)
	// 设置使用多个Cpu
	//runtime.GOMAXPROCS(numCPU - 1)
	//fmt.Println("ok")
}
