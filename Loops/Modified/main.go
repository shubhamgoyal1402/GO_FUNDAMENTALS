package main

import "fmt"

//how many numbers will add up to or greater than n (startinf from 1 and incrementing by 1 )
func calculate(n int) int {
	sum := 0
	for i := 1; ; i++ {
		sum += i
		if sum >= n {
			return i
		}

	}
}
func main() {

	fmt.Println(calculate(11))

}
