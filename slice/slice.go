package main

import "fmt"

func main() {
	//declaring a slice
	x := []int{1, 2, 3, 4, 5}
	fmt.Println(x)
	//for range
	//i = index from slice
	//v = value of the position slice
	for i, v := range x {
		fmt.Println(v)
		fmt.Println(i)
	}
	//dividing a slice
	d := x[2:5]
	fmt.Println(d) //print 3, 4
	//adding a slice into other slice
	x = append(x, 6, 7, 8, 9, 10)
	fmt.Println(x) //print number from 1 to 10
	y := []int{11, 12, 13}
	x = append(x, y...)
	fmt.Println(x) //print number from 1 to 13
	//deleting value from slice
	x = append(x[:3], x[5:]...)
	fmt.Println(x) //should return number from 1 to 13 without number 5
	//using make to create a slice
	m := make([]int, 10, 100)
	fmt.Println(m)
	//deploying a slice using a function
	sum(x...) //this function should print all values from slice one for one how if iterate using a for

}

func sum(x ...int) {
	fmt.Println(x)
}
