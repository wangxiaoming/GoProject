package main

import (
	"fmt"
)

func Add(a, b int) {
	z := a + b
	fmt.Println("z:", z)
}

func main() {
	for i := 0; i < 10; i++ {
		go Add(i, i)
	}

	//time.Sleep(time.Second * 1000)  // 没有这一行会发现什么都没有，因为主线程不等待其他 goroutine 结束
}
