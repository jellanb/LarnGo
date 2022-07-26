package main

import (
	"encoding/json"
	"fmt"
)

type person1 struct {
	Name     string
	LastName string
	Age      int
}

type personFromJson struct {
	Name     string `json:"Name"`
	LastName string `json:LastName`
	Age      int    `json:Age`
}

func main() {
	p1 := person1{
		Name:     "james",
		LastName: "bond",
		Age:      32,
	}
	p2 := person1{
		Name:     "miss",
		LastName: "moneyPenny",
		Age:      27,
	}
	allPerson := []person1{p1, p2}
	fmt.Println(allPerson)
	//convert to JSON:
	bs, err := json.Marshal(allPerson)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs)) //should print object in format json

	//convert from JSON to object:
	j := `[{"Name":"james","LastName":"bond","Age":32},{"Name":"miss","LastName":"moneyPenny","Age":27}]`
	bs = []byte(j)
	var person []personFromJson
	err = json.Unmarshal(bs, person)
	if err != nil {
		fmt.Println("error", err)
	}
	fmt.Printf("%T\n", person)
	fmt.Println(person)
}
