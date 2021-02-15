package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan int, 1) // 创建带一个缓存潮的异步通道

	go func() {
		var ok bool
		select {
		case _, ok = <-ch:
		default:
			ch <- 1
		}
		if !ok {
			fmt.Println(<-ch)
			return
		}
		i := <-ch
		fmt.Println("Value Received:", i)
	}()

	time.Sleep(time.Second * 10)

}
