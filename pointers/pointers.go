package main

import (
	"fmt"
	"math"
)

type circulo struct {
	radio float64
}

type forma interface {
	area() float64
}

func (c *circulo) area() float64 {
	return math.Pi * c.radio * c.radio
}

func info(s forma) {
	fmt.Println("area", s.area())
}

func main() {
	//when use pointers??
	//1: when you need change big data
	//2: when you need change value from a value
	a := 42
	fmt.Println(a)
	fmt.Println(&a)        //when we use $ this return the memory direction from memory RAM
	fmt.Printf("%T\n", &a) //this should return data type from variable a "pointer"

	var b *int = &a
	fmt.Println(b) //this should return a memory direction from pointer b

	//reference
	fmt.Println(*b) //return value from pointer b "42"
	*b = 42         //editing value from pointer, so a and b have value 42 now
	fmt.Println(*b) //42
	fmt.Println(*b) //42

	//calling a function with param pointer
	foo2(&a) //acceding to value from variable in pointer
	foo2(b)  //acceding to value from pointer

	c := circulo{5}
	info(&c) //method sets

}

func foo2(y *int) {
	println(y)
}
