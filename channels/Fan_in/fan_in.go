package main

import (
	"fmt"
	"sync"
)

func main() {
	pair := make(chan int)  //declaring chanel to send with parameter empty
	odd := make(chan int)   //declaring chanel to send with parameter empty
	fanIn := make(chan int) //declaring chanel to send with parameter empty

	//sending
	go send(pair, odd) //sending channel with parameter

	received(pair, odd, fanIn) //sending channel with parameter

	for v := range fanIn {
		fmt.Println(v)
	}

	fmt.Println("Finalizing!")
}

func send(p, o, e chan<- int) { //receiving parameter empty to set values in looping
	for j := 0; j < 100; j++ {
		if j%2 == 0 {
			p <- j
		} else {
			o <- j
		}
	}
	e <- 0
}

func received(p, o, e <-chan int, fanin chan<- int) { //receiving parameter empty to read values in looping
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for v := range p {
			fanin <- v
		}
		wg.Done()
	}()

	go func() {
		for v := range o {
			fanin <- v
		}
		wg.Done()
	}()

	wg.Wait()
	close(fanin)
}
