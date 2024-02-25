package main

import (
	"fmt"
	"tutorial/Testing_exercises/02/quote"
	"tutorial/Testing_exercises/02/word"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso)) 

	for k,v:=range word.UseCount(quote.SunAlso){
		fmt.Println(v , k)
	}
}