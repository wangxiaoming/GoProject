package main

import (
	"fmt"
)

// 人
type Person struct {
	name string
	sex  string
	age  int
}

// 接口
type People interface {
	GetName() string
	GetAge() int
}

func (p *Person) GetName() string {
	return fmt.Sprintf("user: %s, %s", p.name, p.sex)
}

func (p *Person) GetAge() int {
	return p.age
}

func main() {
	var t People = &Person{"xiaoming", "man", 18}
	fmt.Println(t.GetAge())
	fmt.Println(t.GetName())
}
