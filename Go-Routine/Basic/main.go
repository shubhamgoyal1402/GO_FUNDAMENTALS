package main

import (
	"fmt"
	"time"
)

func test(x int) {

	go func() {
		time.Sleep(time.Millisecond * 100) // execute this command after 1 sec makes the sent statemnt work
		fmt.Printf("received %v\n", x)
	}()
	fmt.Printf("sent %v\n", x)
}
func main() {

	arr := []int{1, 2, 3, 4, 5, 6}

	for num := range arr {
		test(num)
		time.Sleep(time.Millisecond * 150) // it ensures it doesn't move to the second immediately
	}

}
