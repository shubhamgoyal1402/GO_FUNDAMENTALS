package main

import "fmt"

func main() {

	// Declaring a slice without an array

	mySlice := make([]int, 9) // make keyword make(type of slice , length of slice)

	for i := 0; i < len(mySlice); i++ {
		mySlice[i] = i

	}

	fmt.Println(mySlice)
	arr := [8]int{1, 2, 3, 4, 5, 6, 7, 8}
	// declaration of a slice from an array
	// from index arr[i:j] return from i to j-1
	slice_arr := arr[2:7] //b={3,4,5,6,7}

	// from index arr[i:] return from i to end of array
	slice_arr2 := arr[2:] //b={3,4,5,6,7,8}

	// from index arr[:j] return from 0 to j-1
	slice_arr3 := arr[:5] //b={1,2,3,4,5}

	// from index arr[:] return entire array
	slice_arr4 := arr[:] //b={1,2,3,4,5,6,7,8}

	fmt.Println(slice_arr)
	fmt.Println(slice_arr2)
	fmt.Println(slice_arr3)
	fmt.Println(slice_arr4)
}
