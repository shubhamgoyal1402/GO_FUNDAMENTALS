package main

import "fmt"

type user struct {
	name string
	age  int
}

func main() {
	detail := map[int]user{
		1: {
			name: "shubham",
			age:  12,
		},
		2: {
			name: "Tisha",
			age:  18,
		},
		3: {
			name: "satvik",
			age:  19,
		},
		4: {"kiran", 52},
	}

	value, ok := detail[7] // prints empty name with age 0  with false

	fmt.Println(value, ok)
	value, ok = detail[3]

	fmt.Println(value, ok) // prints satvik and age 19 with true
}
