package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

//race condition
func main() {

	a := 100

	var count int64

	var wg sync.WaitGroup

	wg.Add(a)

	for i := 0; i < a; i++ {
		go func() {
			atomic.AddInt64(&count, 1)
			fmt.Println(atomic.LoadInt64(&count))
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(count)

	fmt.Println("exit")

}
