package main

import "fmt"

func main() {
	pair := make(chan int) //declaring chanel to send with parameter empty
	odd := make(chan int)  //declaring chanel to send with parameter empty
	exit := make(chan int) //declaring chanel to send with parameter empty

	//sending
	go send(pair, odd, exit) //sending channel with parameter

	received(pair, odd, exit) //sending channel with parameter

	fmt.Println("Finalizing!")
}

func send(p, o, e chan<- int) { //receiving parameter empty to set values in looping
	for j := 0; j < 100; j++ {
		if j%2 == 0 {
			p <- j
		} else {
			o <- j
		}
	}
	e <- 0
}

func received(p, o, e <-chan int) { //receiving parameter empty to read values in looping
	for {
		select { //selecting values if are pair or odd or exit
		case v := <-p:
			fmt.Println("From channel pair:", v)
		case v := <-o:
			fmt.Println("From channel pair:", v)
		case v := <-e:
			fmt.Println("From channel pair:", v)
			return
		}
	}
}
