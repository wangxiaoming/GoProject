package main

import (
	"container/list"
	"fmt"
)

func main() {
	l := list.New()
	l1 := list.New()
	l1.PushBack("hello")
	l1.PushBack("World!")
	l1.PushBack("Test")
	l.PushBack(l1)

	for e := l.Front(); e != nil; e = e.Next() {
		curl := e.Value.(*list.List)
		for p := curl.Front(); p != nil; p = p.Next() {
			fmt.Println(p.Value)
		}
	}
}
