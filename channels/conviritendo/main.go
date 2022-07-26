package main

import "fmt"

func main() {
	c := make(chan int)
	cr := make(<-chan int) //receive
	cs := make(chan<- int) //send

	fmt.Println("---------")
	fmt.Printf("c\t%T\n", c)
	fmt.Printf("c\t%T\n", cr)
	fmt.Printf("c\t%T\n", cs)

	fmt.Println("---------")
	fmt.Printf("c\t%T\n", (<-chan int)(c)) //convert from bi-direction to receive only
	fmt.Printf("c\t%T\n", (chan<- int)(c)) //convert from bi-direction to send only

	//fmt.Printf("c\t%T\n", (chan int)(cr)) //cannot use convert in this situation
	//fmt.Printf("c\t%T\n", (chan int)(cs)) //cannot use convert in this situation
}
