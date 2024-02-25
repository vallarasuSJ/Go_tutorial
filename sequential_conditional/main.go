package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Intn(250)

	if(x<=100){
		fmt.Println("Less than 100");
	}else if x>=100 && x<=200{
		fmt.Println("between 101 and 200")
	}else{
		fmt.Println("equal or greater than 200")
	}
	y:=rand.Intn(3)
	fmt.Println(y)

}