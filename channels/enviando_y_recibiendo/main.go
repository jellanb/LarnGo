package main

import "fmt"

func main() {
	c := make(chan int)

	go sending(c)

	receiving(c)

	fmt.Println("Finalised!")
}

func sending(c chan<- int) {
	c <- 42
}

func receiving(c <-chan int) {
	fmt.Println(<-c)
}
