package main

import (
	"fmt"
	"math/rand"
)

func init(){
	fmt.Println("This is starts from initialization")
}

func main() {
	/*for i:=0;i<=99;i++{
		fmt.Println(i)
	}

	for j:=0;j<100;j++{
		x:=rand.Intn(100)
		y:=rand.Intn(100)
		fmt.Println(j, " " ,x, " " , y)
	}*/

	// for i:=0;i<10;i++{
	// 	x:=rand.Intn(5)
	// 	switch x{
	// 	case 1:
	// 		fmt.Println(x)
	// 	case 2:
	// 		fmt.Println(x)
	// 	default:
	// 		fmt.Println("None")
	// 	}
	// }

	i:=rand.Intn(19)
	for i<=20{
		fmt.Println(i)
		i++
	}   

	// for{
	// 	if i>=10{
	// 		break
	// 	}
	// 	fmt.Println()
	
	// } 

	//for range
	// a:=[]int{12,3,5,7,8}
	// for b,c := range a{
	// 	fmt.Println("index ",b,"value ", c)
	// }

	m:=map[string]int{
		"kv":21,
		"chaitanya":22,
	}

	for k,v:=range m{
		fmt.Println("key ", k , "value ",v)
	}

	age:=m["kv"]
	fmt.Println(age)
//comma ok idiom
	if v,ok:=m["q"];!ok{
		fmt.Println("there is no q",v)
	}

	if v,ok:=m["kv"];ok{
		fmt.Println("kv is there",v)
	}


	
}