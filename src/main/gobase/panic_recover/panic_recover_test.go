package panic

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

func TestPanicVxExit(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("revover from ", err)
		}
	}()
	fmt.Println("Start")
	panic(errors.New("Something wrong"))
	//os.Exit(-1)
}

func TestPanic(t *testing.T) {
	defer fmt.Println("in main")
	go func() {
		defer fmt.Println("in  goroutine")
		panic("")
	}()
	time.Sleep(1 * time.Second)
	defer fmt.Println("in end main")
}
