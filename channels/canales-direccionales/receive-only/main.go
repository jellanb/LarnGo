package main

import "fmt"

func main() {
	channel := make(<-chan int, 2)

	channel <- 42 //this error is for my channel is only create to received
	channel <- 43 //this error is for my channel is only create to received

	fmt.Println(<-channel)
	fmt.Println(<-channel)

}
