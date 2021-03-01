package main

import (
	"fmt"
	"sync"
	"testing"
)

func dataProducer(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
		wg.Done()
	}()
}

func dataConsumer(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 11; i++ {
			data := <-ch
			fmt.Println(data)
		}
		wg.Done()
	}()
}

func TestCloseChanner(t *testing.T) {
	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(1)
	dataProducer(ch, &wg)
	wg.Add(1)
	dataConsumer(ch, &wg)
	wg.Wait()
}
