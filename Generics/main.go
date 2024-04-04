package main

import (
	"fmt"
	// "golang.org/x/exp/constraints"
)

//type set interface
type numbers interface {
	int | float64
}

//type constraint
func addT[T numbers](a, b T) T {
	return a + b
}

func sub[T int | float32](a, b T) T {
	return a - b

}

func main() {

	a := 4.3
	b := 3.5
	fmt.Println(addT(a, b))
	fmt.Println(addT(4, 5))
	fmt.Println(sub(3, 2))

}
