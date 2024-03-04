package main

import "fmt"

//nums is a slice of integer
func sum(nums ...int) int {
	sum2 := 0
	for i := 0; i < len(nums); i++ {
		sum2 += nums[i]
	}
	return sum2
}

func main() {
	// by using nums ...int we can manually pass the slice
	fmt.Println(sum(1, 2, 3, 4))
}
