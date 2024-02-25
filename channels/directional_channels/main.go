package main

import "fmt"

func main() {
	c:=make(chan int)
	cs := make(chan<- int)   //send channel
	cr := make(<-chan int)   //receive channel

	fmt.Printf("c  %T\n",c)
	fmt.Printf("cs %T\n",cs)
	fmt.Printf("cr %T\n",cr)

}