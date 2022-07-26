package main

import "fmt"

func main() {
	channel := make(chan int)

	go func() {
		channel <- 42
	}()

	fmt.Println(<-channel)
}
