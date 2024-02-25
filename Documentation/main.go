package main

import (
	"fmt"
	"tutorial/Documentation/dog"
)

type canine struct {
	name string
	age  int
}

func main() {

	fido := canine{
		name: "Fide",
		age:  dog.Years(10),
	}
	fmt.Println(fido)

}