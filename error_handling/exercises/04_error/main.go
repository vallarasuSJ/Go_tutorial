package main

import (
	"errors"
	"fmt"
	"log"
)

type sqrtError struct{
	lat string
	long string
	err error
}

//custom error
func (se sqrtError) Error() string{
	return fmt.Sprintf("Here is the error: %v %v %v",se.lat,se.long,se.err)
}

func main() {
	_,err:=sqrt(-10.23)
	if err!=nil{
		log.Println(err)
	}
}

func sqrt(f float64)(float64,error){
	if f<0{
		e:=errors.New("need more coffee")
		return 0,sqrtError{"30,3434 N","99.3434 W",e}
	}
	return 43,nil

}