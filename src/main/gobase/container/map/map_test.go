package main

import "testing"

func TestInitMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 2, 3: 3}
	t.Log(m1[2])
	t.Logf("len m1 = %d", len(m1))

	m2 := map[int]int{}
	m2[4] = 16
	t.Logf("len m2 = %d", len(m2))

	m3 := make(map[int]int, 10)
	t.Logf("len m3 = %d", len(m3))
}

func TestAccessNotExitstringkey(t *testing.T) {
	m1 := map[int]int{}
	t.Log(m1[1])

	m1[2] = 0
	t.Log(m1[2])

	m1[3] = 0

	if v, ok := m1[3]; ok {
		t.Logf("key 3 value is %d", v)
	} else {
		t.Log("key 3 is not existing")
	}

}
