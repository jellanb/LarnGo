package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println("CPUs number", runtime.NumCPU())
	fmt.Println("CPUs number", runtime.NumGoroutine())
	var counter int64
	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			runtime.Gosched()
			fmt.Println("Counter :", atomic.LoadInt64(&counter))
			wg.Done()
		}()
		fmt.Println("CPUs number", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("counter;", counter)
}
