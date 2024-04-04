package main

import "fmt"



func main() {

	x := [5]int{}
	for i := 0; i < 5; i++ {
		x[i] = i
	}
	for _, v := range x {
		fmt.Println(v)
	}
}
