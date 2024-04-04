package main

import "fmt"

func main() {

	c := gen()

	receive(c)

	fmt.Println("About to exit")

}

func receive(c <-chan int) {
	for v := range c {
		fmt.Println(v)
	}
}

func gen() <-chan int {
	c1 := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c1 <- i
		}

		close(c1)
	}()
	return c1
}
