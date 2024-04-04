package main

import (
	"fmt"
	"math/rand"
)

func init() {
	fmt.Println("This is starts from initialization")
}

func main() {
	x := rand.Intn(250)
	switch {
	case x <= 100:
		fmt.Println("less than 100")
	case x >= 100 && x <= 200:
		fmt.Println("between 100 and 200")
	default:
		fmt.Println("greater than 200")
	}

	//if condition
	a := rand.Intn(10)
	b := rand.Intn(10)
	if a < 4 && b < 4 {
		fmt.Println("Both are less than 4")
	} else if a > 6 && b > 6 {
		fmt.Println("Both are greater than 6")
	} else {
		fmt.Println("None")
	}

	//switch
	switch {
	case a < 4 && b < 4:
		fmt.Println("Both are less than 4")
		fallthrough
	case a > 6 && b > 6:
		fmt.Println("Both are greater than 6")
	default:
		fmt.Println("None")
	}

}
