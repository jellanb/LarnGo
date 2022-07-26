package main

import "fmt"

func main() {
	//Map declaration
	m := map[string]int{
		"Batman": 32,
		"Robin":  27,
	}
	fmt.Println(m)

	fmt.Println(m["Batman"])

	fmt.Println(m["Superman"]) // When the key not fount then return 0 by default

	//language ok
	v, ok := m["Superman"] //When they key is found ok is true else ok is false
	fmt.Println(v)
	fmt.Println(ok)

	if v, ok := m["Superman"]; ok {
		fmt.Println("Printing from if", v) // This should never execute
	}

	//Add new value to map
	m["Gatuvela"] = 31
	fmt.Println(m)
	//printing all values from map using for
	for key, value := range m {
		fmt.Println(key, value)
	}

	//Delete a value from map
	delete(m, "Gatuvela")
	fmt.Println(m)
	//Using language OK
	if v, ok := m["Gatuvela"]; ok {
		fmt.Println("Se borro la llave con valor", v) //This code never execute why Gatuvela is false
		delete(m, "Gatuvela")
	}
	fmt.Println(m)
}
