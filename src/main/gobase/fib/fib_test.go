package main

import "testing"

func TestFibList(t *testing.T) {
	a := 1
	b := 1

	t.Log(a)

	for i := 0; i < 5; i++ {
		t.Log("", b)
		tmp := a
		a = b
		b = tmp + a
	}
	//t.Log(b)

}

// 交换两个变量的值
func TestExchange(t *testing.T) {
	a := 1
	b := 2
	// tmp := a
	// a = b
	// b = a
	a, b = b, a
	t.Log(a, b)
}
