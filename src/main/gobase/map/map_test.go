package main

import "testing"

func TestInitMap(t *testing.T) {
	m := map[string]int{"one": 1, "two": 2, "three": 3}
	t.Log(m)

	m2 := map[int]int{}
	m2[4] = 16
	t.Log("len m2", len(m2))

	m3 := make(map[int]int, 10)
	t.Log("len m3", len(m3))
}

// 测试访问
func TestAccessNotExit(t *testing.T) {
	m1 := map[int]int{}
	t.Log(m1[1])
	m1[2] = 1
	t.Log(m1[2])
	m1[3] = 3
	if v, ok := m1[3]; ok {
		t.Log("key 3 value is", v)
	} else {
		t.Log("key 3 is not existing")
	}
}

func TestTravelMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 2, 3: 3}
	// k 数组表示的是index， map 表示的是 key
	for k, v := range m1 {
		t.Log(k, v)
	}
}
