package main

import (
	"fmt"
	"runtime"
	"tutorial/Deck"
)

type super struct{
	name string
	main1.Hello
}



func main() {
	a:=10
	b:=float32(a)
	fmt.Println(b) 
    main1.NewDeck()
	fmt.Println(main1.GetName()) 
	fmt.Println(runtime.GOARCH)
}
