package main

import (
	"testing"
)

type MyInt int64

// Go 语言不允许隐式类型转换
// 别名和原有类型不能进行隐式类型转换
func TestImplicit(t *testing.T) {
	var a int32 = 1
	var b int64
	// a = b 编译错误
	// b = a 编译错误 不支持隐式类型转换
	b = int64(a)
	var c MyInt
	c = MyInt(b) // 显示类型转换可以
	// c = b
	t.Log(a, b, c)
}

func TestPoint(t *testing.T) {
	a := 1
	aPtr := &a
	// aPtr = aPtr + 1 go 不支持指针运算
	t.Log(a, aPtr)
	t.Logf("%T %T", a, aPtr)
}

func TestString(t *testing.T) {
	var s string // 初始化零值 空字符串
	t.Log("*" + s + "*")
	t.Log(len(s))
}
