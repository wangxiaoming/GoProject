package main

import "fmt"
import "math"

func main() {
	// 匿名函数
	getSqrt := func(a float64) float64 {
		return math.Sqrt(a)
	}

	fmt.Println(getSqrt(4))
}
