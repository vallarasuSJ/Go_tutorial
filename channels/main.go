package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		c <- 42
		c <- 43
		c <- 45
		c <- 46

	}()

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
}
