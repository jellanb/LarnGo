package main

import "fmt"

func main() {
	channel := make(chan int)

	channel <- 42

	fmt.Println(channel)
}