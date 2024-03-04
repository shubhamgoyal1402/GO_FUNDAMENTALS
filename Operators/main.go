package main

import "fmt"

func fizzbuzz() {
	for i := 1; i <= 100; i++ {

		if i%3 == 0 && i%5 == 0 { // and logical opeartor
			fmt.Println("FIZZBUZZ")
			continue
		}
		if i%3 == 0 {
			fmt.Println("FIZZ")
			continue
		}
		if i%5 == 0 { //modulo operator
			fmt.Println("BUZZ")
			continue
		}

		fmt.Println(i)

	}
}
func main() {
	fizzbuzz()
}
