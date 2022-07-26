package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1) //this method watch how many go routines run and avoid that program finish without the number of go routines run!
	go sayHello("Hi!")
	sayBye("Bye!")
	wg.Wait() //this method watch the number of go routines we have and when this is 0 finish the program!
}

func sayHello(s string) {
	fmt.Println("s")
	wg.Done() //this method subtract 1 to number the go routines set in the add method
}

func sayBye(s string) {
	fmt.Println(s)
}
