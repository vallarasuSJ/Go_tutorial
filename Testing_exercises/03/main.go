package main

import (
	"fmt"
	"tutorial/Testing_exercises/03/mymath"
)

func main() {
	x := gen()
	fmt.Println(x)
	for _, v := range x {
		fmt.Println(mymath.CenterAvg(v))
	}
}

func gen() [][]int {
	a := []int{3, 2, 4, 5, 3}
	b := []int{4, 55, 4, 5, 4, 5}
	c := []int{54, 234, 5, 33, 44, 3}
	e := [][]int{a, b, c}
	return e
}
