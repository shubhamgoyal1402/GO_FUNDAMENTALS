package main

import "fmt"

func main() {
	// initializing arrays

	nums := [10]int{25, 56, 1, 2, 7, 46, 3, 6, 85, 106}

	// check number even or odd in an array
	for i := 0; i < len(nums); i++ {
		if nums[i]%2 == 0 {
			fmt.Printf("%v is even number\n", nums[i])
		} else {
			fmt.Printf("%v is odd number\n", nums[i])
		}
	}

}
