package main

import (
	"fmt"
	"strings"
)

func main() {
	result := repeatedSubstringPattern("abcabc")
	fmt.Println(result)
}

func repeatedSubstringPattern(s string) bool {

	str := s + s

	fmt.Println(str)
	content := str[1 : len(str)-1]
	return strings.Contains(content, s)
	//fmt.Println(content)
}
