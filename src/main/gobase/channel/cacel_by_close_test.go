package main

import (
	"fmt"
	"testing"
	"time"
)

func cancel_1(cancelChan chan struct{}) {
	cancelChan <- struct{}{}
}

func cancel_2(cancelChan chan struct{}) {
	close(cancelChan)
}

func isCaneled(cancelChan chan struct{}) bool {
	select {
	case <-cancelChan:
		return true
	default:
		return false
	}
}

func TestCancel(t *testing.T) {
	cancelChan := make(chan struct{}, 0)
	for i := 0; i < 5; i++ {
		go func(i int, cancelChan chan struct{}) {
			for {
				if isCaneled(cancelChan) {
					break
				}
				time.Sleep(time.Second * 1)
			}
			fmt.Println(i, "Done")
		}(i, cancelChan)
	}
	cancel_1(cancelChan)
	time.Sleep(time.Second * 1)
}
