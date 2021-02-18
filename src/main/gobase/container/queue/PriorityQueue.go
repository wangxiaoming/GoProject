package main

import (
	"container/heap"
	"fmt"
)

type QItem interface {
	Less(item QItem) bool
}

// priorityQueueImpl 用于优先队列底层实现
type priorityQueueImpl []QItem

func (pqi priorityQueueImpl) Swap(i, j int) {
	pqi[i], pqi[j] = pqi[j], pqi[j]
}

// Len 获取队列长度
func (pqi priorityQueueImpl) Len() int {
	return len(pqi)
}

// Less 用来比较元素
func (qpi priorityQueueImpl) Less(i, j int) bool {
	return qpi[i].Less(qpi[j])
}

// Pop 将一个对象压入队列中
func (qpi *priorityQueueImpl) Push(x interface{}) {
	item := x.(QItem)
	*qpi = append(*qpi, item)
}

func (qpi *priorityQueueImpl) Pop() interface{} {
	old := *qpi
	n := len(old)
	// 弹出最后一个
	item := old[n-1]
	*qpi = old[0 : n-1]
	return item
}

type PriorityQueue struct {
	priorityQueueImpl
}

func NewPriorityQueue() *PriorityQueue {
	var pq PriorityQueue
	heap.Init(&pq.priorityQueueImpl)
	return &pq
}

func (qp *PriorityQueue) Push(x QItem) {
	heap.Push(&qp.priorityQueueImpl, x)
}

func (qp *PriorityQueue) Pop() QItem {
	return heap.Pop(&qp.priorityQueueImpl).(QItem)
}

func (qp *PriorityQueue) Front() QItem {
	return qp.priorityQueueImpl[0]
}

func (pq *PriorityQueue) Length() int {
	return len(pq.priorityQueueImpl)
}

type Int int

func (i Int) Less(j QItem) bool {
	return i < j.(Int)
}
func main() {
	pq := NewPriorityQueue()
	pq.Push(Int(123))
	pq.Push(Int(234))
	pq.Push(Int(135))

	for _, value := range pq.priorityQueueImpl {
		fmt.Println(value)
	}

}
