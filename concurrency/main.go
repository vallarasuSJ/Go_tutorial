package main

import (
	"fmt"
	"runtime"
	"sync"
)

//race condition
func main() {

	a := 100

	count := 0

	var wg sync.WaitGroup

	wg.Add(a)

	for i := 0; i < a; i++ {
		go func() {
			b := count
			runtime.Gosched()
			b++
			count = b
			fmt.Println(count)
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println("exit")

}
