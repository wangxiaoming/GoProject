package main

import (
	list2 "container/list"
	"fmt"
)

func main() {

	list := list2.New()
	list.PushBack("Hello")
	len := list.Len()

	fmt.Println(len, list.Front().Value)

}
