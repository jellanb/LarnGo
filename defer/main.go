package main

import "fmt"

func main() {
	defer foo() //to defer set a function to execute with last function and bar execute first
	bar()
	//so the result is print:
	//bar
	//foo
}

func foo() {
	fmt.Println("foo")
}

func bar() {
	fmt.Println("bar")
}
