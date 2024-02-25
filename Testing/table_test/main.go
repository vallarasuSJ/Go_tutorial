package main

import "fmt"

func main() {
	fmt.Println("3+4 = ", mySum(3, 4))
}

func mySum(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}