package main

import "fmt"

// BINARY SEARCH {1,2,3,4,5,6,7,8}

func binary_search(target int, nums []int) int {

	start := 0
	end := len(nums) - 1

	for start < end {
		mid := start + (end-start)/2

		if nums[mid] == target {
			return mid // element founded { return index of element}
		}
		if nums[mid] > target {
			end = mid - 1
		}
		if nums[mid] < target {
			start = mid + 1
		}

	}
	return -1 // if not found element
}
func main() {

	nums := []int{1, 2, 3, 4, 5, 6, 7, 8}

	fmt.Println(binary_search(4, nums))  // search for element 4 { return 3 }
	fmt.Println(binary_search(50, nums)) // search for element 4 { return -1 }

}
