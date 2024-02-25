package main

import ("fmt"
)

func main() {
	//exercise one
	x:=[5]int{}
	for i:=0;i<5;i++{
		x[i]=i
	}
	for _,v:=range x{
		fmt.Println(v)
	}

	//exercise two
	a:=[]int{5,6,5,64,54,54,75,23,64,23,75,43,65}
	for i,v:=range a{
		fmt.Println(i," ",v)
	}
	fmt.Printf("%T %v\n",a,a)
	
	//exercise three-slice
	fmt.Println(a[2:4])
	fmt.Println(a[:5])
	fmt.Println(a[4:7]) 

	//exercise four-slice
	a=append(a,45,867,97,67,54,3,45,43,4)
	fmt.Println(a)

	//exercise five-slice
	// b:=make([]string,50) //this length is 50 and cap is 50 . It assigns all value to 0.
	b:=make([]string,0,50) //this length is 0 and cap is 50 
	b=append(b, ` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, `
	Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, `
	Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, `
	Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`,
	` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`,
	` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, `
	Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`,) 
	fmt.Println(b)
	fmt.Println(len(b))
	fmt.Println(cap(b)) 
	for _,v:=range b{
		fmt.Println("Values- ",v)
	}

	a1:=[]string{"a","b","c"}
	a2:=[]string{"d","e","f"}
	a3:=[][]string{a1,a2}
	for i:=0;i<len(a3);i++{
		for j:=0;j<len(a2);j++{
			fmt.Println(a3[i][j])
		}
	}

}