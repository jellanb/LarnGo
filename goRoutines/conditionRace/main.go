package main

import (
	"fmt"
	"runtime"
	"sync"
)

//data race is when we have several go routines acceding to a value or variable, each go routines shared this value or variable and access simultaneously,
//so how we can controller that this access not broken our program!! ***using mutex method from sync***!!
//mutex block the access to a value or variable when a go routines is operating in here

func main() {
	fmt.Println("CPUs number", runtime.NumCPU())
	fmt.Println("CPUs number", runtime.NumGoroutine())
	counter := 0
	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)
	var mu sync.Mutex //declare the variable mutex
	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			v := counter
			v++
			counter = v
			runtime.Gosched() //Gosched yields (free variable) the processor, allowing other goroutines to run
			wg.Done()
			mu.Unlock() //mu Lock and Unlock block in the scope that is declare all variables shared by go routines
			//TODO: see in documentation the method RWMutex
		}()
		fmt.Println("CPUs number", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("counter;", counter)
}
