package main

import "fmt"

func main() {
	channel := make(chan int, 2)

	channel <- 42
	channel <- 43

	fmt.Println(<-channel)
	fmt.Println(<-channel)
}
