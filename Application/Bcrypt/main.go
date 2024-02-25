package main

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func main() {

	s:=`password123` 
	bs,err:=bcrypt.GenerateFromPassword([]byte(s),bcrypt.MinCost) 
	if err!=nil{
		log.Fatalf(err.Error())
	}
	fmt.Println(s)
	fmt.Println(bs)

	loginPassword:=`password13`
    err=bcrypt.CompareHashAndPassword(bs,[]byte(loginPassword)) 

	if err!=nil{
		log.Fatalf("You can't login")
	}

	fmt.Println("You logged In")
	
}