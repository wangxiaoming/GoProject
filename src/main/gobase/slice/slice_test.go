package slice

import "testing"

func TestingSlice(t *testing.T) {
	t.Log("da")
}

func TestSliceInit(t *testing.T) {
	var s0 []int // 和数组不一样的地方就是没有指定长度
	t.Log(len(s0), cap(s0))

	s0 = append(s0, 1)
	t.Log(len(s0), cap(s0))

	s1 := []int{1, 2, 3, 4}
	t.Log(len(s1), cap(s1))

	s2 := make([]int, 3, 5)
	t.Log(len(s2), cap(s2))
	t.Log(s2[0], s2[1], s2[3])

}
