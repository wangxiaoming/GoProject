package main

import (
	"GoProject/src/main/gobase"
	"GoProject/src/main/gobase/hello"
	"fmt"
)

func main() {
	fmt.Println(gobase.Add(1, 2))
	hello.Print()
	fmt.Println("Hello World")
}
