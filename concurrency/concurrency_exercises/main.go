package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	//exercise-1
	fmt.Println("CPU: ", runtime.NumCPU())
	fmt.Println("Go routines : ", runtime.NumGoroutine())
	var wg sync.WaitGroup
	wg.Add(7)
	go func() {
		fmt.Println("Hello world!!")
		wg.Done()
	}()
	go func() {
		fmt.Println("Hello world!!")
		wg.Done()
	}()

	go func() {
		fmt.Println("Hello world!!!!")
		wg.Done()
	}()

	go func() {
		fmt.Println("Hello world!!!!")
		wg.Done()
	}()
	go func() {
		fmt.Println("Hello world!!!!")
		wg.Done()
	}()
	go func() {
		fmt.Println("Hello world!!!!")
		wg.Done()
	}()
	go func() {
		fmt.Println("Hello world!!!!")
		wg.Done()
	}()

	fmt.Println("CPU: ", runtime.NumCPU())
	fmt.Println("Go routines : ", runtime.NumGoroutine())
	wg.Wait()

	fmt.Println("CPU: ", runtime.NumCPU())
	fmt.Println("Go routines : ", runtime.NumGoroutine())
	fmt.Println("exit")

}
