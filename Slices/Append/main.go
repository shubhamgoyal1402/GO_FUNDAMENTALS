package main

import "fmt"

func main() {

	slice1 := []int{1, 2, 3, 4, 5}
	fmt.Println(slice1)

	slice1 = append(slice1, 6, 7) // append elements in a slice
	fmt.Println(slice1)

	slice2 := []int{10, 20, 30, 40, 50}

	slice1 = append(slice1, slice2...) // append other slice in other slice

	fmt.Println(slice1)
}
