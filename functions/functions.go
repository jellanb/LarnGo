package main

import (
	"fmt"
	"sort"
)

type person struct {
	name string
	last string
}

type secretAgent struct {
	person
	lpm bool
}

func (a secretAgent) present() {
	fmt.Println("Hi i am", a.name, a.last, "the agent secret is present")
}

func (p person) present() {
	fmt.Println("Hi i am", p.name, p.last, "the person i present")
}

type human interface {
	present()
}

func main() {
	as1 := secretAgent{
		person: person{
			name: "jellan",
			last: "bozo",
		},
		lpm: true,
	}
	as1.present()

	p1 := person{
		name: "jellan",
		last: "bozo",
	}
	p1.present()
	//or
	as1.person.present()

	//anonymous function
	func() {
		fmt.Println("this is a anonymous function!")
	}()
	//anonymous function with params
	func(x int) {
		fmt.Println("this is a anonymous function! with params:", x)
	}(42)
	//expressions function
	fn := func(x int) {
		fmt.Println("this is a anonymous function! with params:", x)
	}
	fn(12)
	//returning a function into of function
	summ()
	//callback
	//a callback is a function this is call with other function how param
	message(callback)
	//closure
	c := incrementor() //closure
	fmt.Println(c())   //this function should return 0
	fmt.Println(c())   //this function should return 1
	fmt.Println(c())   //this function should return 2
	//sort: is a function that is defined to order a slice
	xc := []int{2, 7, 3, 42, 99, 18, 16, 56, 12}
	fmt.Println(xc)
	sort.Ints(xc)   //in this point the function sort order by int in our pointer array
	fmt.Println(xc) //when we print our xc slice show the array order in low to height
	//when we have other data type should use the sort to data type by example string:
	xs := []string{"Jame", "Q", "M", "Moneypenny", "Dr No"}
	fmt.Println(xs)
	sort.Strings(xs)
	fmt.Println(xs) //return the slice xs order by alphabetic

}

func summ() func() int {
	return makeThen
}

func makeThen() int {
	return 10
}

func message(fn func()) {
	fn()
}

func callback() {
	fmt.Println("yo soy un callback")
}

func incrementor() func() int {
	var o int
	return func() int {
		o++
		return o
	}
}
