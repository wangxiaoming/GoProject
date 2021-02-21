package main

import "testing"

const (
	Monday = 1 + iota
	Tuesday
	Wednesday
)

const (
	// 连续常量的定义
	Readble = 1 << iota
	Writeble
	Executable
)

func TestConstantTry(t *testing.T) {
	t.Log(Monday, Tuesday)
}

func TestConstantTyr1(t *testing.T) {
	a := 1 // 0001
	t.Log(Readble, Writeble, Executable)
	t.Log(Readble&a == Readble, a&Writeble == Writeble, a&Executable == Executable)
}
