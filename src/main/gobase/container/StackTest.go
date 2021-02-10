package main

import (
	"container/list"
	"fmt"
)

type Stack struct {
	data *list.List
	size int
}

func (q *Stack) pop() {
	if q.size > 0 {
		//  找到队尾的数
		tmp := q.data.Back()
		q.data.Remove(tmp)
		q.size--
	}
}

func (q *Stack) top() interface{} {
	return q.data.Back().Value
}

func (q *Stack) push(value interface{}) {
	q.data.PushBack(value)
	q.size++
}
func main() {

	var stack = new(Stack)
	stack.data = list.New()
	stack.push(123)

	fmt.Println(stack.top())
	stack.push("hello")
	stack.push("World")
	fmt.Println(stack.size)

}
