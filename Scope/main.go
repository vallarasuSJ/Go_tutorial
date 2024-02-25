package main

import "fmt"

var x=45
const y=41

func main()  {
	z:=42
	fmt.Printf("The value of z is %v and the type is %T\n",z,z)
	fmt.Printf("The value of z is %v and the type is %T\n",x,x)
	fmt.Printf("The value of z is %v and the type is %T\n",y,y)
}