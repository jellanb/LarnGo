package main

import "fmt"

func main() {
	channel := make(chan<- int, 2)

	channel <- 42
	channel <- 43

	fmt.Println(<-channel) //this error is for my channel is only create to send
	fmt.Println(<-channel) //this error is for my channel is only create to send

}
