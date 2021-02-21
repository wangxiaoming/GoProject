package main

import (
	"testing"
)

const (
	Readable = 1 << iota
	Writeable
	Executeable
)

// 相同维数且含有相同个数元素的数组才可以比较
// 每个元素相同的才相等
func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 2, 3, 4}
	// c := [...]int{1, 2, 3, 4, 5}
	d := [...]int{1, 3, 2, 4} // 顺序不同
	t.Log(a == b)
	// t.Log(a == c) 编译错误 长度不同的数组
	t.Log(a == d)
}

func TestBitClear(t *testing.T) {
	a := 7             //0111
	a = a &^ Writeable // 写权限清掉
	t.Log(a)           // 0111 &^ 0010 = 0101
	a = a &^ Readable  // 0101 &^ 0001 = 0100
	t.Log(a)
	t.Log(a&Readable == Readable)
}
