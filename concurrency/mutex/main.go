package main

import (
	"fmt"
	"sync"
)

//race condition
func main() {

	a := 100

	count := 0

	var wg sync.WaitGroup
	var m sync.Mutex

	wg.Add(a)

	for i := 0; i <= a; i++ {
		go func() {
			m.Lock()
			b := count
			b++
			count = b
			m.Unlock()
			fmt.Println(count)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(count)

	fmt.Println("exit")

}
