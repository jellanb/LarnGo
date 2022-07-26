package main

import "fmt"

func main() {
	channel := make(chan int, 1)

	channel <- 42 //when we use a channel with buffer this make a value when finish
	//channel <- 43//this is other buffer bot we do declare a channel only 1 channel, we cannot use other buffer

	fmt.Println(<-channel)
}
