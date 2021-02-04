package gobase

import "fmt"

func Deferprint(arg [5]struct{}) {
	for i := range arg {
		fmt.Print(i, " ")
	}
}

func Deferprint1(arg [5]struct{}) {
	for i := range arg {
		defer func() { fmt.Print(i, " ") }()
	}
}
