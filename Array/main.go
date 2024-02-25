package main

import "fmt"

type Rectangle struct {
	width, height float64
}

func main() {

	rec := Rectangle{
		width:  30,
		height: 20,
	}
	// fmt.Println(rec.area())

	fmt.Println(rec.add(10,10))
}
func (r Rectangle) area() float64 {
	return r.width * r.height
}

func (r Rectangle) add(a, b int) int {
	
    return a + b
}
