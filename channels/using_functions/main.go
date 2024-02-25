package main

import "fmt"

func main() {
	c := make(chan int)

	//send
	go foo(c)

	//receive
	bar(c)

	fmt.Println("About to exit")

}

//converting bidirectional to send only type channel
func foo(c chan<- int){
	c<-43             //sending 43 to channel
}

//converting bidirectional to receive only type channel
func bar(c <-chan int){
	fmt.Println(<-c)           //receiving and print the output
}