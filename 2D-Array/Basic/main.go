package main

import "fmt"

func matrix_fn(row, col int) {

	matrix := [][]int{} // initializing a 2D Array

	for i := 0; i < row; i++ {
		row_line := make([]int, 0)
		for j := 0; j < col; j++ {
			row_line = append(row_line, i*j) // appending elements of a row

		}

		matrix = append(matrix, row_line) // appending the row to the matrix
		fmt.Println(row_line)

	}

}
func main() {
	matrix_fn(10, 10)

}
