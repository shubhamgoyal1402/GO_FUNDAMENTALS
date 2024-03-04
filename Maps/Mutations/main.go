package main

import "fmt"

func main() {

	names := map[int]string{

		1: "shubham",
		2: "Ayush",
		3: "Shreyansh",
		4: "Rudranil",
	}

	// adding an element
	names[5] = "navya"

	// accesing an element
	fmt.Println(names[5]) // should return navya
	fmt.Println(names[7]) // return nothing

	// deleting an element
	delete(names, 1) // deletimg elemnt with key 1

	// check if an element exists
	value, ok := names[2] //value should be ayush and ok should be true

	fmt.Println(value, ok)
	value, ok = names[1] //value should be empty and ok should be false
	fmt.Println(value, ok)
}
