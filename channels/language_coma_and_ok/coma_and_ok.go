package main

import "fmt"

func main() {
	c := make(chan int) //declaring a channel

	go func() {
		c <- 42  //sending a value 42 to channel
		close(c) //closing the channel, is this channel is not closed we cannot use language coma ok in line 16
	}()

	v, ok := <-c       //using language coma and ok to know if the channel is open or closed
	fmt.Println(v, ok) //return ok= true "channel is open" and v= 42

	v, ok = <-c        //using language coma and ok to know if the channel is open or closed  <---- line 16
	fmt.Println(v, ok) //return ok= false "channel is open" and v= 0
}