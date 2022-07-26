package main

import (
	"fmt"
)

//this is pattern design of concurrence!
//when we run our program this is sequential, so our program run line by line, but, a go routines is going run his time
//When we need use a go routines? A: When we need to do different tasks at the same time

func main() {
	go sayHello("Hi!") //The go routines run in sync threads so, how in tha function "say" we have a time sleep our program not wait to finish in this point!
	sayBye("Bye!")     //In this point our program run after the las code line and finish, while our go routines is not finish yet
} //when main finish the program is closed, so  the go routines don´t have time to run and the program finish! in this scenery we need sync control method to controller this problem! see waitGroup!

func sayHello(s string) {
	fmt.Println(s)
}

func sayBye(s string) {
	fmt.Println(s)
}
