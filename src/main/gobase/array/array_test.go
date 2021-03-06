package array

import (
	"fmt"
	"testing"
)

func TestArrayInit(t *testing.T) {
	var arr [3]int                 // 声明并初始化为默认值
	arr1 := [4]int{1, 2, 3, 4}     // 声明同时初始化
	c := [2][2]int{{1, 2}, {3, 4}} // 多维数组初始化
	t.Log(c)

	arr3 := [...]int{1, 2, 3, 5}

	arr[1] = 5

	t.Log(arr[1], arr[2])
	t.Log(arr1, arr3)
}

func TestArrayTravel(t *testing.T) {
	arr3 := [...]int{1, 2, 3, 4}
	for i := 0; i < len(arr3); i++ {
		t.Log(arr3[i])
	}

	for index, v := range arr3 {
		t.Log(index, v)
	}

	for _, v := range arr3 {
		t.Log(v)
	}
}

func TestArraySection(t *testing.T) {
	arr3 := [...]int{1, 2, 3, 4, 5}

	arr3_sec := arr3[1:3]
	t.Log(arr3_sec)

}

func Test2WArray(t *testing.T) {
	var array1 [][]string
	var tmp []string
	var tmp1 []string

	tmp = append(tmp, "001")
	tmp = append(tmp, "002")
	tmp = append(tmp, "003")

	tmp1 = append(tmp1, "11")
	tmp1 = append(tmp1, "22")

	array1 = append(array1, tmp)
	array1 = append(array1, tmp1)
	fmt.Println(array1)
	fmt.Println(array1[0])
	fmt.Println(array1[0][1])
}
