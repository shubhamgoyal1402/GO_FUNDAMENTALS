package main

import "fmt"

//Basic syntax of struct
type student struct {
	id   int
	name string
}

func main() {
	//Initializing a struct
	stu1 := student{
		id:   1,
		name: "Shubham",
	}

	fmt.Println(stu1)
}
