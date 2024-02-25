package main

import (
	"fmt"
	"tutorial/Testing_exercises/01/Finished/cat"
)

type canine struct {
	name string
	age  int
}

func main() {
	fido := canine{
		name: "Fido",
		age:  cat.Years(10),
	}
	fmt.Println(fido)
	fmt.Println(cat.YearsTwo(30))
}