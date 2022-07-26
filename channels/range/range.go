package main

import "fmt"

//range is a function to iterate in channel, but we need close this channel when we don´t have more values!
func main() {
	c := make(chan int)

	//sending
	go func() {
		for i := 0; i < 5; i++ { //adding values to channel
			c <- i //adding
		}
		close(c) //closing channel to allow read and not lock
	}()

	for v := range c { //looping in channel to show values
		fmt.Println(v)
	}

	fmt.Println("Finish.")
}
