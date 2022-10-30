package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var lock sync.Mutex

	go func() {
		fmt.Println("into func1:", time.Now())
		lock.Lock()
		fmt.Println("func1 get lock at", time.Now())
		time.Sleep(time.Second)
		fmt.Println("func1 release lock at", time.Now())
		lock.Unlock()
	}()

	go func() {
		fmt.Println("into func2:", time.Now())
		lock.Lock()
		fmt.Println("func2 get lock at", time.Now())
		time.Sleep(time.Second)
		fmt.Println("func2 release lock at", time.Now())
		lock.Unlock()
	}()

	time.Sleep(time.Second * 3)
	fmt.Println("over ...")
}
