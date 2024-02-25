package main

import (
	"fmt"
	
)

type person struct{
	name string
}

func (p person) speak() {
	fmt.Println(p.name," Can speak")
} 
func (p *person) loud() {
	fmt.Println(p.name," Can loud")
} 


type human interface{
	speak() 
	loud()
}

func saySomething(h human){
	h.speak()
	h.loud()

}

func main() {
	p1:=person{
		"luffy",
	}
	p1.speak()
	p1.loud()
	saySomething(&p1)                //cannot pass type value person, can only send pointer type value
	
}

