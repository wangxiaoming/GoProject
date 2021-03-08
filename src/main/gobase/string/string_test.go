package main

import (
	"fmt"
	"testing"
)

func TestStrig(t *testing.T) {
	str := "aba"
	fmt.Println(isPa(str))
}
func isPa(s string) bool {
	start := 0
	end := len(s) - 1

	for start < end {
		if s[start] == s[end] {
			start++
			end--
		} else {
			return false
		}
	}
	return true
}
