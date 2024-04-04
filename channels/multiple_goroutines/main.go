package main

import (
	"fmt"
	"runtime"
)

func main() {

	c := make(chan int)

	for j := 0; j < 10; j++ {
		go func() {
			for i := 0; i < 10; i++ {
				c <- i
			}
		}()
		fmt.Println("Routines : ", runtime.NumGoroutine())

	}

	for k := 0; k < 100; k++ {
		fmt.Println(k, " ", <-c)
	}
	fmt.Println("Routines : ", runtime.NumGoroutine())

	fmt.Println("About to exit")

}
