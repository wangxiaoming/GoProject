package main

import (
	"fmt"
	"sort"
)

func main() {

	//var  arr = [4][2]int{{3, 4}, {1, 2}, {5, 6}, {7, 8}}

	arr := [...][]int{{3, 4}, {1, 2}, {5, 6}, {7, 8}}
	fmt.Println(arr)
	var array []int = make([]int, 5)
	array[0] = 2
	array[1] = 3
	array[2] = 1
	array[3] = 5
	array[4] = 4

	/*	sort.Slice(arr, func(i , j int) bool {
		 if arr[i][0] == arr[j][0] {
		 	return arr[i][1] > arr[j][1]
		 } else {
		 	return arr[i][0] > arr[j][0]
		 }
	})*/
	sort.Slice(array, func(i, j int) bool {
		return array[i] > array[j]
	})

	fmt.Println(array)

}
