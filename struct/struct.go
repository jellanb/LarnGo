package main

import "fmt"

type persona struct { //define a type of struct
	name     string
	lastName string
}

func main() {
	p1 := persona{
		name:     "Jellan",
		lastName: "Bozo",
	}
	p2 := persona{
		name:     "Maily",
		lastName: "Alvarez",
	}

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p1.name)
	fmt.Println(p1.lastName, p1.name)

	//Anonimus struct
	p3 := struct {
		fist string
		last string
		age  int
	}{
		fist: "jellan",
		last: "bozo",
		age:  31,
	}

	fmt.Println(p3)
}
