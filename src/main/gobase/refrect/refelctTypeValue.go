package main

import (
	"fmt"
	"reflect"
)

func main() {
	author := "xiaoming"
	fmt.Println("TypeOf author:", reflect.TypeOf(author))
	fmt.Println("ValueOf author:", reflect.ValueOf(author))
}
