package main

import "fmt"

func main() {
	c1 := make(chan int)

	//buffer
	c2:=make(chan int,1)
	c2<-45

	//anonymous func
	go func(){
		c1<-43
	}()



	fmt.Println(<-c1)
	fmt.Println(<-c2)

}

