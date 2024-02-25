package main

import (
	"fmt"
)

type customError struct{
	info string
}

//custom error
func (ce customError) Error() string{
	return fmt.Sprintf("Here is the error: %v",ce.info)
}

func main() {
	c1:=customError{
		"need more coffee",
	}
	foo(c1)
}

func foo(e error){
	fmt.Println("foo ran -",e,"\n",e.(customError).info)

}