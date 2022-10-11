package main

import (
	"GoRedis/db"
	"fmt"
)

// brew 启动软件
// brew services start redis
// Redis支持数据的持久化，可以将内存中的数据保存在磁盘中，重启的时候可以再次加载进行使用。
// Redis不仅仅支持简单的key-value类型的数据，同时还提供string、list（链表）、set（集合）、hash表等数据结构的存储。
// Redis支持数据的备份，即master-slave模式的数据备份。主设备负责分配工作并整合结果，或作为指令的来源；从设备负责完成工作，一般只能和主设备通信
// 性能极高
// 丰富的数据类型
// 原子--Redis的所有操作都是原子性的，意思就是要么成功执行要么失败完全不执行。单个操作是原子性的。多个操作也支持事务，即原子性，通过MULTI和EXEC指令包起来
// 丰富的特性 – Redis还支持 publish/subscribe, 通知, key 过期等等特性
// Redis的数据类型都是基于基本数据结构的同时对程序员透明，无需进行额外的抽象
// Redis运行在内存中但是可以持久化到磁盘，所以在对不同数据集进行高速读写时需要权衡内存，因为数据量不能大于硬件内存
// 在内存数据库方面的另一个优点是，相比在磁盘上相同的复杂的数据结构，在内存中操作起来非常简单，这样Redis可以做很多内部复杂性很强的事情
// 在磁盘格式方面他们是紧凑的以追加的方式产生的，因为他们并不需要进行随机访问
func main() {
	conn, err := db.InitRedis()
	if err != nil {
		fmt.Println("conn")
	}
	defer conn.Close()
}
