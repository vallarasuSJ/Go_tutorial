package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {

	//prints no of arch,os,cpu,goroutine
	fmt.Println("Os\t", runtime.GOOS)
	fmt.Println("ARCH\t", runtime.GOARCH)
	fmt.Println("CPU\t", runtime.NumCPU())
	fmt.Println("Goroutine\t",runtime.NumGoroutine())

	fmt.Println("---------------------------------------------") 

	wg.Add(1)
	go foo()
	// foo()
	bar()

	fmt.Println("CPU\t", runtime.NumCPU())
	fmt.Println("Goroutine\t",runtime.NumGoroutine())
	wg.Wait()

}

func foo(){
	for i:=0;i<10;i++{
		fmt.Println("foo: ",i)
	}
	wg.Done()
}

func bar(){
	for i:=0;i<10;i++{
		fmt.Println("bar: ",i)
	}
}